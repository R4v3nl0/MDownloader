package cmd

import (
	"fmt"
	"strings"

	"github.com/R4v3nl0/MDownloader/config"
	"github.com/R4v3nl0/MDownloader/downloader"
	"github.com/urfave/cli/v2"
)

func NewUrlsCommand() *cli.Command {
	return &cli.Command{
		Name:      "urls",
		Usage:     "download videos from urls",
		UsageText: `MDownloader urls "url1,url2,url3..."`,
		Action: func(c *cli.Context) error {
			if c.NArg() == 0 {
				return cli.ShowSubcommandHelp(c)
			}

			urlsString := c.Args().First()
			urls := strings.Split(urlsString, ",")

			cfg := c.Context.Value("cfg").(*config.Config)

			for _, url := range urls {
				err := downloader.Download(url, cfg)
				if err != nil {
					fmt.Printf("Download %s failed: %s, skip.\n", url, err)
					continue
				}
			}

			return nil
		},
	}
}
