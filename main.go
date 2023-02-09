package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	g "github.com/serpapi/google-search-results-golang"
)

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	parameter := map[string]string{
		"api_key":       os.Getenv("API_KEY"),
		"engine":        os.Getenv("ENGINE"),
		"q":             os.Getenv("Q"),
		"location":      os.Getenv("LOCATION"),
		"google_domain": os.Getenv("GOOGLE_DOMAIN"),
		"gl":            os.Getenv("GL"),
		"hl":            os.Getenv("HL"),
	}

	search := g.NewGoogleSearch(parameter, os.Getenv("API_KEY"))
	results, err := search.GetJSON()

	file, err := os.Create(os.Getenv("RESULTS_FILE"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(file, "%v\n", results)
}
