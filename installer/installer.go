package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type filter struct {
	DisplayName string `json:"displayName"`
	FilterURL   string `json:"filterURL"`
	FilterName  string `json:"filterName"`
}

func main() {
	// path := getInstallPath()

	filterList := getFilters()

	printMenu(filterList)
	// filter := getUserSelection()

	// filterName := "ThioleLootFilter.filter"
	// filterURL := "https://raw.githubusercontent.com/icbat/LootFilter/master/" + filterName
	// fmt.Println("Grabbing " + filterName + " from: " + filterURL)
	// downloadTo(filterURL, path+"/"+filterName)
	//
	// fmt.Println(`Done!
	//
	// To finish the install:
	//
	// 1) Restart Path of Exile Beta if you have it open
	// 2) Start game
	// 3) Login to Beta
	// 4) Esc > Options
	// 5) GO to UI tab, at the bottom select the filter, once it says "Filter loaded successfully" no restart required, you are good to go
	// 6) Any updates to the filter can be reloaded without restarting the game by clicking the "reload" button in options
	//
	//
	// Press any key to exit!`)
	// reader := bufio.NewReader(os.Stdin)
	// reader.ReadString('\n')
}

func getInstallPath() string {
	fmt.Println("Checking environment")
	basePath := os.Getenv("UserProfile")
	if basePath == "" {
		panic(" -- Couldn't resolve environment variables!")
	}
	fmt.Println("Found base path of: " + basePath)

	fmt.Println("Verifying directories")
	path := basePath + "/Documents/My Games/Path of Exile"
	if !pathExists(path) {
		panic(" -- Couldn't find a folder at " + path)
	}
	fmt.Println("Found PoE folder at: " + path)

	fmt.Println("Reticulating splines")
	return path
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}

func getFilters() []filter {

	rawJSON, err := ioutil.ReadFile("filters.json")
	if err != nil {
		panic(err)
	}

	var filterList []filter
	err = json.Unmarshal(rawJSON, &filterList)
	if err != nil {
		panic(err)
	}
	return filterList
}

func printMenu(filters []filter) {
	fmt.Println("\nAvailable filters:")
	for index, filter := range filters {
		fmt.Printf("\t%d : %s\n", index+1, filter.DisplayName)
	}
}

func cleanText(text string) string {
	text = strings.Replace(text, "\n", "", -1)
	text = strings.Replace(text, "\r", "", -1)
	return text
}

func downloadTo(url string, targetFile string) {
	response, getError := http.Get(url)
	if getError != nil {
		fmt.Println(" -- Couldn't GET: " + url)
		return
	}
	defer response.Body.Close()

	bytes, readError := ioutil.ReadAll(response.Body)
	if readError != nil {
		fmt.Println(" -- Couldn't read file at: " + url)
		return
	}

	fileError := ioutil.WriteFile(targetFile, bytes, 0644)
	if fileError != nil {
		fmt.Println(" -- Couldn't write to file at: " + targetFile)
		return
	}
}
