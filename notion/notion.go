package notion

import (
	"bufio"
	"fmt"
	"github.com/sclevine/agouti"
	"log"
	"os"
	"time"
)

func Login(page *agouti.Page) {
	// 開発タスクのURl
	err := page.Navigate("https://www.notion.so/e0a9b6c57872431a9678bceb7fcf8e5a?v=7b60e20ae85d4220a552c9c4afaaeed5")
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(2 * time.Second)
	page.Find("div.notion-focusable > svg.googleLogo").Click()
	fmt.Println("notionにログインしてください。ログインできたらエンターを押してください。:")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	time.Sleep(2 * time.Second)
}

func PasteTasks(page *agouti.Page, tasks [][]string) {
	for _, task := range tasks {
		addTask(page)
		//タイトルを埋める
		inputTitle := page.Find("#notion-app > div > div.notion-cursor-listener > div:nth-child(2) > div.notion-frame > div.notion-scroller.vertical.horizontal > div:nth-child(3) > div > div > div > div:nth-child(2) > div:nth-child(2) > div:nth-child(1) > div:nth-child(2) > a > div > div:nth-child(2) > div")
		inputTitle.Fill(task[0])
		//適当なところクリックして入力状態を解除
		page.Find("#notion-app > div > div.notion-cursor-listener > div:nth-child(2) > div.notion-frame > div.notion-scroller.vertical.horizontal > div:nth-child(1) > div:nth-child(2) > div:nth-child(2) > div:nth-child(1) > div.notion-selectable.notion-collection_view_page-block > div").Click()
		// modal開く
		title := page.Find("#notion-app > div > div.notion-cursor-listener > div:nth-child(2) > div.notion-frame > div.notion-scroller.vertical.horizontal > div:nth-child(3) > div > div > div > div:nth-child(2) > div:nth-child(2) > div:nth-child(1) > div:nth-child(2)")
		title.Click()
		time.Sleep(1 * time.Second)

		//説明欄を埋める
		desc := page.Find("#notion-app > div > div.notion-overlay-container.notion-default-overlay-container > div:nth-child(2) > div > div:nth-child(2) > div.notion-scroller.vertical > div.whenContentEditable > div:nth-child(3) > div > div > div:nth-child(2) > div:nth-child(1) > div")
		desc.Click()
		editableDesc := page.Find("#notion-app > div > div.notion-overlay-container.notion-default-overlay-container > div:nth-child(2) > div > div:nth-child(2) > div.notion-scroller.vertical > div.whenContentEditable > div:nth-child(3) > div > div > div > div > div > div")
		editableDesc.Fill(task[1])

		//modalを閉じる
		corner := page.Find("#notion-app > div > div.notion-overlay-container.notion-default-overlay-container > div:nth-child(2) > div > div:nth-child(2) > div:nth-child(1) > div:nth-child(1) > div:nth-child(6) > div.notion-topbar-more-button.notion-focusable > svg > g > path:nth-child(2)")
		corner.MouseToElement()
		page.MoveMouseBy(80, 0)
		page.DoubleClick()
		time.Sleep(1 * time.Second)
	}
}

func addTask(page *agouti.Page) {
	newButton := page.Find("#notion-app > div > div.notion-cursor-listener > div:nth-child(2) > div.notion-frame > div.notion-scroller.vertical.horizontal > div:nth-child(3) > div > div > div > div:nth-child(1) > div:nth-child(2) > div:nth-child(1) > div > div:nth-child(5)")
	err := newButton.MouseToElement()
	if err != nil {
		log.Fatal(err)
	}
	err = newButton.Click()
	if err != nil {
		log.Fatal(err)
	}
}
