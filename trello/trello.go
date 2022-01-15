package trello

import (
	"github.com/sclevine/agouti"
	"log"
	"os"
	"time"
)

func Login(page *agouti.Page) {
	// 画面遷移に時間がかかるためsleep入れてる。
	page.Navigate("https://trello.com/login")
	page.FindByID("googleButton").Click()
	page.FindByName("identifier").Fill(os.Getenv("EMAIL"))
	page.FindByID("identifierNext").Click()
	time.Sleep(3 * time.Second)
	page.FindByName("password").Fill(os.Getenv("PASSWORD"))
	page.FindByID("passwordNext").Click()
	time.Sleep(7 * time.Second)
}

func DrainTasks(page *agouti.Page) [][]string {
	//現在開発中ボード
	page.Navigate("https://trello.com/b/joePfsTs/%E7%8F%BE%E5%9C%A8%E9%96%8B%E7%99%BA%E4%B8%AD")
	time.Sleep(5 * time.Second)

	// 要企画、企画、開発・
	rows := page.Find("div.board-canvas > div#board").All("div.js-list.list-wrapper").At(1)

	tasksElement := rows.Find("div.list.js-list-content > div.list-cards").All("a.list-card")

	tasksCount, err := tasksElement.Count()
	if err != nil {
		log.Fatal(err)
	}
	var tasks [][]string
	for i := 0; i < tasksCount ; i++ {
		//タスクを開く
		tasksElement.At(i).Click()
		time.Sleep(2 * time.Second)
		//タイトル、説明を取得
		title, description := getTitleAndDescription(page)
		if title == "" {
			log.Fatal("Error: タイトルがない")
		}
		task := []string{title, description}
		tasks = append(tasks, task)
		//タスク閉じる
		err = page.Find("a.icon-md.icon-close.dialog-close-button.js-close-window").Click()
		if err != nil {
			log.Fatal(err)
		}
	}
	return tasks
}

func getTitleAndDescription(page *agouti.Page) (string, string) {
	title, err := page.Find("textarea.mod-card-back-title.js-card-detail-title-input").Attribute("value")
	if err != nil {
		log.Fatal(err)
	}
	page.Find("div.editable > a").Click()
	description, err := page.Find("textarea.description.card-description").Attribute("value")
	if err != nil {
		log.Fatal(err)
	}
	return title, description
}
