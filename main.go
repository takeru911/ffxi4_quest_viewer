package main

import (
	"fmt"
	"github.com/takeru911/ffxiv_quest_viewer/scraper"
)

func main() {

	questList, err := scraper.ScrapeQuests()
    if err != nil {
        fmt.Println(err)
    }
    for _, quest := range questList {
        fmt.Println(quest)
    }
}
