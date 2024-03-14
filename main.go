package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

const confFile = "./conf.json"

type Config struct {
	Urls            []string
	IntervalMinutes int
	EmailTo         string
}

func checkUrl(url string) (bool, error) {

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return false, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("received status code not equal to 200: %d", resp.StatusCode)
	}
	return true, nil
}

func main() {
	fmt.Println("Starting web-go-check.")
	conf, err := readConf()

	if err != nil {
		fmt.Printf("The configuration file '%s' is incorrect. \n %s", confFile, err)
	}

	erroredUrls := map[string]bool{}

	for {
		fmt.Printf("%s Checking all the urls\n", time.Now().Format("2006-01-02 15:04:05"))

		for _, url := range conf.Urls {
			success, err := checkUrl(url)
			if !success {

				errorString := fmt.Sprintf("☒ %s\n", url)
				if err != nil {
					errorString += fmt.Sprintf("the exception was:\n %s \n\n", err)
				}
				fmt.Println("\033[31m", errorString, "\033[0m")

				if _, exists := erroredUrls[url]; !exists {
					err = SendMail(FormatErrorEmail(url, errorString), conf.EmailTo, url+" is DOWN")
					if err != nil {
						fmt.Println(err)
					}
					erroredUrls[url] = true
				}

			} else {

				if _, exists := erroredUrls[url]; exists {
					delete(erroredUrls, url)
					err = SendMail(FormatUpEmail(url), conf.EmailTo, url+" is UP")
					if err != nil {
						fmt.Println(err)
					}

				}

				fmt.Printf(" \033[32m☑ %s \n\033[0m", url)
			}
		}

		time.Sleep(time.Duration(conf.IntervalMinutes) * time.Minute)
	}

}

func readConf() (*Config, error) {
	data, err := os.ReadFile(confFile)
	if err != nil {
		return nil, err
	}

	var obj Config
	err = json.Unmarshal(data, &obj)
	if err != nil {
		return nil, err
	}

	return &obj, nil
}
