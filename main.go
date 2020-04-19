package main

import (
	"os"

	"github.com/LeeHoSung/learngo/scrapper"
	"github.com/labstack/echo"
)

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScrap(c echo.Context) error {
	defer os.Remove("JobKorea.csv")
	keyword := scrapper.CleanString(c.FormValue("keyword"))
	scrapper.Scrape(keyword)
	return c.Attachment("JobKorea.csv", "JobKorea.csv")
}

func main() {

	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrap", handleScrap)
	e.Logger.Fatal(e.Start(":1323"))
}
