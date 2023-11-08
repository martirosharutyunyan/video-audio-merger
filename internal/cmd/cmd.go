package cmd

import (
	"context"
	"os"
)

func Execute(args []string, stdin, stdout, stderr *os.File) {
	rootCmd.AddCommand(mergeDirCmd)
	rootCmd.AddCommand(mergeFileCmd)

	mergeFileCmd.PersistentFlags().String("input-video", "", "video-audio-merger file --input-video {source_path} --input-audio {source_path} --output {output_path}")
	mergeFileCmd.PersistentFlags().String("input-audio", "", "video-audio-merger file --input-video {source_path} --input-audio {source_path} --output {output_path}")
	mergeFileCmd.PersistentFlags().String("output", "", "video-audio-merger file --input-video {source_path} --input-audio {source_path} --output {output_path}")

	mergeDirCmd.PersistentFlags().String("input", "", "video-audio-merger dir --input {source_path} --output {output_path}")
	mergeDirCmd.PersistentFlags().String("output", "", "video-audio-merger dir --input {source_path} --output {output_path}")
	mergeDirCmd.PersistentFlags().Int("p", 2, "parallel cpu core count")

	rootCmd.SetArgs(args)
	rootCmd.SetIn(stdin)
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	ctx := context.Background()
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}
