package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/reujab/wallpaper"
)

func main() {
	url := "https://source.unsplash.com/featured/1600x900"

	if len(os.Args) > 1 {
		search := strings.Join(os.Args[1:], ",")
		url = fmt.Sprintf("%s?%s", url, search)
	}

	err := wallpaper.SetFromURL(url)
	if err != nil {
		log.Println("Error has occurred", err.Error())
	}
}