package main

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func getTickers() []string {
	// returns slice with list of tickers from file
	file, error := os.Open("tickers.txt")
	tickers := []string{}

	if error != nil {
		log.Fatal(error)
	}

	// defer causes function to be executed after all of the
	// surrounding methods return or panic with an error
	// this ensures that file is closed even if an error occurs
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		tickers = append(tickers, line)
	}

	// check if scanner picked up error
	if error := scanner.Err(); error != nil {
		log.Fatal(error)
	}

	return tickers
}

func getJson(url string, target interface{}) error {
	resp, err := http.Get(url)

	// if there is an error in the response, then
	// log the error and return dictionary with error
	if err != nil {
		log.Fatal(err)
		return err
	}

	// ensure that response streaming is stoped
	// before function is exited
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}
