package main

import (
	"fmt"

	"github.com/jleight/omxplayer"
)

func main() {
omxplayer.SetUser("HOME", "/media")

player, err := omxplayer.New("/test.mp4", "--no-osd")
if err != nil {
	panic(err)
}

if player.IsReady() {
	fmt.Printf("Hello")
}
// player.WaitForReady()
// err = player.PlayPause()

// time.Sleep(5 * time.Second)
// err = player.ShowSubtitles()

// time.Sleep(5 * time.Second)
// err = player.Quit()

}