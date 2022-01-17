package main

import (
	"bufio"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sclevine/agouti"
	"log"
	"os"
	"transfer/notion"
	"transfer/trello"
)

func main() {
	loadEnv()

	driver := startChromeDriver()
	defer driver.Stop()

	page, err := driver.NewPage() //Driverに対応したページを返す。（今回はChrome）
	if err != nil {
		log.Fatal(err)
	}

	trello.Login(page)
	tasks := trello.DrainTasks(page)

	confirm(tasks)

	notion.Login(page)

	notion.PasteTasks(page, tasks)

	fmt.Println("タスクの移行が完了しました。savingが終わったらエンターしてください。:")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
}

func confirm(tasks [][]string) {
	for _, task := range tasks {
		fmt.Print("タイトル:　")
		fmt.Println(task[0])
		fmt.Println("------------説明------------")
		_, err := fmt.Println(task[1])
		if err != nil {
			log.Fatal("Error: 説明がない？")
		}
		fmt.Println("---------------------------")
	}

	fmt.Println("タスクの取得が完了しました。ノーションに移行してよければエンターを押してください。:")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}
}

func startChromeDriver() *agouti.WebDriver {
	//ChromeDriverを使用するための記述
	driver := agouti.ChromeDriver()

	err := driver.Start()
	if err != nil {
		log.Fatal(err)
	}

	return driver
}
