package main

import (
	"context"
	"io/ioutil"
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
	var buffer []byte
	username_selector := `#clogs > input:nth-child(2)`
	password_selector := `#clogs > input:nth-child(3)`
	submit_selector := `#clogs-captcha-button`
	err = chromedp.Run(ctx, submit(`https://noip.com/login`, username_selector,
		password_selector, submit_selector,
		username, password, &result, buffer))

	if err := ioutil.WriteFile("screenshot.png", buffer, 0o644); err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Done with chromedp.\n")
}

func submit(urlstr, username_select,
	password_select, submit_selector,
	username, password string,
	result *string, screenshot_buffer []byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.WaitVisible(username_select),
		chromedp.SendKeys(username_select, username),
		chromedp.SendKeys(password_select, password),
		chromedp.Click(submit_selector),
		chromedp.WaitNotPresent(username_select),
		chromedp.FullScreenshot(&screenshot_buffer, 90),
		//chromedp.WaitVisible(`#content-wrapper > div:nth-child(1) > div:nth-child(1) > div > div > div > h1`),
		//chromedp.Text(`#content-wrapper > div:nth-child(1) > div:nth-child(1) > div > div > div > h1`, result),
	}
}
