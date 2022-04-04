package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/joho/godotenv"
)

func main() {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithDebugf(log.Printf),
	)

	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file!")
	}

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	var result string
	username_selector := `#clogs > input:nth-child(2)`
	password_selector := `#clogs > input:nth-child(3)`
	submit_selector := `#clogs-captcha-button`
	err = chromedp.Run(ctx, submit(`https://noip.com/login`, username_selector, password_selector, submit_selector, username, password, &result))

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Go's time.After example\n%s", result)
}

func submit(urlstr, username_select, password_select, submit_selector, username, password string, result *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.WaitVisible(username_select),
		chromedp.SendKeys(username_select, username),
		chromedp.SendKeys(password_select, password),
		chromedp.Submit(submit_selector),
		//chromedp.Text(`#main-navbar > div > div.navbar-header > a`, result),
		chromedp.WaitVisible(`#content-wrapper > div:nth-child(1) > div:nth-child(1) > div > div > div > h1`, chromedp.ByID),
		chromedp.Text(`#content-wrapper > div:nth-child(1) > div:nth-child(1) > div > div > div > h1`, result),
		//chromedp.Text(`#dashboard-nav > a > span`, result),
	}
}
