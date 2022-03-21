package streamer

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
)

func PushStream(ffmpegPath, videoPath, rtmpURL string) error {
	ffmpegCommand := "FFMPEG_PATH -re -stream_loop -1 -i VIDEO_PATH -c:v libx264 -c:a aac -f flv RTMP_URL"
	ffmpegCommand = strings.Replace(ffmpegCommand, "FFMPEG_PATH", ffmpegPath, 1)
	ffmpegCommand = strings.Replace(ffmpegCommand, "VIDEO_PATH", videoPath, 1)
	ffmpegCommand = strings.Replace(ffmpegCommand, "RTMP_URL", rtmpURL, 1)

	commands := strings.Split(ffmpegCommand, " ")

	var stderr bytes.Buffer
	cmd := exec.Command(commands[0], commands[1:]...)
	cmd.Stdout = &stderr

	err := cmd.Run()
	log.Println(stderr.String())
	if err != nil {
		return err
	}
	return nil
}
