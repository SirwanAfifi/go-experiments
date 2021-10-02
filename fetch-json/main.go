package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type user struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

type response struct {
  PageCount int `json:"pageCount"`
  TotalCounts int `json:"totalCounts"`
  Posts []user `json:"posts"`
}

func main() {
	url := flag.String("url", "localhost:8888/api/users", "URL To Resource")
	flag.Parse()

	result, err := getUsers(*url)

	if err != nil {
		log.Fatal(err)
	}

	for _, u := range result.Posts {
		fmt.Printf("User ID: %s - Full Name: %s\n", u.Id, u.Name)
	}
}

func getUsers(apiURL string) (response, error) {
	client := http.Client {
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "sirwan")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	result := response{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal(err)
	}

	return result, nil
}