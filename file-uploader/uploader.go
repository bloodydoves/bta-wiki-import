package uploader

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

/* Examples:
https://gist.github.com/mattetti/5914158/f4d1393d83ebedc682a3c8e7bdc6b49670083b84
https://stackoverflow.com/questions/72735728/http-post-containing-binary-data-in-golang

*/

func createMultipartFormData(fieldName, fileName string, params map[string]string) (bytes.Buffer, *multipart.Writer) {
	var buff bytes.Buffer
	writer := multipart.NewWriter(&buff)
	var fileWriter io.Writer
	file, err := os.Open(fileName)
	if fileWriter, err = writer.CreateFormFile(fieldName, file.Name()); err != nil {
		fmt.Println("Error: ", err)
	}
	if _, err = io.Copy(fileWriter, file); err != nil {
		fmt.Println("Error: ", err)
	}
	for key, value := range params {
		writer.WriteField(key, value)
	}
	writer.Close()
	return buff, writer
}

func Upload(url, targetFile string, params map[string]string) error {
	buff, writer := createMultipartFormData("file", targetFile, params)

	req, err := http.NewRequest("POST", url, &buff)
	if err != nil {
		return
	}
	// Don't forget to set the content type, this will contain the boundary.
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	response, error := client.Do(req)
	if err != nil {
		panic(error)
	}
	defer response.Body.Close()

	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))
	return body
}
