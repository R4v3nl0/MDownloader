package downloader

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/R4v3nl0/MDownloader/config"
	"github.com/R4v3nl0/MDownloader/utils"
)

func Download(url string, cfg *config.Config) error {
	urlSplit := strings.Split(url, "/")
	movieName := urlSplit[len(urlSplit)-1]

	// get uuid and movie title, if title not found, use movieName instead
	uuid, title, err := getInfo(url, movieName, cfg)
	if err != nil {
		if err.Error() != "title not found" {
			return err
		} else {
			title = movieName
		}
	}

	if _, err := os.Stat(filepath.Join(cfg.SavePath, fmt.Sprintf("%s.mp4", title))); !os.IsNotExist(err) {
		return fmt.Errorf("%s already exists, skip download", movieName)
	}

	fmt.Printf("Downloading %s\n", title)

	// make directory to save temp files
	os.MkdirAll(filepath.Join(cfg.SavePath, title), os.ModePerm)

	sliceNum, quality, err := getVideoSliceNum(uuid, cfg)
	if err != nil {
		return err
	}

	qualityStr := fmt.Sprintf("%dp", quality)

	// download video slices
	err = downloadVideoSlice(uuid, title, qualityStr, sliceNum, cfg)
	if err != nil {
		return err
	}

	return nil
}

func downloadVideoSlice(uuid, title, quality string, sliceNum int, cfg *config.Config) error {
	fmt.Printf("uuid: %s, title: %s, quality: %s, sliceNum: %d\n", uuid, title, quality, sliceNum)
	var wg sync.WaitGroup

	cpuNum := runtime.NumCPU()
	wg.Add(cpuNum)

	sliceChan := make(chan int, cpuNum)
	go func() {
		// slice 0~sliceNum
		for i := 0; i <= sliceNum; i++ {
			sliceChan <- i
		}
		close(sliceChan)
	}()

	for i := 0; i < cpuNum; i++ {
		go func() {
			defer wg.Done()
			var retry int = 0

			for {
				sliceIndex, ok := <-sliceChan
				if !ok {
					// There is no slice to download, quit routine
					break
				}

				url := fmt.Sprintf("%s%s/%s/video%d.jpeg", cfg.Prefixes.Video, uuid, quality, sliceIndex)

				// Check if the slice file already exists
				filePath := filepath.Join(cfg.SavePath, title, fmt.Sprintf("video%d.jpeg", sliceIndex))
				if _, err := os.Stat(filePath); !os.IsNotExist(err) {
					// Slice file already exists, skip download
					continue
				}

				for retry = 0; retry < cfg.Requests.Retry; retry++ {
					// Request video slice
					content, err := getRequest(url, cfg)
					if err != nil {
						fmt.Printf("Download %s failed: %s, retry %d\n", url, err, retry)
						time.Sleep(time.Duration(cfg.Requests.Delay) * time.Second)
						continue
					}

					// Save video slice to file
					err = utils.SaveSliceFile(filePath, content)
					if err != nil {
						fmt.Printf("Save %s failed: %s, retry %d\n", filePath, err, retry)
						time.Sleep(time.Duration(cfg.Requests.Delay) * time.Second)
						continue
					}

					// Save success, break retry loop, retry < cfg.Requests.Retry
					break
				}

				// lost slice, quit routine
				if retry == cfg.Requests.Retry {
					fmt.Printf("Download %s failed: retry limit reached\n", url)
					break
				}
			}
		}()
	}

	wg.Wait()

	// Check if all slices are downloaded
	entries, err := os.ReadDir(filepath.Join(cfg.SavePath, title))
	if err != nil {
		return err
	}

	if len(entries) != sliceNum+1 {
		return fmt.Errorf("lost slice, please retry download")
	}

	// Merge slices to video
	err = utils.MergeSliceToVideo(title, sliceNum, cfg)
	if err != nil {
		return err
	}

	// Clean temp files
	os.RemoveAll(filepath.Join(cfg.SavePath, title))

	return nil
}
