package scraper

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"strconv"
	"strings"
)

type Quest struct {
	Url       string
	Name      string
	Selection *goquery.Selection
}

//これはいずれ別のとこに行くはず
//Itemのスクレイピングは現状行わないのでいったんここで
type Item struct {
	Url  string
	Name string
}

type QuestDetail struct {
	Url                   string
	Name                  string
	QuestType             string
	Client                string
	Place                 string
	X                     float64
	Y                     float64
	InitJobCondition      string
	ClassCondition        string
	GrandCompanyCondition string
	ContentCondition      string
	Exp                   int
	Gil                   int
	PremisQuests          []Quest
	UnlockQuests          []Quest
	SelectRewards         []Item
}

func (q *Quest) FetchQuestDetail() error {
	url := HOST + q.Url
	doc, err := FetchDocument(url, CONTENT_TAGS)

	if err != nil {
		return err
	}
	q.Selection = doc

	return nil
}

func (q *Quest) ParseQuestDetail() (QuestDetail, error) {
	if q.Selection == nil {
		return QuestDetail{}, errors.New("quest detail doc is nil, please exec Quest.FetchQuestDetail")
	}

	questContent := q.Selection.Find("div.db__l_main")

	questName, err := parseQuestName(questContent)
	if err != nil {
		return QuestDetail{}, err
	}
	fmt.Println("parsed QuestName")
	questType, err := parseQuestType(questContent)
	if err != nil {
		return QuestDetail{}, err
	}
	fmt.Println("parsed QuestType")
	questClient, err := parseQuestClient(questContent)
	if err != nil {
		return QuestDetail{}, err
	}
	fmt.Println("parsed QuestClient")
	x, y, err := parseQuestXY(questContent)
	if err != nil {
		return QuestDetail{}, err
	}
	fmt.Println("parsed QuestXY")
	questPlace, err := parseQuestPlace(questContent)
	if err != nil {
		return QuestDetail{}, err
	}
	fmt.Println("parsed QuestPlace")
	questCondition, err := parseQuestConditions(questContent)
	if err != nil {
		return QuestDetail{}, err
	}
	fmt.Println("parsed QuestCondition")
	questReward, err := parseQuestReward(questContent)
	if err != nil {
		return QuestDetail{}, err
	}
	fmt.Println("parsed QuestReward")
	premisQuests, err := parsePremisQuests(questContent)
	if err != nil {
		return QuestDetail{}, err
	}
	unlockQuests, err := parseUnlockQuests(questContent)
	if err != nil {
		return QuestDetail{}, err
	}
	selectReward, err := parseSelectReward(questContent)
	if err != nil {
		return QuestDetail{}, err
	}
	return QuestDetail{
		q.Url,
		questName,
		questType,
		questClient,
		questPlace,
		x,
		y,
		questCondition["initJob"],
		questCondition["class"],
		questCondition["grandCompany"],
		questCondition["content"],
		questReward["exp"],
		questReward["gil"],
		premisQuests,
		unlockQuests,
		selectReward,
	}, nil

}

func parseQuestName(selection *goquery.Selection) (string, error) {
	questNameTag := selection.Find("div.db-view__icon > h2.db-view__detail__lname_name")
	if len(questNameTag.Nodes) < 1 {
		return "", errors.New("unexpected, html")
	}
	return strings.TrimSpace(questNameTag.Text()), nil
}

func parseQuestType(selection *goquery.Selection) (string, error) {
	questTypeTag := selection.Find("div.db-view__icon > p > span.db-view__detail__content_type")
	if len(questTypeTag.Nodes) < 1 {
		return "", errors.New("unexpected, html")
	}
	return strings.TrimSpace(questTypeTag.Text()), nil
}

func parseQuestClient(selection *goquery.Selection) (string, error) {
	questClientTag := selection.Find("div.db-table__wrapper > table.db-table > tbody > tr > td.db-table__body--light > a > strong")
	if len(questClientTag.Nodes) < 1 {
		return "", errors.New("unexpected, html")
	}
	return strings.TrimSpace(questClientTag.Text()), nil
}

