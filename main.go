package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/reujab/wallpaper"
)

func main() {
	u := "https://source.unsplash.com/featured/1600x900"

	if len(os.Args) > 1 {
		search := strings.Join(os.Args[1:], ",")
		search = url.QueryEscape(search)
		u = fmt.Sprintf("%s?%s", u, search)
	}

	err := wallpaper.SetFromURL(u)
	if err != nil {
		log.Println("Error has occurred", err.Error())
	}
}
