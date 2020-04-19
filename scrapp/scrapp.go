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

var baseURL = "https://kr.indeed.com/jobs?q=JAVA"

type extractedJob struct {
	id       string
	title    string
	location string
	salary   string
	company  string
}

func main() {
	var jobs []extractedJob
	getTotalPage := getPages()
	fmt.Println("총 페이지 개수 : ", getTotalPage)
	for i := 0; i < getTotalPage; i++ {
		extractedJob := getPage(i)
		jobs = append(jobs, extractedJob...)
	}
	WriteJobs(jobs)
}

//Write Jobs
func WriteJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	CheckErr(err)
	w := csv.NewWriter(file)
	defer w.Flush()

	header := []string{"주소", "공고제목", "회사위치", "연봉", "회사이름"}
	wErr := w.Write(header)
	CheckErr(wErr)

	for _, val := range jobs {
		JobSlice := []string{"https://kr.indeed.com/%EC%B1%84%EC%9A%A9%EB%B3%B4%EA%B8%B0?jk=" + val.id, val.title, val.location, val.salary, val.company}
		jwErr := w.Write(JobSlice)
		CheckErr(jwErr)
	}

}

//URL Error check
func CheckErr(err error) {

	if err != nil {
		log.Fatalln(err)
	}
}

//URL StatusCode check
func CheckCode(resp *http.Response) {

	if resp.StatusCode >= 400 {
		log.Fatalln(resp.StatusCode)
	}
}

//All pages count
func getPages() int {
	getTotalPage := 0
	resp, err := http.Get(baseURL)
	CheckErr(err)
	CheckCode(resp)
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	CheckErr(err)
	doc.Find(".pagination").Each(func(i int, card *goquery.Selection) {

		getTotalPage = card.Find("a").Length()
	})

	return getTotalPage
}

//get page info
func getPage(page int) []extractedJob {
	var jobs []extractedJob
	pageURL := baseURL + "&start=" + strconv.Itoa(page*10)
	fmt.Println(pageURL)
	resp, err := http.Get(pageURL)
	CheckErr(err)
	CheckCode(resp)
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	CheckErr(err)
	doc.Find(".jobsearch-SerpJobCard").Each(func(i int, card *goquery.Selection) {
		job := extractJob(card)
		jobs = append(jobs, job)
	})

	return jobs
}

//extractJob
func extractJob(card *goquery.Selection) extractedJob {
	id, _ := card.Attr("data-jk")
	id = CleanString(id)
	title, _ := card.Find(".title > a").Attr("title")
	title = CleanString(title)
	company := card.Find(".sjcl > div > .company").Text()
	company = CleanString(company)
	location := card.Find(".sjcl > span").Text()
	salary := card.Find(".salarySnippet > .salary > span").Text()
	salary = CleanString(salary)
	//fmt.Println("공고 제목 : ", title)
	//fmt.Println("회사 이름 : ", company)
	//fmt.Println("회사 위치 : ", location)
	//fmt.Println("연봉 : ", salary)

	return extractedJob{id: id, title: title, company: company, location: location, salary: salary}

}

//String Trim
func CleanString(str string) string {
	return strings.TrimSpace(str)
}