func parseQuestPlace(selection *goquery.Selection) (string, error) {
	//Removeには副作用があるので、これでとるのをやめたほうがいい
	//いまはTextで取れるのが子要素のもとれちゃうのでRemoveしてる。。。
	//呼び出し順序で回避してるけど、根本回避したい・・・
	questPlaceTag := selection.Find("div.db-table__wrapper > table.db-table > tbody > tr > td.db-table__body--light > ul.db-view__npc__location__list > li").Children().Remove().End()

	if len(questPlaceTag.Nodes) < 1 {
		return "", errors.New("unexpected, html")
	}
	return strings.TrimSpace(questPlaceTag.Text()), nil
}

func parseQuestXY(selection *goquery.Selection) (float64, float64, error) {
	questXYTag := selection.Find("div.db-table__wrapper > table.db-table > tbody > tr > td.db-table__body--light > ul.db-view__npc__location__list > li > ul > li")
	if len(questXYTag.Nodes) < 1 {
		return 0, 0, errors.New("unexpected, html")
	}

	tagText := strings.TrimSpace(questXYTag.Text())
	r := regexp.MustCompile(`X：(\d{1,2}\.\d{1}) Y：(\d{1,2}\.\d{1})`)
	xy := r.FindAllStringSubmatch(tagText, -1)

	if len(xy[0]) != 3 {
		errorMessage := fmt.Sprintf("unexpected coordinate text, %v", xy)
		return 0, 0, errors.New(errorMessage)
	}

	strX := xy[0][1]
	x, err := strconv.ParseFloat(strX, 64)

	if err != nil {
		errorMessage := fmt.Sprintf("convert failed, %v -> float64", strX)
		return 0, 0, errors.New(errorMessage)
	}

	strY := xy[0][2]
	y, err := strconv.ParseFloat(strY, 64)

	if err != nil {
		errorMessage := fmt.Sprintf("convert failed, %v -> float64", strY)
		return 0, 0, errors.New(errorMessage)
	}

	return x, y, nil
}

func parseQuestConditions(selection *goquery.Selection) (map[string]string, error) {
	questConditionTag := selection.Find("div.db-view__data > div.db-view__data__detail_list__wrapper > dl > dd")
	if len(questConditionTag.Nodes) < 1 {
		return nil, errors.New("unexpected, html")
	}

	conditions := map[string]string{}

	questConditionTag.Each(func(index int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())
		switch index {
		case 0:
			conditions["initJob"] = text
		case 1:
			conditions["class"] = text
		case 2:
			conditions["grandCompany"] = text
		case 3:
			conditions["content"] = text
		default:
		}
	})

	return conditions, nil
}

func parseQuestReward(s *goquery.Selection) (map[string]int, error) {
	questRewardTag := s.Find("div.db-view__data > ul.db-view__quest__reward__list > li.db-view__quest__reward__box > div.db-view__quest__reward__value")
	if len(questRewardTag.Nodes) < 1 {
		return nil, errors.New("unexpected, html")
	}
	rewards := map[string]int{}
	var err error
	questRewardTag.Each(func(index int, s *goquery.Selection) {
		text := s.Text()
		value, err := strconv.ParseInt(text, 10, 64)
		if err != nil {
			errorMessage := fmt.Sprintf("convert failed, %v -> int", text)
			err = errors.New(errorMessage)
		}
		switch index {
		case 0:
			if err == nil {
				rewards["exp"] = int(value)
			}
		case 1:
			if err == nil {
				rewards["gil"] = int(value)
			}
		}
	})

	if err != nil {
		return nil, err
	}

	return rewards, nil
}

func parsePremisQuests(s *goquery.Selection) ([]Quest, error) {
	premisQuestsParentTag := s.FilterFunction(func(index int, s *goquery.Selection) bool {
		return s.Find("div.clearfix > h3.db-category__title__diamond").Text() == "前提クエスト"
	})
	//前提クエストがないクエストもある
	if len(premisQuestsParentTag.Nodes) == 0 {
		return []Quest{}, nil
	}
	premisQuestsTag := premisQuestsParentTag.Find("div.db-table__wrapper > table.db-table > tbody > tr > td.db-table__body--light > a")
	premisQuests := make([]Quest, len(premisQuestsTag.Nodes), len(premisQuestsTag.Nodes))
	premisQuestsTag.Each(func(index int, s *goquery.Selection) {
		questName := strings.TrimSpace(s.Text())
		tmp, _ := s.Attr("href")
		questUrl := strings.TrimSpace(tmp)
		premisQuests[index] = Quest{
			questUrl,
			questName,
			nil,
		}
	})
	return premisQuests, nil
}

