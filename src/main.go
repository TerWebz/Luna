package main

import (
	"fmt"
	"io"
	"os"
	"net/http"
) 

func download() (err error) {
	var name string
	var base string	
	var game string
	var uri string
	
	base = "https://javier06.ejemplo.me/"

	fmt.Printf("Enter a game name: ")
	fmt.Scanf("%s", &name)

	game = name + ".ciso"
	uri = base + game

	response, err := http.Get(uri)
	if err != nil {
		return err
		os.Exit(1)
	}

	if response.StatusCode != 200 {
   		fmt.Println("Error: The servers are down or the URL/Game that has been entered is invalid.")
		os.Exit(1)
    	} else {
		fmt.Printf("Downloading %s... \n", game)
	} // This is before anything because we want to determine if the user put a valid game before we make the file, to not make garbage or filler files.
    	defer response.Body.Close()

	file, err := os.Create(game)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	var title string = "Luna v1.0.1.1"
	fmt.Println(title)
	download();
} 
