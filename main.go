package main

import (
	"embed"
	"fmt"
	"os"

	"github.com/R4v3nl0/MDownloader/cmd"
	"github.com/R4v3nl0/MDownloader/config"
	"github.com/urfave/cli/v2"
	"golang.org/x/net/context"
)

//go:embed config.yaml
var configFile embed.FS

func main() {
	app := &cli.App{
		HideHelpCommand: true,
		Name:            "MDownloader",
		Version:         "0.0.1",
		Description:     "A tool for downloading videos from the \"MissAV\" website that can help you build your own private video library.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Value:   "config.yaml",
				Usage:   "set the common config file path",
			},
		},
		Commands: []*cli.Command{
			cmd.NewUrlsCommand(),
		},
		Before: func(c *cli.Context) error {
			configPath := c.String("config")
			if _, err := os.Stat(configPath); os.IsNotExist(err) {
				configData, err := configFile.ReadFile("config.yaml")
				if err != nil {
					return err
				}

				err = os.WriteFile(configPath, configData, 0644)
				if err != nil {
					return err
				}

				return fmt.Errorf("config file not found, created a new one, please restart")
			}

			cfg, err := config.LoadConfig(configPath)
			if err != nil {
				return err
			}

			c.Context = context.WithValue(c.Context, "cfg", cfg)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
