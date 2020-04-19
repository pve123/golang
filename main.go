package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	title   string
	loc     string
	company string
}

var baseURL = "http://www.jobkorea.co.kr/Search/?stext=Python"

func main() {

	var jobs []extractedJob
	getTotalPages := getPages()
	for i := 0; i <= getTotalPages; i++ {
		job := getPage(i + 1)
		jobs = append(jobs, job...)
	}
	WriteJob(jobs)
}

//WriteJob
func WriteJob(jobs []extractedJob) {
	file, err := os.Create("JobKoreaPython.csv")
	CheckErr(err)

	w := csv.NewWriter(file)

	defer w.Flush()

	header := []string{"공고제목", "회사위치", "회사이름"}
	wErr := w.Write(header)
	CheckErr(wErr)

	for _, val := range jobs {
		JobsCon := []string{val.title, val.loc, val.company}
		wErr := w.Write(JobsCon)
		CheckErr(wErr)
	}

}

//get Page
func getPage(page int) []extractedJob {
	var first, second, third []string
	var jobs []extractedJob
	PageURL := baseURL + "&Page_No=" + strconv.Itoa(page)
	fmt.Println(PageURL)

	resp, err := http.Get(PageURL)
	CheckErr(err)
	CheckCode(resp)

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	CheckErr(err)

	doc.Find(".title.dev_view").Each(func(i int, card *goquery.Selection) {

		title, _ := card.Attr("title")
		CleanString(title)
		first = append(first, title)
	})

	doc.Find(".loc.long").Each(func(i int, card *goquery.Selection) {
		loc := card.Text()
		CleanString(loc)
		second = append(second, loc)
	})
	doc.Find(".post-list-corp > .name.dev_view").Each(func(i int, card *goquery.Selection) {
		company := card.Text()
		CleanString(company)
		third = append(third, company)
	})

	for i := 0; i < 20; i++ {
		job := extractedJob{title: first[i], loc: second[i], company: third[i]}
		jobs = append(jobs, job)
	}
	return jobs
}

//get Pages All
func getPages() int {
	getTotalPages := 0
	resp, err := http.Get(baseURL)
	CheckErr(err)
	CheckCode(resp)

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	CheckErr(err)
	doc.Find(".recruit-info > .lists > .lists-cnt.dev_list> .tplPagination.newVer.wide").Each(func(i int, card *goquery.Selection) {
		getTotalPages = card.Find("ul > li >a").Length()

	})
	return getTotalPages
}

//Error Check
func CheckErr(err error) {

	if err != nil {
		log.Fatalln(err)
	}
}

//Code Check
func CheckCode(resp *http.Response) {

	if resp.StatusCode >= 400 {
		log.Fatalln(resp)
	}
}

//Clean String
func CleanString(str string) string {

	return strings.TrimSpace(str)
}
