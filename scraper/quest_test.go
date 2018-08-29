package scraper

import (
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"reflect"
	"strings"
	"testing"
)

func TestParseQuestName(t *testing.T) {
	testHtml := "../resources/scraper/quest.html"
	file, _ := ioutil.ReadFile(testHtml)
	stringReader := strings.NewReader(string(file))
	doc, err := goquery.NewDocumentFromReader(stringReader)
	if err != nil {
		t.Fatal("tailed test, cannot read " + testHtml)
	}
	selection := doc.Find("div#main > div#eorzea_db > div.clearfix > div.db_cnts > div.db__l_main")
	questName, err := parseQuestName(selection)
	if err != nil {
		t.Fatalf("failed test, %v", err)
	}
	expected := "冒険者への手引き"
	if questName != expected {
		t.Fatalf("failed test, incorrect result, actual = %v, expected = %v", questName, expected)
	}
}

func TestParseQuestType(t *testing.T) {
	testHtml := "../resources/scraper/quest.html"
	file, _ := ioutil.ReadFile(testHtml)
	stringReader := strings.NewReader(string(file))
	doc, err := goquery.NewDocumentFromReader(stringReader)
	if err != nil {
		t.Fatal("tailed test, cannot read " + testHtml)
	}
	selection := doc.Find("div#main > div#eorzea_db > div.clearfix > div.db_cnts > div.db__l_main")
	questType, err := parseQuestType(selection)
	if err != nil {
		t.Fatalf("failed test, %v", err)
	}
	expected := "新生エオルゼア"
	if questType != expected {
		t.Fatalf("failed test, incorrect result, actual = %v, expected = %v", questType, expected)
	}
}

func TestParseQuestClient(t *testing.T) {
	testHtml := "../resources/scraper/quest.html"
	file, _ := ioutil.ReadFile(testHtml)
	stringReader := strings.NewReader(string(file))
	doc, err := goquery.NewDocumentFromReader(stringReader)
	if err != nil {
		t.Fatal("tailed test, cannot read " + testHtml)
	}
	selection := doc.Find("div#main > div#eorzea_db > div.clearfix > div.db_cnts > div.db__l_main")
	questClient, err := parseQuestClient(selection)
	if err != nil {
		t.Fatalf("failed test, %v", err)
	}
	expected := "ミューヌ"
	if questClient != expected {
		t.Fatalf("failed test, incorrect result, actual = %v, expected = %v", questClient, expected)
	}
}

func TestParseQuestPlace(t *testing.T) {
	testHtml := "../resources/scraper/quest.html"
	file, _ := ioutil.ReadFile(testHtml)
	stringReader := strings.NewReader(string(file))
	doc, err := goquery.NewDocumentFromReader(stringReader)
	if err != nil {
		t.Fatal("tailed test, cannot read " + testHtml)
	}
	selection := doc.Find("div#main > div#eorzea_db > div.clearfix > div.db_cnts > div.db__l_main")
	questClient, err := parseQuestPlace(selection)
	if err != nil {
		t.Fatalf("failed test, %v", err)
	}
	expected := "グリダニア：新市街"
	if questClient != expected {
		t.Fatalf("failed test, incorrect result, actual = %v, expected = %v", questClient, expected)
	}
}

func TestParseQuestXY(t *testing.T) {
	testHtml := "../resources/scraper/quest.html"
	file, _ := ioutil.ReadFile(testHtml)
	stringReader := strings.NewReader(string(file))
	doc, err := goquery.NewDocumentFromReader(stringReader)
	if err != nil {
		t.Fatal("tailed test, cannot read " + testHtml)
	}
	selection := doc.Find("div#main > div#eorzea_db > div.clearfix > div.db_cnts > div.db__l_main")
	x, y, err := parseQuestXY(selection)
	if err != nil {
		t.Fatalf("failed test, %v", err)
	}
	expectedX := 11.7
	expectedY := 13.5
	if x != expectedX {
		t.Fatalf("failed test, incorrect result, actual = %v, expected = %v", x, expectedX)
	}

	if y != expectedY {
		t.Fatalf("failed test, incorrect result, actual = %v, expected = %v", y, expectedY)
	}

}

