package main

import (
	"fmt"
	"io"
	"os"
	"net/http"
) 

func download() (err error) {
	var name string
	fmt.Printf("Enter a Game: ")
	fmt.Scanf("%s", &name)	

	var base string = "https://javier06.ejemplo.me/"
	var game string = name + ".ciso"
	var uri string = base + game

	response, err := http.Get(uri)
	if err != nil {
		return err
		os.Exit(1)
	}

	if response.StatusCode != 200 {
   		fmt.Println("Error: Invalid URI!")
		os.Exit(1)
    	}

    	defer response.Body.Close()

	file, err := os.Create(game)

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	var title string = "Luna v1.0.0"
	fmt.Println(title)
	download();
} 
