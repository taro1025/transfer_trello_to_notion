package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sclevine/agouti"
	"log"
	"time"
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

	notion.Login(page)

	notion.PasteTasks(page, tasks)

	time.Sleep(10 * time.Second)
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
