package downloader

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/R4v3nl0/MDownloader/config"
)

func reverseSlice(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func getInfo(url, movieName string, cfg *config.Config) (uuid, title string, err error) {
	uuid = ""
	title = ""
	err = nil

	htmlData, err := getRequest(url, cfg)
	if err != nil {
		return
	}

	htmlText := string(htmlData)

	regexpUUID := regexp.MustCompile(cfg.Regexes.Data.Uuid)
	submatch := regexpUUID.FindStringSubmatch(htmlText)
	if len(submatch) < 2 {
		err = fmt.Errorf("uuid not found")
		return
	}

	// get uuid
	uuidReversed := strings.Split(submatch[1], "|")
	reverseSlice(uuidReversed)
	uuid = strings.Join(uuidReversed, "-")

	regexpTitle := regexp.MustCompile(cfg.Regexes.Data.Title)
	submatch = regexpTitle.FindStringSubmatch(htmlText)
	if len(submatch) < 2 {
		err = fmt.Errorf("title not found")
		return
	}

	title = strings.Replace(submatch[1], "&#039;", "'", -1)
	title = strings.Replace(title, "/", "_", -1)
	title = strings.Replace(title, "\\", "_", -1)

	if strings.Contains(movieName, "uncensored-leak") {
		title = fmt.Sprintf("[Uncensored] %s", title)
	}

	return
}

func getVideoSliceNum(uuid string, cfg *config.Config) (sliceNum, quality int, err error) {
	sliceNum = 0
	quality = 0
	err = nil

	playlistUrl := fmt.Sprintf("%s%s%s", cfg.Prefixes.Video, uuid, cfg.Suffixes.Playlist)
	htmlData, err := getRequest(playlistUrl, cfg)
	if err != nil {
		return
	}

	htmlText := string(htmlData)

	regexpResolution := regexp.MustCompile(cfg.Regexes.Data.Resolution)
	submatch := regexpResolution.FindAllStringSubmatch(htmlText, -1)
	if len(submatch) == 0 {
		err = fmt.Errorf("resolution not found")
		return
	}

	// get quality map
	// key: qualityY, value: qualityX
	// 1920x1080 -> 1080:1920
	qualityMap := make(map[int]int)
	for i := 0; i < len(submatch); i++ {
		qualityX, err := strconv.Atoi(submatch[i][1])
		if err != nil {
			fmt.Printf("convert resolution to int failed")
			return 0, 0, err
		}

		qualityY, err := strconv.Atoi(submatch[i][2])
		if err != nil {
			fmt.Printf("convert resolution to int failed")
			return 0, 0, err
		}

		qualityMap[qualityY] = qualityX
	}

	// get quality
	if cfg.Quality != 0 {
		// find the nearest quality
		if _, ok := qualityMap[cfg.Quality]; ok {
			quality = cfg.Quality
		} else {
			difference := 0
			for k := range qualityMap {
				if difference == 0 || difference > k-cfg.Quality {
					difference = k - cfg.Quality
					quality = k
				}
			}
		}
	} else {
		// get the highest quality
		for k := range qualityMap {
			if quality < k {
				quality = k
			}
		}
	}

	// get slice num
	htmlData, err = getRequest(fmt.Sprintf("%s%s/%dp/video.m3u8", cfg.Prefixes.Video, uuid, quality), cfg)
	if err != nil {
		return
	}

	htmlText = string(htmlData)

	regexpSliceNum := regexp.MustCompile(cfg.Regexes.Data.VideoSlice)
	// 倒着匹配
	submatch = regexpSliceNum.FindAllStringSubmatch(htmlText, -1)
	if len(submatch) == 0 {
		err = fmt.Errorf("slice regex not match")
		return
	}

	sliceNum, err = strconv.Atoi(submatch[len(submatch)-1][1])
	if err != nil {
		fmt.Println("convert slice num to int failed")
		return
	}

	return
}
