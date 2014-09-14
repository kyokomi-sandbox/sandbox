package main

import (
	"fmt"
	"net/http"
	"net/url"
	"log"
	"io/ioutil"
	"strings"
)


func doExcel() {
	readExcel()
}

func doTemplate() {
	text := CreateTemplateTree("Hello", "golang").Execute()
	fmt.Println("Result:", text)
}

func doString() {
	Sample()
}

func doScan() {
	c := CliScan {
		Scans: []Scan{
			{
				Key: "answer",
				Message: "Do you want to create one? [Y/n]",
			},
		},
	}
	fmt.Println("answer >", c.scan("answer"))
}

func doHttpPUT() {
	params := make(url.Values)
	params.Set("title", "hogefuga11111111")

	// http put
	id      := "67040"
	issueID := "67696"
	token   := "ssssssssssssssssssssss"

	urlStr := fmt.Sprintf("https://gitlab.com/api/v3/projects/%s/issues/%s?private_token=%s", id, issueID, token)

	req, err := http.NewRequest("PUT", urlStr, strings.NewReader(params.Encode()))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}

func main() {
	fmt.Println("Hello Go Sandbox!")


}

