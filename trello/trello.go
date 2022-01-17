package trello

import (
	"fmt"
	"github.com/sclevine/agouti"
	"log"
	"os"
	"strconv"
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

	// 要企画。他の場所を取りたければAtをかえる
	rows := page.Find("div.board-canvas > div#board").All("div.js-list.list-wrapper").At(1)

	tasksElement := rows.Find("div.list.js-list-content > div.list-cards").All("a.list-card")

	tasksCount, err := tasksElement.Count()
	fmt.Println("タスク数：" + strconv.Itoa(tasksCount))
	if err != nil {
		log.Fatal(err)
	}
	var tasks [][]string
	for i := 0; i < tasksCount ; i++ {
		//タスクを開く
		tasksElement.At(i).Click()
		time.Sleep(1 * time.Second)

		//タイトル、説明を取得。エラーが出てもリトライで治る場合があるのでリトライしてる。
		title, description, err := getTitleAndDescription(page)
		if err != nil {
			title, description = retryGetTitleAndDescription(page, i, rows)
		}
		fmt.Println(title + ": " + strconv.Itoa(i))

		task := []string{title, description}
		tasks = append(tasks, task)
		//タスク閉じる
		err = page.Find("a.icon-md.icon-close.dialog-close-button.js-close-window").Click()
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(500 * time.Millisecond)
	}
	return tasks
}

func getTitleAndDescription(page *agouti.Page) (string, string, error) {
	title, err := page.Find("textarea.mod-card-back-title.js-card-detail-title-input").Attribute("value")
	if err != nil {
		fmt.Println("titleが取得できませんでした")
	}
	err = page.Find("div.editable > a").Click()
	if err != nil {
		fmt.Println("descのクリック失敗")
	}
	description, err := page.Find("textarea.description.card-description").Attribute("value")
	if err != nil {
		fmt.Println("desc入力失敗")
	}
	return title, description, err
}

func retryGetTitleAndDescription(page *agouti.Page, i int, rows *agouti.Selection) (string, string) {
	tasksElement := rows.Find("div.list.js-list-content > div.list-cards").All("a.list-card")
	tasksElement.At(i).Click()
	title, description, err := getTitleAndDescription(page)
	if err != nil {
		log.Fatal("title取得失敗")
	}
	return title, description
}
