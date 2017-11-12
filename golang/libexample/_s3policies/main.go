package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/kyokomi/s3gopolicy"
)

func main() {
	policies := s3gopolicy.UploadPolicies{
		URL: "",
		Form: s3gopolicy.PoliciesForm{
			AWSAccessKeyID: "",
			ContentType:    "image/jpg",
			ObjectKey:      "uploads/user_TaW-06CdSh2dyP34mCyxRA/file_OQTKEFNFNL165E6PCVGGE9T140.jpg",
			Policy: "",
			Signature:      "",
		},
	}

	//policies := getPolicies()
	sessionToken := ""
	err := Upload(
		policies.URL,
		"./images/me.jpg", policies, sessionToken)
	log.Println(err.Error())
}

func getPolicies() s3gopolicy.UploadPolicies {
	policies, _ := s3gopolicy.CreatePolicies(s3gopolicy.AWSCredentials{
		AWSAccessKeyID: "",
		AWSSecretKeyID: "",
	}, s3gopolicy.UploadConfig{
		UploadURL:   "",
		// 
		ObjectKey:   "uploads/user_TaW-06CdSh2dyP34mCyxRA/file_RQ4DL0TPG532PD9GA0CJQ49SL0.jpg",
		ContentType: "image/jpg",
		FileSize:    57779,
	})
	return policies
}

func Upload(url, file string, policies s3gopolicy.UploadPolicies, sessionToken string) (err error) {
	// Add your image file
	f, err := os.Open(file)
	if err != nil {
		return
	}
	defer f.Close()

	fileInfo, _ := f.Stat()
	log.Println(fileInfo.Size())

	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	if err := w.WriteField("key", policies.Form.ObjectKey); err != nil {
		return err
	}
	if err := w.WriteField("Content-Type", policies.Form.ContentType); err != nil {
		return err
	}
	if err := w.WriteField("AWSAccessKeyId", policies.Form.AWSAccessKeyID); err != nil {
		return err
	}
	if err := w.WriteField("policy", policies.Form.Policy); err != nil {
		return err
	}
	if err := w.WriteField("signature", policies.Form.Signature); err != nil {
		return err
	}

	fw, err := w.CreateFormFile("file", file)
	if err != nil {
		return
	}
	if _, err = io.Copy(fw, f); err != nil {
		return
	}

	// Don't forget to close the multipart writer.
	// If you don't close it, your request will be missing the terminating boundary.
	if err := w.Close(); err != nil {
		return err
	}

	// Now that you have a form, you can submit it to your handler.
	req, err := http.NewRequest("POST", url, &b)
	if err != nil {
		return
	}
	// Don't forget to set the content type, this will contain the boundary.
	req.Header.Set("Content-Type", w.FormDataContentType())
	if sessionToken != "" {
		req.Header.Set("x-amz-security-token", sessionToken)
	}

	// Submit the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return
	}

	// Check the response
	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("bad status: %s", res.Status)
		data, _ := ioutil.ReadAll(res.Body)
		log.Println(string(data))
	}
	return
}
