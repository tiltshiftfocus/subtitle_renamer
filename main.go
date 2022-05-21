package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	dirFlag := flag.String("dir", "", "Source folder for conversion")
	flag.Parse()

	if *dirFlag == "" {
		fmt.Println("No source specified, exiting")
		return
	}

	list, _ := os.ReadDir(*dirFlag)

	var subtitles = make(map[int]string)
	var videofiles = make(map[int]string)

	for _, l := range list {
		name := l.Name()
		if strings.Contains(name, ".mkv") || strings.Contains(name, ".mp4") || strings.Contains(name, ".avi") {
			episodeFinder, _ := regexp.Compile(`Ep(\d{0,2})|E(\d{0,2})`)
			episode := episodeFinder.Find([]byte(name))
			numberFinder, _ := regexp.Compile(`\d+`)
			number := numberFinder.Find([]byte(episode))

			converted, _ := strconv.Atoi(string(number))
			videofiles[converted] = name
		}
		if strings.Contains(name, ".srt") || strings.Contains(name, ".ass") {
			episodeFinder, _ := regexp.Compile(`Ep(\d{0,2})|E(\d{0,2})`)
			episode := episodeFinder.Find([]byte(name))
			numberFinder, _ := regexp.Compile(`\d+`)
			number := numberFinder.Find([]byte(episode))

			converted, _ := strconv.Atoi(string(number))
			subtitles[converted] = name
		}
	}

	if len(videofiles) == len(subtitles) {
		for k, v := range videofiles {
			properPathDir := strings.ReplaceAll(*dirFlag, "\\", "/")
			originalSubName := subtitles[k]
			originalSubExt := originalSubName[len(originalSubName)-4:]

			videoFileName := v[:len(v)-4]

			old := path.Join(properPathDir, originalSubName)
			new := fmt.Sprintf("%s/%s%s", properPathDir, videoFileName, originalSubExt)

			err := os.Rename(old, new)
			if err != nil {
				log.Printf("Error renaming: %s\r\n", err)
			}
		}
	}

}
