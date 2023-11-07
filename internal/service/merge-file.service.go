package service

import (
	"os"
	"os/exec"
)

type IMergeVideoAudioFileService interface {
	Merge(sourceVideo, sourceAudio, output string) error
}

type MergeVideoAudioFileService struct{}

var _ IMergeVideoAudioFileService = MergeVideoAudioFileService{}

func (MergeVideoAudioFileService) Merge(sourceVideo, sourceAudio, output string) error {
	cmd := exec.Command("ffmpeg",
		"-i", sourceVideo,
		"-i", sourceAudio,
		"-c:v", "copy",
		"-c:a", "aac",
		//"-strict", "experimental",
		//"-map", "0:v:0",
		//"-map", "1:a:0",
		output,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		return err
	}

	//reader, err := cmd.StdoutPipe()
	//if err != nil {
	//	return err
	//}
	//
	//scanner := bufio.NewScanner(reader)
	//for scanner.Scan() {
	//	log.Println(scanner.Text())
	//}

	return cmd.Wait()
}

func NewMergeAudioFileService() *MergeVideoAudioFileService {
	return &MergeVideoAudioFileService{}
}
