package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/jasonlvhit/gocron"
	"github.com/jinzhu/gorm"
	"os"
	"strconv"
	"strings"
)

var citiesUrl = map[string]string{
	"Rotterdam": "/s/Rotterdam--Netherlands?sorting=mostRecent",
}

type Room struct {
	gorm.Model
	Price         string
	PricePeriod   string
	PriceCurrency string
	Type          string
	City          string
	Districts     string
	Url           string
	ScraperName   string
}

var DB *gorm.DB = nil
var err error = nil

//var rooms = make([]Room, 0)

/**
* connect with data base with ..env file params
* just edit all data in ..env file
 */
func ConnectToDatabase() {
	if DB == nil {
		DB, err = gorm.Open("mysql", os.Getenv("DATABASE_USERNAME")+":"+os.Getenv("DATABASE_PASSWORD")+"@tcp("+os.Getenv("DATABASE_HOST")+":"+os.Getenv("DATABASE_PORT")+")/"+os.Getenv("DATABASE_NAME")+"?charset=utf8mb4&parseTime=True&loc=Local&character_set_server=utf8mb4")
	}
	if err != nil {
		fmt.Println("Connect To Database Error:", err.Error())
		return
	}
	debug, _ := strconv.ParseBool(os.Getenv("DEBUG_DATABASE"))
	if os.Getenv("APP_ENV") == "local" {
		DB.LogMode(debug)
	}
}

func grabWithMap() {
	for _, url := range citiesUrl {
		c := colly.NewCollector()
		c.OnHTML(`div[id=app-root]`, func(e *colly.HTMLElement) {
			e.ForEachWithBreak("div ", func(i int, el *colly.HTMLElement) bool {
				if el.Attr("class") == "MuiGrid-root MuiGrid-container MuiGrid-spacing-xs-3" {
					el.ForEach("div ", func(i int, el *colly.HTMLElement) {
						if el.Attr("class") == "MuiGrid-root MuiGrid-item MuiGrid-grid-xs-12 MuiGrid-grid-sm-12 MuiGrid-grid-md-4" {
							room := Room{
								ScraperName: os.Getenv("SCRAPER_NAME"),
							}
							url := os.Getenv("BASE_SCRAPER_URL") + el.ChildAttr("a", "href")
							DB.Where("url = ?", url).First(&room)
							if room.ID != 0 {
								goto endOfIteration
							}
							room.Url = url
							urlParts := strings.Split(strings.Split(room.Url, "?")[0], "/")
							room.City = urlParts[len(urlParts)-2]
							room.Districts = urlParts[len(urlParts)-1]
							el.ForEachWithBreak("ul ", func(i int, el *colly.HTMLElement) bool {
								if el.Attr("class") == "MuiList-root makeStyles-infoContainer-221 MuiList-padding" {
									el.ForEachWithBreak("li ", func(i int, el *colly.HTMLElement) bool {
										switch i {
										case 0:
											room.Price = el.ChildAttr("meta[itemprop=price]", "content")
											room.PriceCurrency = el.ChildAttr("meta[itemprop=priceCurrency]", "content")
											el.ForEach("span ", func(i int, el *colly.HTMLElement) {
												if el.Attr("class") == "MuiTypography-root makeStyles-billsContainer-226 makeStyles-caption-52 makeStyles-overflow-initial-35 makeStyles-color-default-38 MuiTypography-caption MuiTypography-displayBlock" {
													room.PricePeriod = strings.Split(strings.Split(el.Text, "/")[1], " ")[0]
												}
											})
											break
										case 1:
											room.Type = strings.Split(el.Text, "•")[0]
											return false

										}
										return true
									})
									return false
								}
								return true
							})
							DB.Create(&room)
							//x, _ := json.Marshal(room)
							//fmt.Println("room", i+1, string(x))
							//rooms = append(rooms, room)
						}
						endOfIteration:
					})
					return false
				}
				return true
			})
		})
		c.OnError(func(r *colly.Response, err error) {
			fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
		})
		errr := c.Visit(os.Getenv("BASE_SCRAPER_URL") + url)
		if errr != nil {
			fmt.Println("errr", errr.Error())
		}
	}
}

func jobs() {
	gocron.Every(2).Hours().From(gocron.NextTick()).Do(grabWithMap)
	gocron.Start()
}