func TestParseQuestConditions(t *testing.T) {
	testHtml := "../resources/scraper/quest.html"
	file, _ := ioutil.ReadFile(testHtml)
	stringReader := strings.NewReader(string(file))
	doc, err := goquery.NewDocumentFromReader(stringReader)
	if err != nil {
		t.Fatal("tailed test, cannot read " + testHtml)
	}
	selection := doc.Find("div#main > div#eorzea_db > div.clearfix > div.db_cnts > div.db__l_main")
	condition, err := parseQuestConditions(selection)
	if err != nil {
		t.Fatalf("failed test, %v", err)
	}

	expected := map[string]string{
		"initJob":      "槍術士",
		"class":        "いずれかのクラス・ジョブ Lv 1～",
		"grandCompany": "指定なし",
		"content":      "指定なし",
	}
	if !reflect.DeepEqual(condition, expected) {
		t.Fatalf("failed test, incorrect result, actual = %v, expected = %v", condition, expected)
	}
}

func TestParseQuestReward(t *testing.T) {
	testHtml := "../resources/scraper/quest.html"
	file, _ := ioutil.ReadFile(testHtml)
	stringReader := strings.NewReader(string(file))
	doc, err := goquery.NewDocumentFromReader(stringReader)
	if err != nil {
		t.Fatal("tailed test, cannot read " + testHtml)
	}
	selection := doc.Find("div#main > div#eorzea_db > div.clearfix > div.db_cnts > div.db__l_main")
	rewards, err := parseQuestReward(selection)
	if err != nil {
		t.Fatalf("failed test, %v", err)
	}

	expected := map[string]int{
		"exp": 100,
		"gil": 107,
	}
	if !reflect.DeepEqual(rewards, expected) {
		t.Fatalf("failed test, incorrect result, actual = %v, expected = %v", rewards, expected)
	}
}

func TestParsePremisQuests(t *testing.T) {
	testHtml := "../resources/scraper/quest.html"
	file, _ := ioutil.ReadFile(testHtml)
	stringReader := strings.NewReader(string(file))
	doc, err := goquery.NewDocumentFromReader(stringReader)
	if err != nil {
		t.Fatal("tailed test, cannot read " + testHtml)
	}
	selection := doc.Find("div#main > div#eorzea_db > div.clearfix > div.db_cnts > div.db__l_main")
	premisQuests, _ := parsePremisQuests(selection)
	expected := []Quest{
		{
			"/lodestone/playguide/db/quest/298088846dc/",
			"森の都グリダニアへ",
			nil,
		},
		{
			"/lodestone/playguide/db/quest/298088846dc/",
			"森の都グリダニアへ2",
			nil,
		},
	}

	if !reflect.DeepEqual(premisQuests, expected) {
		t.Fatalf("failed test, incorrect result, actual = %v, expected = %v", premisQuests, expected)
	}

}

func TestParseUnlockQuests(t *testing.T) {
	testHtml := "../resources/scraper/quest.html"
	file, _ := ioutil.ReadFile(testHtml)
	stringReader := strings.NewReader(string(file))
	doc, err := goquery.NewDocumentFromReader(stringReader)
	if err != nil {
		t.Fatal("tailed test, cannot read " + testHtml)
	}
	selection := doc.Find("div#main > div#eorzea_db > div.clearfix > div.db_cnts > div.db__l_main")
	unlockQuests, _ := parseUnlockQuests(selection)
	expected := []Quest{
		{
			"/lodestone/playguide/db/quest/088a43daa15/",
			"バノック練兵所へ",
			nil,
		},
		{
			"/lodestone/playguide/db/quest/2ee087b785d/",
			"ツリースピーク厩舎の金具",
			nil,
		},
		{
			"/lodestone/playguide/db/quest/8d6c6da2282/",
			"小さな預かり物",
			nil,
		},
	}

	if !reflect.DeepEqual(unlockQuests, expected) {
		t.Fatalf("failed test, incorrect result, actual = %v, expected = %v", unlockQuests, expected)
	}

}
