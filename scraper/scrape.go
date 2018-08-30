package scraper

import (
	"fmt"
	"time"
)

func ScrapeQuests() ([]QuestDetail, error) {
	questList, err := scrapeQuestList()
	if err != nil {
		return nil, err
	}
	questDetails := []QuestDetail{}
	for _, quest := range questList {
        fmt.Println(quest.Name)
		err := quest.FetchQuestDetail()
		if err != nil {
			return nil, err
		}
		questDetail, err := quest.ParseQuestDetail()
		if err != nil {
			return nil, err
		}
		questDetails = append(questDetails, questDetail)
		time.Sleep(3 * time.Second)
	}

	return questDetails, nil
}

func scrapeQuestList() ([]Quest, error) {

	page := 1
	hasNext := true
	questList := []Quest{}
	for hasNext {
		tmp, hasNext, err := scrapeQuestListPage(page)
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
		time.Sleep(3 * time.Second)
		if page > 1 {
			break
		}
	}

	return questList, nil
}

func scrapeQuestListPage(page int) ([]Quest, bool, error) {
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
