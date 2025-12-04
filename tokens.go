package function

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

type Data struct {
	Policy         string `json:"policy"`
	Asset_name     string `json:"asset_name"`
	Asset_name_hex string `json:"asset_name_hex"`
	Fingerprint    string `json:"fingerprint"`
	Supply         string `json:"supply"`
}

// Define a struct to represent the JSON data structure
type Token struct {
	Data Data `json:"data"`
}

func getToken() Token {
	resp, err := http.Get("https://api.adastat.net/rest/v1/tokens/0691b2fecca1ac4f53cb6dfb00b7013e561d1f34403b957cbb5af1fa4e49474854.json?rows=minting&dir=asc") // Replace with your actual API endpoint
	if err != nil {
		log.Fatalf("Error making HTTP request: %v", err)
	}
	defer resp.Body.Close() // Ensure the response body is closed

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Unexpected status code: %d", resp.StatusCode)
	}

	// ... (decode JSON in the next step)
	var token Token
	err = json.NewDecoder(resp.Body).Decode(&token)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}

	fmt.Printf("Decoded Token: %+v\n", token)
	return token
}

func init() {
	functions.HTTP("Supply", supply)
}

// helloWorld writes "Hello, World!" to the HTTP response.
func supply(w http.ResponseWriter, r *http.Request) {
	var supply = getToken().Data.Supply
	var b4Decimal = supply[:len(supply)-6]
	var afterDecimal = supply[len(supply)-6:]
	// fmt.Fprintln(w, supply)
	fmt.Fprintln(w, b4Decimal+"."+afterDecimal)
}
