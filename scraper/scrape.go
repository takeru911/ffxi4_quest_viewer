package scraper

import (
	"fmt"
	"time"
)

func ScrapeQuestList() ([]Quest, error) {

	page := 1
	hasNext := true
	questList := []Quest{}
	for hasNext {
		tmp, hasNext, err := scrapeQuestList(page)
		if err != nil {
			return nil, err
		}
		questList = append(questList, tmp...)
		if page > 3 {
			break
		}
		fmt.Printf("page: %v completed.\n", page)
		if hasNext {
			page++
		}
		time.Sleep(2 * time.Second)
	}

	return questList, nil
}

func scrapeQuestList(page int) ([]Quest, bool, error) {
	questPage := QuestPage{
		fmt.Sprintf("https://jp.finalfantasyxiv.com/lodestone/playguide/db/quest/?page=%v", page),
		nil,
	}
	fmt.Println(questPage.Url)
	err := questPage.FetchQuestPage()

	if err != nil {
		return nil, false, err
	}

	questList, err := questPage.ParseQuestList()
	if err != nil {
		return nil, false, err
	}
	hasNext, err := questPage.HasNextPage()
	if err != nil {
		return nil, false, err
	}

	return questList, hasNext, nil
}
