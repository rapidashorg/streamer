package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/rapidashorg/streamer"
)

const (
	workerNum = 2
)

func main() {
	unixNow := time.Now().Unix()

	wg := sync.WaitGroup{}

	for i := 1; i <= workerNum; i++ {
		wg.Add(1)
		go func(wn int) {
			log.Println("Worker", wn, "start")
			err := streamer.PushStream("bin/ffmpeg-darwin", "sample-video/test-video.mp4", fmt.Sprintf("rtmp://live-ingest.tokopedia.net/streamv/%d-%d", unixNow, wn))
			if err != nil {
				log.Println(err)
				return
			}
			log.Println("Worker", wn, "done")
			wg.Done()
		}(i)
	}
	wg.Wait()
}
