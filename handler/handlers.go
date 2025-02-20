package handler

import (
	"bytes"
	"reflect"

	//"reflect"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	//"time"
)

/*
 * Convert request to body
 */
func ReqToBody(r *http.Request) []byte {
	// Create a new HTTP client
	client := &http.Client{}

	// Send the request using the client
	res, err := client.Do(r)
	if err != nil {
		fmt.Println("error sending request: %s\n", err)
		os.Exit(1)
	}

	defer res.Body.Close() // Closing response body after reading
	// Read response body
	bodyR, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	return bodyR
}

/*
* General country info
 */
func GeneralInfo(w http.ResponseWriter, r *http.Request) {
	// Construct the request URL
	RESTURL := "http://129.241.150.113:8080/v3.1/alpha/" + r.PathValue("two_letter_country_code")
	reqR, err := http.NewRequest(http.MethodGet, RESTURL, nil)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	bodyR := ReqToBody(reqR)

	var dataR []interface{}
	err4 := json.Unmarshal(bodyR, &dataR)
	if err4 != nil {
		fmt.Println(err4)
	}

	// decoded data
	mappR := dataR[0]

	// reference{} to map
	mR, err2 := mappR.(map[string]interface{})
	if !err2 {
		fmt.Println("want type map[string]interface{};  got %T", dataR)
	}

	CCURL := "http://129.241.150.113:3500/api/v0.1/countries/cities"
	country := mR["name"].(map[string]interface{})["common"].(string) // Use the country name from the first REST response
	bodyPost := map[string]string{
		"country": country,
	}
	jsonStr, err := json.Marshal(bodyPost)
	if err != nil {
		log.Println("Error marshaling cities request: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	reqC, err := http.NewRequest("POST", CCURL, bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println("error making http request: %s\n", err)
		os.Exit(1)
	}
	reqC.Header.Set("Content-Type", "application/json")

	bodyC := ReqToBody(reqC)

	var dataC interface{}
	err5 := json.Unmarshal(bodyC, &dataC)
	if err5 != nil {
		fmt.Println(err5)
	}

	mappC := dataC
	mC, err3 := mappC.(map[string]interface{})
	if !err3 {
		fmt.Errorf("want type map[string]interface{};  got %T", dataC)
	}

	// defining a struct instance
	var info Info

	// From m interface to info struct
	info.Name = country
	//_______________________
	continents := mR["continents"].([]interface{})
	// Convert the []interface{} to []string
	info.Continents = make([]string, len(continents))
	for i, v := range continents {
		continent := v.(string)
		info.Continents[i] = continent
	}
	//_______________________
	info.Population = int(mR["population"].(float64))
	//_______________________
	languages := mR["languages"].(map[string]interface{})
	// Convert the languages map to map[string]string
	info.Languages = make(map[string]string)
	for key, value := range languages {
		strValue, ok := value.(string)
		if !ok {
			log.Printf("Skipping non-string value in languages map for key %s", key)
			continue
		}
		info.Languages[key] = strValue
	}
	//_______________________
	borders := mR["borders"].([]interface{})
	info.Borders = make([]string, len(borders))
	for i, v := range borders {
		borders := v.(string)
		info.Borders[i] = borders
	}
	//_______________________
	info.Flag = mR["flag"].(string)
	//_______________________
	info.Capital = mR["capital"].([]interface{})[0].(string)
	//_______________________
	limit := 10 // Set limit
	if len(r.URL.RawQuery) != 0 {
		l := r.URL.Query()["limit"] // Extract limit url
		if len(l) != 0 {            // If limit is found
			limit, err = strconv.Atoi(l[0])
			if limit > 100 {
				limit = 100
			}
			if limit < 1 { // Checks for negative numbers or non-numbers
				limit = 1
			}
		}
	}

	cities := mC["data"].([]interface{})
	info.Cities = make([]string, limit)
	for i := 0; i < limit; i++ {
		cities := cities[i].(string)
		info.Cities[i] = cities
	}

	infoJ, _ := json.MarshalIndent(info, "", "    ")
	output := string(infoJ)

	_, err6 := fmt.Fprintf(w, "%v", output)
	if err6 != nil {
		http.Error(w, "Error when returning output", http.StatusInternalServerError)
	}
}

/*
*
Handler matching on advanced structure with two elements.
*/
func PopulationLevel(w http.ResponseWriter, r *http.Request) {
	// Construct the request URL
	RESTURL := "http://129.241.150.113:8080/v3.1/alpha/" + r.PathValue("two_letter_country_code")
	reqR, err := http.NewRequest(http.MethodGet, RESTURL, nil)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	bodyR := ReqToBody(reqR)

	var dataR []interface{}
	err4 := json.Unmarshal(bodyR, &dataR)
	if err4 != nil {
		fmt.Println(err4)
	}

	// decoded data
	mappR := dataR[0]

	// reference{} to map
	mR, err2 := mappR.(map[string]interface{})
	if !err2 {
		fmt.Errorf("want type map[string]interface{};  got %T", dataR)
	}

	CCURL := "http://129.241.150.113:3500/api/v0.1/countries/population"
	country := mR["name"].(map[string]interface{})["common"].(string) // Use the country name from the first REST response
	bodyPost := map[string]string{
		"country": country,
	}
	jsonStr, err := json.Marshal(bodyPost)
	if err != nil {
		log.Printf("Error marshaling cities request: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	reqC, err := http.NewRequest("POST", CCURL, bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}
	reqC.Header.Set("Content-Type", "application/json")

	bodyC := ReqToBody(reqC)

	var dataC interface{}
	err5 := json.Unmarshal(bodyC, &dataC)
	if err5 != nil {
		fmt.Println(err5)
	}

	mappC := dataC
	mC, err3 := mappC.(map[string]interface{})
	if !err3 {
		fmt.Errorf("want type map[string]interface{};  got %T", dataC)
	}

	// defining a struct instance
	var pop Pop

	// From m interface to Pop struct
	pop.Mean = 0 // Initiate mean
	dataPop := mC["data"].(map[string]interface{})
	count := dataPop["populationCounts"].([]interface{})

	years := [2]int{1900, 2100}
	if len(r.URL.RawQuery) != 0 {
		l := r.URL.Query()["limit"] // Extract limit url
		if len(l) != 0 {            // If limit is found
			if len(l) == 9 {
				fmt.Println(reflect.TypeOf(l[0]))
				//years = l[0]
				//fmt.Println(years)
				/*
					if limit > 100 {
						limit = 100
					}
					if limit < 1 { // Checks for negative numbers or non-numbers
						limit = 1
					}*/
			}
		}
	}

	startYear := years[0]
	endYear := years[1]

	amount := endYear - startYear + 1 // Amount of years to be listed
	not := 0                          // Values that aren't within the range
	pop.Values = make([]map[string]interface{}, amount)
	var pops map[string]interface{}
	for i, v := range count {
		popus := v.(map[string]interface{})
		if popus["year"].(float64) >= 2010 && popus["year"].(float64) <= 2015 {
			// Convert the languages map to map[string]string
			pops = make(map[string]interface{})
			for key, value := range popus {
				floatValue, ok := value.(float64)
				if !ok {
					log.Printf("Skipping non-float value in languages map for key %s", key)
					continue
				}
				pop.Mean += int(floatValue) // Sum populations
				pops[key] = floatValue
			}
			pop.Values[i-not] = pops
		} else {
			not++ // +1 not in range
		}
	}
	//_______________________
	pop.Mean = pop.Mean / amount

	popJ, _ := json.MarshalIndent(pop, "", "    ")
	output := string(popJ)

	_, err6 := fmt.Fprintf(w, "%v", output)
	if err6 != nil {
		http.Error(w, "Error when returning output", http.StatusInternalServerError)
	}
}

/*
* Overview of api
 */
func Overview(w http.ResponseWriter, r *http.Request) {
	// Construct the request URL
	RESTURL := "http://129.241.150.113:8080/v3.1/independent?status=true" + r.PathValue("two_letter_country_code")
	reqR, err := http.NewRequest(http.MethodGet, RESTURL, nil)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	CCURL := "http://129.241.150.113:3500/api/v0.1/countries/population/cities"
	reqC, err := http.NewRequest(http.MethodGet, CCURL, nil)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	// Create a new HTTP client
	client := &http.Client{}

	// Send the request using the client
	resR, err := client.Do(reqR)
	if err != nil {
		fmt.Println("error sending request: %s\n", err)
		os.Exit(1)
	}
	resC, err := client.Do(reqC)
	if err != nil {
		fmt.Println("error sending request: %s\n", err)
		os.Exit(1)
	}

	// defining a struct instance
	var status Status

	// From m interface to Pop struct
	status.Countriesnowapi = resC.StatusCode
	status.Testcountriesapi = resR.StatusCode
	status.Version = "v1"
	status.Uptime = 1

	statusJ, _ := json.MarshalIndent(status, "", "    ")

	output := string(statusJ)

	_, err6 := fmt.Fprintf(w, "%v", output)
	if err6 != nil {
		http.Error(w, "Error when returning output", http.StatusInternalServerError)
	}
}
