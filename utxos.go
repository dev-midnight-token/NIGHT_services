package function

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

const Asset_name = "NIGHT"
const Asset_name_hex = "4e49474854"
const Policy = "0691b2fecca1ac4f53cb6dfb00b7013e561d1f34403b957cbb5af1fa"
const Fingerprint = "asset1wd3llgkhsw6etxf2yca6cgk9ssrpva3wf0pq9a"

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

// Returns the circulating supply of the token as a float with 6 decimal places
func circulatingSupply(w http.ResponseWriter, r *http.Request) {
	/*	var tokenRows = getAddress().Rows

		//	T_LF_R Treasury, Lost and Found, Reserve
		var T_LF_R TokenInfo
		for _, row := range tokenRows {
			if row.Asset_name == Asset_name && row.Policy == Policy &&
				row.Asset_name_hex == Asset_name_hex && row.Fingerprint == Fingerprint {
				T_LF_R = row
				break
			}
		}

		// Convert supply and quantity from string to int
		supply, err := strconv.Atoi(T_LF_R.Supply)
		if err != nil {
			log.Fatalf("Error converting supply to int: %v", err)
		}

		uncirculatingSupply, err := strconv.Atoi(T_LF_R.Quantity)
		if err != nil {
			log.Fatalf("Error converting circulating supply to int: %v", err)
		}
		circulatingSupply := float64(supply-uncirculatingSupply) / math.Pow10(6)
	*/
	circulatingSupply := float64(16607399400.634954)

	fmt.Fprintf(w, "%.6f\n", circulatingSupply)
}
