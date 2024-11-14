package ytdl

import (
	"encoding/json"
	"errors"
	"flag"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"path"
)

type Request struct {
	URL    string
	Action func()
}

var (
	ErrNotAvaliable = errors.New("cannot download")
	Handler         http.Handler
	executable      string
	outputDir       string
)

func init() {
	var err error
	outputDir, err = os.MkdirTemp("", "ijod*")
	if err != nil {
		log.Print(err)
		return
	}
	Handler = http.FileServer(http.Dir(outputDir))

	// Attempt to clean up when interrupted
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)

		<-c
		err := os.RemoveAll(outputDir)
		if err != nil {
			log.Print(err)
		}

		os.Exit(0)
	}()

	flag.StringVar(&executable, "dl", "yt-dlp",
		"Name or path of the youtube-dl script.")
}

func avaliable() bool {
	return exec.Command(executable, "--help").Run() == nil
}

func Download(uri string) (string, error) {
	if outputDir == "" {
		return "", os.ErrNotExist
	}
	// dir := path.Join(outputDir, room)
	// err := os.Mkdir(dir, 0700)
	// if err != nil {
	// 	if !os.IsExist(err) {
	// 		return "", err
	// 	}
	// }

	if !avaliable() {
		return "", ErrNotAvaliable
	}

	// Start youtube-dl
	log.Println("Request to download", uri)
	cmd := exec.Command(executable,
		"-o", path.Join(outputDir, "%(title)s-%(id)s.%(ext)s"),
		"--format", "mp4",
		"--no-playlist", "--print-json", "--quiet", uri)
	cmd.Stderr = os.Stderr
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}
	err = cmd.Start()
	if err != nil {
		return "", err
	}

	// Prepare output parsing
	var result struct {
		Filename string `json:"_filename"`
	}
	err = json.NewDecoder(stdout).Decode(&result)
	if err != nil {
		return "", err
	}
	if err := cmd.Wait(); err != nil {
		return "", err
	}

	// Return the resulting filename
	log.Print("Finished downloading", result.Filename)
	return result.Filename, nil
}
