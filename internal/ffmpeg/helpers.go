package ffmpeg

import (
	"bytes"
	"os/exec"
)

func TranscodeToJPEG(b []byte) ([]byte, error) {
	cmd := exec.Command(defaults["bin"], "-hide_banner", "-i", "-", "-f", "mjpeg", "-")
	cmd.Stdin = bytes.NewBuffer(b)
	return cmd.Output()
}
