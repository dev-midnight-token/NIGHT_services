package function

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

type AddressData struct {
	Address string `json:"address"`
}

type TokenInfo struct {
	Policy         string `json:"policy"`
	Asset_name     string `json:"asset_name"`
	Asset_name_hex string `json:"asset_name_hex"`
	Fingerprint    string `json:"fingerprint"`
	Name           string `json:"name"`
	Ticker         string `json:"ticker"`
	Verified       int    `json:"verified"`
	Decimals       int    `json:"decimals"`
	Quantity       string `json:"quantity"`
	Supply         string `json:"supply"`
}

type Rows struct {
	Tokens []TokenInfo
}

// Define a struct to represent the JSON data structure
type Address struct {
	Data AddressData `json:"data"`
	Rows []TokenInfo `json:"rows"`
	Code int         `json:"code"`
}

func getAddress() Address {
	resp, err := http.Get("https://api.adastat.net/rest/v1/addresses/addr1w9m77uujyd6069nfdly22ejtslcdkf8w6z3f4603pzf42tqt6a398.json?rows=tokens&dir=desc") // Replace with your actual API endpoint
	if err != nil {
		log.Fatalf("Error making HTTP request: %v", err)
	}
	defer resp.Body.Close() // Ensure the response body is closed

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Unexpected status code: %d", resp.StatusCode)
	}

	// ... (decode JSON in the next step)
	var address Address
	err = json.NewDecoder(resp.Body).Decode(&address)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}

	fmt.Printf("Decoded Address: %+v\n", address)
	return address
}

func init() {
	functions.HTTP("CirculatingSupply", circulatingSupply)
}

// helloWorld writes "Hello, World!" to the HTTP response.
func circulatingSupply(w http.ResponseWriter, r *http.Request) {
	var rowNIGHT = getAddress().Rows[0]
	supply, err := strconv.Atoi(rowNIGHT.Supply)
	if err != nil {
		log.Fatalf("Error converting supply to int: %v", err)
	}

	uncirculatingSupply, err := strconv.Atoi(rowNIGHT.Quantity)
	if err != nil {
		log.Fatalf("Error converting circulating supply to int: %v", err)
	}
	circulatingSupply := float64(supply-uncirculatingSupply) / math.Pow10(6)

	fmt.Fprintf(w, "%.6f\n", circulatingSupply)
}
