package scraper

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

type QuestPage struct {
	Url string
	Doc *goquery.Selection
}

func (p *QuestPage) FetchQuestPage() error {
	url := p.Url
	doc, err := FetchDocument(url, CONTENT_TAGS)

	if err != nil {
		return err
	}
	p.Doc = doc

	return nil
}

func (p *QuestPage) ParseQuestList() ([]Quest, error) {
	questTable := p.Doc.Find("div.db__l_main > div.db-table__wrapper > table.db-table > tbody > tr")
	questList := []Quest{}

	questTable.Each(func(index int, s *goquery.Selection) {
		aTag := s.Find("td.db-table__body--light > a")
		href, isExist := aTag.Attr("href")
		if !isExist {
			fmt.Println("skipped, not have href attribute")
		} else {
			name := aTag.Text()
			questList = append(questList, Quest{
				href,
				name,
				nil,
			})
		}
	})

	if len(questList) < 1 {
		return nil, errors.New("quest list is not found.")
	}

	return questList, nil
}

func (p *QuestPage) HasNextPage() (bool, error) {
	pager := p.Doc.Find("div.db-filter__row > div.pager > div.pagination > ul > li")
	if len(pager.Nodes) < 1 {
		return false, errors.New("unexpected html source")
	}

	var isExist = true
	pager.Each(func(index int, s *goquery.Selection) {
		aTag := s.Find("a")
		_, isExist = aTag.Attr("href")
	})

	return isExist, nil
}
