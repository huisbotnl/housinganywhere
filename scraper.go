package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/jasonlvhit/gocron"
	"os"
	"strings"
	"time"
)

var roomsUrl = make(map[string]string)

func setupDataMaps() {
	setupAmenitiesMap()
	setupBedroomsMap()
	setupCitiesMap()
	setupFacilitiesMap()
	setupFurnituresMap()
	setupNeighborhoodsMap()
	setupTypesMap()
}

func grabWithMap() {
	startTime := time.Now()
	for _, url := range citiesUrl {
		c := colly.NewCollector()
		c.OnHTML(`div[id=app-root]`, func(e *colly.HTMLElement) {
			e.ForEachWithBreak("div ", func(i int, el *colly.HTMLElement) bool {
				if el.Attr("class") == "MuiGrid-root MuiGrid-container MuiGrid-spacing-xs-3" {
					el.ForEach("div ", func(i int, ell *colly.HTMLElement) {
						if ell.Attr("class") == "MuiGrid-root MuiGrid-item MuiGrid-grid-xs-12 MuiGrid-grid-sm-12 MuiGrid-grid-md-4" {
							room := Room{}
							roomUrl := os.Getenv("BASE_SCRAPER_URL") + ell.ChildAttr("a", "href")
							DB.Where("url = ?", roomUrl).First(&room)
							if room.ID == 0 {
								roomsUrl[roomUrl] = roomUrl
							}
						}
					})
					return false
				}
				return true
			})
		})
		c.OnError(func(r *colly.Response, err error) {
			fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
		})
		c.OnRequest(func(request *colly.Request) {
			(*request.Headers)["User-Agent"] = []string{"*"}
			fmt.Println("Visiting", request.URL)
		})
		errr := c.Visit(os.Getenv("BASE_SCRAPER_URL") + url)
		if errr != nil {
			fmt.Println("errr", errr.Error())
		}
		for _, roomUrl := range roomsUrl {
			c := colly.NewCollector()
			c.OnHTML(`div[class=makeStyles-contentContainer-1]`, func(e *colly.HTMLElement) {
				room := Room{
					Url:         roomUrl,
					ScraperName: os.Getenv("SCRAPER_NAME"),
				}
				e.ForEachWithBreak("div", func(i int, el *colly.HTMLElement) bool {
					if i == 3 {
						el.ForEachWithBreak("span", func(i int, ell *colly.HTMLElement) bool {
							if i == 6 {
								room.PricePeriod = strings.Split(ell.Text, " ")[0]
								return false
							}
							return true
						})
						el.ForEachWithBreak("li", func(i int, ell *colly.HTMLElement) bool {
							if i < 2 || (i > 2 && i < 6) {
								return true
							}
							if i > 6 {
								return false
							}
							if i == 2 {
								room.SetCity(ell.Text)
							}
							if i == 6 {
								room.SetDistrict(ell.Text)
							}
							return true
						})
						el.ForEachWithBreak("p", func(i int, ell *colly.HTMLElement) bool {
							if i > 0 && i < 8 {
								return true
							}
							if i > 8 {
								return false
							}
							if i == 0 {
								priceData := strings.Split(ell.Text, "€")
								room.PriceCurrency = "EUR"
								room.Price = priceData[0]
							}
							if i == 8 {
								room.SetType(ell.Text)
							}
							return true
						})
						el.ForEachWithBreak("div", func(i int, ell *colly.HTMLElement) bool {
							if ell.Attr("data-test-locator") == "Listing/ListingHighlights/Furnished" {
								room.SetFurniture(ell.ChildText("p"))
							}
							if ell.Attr("data-test-locator") == "Listing/ListingHighlights/Bedroom Count" {
								room.SetBedroom(ell.ChildText("p"))
							}
							if ell.Attr("class") == "MuiGrid-root MuiGrid-item MuiGrid-grid-xs-10 MuiGrid-grid-md-10" {
								room.AddFacility(ell.Text)
								room.AddAmenity(ell.Text)
							}
							return true
						})
						return false
					}
					return true
				})
				room.Create()
			})
			c.OnError(func(r *colly.Response, err error) {
				fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
			})
			c.OnRequest(func(request *colly.Request) {
				(*request.Headers)["User-Agent"] = []string{"*"}
				fmt.Println("Visiting", request.URL)
			})
			errr := c.Visit(roomUrl)
			if errr != nil {
				fmt.Println("errr", errr.Error())
			}
		}
		roomsUrl = make(map[string]string)
		time.Sleep(30 * time.Second)
	}
	EndTime := time.Now()
	fmt.Println("job finished")
	fmt.Println("start at")
	fmt.Println(startTime.String())
	fmt.Println("end at")
	fmt.Println(EndTime.String())
}

func jobs() {
	gocron.Every(1).Hours().From(gocron.NextTick()).Do(grabWithMap)
	gocron.Start()
}
