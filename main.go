package main

import (
	"fmt"
	"github.com/takeru911/ffxiv_quest_viewer/scraper"
)

func main() {

	questList, _ := scraper.ScrapeQuestList()
    for _, quest := range questList {
        fmt.Println(quest)
    }
}
