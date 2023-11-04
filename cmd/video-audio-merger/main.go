package main

import (
	"github.com/martirosharutyunyan/video-audio-merger/internal/cmd"
	"os"
)

func main() {
	cmd.Execute(os.Args[1:], os.Stdin, os.Stdout, os.Stderr)
}
