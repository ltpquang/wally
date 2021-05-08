package main

import (
	"github.com/reujab/wallpaper"
)

func main() {
	err := wallpaper.SetFromURL("https://source.unsplash.com/random/1600x900")
	if err != nil {
		panic(err)
	}
}