package cmd

import (
	"github.com/martirosharutyunyan/video-audio-merger/internal/service"
	"github.com/martirosharutyunyan/video-audio-merger/internal/utils"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strconv"
)

var rootCmd = &cobra.Command{
	Use:   "video-audio-merger",
	Short: "Merge video and audio files with directory option file names must be {video_name}.mp4 and {video_name}_audio.mp4",
}

var mergeDirCmd = &cobra.Command{
	Use:   "dir",
	Short: "Merge directory of videos and audios",
	RunE: func(cmd *cobra.Command, args []string) error {
		input := cmd.Flag("input").Value.String()
		output := cmd.Flag("output").Value.String()
		parallelCoreCountString := cmd.Flag("p").Value.String()
		parallelCoreCount, err := strconv.Atoi(parallelCoreCountString)
		if err != nil {
			return err
		}

		if output == "" {
			output, err = utils.GenCopyPath(input)
			if err != nil {
				return err
			}
		}
		err = os.Mkdir(output, 0777)
		if err != nil {
			return err
		}

		input, err = filepath.Abs(input)
		if err != nil {
			return err
		}

		mergeDirService := service.NewMergeVideoAudioDirService(service.NewMergeAudioFileService())

		return mergeDirService.Merge(input, output, parallelCoreCount)
	},
}

var mergeFileCmd = &cobra.Command{
	Use:   "file",
	Short: "Merge video and audio file",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		inputVideo := cmd.Flag("input-video").Value.String()
		inputAudio := cmd.Flag("input-audio").Value.String()
		output := cmd.Flag("output").Value.String()

		if output == "" {
			output, err = utils.GenCopyPath(inputVideo)
			if err != nil {
				return err
			}
		}

		mergeFileService := service.NewMergeAudioFileService()
		return mergeFileService.Merge(inputVideo, inputAudio, output)
	},
}
