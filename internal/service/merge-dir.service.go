package service

import (
	"golang.org/x/sync/errgroup"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type IMergeVideoAudioDirService interface {
	Merge(source, output string, parallelCoreCount int) error
}

type MergeVideoAudioDirService struct {
	mergeVideoAudioFileService IMergeVideoAudioFileService
}

var _ IMergeVideoAudioDirService = MergeVideoAudioDirService{}

func (s MergeVideoAudioDirService) Merge(source, output string, parallelCoreCount int) error {
	group := errgroup.Group{}
	group.SetLimit(parallelCoreCount)
	err := filepath.WalkDir(source, func(path string, _ fs.DirEntry, err error) error {
		if source == path {
			return nil
		}

		if strings.Contains(path, "_audio.mp4") {
			return nil
		}

		info, err := os.Stat(path)
		if err != nil {
			return err
		}

		if runtime.GOOS == "windows" {
			path = strings.ReplaceAll(path, "\\", "/")
		}

		fileRelativePath := strings.Split(path, source)[1]

		var outputBuilder strings.Builder
		outputBuilder.WriteString(output)
		outputBuilder.WriteString(fileRelativePath)

		audioPath := strings.Replace(path, ".mp4", "_audio.mp4", -1)

		if info.IsDir() {
			err = os.Mkdir(outputBuilder.String(), 0777)
			if err != nil {
				return err
			}
		} else {
			group.Go(func() error {
				return s.mergeVideoAudioFileService.Merge(path, audioPath, outputBuilder.String())
			})
		}

		return err
	})

	if err != nil {
		return err
	}

	return group.Wait()
}

func NewMergeVideoAudioDirService(mergeVideoAudioFileService *MergeVideoAudioFileService) *MergeVideoAudioDirService {
	return &MergeVideoAudioDirService{mergeVideoAudioFileService: mergeVideoAudioFileService}
}
