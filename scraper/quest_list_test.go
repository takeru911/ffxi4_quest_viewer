package scraper

import (
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"reflect"
	"strings"
	"testing"
)

func TestParseQuestList(t *testing.T) {
	testHtml := "../resources/scraper/quest_list.html"
	file, _ := ioutil.ReadFile(testHtml)
	stringReader := strings.NewReader(string(file))
	doc, err := goquery.NewDocumentFromReader(stringReader)
	selection := doc.Find("div#main > div#eorzea_db > div.clearfix > div.db_cnts")

	if err != nil {
		t.Fatal("failed test, cannot read " + testHtml)
	}
	questPage := QuestPage{
		testHtml,
		selection,
	}
	questList, err := questPage.ParseQuestList()
	expected := []Quest{
		Quest{
			"/lodestone/playguide/db/quest/ca56a09bbfa/",
			"モラビー造船廠へ",
            nil,
		},
		Quest{
			"/lodestone/playguide/db/quest/ee334237df0/",
			"五月蝿いヤツら",
            nil,
		},
		Quest{
			"/lodestone/playguide/db/quest/d84e586bc0a/",
			"貧民街の連隊長",
            nil,
		},
	}
	if !reflect.DeepEqual(questList, expected) {
		t.Fatalf("failed test, incorrect result, actual = %v, expected = %v", questList, expected)
	}
}

func TestHasNextPage(t *testing.T) {
	testHtml := "../resources/scraper/quest_list.html"
	file, _ := ioutil.ReadFile(testHtml)
	stringReader := strings.NewReader(string(file))
	doc, err := goquery.NewDocumentFromReader(stringReader)
	selection := doc.Find("div#main > div#eorzea_db > div.clearfix > div.db_cnts")

	if err != nil {
		t.Fatal("tailed test, cannot read " + testHtml)
	}
	questPage := QuestPage{
		testHtml,
		selection,
	}
	hasNext, err := questPage.HasNextPage()

	if err != nil {
		t.Fatalf("tailed test, %v", err)
	}

	if hasNext {
		t.Fatalf("failed test, incorrect result(this page doesn't have next page), actual = %v, expected = %v", hasNext, true)
	}
}
