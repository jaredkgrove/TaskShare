package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://www.googleapis.com/tasks/v1/lists/RW9yQXktOXBOZ09rQS1rUw/tasks", nil)

	//get the new access token from google api playground or something like that
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %a", getHardCodedAccessToken()))
	response, err := client.Do(req)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))
}