func parseUnlockQuests(s *goquery.Selection) ([]Quest, error) {
	unlockQuestsParentTag := s.FilterFunction(func(index int, s *goquery.Selection) bool {
		return s.Find("div.clearfix > h3.db-category__title__diamond").Text() == "発生クエスト"
	})
	//発生クエストがないクエストもある
	if len(unlockQuestsParentTag.Nodes) == 0 {
		return []Quest{}, nil
	}
	unlockQuestsTag := unlockQuestsParentTag.Find("div.db-table__wrapper > table.db-table > tbody > tr > td.db-table__body--light > a")
	unlockQuests := make([]Quest, len(unlockQuestsTag.Nodes), len(unlockQuestsTag.Nodes))
	unlockQuestsTag.Each(func(index int, s *goquery.Selection) {
		questName := strings.TrimSpace(s.Text())
		tmp, _ := s.Attr("href")
		questUrl := strings.TrimSpace(tmp)
		unlockQuests[index] = Quest{
			questUrl,
			questName,
			nil,
		}
	})
	return unlockQuests, nil
}

func parseSelectReward(s *goquery.Selection) ([]Item, error) {
	selectRewardParentTag := s.FilterFunction(func(index int, s *goquery.Selection) bool {
		return s.Find("div.db-view__data > div.db-view__data__inner--select_reward > div.db-view__data__inner__wrapper > h4.db-view__data__title--quest_reward_list").Text() == "選択報酬"
	})

	if len(selectRewardParentTag.Nodes) == 0 {
		return []Item{}, nil
	}
	selectRewardTag := selectRewardParentTag.Find("div.db-view__data > div.db-view__data__inner--select_reward > div.db-view__data__inner__wrapper > ul.db-view__data__item_list")
	selectRewardItems := make([]Item, len(selectRewardTag.Nodes), len(selectRewardTag.Nodes))

	selectRewardTag.Each(func(index int, s *goquery.Selection) {
		aTag := s.Find("li > div.db-view__data__reward__item__name > div.db-view__data__reward__item__name__wrapper > a")
		itemName := strings.TrimSpace(aTag.Text())
		url, _ := aTag.Attr("href")
		selectRewardItems[index] = Item{
			url,
			itemName,
		}
	})

	return selectRewardItems, nil
}

func (q Quest) String() string {
	return fmt.Sprintf(`
{
	QuestUrl: %v,
	QuestName: %v
}
`,
		q.Url,
		q.Name,
	)
}

func (i Item) String() string {
	return fmt.Sprintf(`
{
	ItemUrl: %v,
	ItemName: %v
}
`,
		i.Url,
		i.Name,
	)
}

func (q QuestDetail) String() string {
	return fmt.Sprintf(`
{
	QuestUrl : %v,
	QuestName: %v,
	QuestType: %v,
	QuestClient: %v,
	QuestPlace: %v,
	x: %v,
	y: %v,
	QuestCondition(InitJob): %v,
	QuestCondition(class): %v,
	QuestCondition(grandCompany): %v,
	QuestCondition(content): %v,
	Reward(Exp): %v,
	Reward(Gil): %v,
	PromisQuests: %v,
	UnlockQuests: %v,
	SelectReward: %v
}
	`,
		q.Url,
		q.Name,
		q.QuestType,
		q.Client,
		q.Place,
		q.X,
		q.Y,
		q.InitJobCondition,
		q.ClassCondition,
		q.GrandCompanyCondition,
		q.ContentCondition,
		q.Exp,
		q.Gil,
		q.PremisQuests,
		q.UnlockQuests,
		q.SelectReward,
	)
}
