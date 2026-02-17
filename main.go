package main

import (
	"image/png"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/corona10/goimagehash"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/launcher"
	_ "github.com/joho/godotenv/autoload"
	"gopkg.in/telebot.v4"
)

func main() {
	defer os.RemoveAll("tmp")

	hash6, hash20 := getHash()  

	pref := telebot.Settings{
		Token: os.Getenv("TELEGRAM_TOKEN"),
	}

	b, err := telebot.NewBot(pref)
	if err != nil {
		panic(err)
	}

	u := launcher.NewUserMode().
		Leakless(true).
		UserDataDir("tmp/t").
		Set("disable-default-apps").
		Set("no-first-run").
		Set("no-sandbox").
		Set("window-size", "768,1080").
		Headless(false).
		Bin(os.Getenv("CHROME_PATH")). 
		MustLaunch()
	page := rod.New().
		ControlURL(u).
		MustConnect().
		NoDefaultDevice().
		MustPage(os.Getenv("LINK"))

	// <-time.After(time.Second * 15)
	// page.MustScreenshot("./loadedpage.png")
	// panic("look at the loaded page")

	println("awaiting for max 40 seconds for page to load...")
	page.Timeout(40 * time.Second).
		MustElement("div.form-item.form__item.form__item--radio.form__item--search-type.form__item--radio--2 > label").
		MustClick()

	<-time.After(time.Second * 3)

	page.Keyboard.Press(input.Tab)
	<-time.After(time.Millisecond * 600)
	page.Keyboard.Press(input.Tab)
	<-time.After(time.Millisecond * 700)

	page.InsertText("400720714")
	<-time.After(time.Millisecond * 300)
	page.Keyboard.Press(input.Enter)

	<-time.After(time.Second * 5)
	page.MustElement("body > div.dialog-off-canvas-main-canvas > div > header > div.site-header-middle > button").MustRemove()
	page.MustElement(".disconnection-detailed-table-container").MustScreenshot("./img/638.png")

	page.Keyboard.Press(input.Tab)
	<-time.After(time.Millisecond * 600)
	page.Keyboard.Press(input.Tab)
	<-time.After(time.Millisecond * 700)

	page.InsertText("400910046")
	<-time.After(time.Millisecond * 300)
	page.Keyboard.Press(input.Enter)

	<-time.After(time.Second * 3)
	page.MustElement(".disconnection-detailed-table-container").MustScreenshot("./img/2037.png")
	
	println("screenshots taken, comparing with the previous ones...")

	hash6new, hash20new := getHash()

	ids := ParseIDs(os.Getenv("id1"))
	recipients := make([]*telebot.Chat, len(ids))
	for i, id := range ids {
		recipients[i] = &telebot.Chat{ID: int64(id)}
	}


	photo6 := &telebot.Photo{
		File:    telebot.FromDisk("./img/638.png"),
		Caption: "6/38",
	}

	photo20 := &telebot.Photo{
		File:    telebot.FromDisk("./img/2037.png"),
		Caption: "20/37",
	}

	if v, _ := hash6.Distance(hash6new); v > 5 {
		for _, recipient := range recipients {
			_, err = b.Send(recipient, photo6, telebot.ModeHTML)
			if err != nil {
				log.Printf("Failed to send photo: %v", err)
			}
		}
	}

	if v, _ := hash20.Distance(hash20new); v > 5 {
		for _, recipient := range recipients {
			_, err = b.Send(recipient, photo20, telebot.ModeHTML)
			if err != nil {
				log.Printf("Failed to send photo: %v", err)
			}

		}
	}
	println("done")
}


func getHash() (*goimagehash.ImageHash, *goimagehash.ImageHash) {
	img6, err := os.Open("img/638.png")
	if err != nil {
		panic(err)
	}
	img20, err := os.Open("img/2037.png")
	if err != nil {
		panic(err)
	}
	defer img6.Close()
	defer img20.Close()

	i6, err := png.Decode(img6)
	if err != nil {
		panic(err)
	}
	i20, err := png.Decode(img20)
	if err != nil {
		panic(err)
	}
	hash6, err := goimagehash.AverageHash(i6)
	if err != nil {
		panic(err)
	}
	hash20, err := goimagehash.AverageHash(i20)
	if err != nil {
		panic(err)
	}
	return hash6, hash20
}

func ParseIDs(ids string) []int {
	idStrings := strings.Split(ids, ",")
	idInts := make([]int, len(idStrings))

	for i, idStr := range idStrings {
		id, err := strconv.Atoi(strings.TrimSpace(idStr))
		if err != nil {
			log.Panic(err)
		}
		idInts[i] = id
	}

	return idInts
}
