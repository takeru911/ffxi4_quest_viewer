package main

import (
	"fmt"
	"github.com/takeru911/ffxiv_quest_viewer/scraper"
)

func main() {

	quest := scraper.Quest{
		"/lodestone/playguide/db/quest/dabd7659695/",
		"ただ盟友のため",
		nil,
	}
	quest.FetchQuestDetail()
	questDetail, err := quest.ParseQuestDetail()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(questDetail)
}
