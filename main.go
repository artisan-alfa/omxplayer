package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"time"
)

func main() {
	v := true
	if _, err := os.Stat("concat.mp4"); err == nil {
		fmt.Printf("File exists\n")
		// cmd := exec.Command("omxplayer", "-o", "hdmi", "test.mp4")
		// err := cmd.Run()
		// if err != nil {
		// 	log.Fatal(err)
		// }
		v = true
	} else if checkExt(".mp4") > 0 {
		fmt.Println("files visible")
		// cmd := exec.Command("omxplayer", "-o", "hdmi", "test.mp4")
		// err := cmd.Run()
		// if err != nil {
		// 	log.Fatal(err)
		// }
		v = true
	} else {
		fmt.Println("no file")
		v = false
	}
	fmt.Println(v)
	for checkExt(".mp4") > 0 {
		cmd := exec.Command("omxplayer", "-o", "hdmi", "test.mp4")
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	}

}

func checkExt(ext string) int {

	// pathS := "/home/watcharaphan/omxplayer/media"
	pathS, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	currentTime := time.Now()
	TODAY := currentTime.Format("2006-01-02")
	NOW := currentTime.Format("15:04")
	NOW_DIR := pathS + "/" + TODAY + "/" + NOW
	var files []string
	filepath.Walk(pathS, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			r, err := regexp.MatchString(ext, f.Name())
			if err == nil && r {
				files = append(files, f.Name())
			}
		}
		return nil
	})
	fmt.Println(TODAY, NOW, NOW_DIR)
	return len(files)
}
