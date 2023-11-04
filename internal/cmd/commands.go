package cmd

import (
	"github.com/martirosharutyunyan/video-audio-merger/internal/service"
	"github.com/martirosharutyunyan/video-audio-merger/internal/utils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "video-audio-merger",
	Short: "Merge video and audio files with directory option file names must be {video_name}.mp4 and {video_name}_audio.mp4",
}

var mergeDirCmd = &cobra.Command{
	Use:   "dir",
	Short: "Merge directory of videos and audios",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
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
