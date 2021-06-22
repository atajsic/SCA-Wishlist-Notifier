package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gregdel/pushover"
	"github.com/robfig/cron"
)

func main() {

	log.Printf("Wishlist Notifier loading...")

	//get environments:
	w, ok := os.LookupEnv("WL_SCA_WLID")
	if !ok {
		log.Fatal("WL_SCA_WLID not set")
	}

	a, ok := os.LookupEnv("WL_PUSHOVER_APP")
	if !ok {
		log.Fatal("WL_PUSHOVER_APP not set")
	}

	r, ok := os.LookupEnv("WL_PUSHOVER_RECIPIENT")
	if !ok {
		log.Fatal("WL_PUSHOVER_RECIPIENT not set")
	}

	c, ok := os.LookupEnv("WL_CRON")
	if !ok {
		log.Fatal("WL_CRON not set")
	}

	pa := pushover.New(a)
	pr := pushover.NewRecipient(r)

	timer := cron.New()
	err := timer.AddFunc(c, func() { checkWishlist(w, pa, pr) })
	if err != nil {
		log.Printf("Cron failed to initiate: \n%s", err)
	}
	timer.Start()
	log.Printf("Cron started.")

	select {}

}

func checkWishlist(w string, pa *pushover.Pushover, pr *pushover.Recipient) {

	log.Println("Checking Wishlist...")

	res, err := http.Get("https://www.supercheapauto.com.au/showotherwishlist?WishListID=" + w)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	//load the HTML doc
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the wishlist items
	doc.Find(".wishlist-row").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title, price, and sale price if applicable
		title := s.Find(".name a").Text()
		title = strings.Trim(title, " ")
		title = strings.Trim(title, "\n")
		link, _ := s.Find(".name a").Attr("href")
		sales := s.Find(".price-sales").Text()
		standard := s.Find(".price-standard").Text()

		sales_f, _ := strconv.ParseFloat(strings.Trim(sales, "$"), 32)
		standard_f, _ := strconv.ParseFloat(strings.Trim(standard, "$"), 32)
		perc := (standard_f - sales_f) / standard_f * 100

		//item on sale
		if len(standard) > 0 {

			message := &pushover.Message{
				Message:  fmt.Sprintf("%s\n%s (Usually %s)\n%.2f%% off", title, sales, standard, perc),
				Title:    "Wishlist Notifier",
				URL:      link,
				URLTitle: "View Product",
			}

			_, err := pa.SendMessage(message, pr)
			if err != nil {
				log.Panic(err)
			}

			log.Println(title + " - " + sales)
		}

	})

}
