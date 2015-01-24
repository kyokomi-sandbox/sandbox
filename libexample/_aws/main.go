package main

import (
	"github.com/stripe/aws-go/aws"
	"github.com/stripe/aws-go/gen/s3"
	"log"
	"io/ioutil"
	"os"
	"fmt"
	"encoding/json"
	"github.com/k0kubun/pp"
)

type Media struct {
	DisplayURL    string `json:"display_url"`
	ExpandedURL   string `json:"expanded_url"`
	ID            int    `json:"id"`
	IDStr         string `json:"id_str"`
	Indices       []int  `json:"indices"`
	MediaURL      string `json:"media_url"`
	MediaURLHTTPS string `json:"media_url_https"`
	Sizes         Sizes  `json:"sizes"`
	Type string `json:"type"`
	URL  string `json:"url"`
}

type Sizes struct {
	Large Size  `json:"large"`
	Medium Size `json:"medium"`
	Small Size  `json:"small"`
	Thumb Size  `json:"thumb"`
}

type Size struct {
	H      int    `json:"h"`
	Resize string `json:"resize"`
	W      int    `json:"w"`
}

func main() {

	fmt.Println("============== Start1 ================")
	c1 := accessKeyCreds()
	printGetObject(c1)

	fmt.Println("============== Start2 ================")
	c2 := credentials()
	printGetObject(c2)

	fmt.Println("============== Start3 ================")
	c3 := iamCreds()
	printGetObject(c3)
}

func accessKeyCreds() *s3.S3 {
	accessKey := os.Getenv("AWS_ACCESS_KEY")
	secretKey := os.Getenv("AWS_SECRET_KEY")
	creds := aws.Creds(accessKey, secretKey, "")
	return s3.New(creds, "ap-northeast-1", nil)
}

func iamCreds() *s3.S3 {
	return s3.New(aws.IAMCreds(), "ap-northeast-1", nil)
}

func credentials() *s3.S3 {
	// 指定なしで、~/.aws/credentialsのdefaultになる
	creds, err := aws.ProfileCreds("", "", 0)
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Fatalln(err)
	}

	return s3.New(creds, "ap-northeast-1", nil)
}

func printGetObject(client *s3.S3) {
	req := s3.GetObjectRequest{}
	req.Bucket = aws.String("kyokomi-foo")
	req.Key = aws.String("bar/media.json")

	res, err := client.GetObject(&req)
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(data))

	var m Media
	if err := json.Unmarshal(data, &m); err != nil {
		log.Fatalln(err)
	}

	pp.Println(m)
}
