package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/R4v3nl0/MDownloader/config"
)

func SaveSliceFile(path string, data []byte) error {
	// Save the slice file to the specified path
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func MergeSliceToVideo(title string, sliceNum int, cfg *config.Config) error {
	if cfg.Ffmpeg != "" {
		// generate list.txt
		listFile, err := os.Create(filepath.Join(cfg.SavePath, title, "list.txt"))
		if err != nil {
			return err
		}
		defer listFile.Close()

		for i := 0; i < sliceNum; i++ {
			_, err = listFile.WriteString(fmt.Sprintf("file '%s'\n", filepath.Join(cfg.SavePath, title, fmt.Sprintf("video%d.jpeg", i))))
			if err != nil {
				return err
			}
		}

		ffmpegArg := []string{
			"-loglevel", "error",
			"-f", "concat",
			"-safe", "0",
			"-i", filepath.Join(cfg.SavePath, title, "list.txt"),
			"-c", "copy",
			filepath.Join(cfg.SavePath, title+".mp4"),
		}

		err = RunCmd(cfg.Ffmpeg, ffmpegArg)
		if err != nil {
			return err
		}
	} else {
		// Read all slice files and merge them into a video file
		videoFile, err := os.Create(filepath.Join(cfg.SavePath, title+".mp4"))
		if err != nil {
			return err
		}
		defer videoFile.Close()

		for i := 0; i < sliceNum; i++ {
			sliceFile, err := os.Open(filepath.Join(cfg.SavePath, title, fmt.Sprintf("video%d.jpeg", i)))
			if err != nil {
				return err
			}
			defer sliceFile.Close()

			_, err = sliceFile.WriteTo(videoFile)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func RunCmd(cmd string, args []string) error {
	// Run the command and return the error
	if _, err := os.Stat(cmd); os.IsNotExist(err) {
		return fmt.Errorf("command file %s not found", cmd)
	}

	command := exec.Command(cmd, args...)
	err := command.Run()
	if err != nil {
		return err
	}
	return nil
}
