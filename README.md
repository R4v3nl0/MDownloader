# MDownloader

A tool for downloading videos from the "MissAV" website. The Go language version of "MissAV-Downloader".

MiyukiQAQ is the original author of this project. The original project is written in Python. I rewrote it in Go language.

Python Version: [MiyukiQAQ/MissAV-Downloader](https://github.com/MiyukiQAQ/MissAV-Downloader/tree/master)

## 📖 Instructions

```
Usage of MDownloader:
    MDownloader [-config/-c <config file path>] urls "url1,url2,url3..."

Global Options:
    -config, -c <config file path>    Specify the configuration file path. (default: "config.yaml")

Commands:
    urls
        "url1,url2,url3..."            The URLs of the videos you want to download.

Examples:
    MDownloader -c /mnt/mdownloader/config.yaml urls "https://missav.ai/vrtm-381-uncensored-leak"

```

## 📦 Installation

**Prerequisites**: [Go 1.23+](https://go.dev/dl/) is required to build the project.

```bash
git clone https://github.com/R4v3nl0/MDownloader.git && cd MDownloader
go build -ldflags "-s -w" -o MDownloader
```

## 📖 Config Example

<details><summary>config.yaml</summary>

This is an example of the configuration file. The program will create a new configuration file if it does not exist.

```yaml
prefix:
  cover: https://fourhoi.com/
  video: https://surrit.com/

suffix:
  playlist: /playlist.m3u8

# http request settings
requests:
  timeout: 10
  retry: 5
  delay: 2
  proxy: 
  headers:
    User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0

# if you don't know how to use regex, please don't change it
regex:
  href:
    movieCollection: <a class="text-secondary group-hover:text-primary" href="([^"]+)" alt="
    publicPlaylist: <a href="([^"]+)" alt="
    nextPage: <a href="([^"]+)" rel="next"
  
  data:
    uuid: m3u8\|([a-f0-9\|]+)\|com\|surrit\|https\|video
    title: <title>([^"]+)</title>
    resolution: RESOLUTION=(\d+)x(\d+)
    videoSlice: video(\d+).jpeg

ffmpeg: ""    # input ffmpeg binary path if want to use ffmpeg
quality: 0    # 1080, 720, 480, 360, 240, 144. 0 for best quality

savePath: downloads
```

</details>

## 👀 About FFmpeg

1. If you want miyuki to use ffmpeg to process the video, use the -ffmpeg option.
2. Please check whether the ffmpeg command is valid before using the -ffmpeg option. (e.g. ```ffmpeg -version```)
3. To install FFmpeg, please refer to https://ffmpeg.org/

## 📄 Disclaimer

This project is licensed under the [MIT License](LICENSE). The following additional disclaimers and notices apply:

### 1. Legal Compliance
- This software is provided solely for **communication, research, learning, and personal use**.  
- Users are responsible for ensuring that their use of this software complies with all applicable laws and regulations in their jurisdiction.  
- The software must not be used for any unlawful, unethical, or unauthorized purposes, including but not limited to violating third-party rights or legal restrictions.

### 2. No Warranty
As stated in the MIT License:  
> "THE SOFTWARE IS PROVIDED 'AS IS', WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT."

### 3. Limitation of Liability
- The author(s) shall not be held liable for any claims, damages, or other liabilities arising from or in connection with the use or performance of this software.  
- Users bear all risks and responsibilities for the use of this software, including but not limited to data loss, system damage, or legal consequences.

### 4. Third-Party Dependencies
- This project may include or depend on third-party libraries or tools. Users are responsible for reviewing and complying with the licenses and terms of these dependencies.

### 5. Security and Privacy
- This software may interact with user systems, networks, or data. Users should implement appropriate security measures to protect sensitive information and infrastructure.  
- The authors are not responsible for any security vulnerabilities or data breaches resulting from the use of this software.

## 🛠️ TODO

<details><summary>Commands</summary>

- [x] support urls command - option to specify the video URLs to download.
- [ ] support auth command - option to specify the username and password to download the videos collected by the account.
- [ ] support plist command - option to specify the public playlist URL to download all videos in the list.
- [ ] support search command - option to specify the search keyword to download the search results.
- [ ] support file command - option to specify the file path to download the URLs in the file.

</details>

<details><summary>Options</summary>

- [ ] limit option - option using in `plist` command to specify the number of videos to download.
- [ ] cover option - option to save the cover image of the video. (default: false)
- [ ] ffcover option - option to set the cover as the video preview while using ffmpeg. (default: false)

</details>

<details><summary>Others</summary>

- [ ] support displaying the download progress.

</details>

## 📈 Star History

[![Star History Chart](https://api.star-history.com/svg?repos=R4v3nl0/MDownloader&type=Date)](https://star-history.com/#R4v3nl0/MDownloader&Date)