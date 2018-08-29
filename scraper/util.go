package scraper

import (
	"github.com/PuerkitoBio/goquery"
)

const (
	HOST         = "https://jp.finalfantasyxiv.com"
	BASE_URL     = "https://jp.finalfantasyxiv.com/lodestone/playguide/db/quest/"
	CONTENT_TAGS = "div#main > div#eorzea_db > div.clearfix > div.db_cnts"
)

func FetchDocument(url string, findTag string) (*goquery.Selection, error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}
	contents := doc.Find(findTag)
	numOfNodes := len(contents.Nodes)

	if numOfNodes < 1 {
		return nil, NoSuchContentsError{
			url,
			findTag,
		}
	}

	return contents, nil
}
