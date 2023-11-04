package service

import (
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

	return cmd.Run()
}

func NewMergeAudioFileService() *MergeVideoAudioFileService {
	return &MergeVideoAudioFileService{}
}
