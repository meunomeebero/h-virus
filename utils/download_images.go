package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

func DownloadImages(images []string, path string) {
	wg := sync.WaitGroup{}

	for i, r := range images {
		wg.Add(1)

		go func(index int, url string) {
			defer wg.Done()

			res, err := http.Get(url)

			if err != nil {
				return
			}

			file, err := os.Create(fmt.Sprintf("%s/hentai-%d.jpeg", path, index))

			if err != nil {
				return
			}

			io.Copy(file, res.Body)

			defer file.Close()
		}(i, r)
	}

	wg.Wait()
}
