package main

import (
	"01-REST-diag/handler"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const LINEBREAK = "\n"

var StartTime = time.Now()

/*
Diagnostic handler to showcases access to request content (headers, body, method, parameters, etc.)
*/
func diagHandler(w http.ResponseWriter, r *http.Request) {

	// Prepares return info with general request information
	output := "Request Information:" + LINEBREAK
	output += "--------------------" + LINEBREAK
	output += "URL (Path): " + r.URL.Path + LINEBREAK
	output += "Path value: " + r.PathValue("p1") + LINEBREAK
	output += "Method: " + r.Method + LINEBREAK
	output += "URL Parameters: " + LINEBREAK

	// Decompose parameters only if present
	if len(r.URL.RawQuery) != 0 {
		output += " Raw parameter content: " + r.URL.RawQuery + LINEBREAK

		output += "Decomposed parameters: " + LINEBREAK
		// Print individual parameters (here, they are decomposed as key-value pairs)
		for parKey, parVal := range r.URL.Query() {
			output += "- " + parKey + "=" + strings.Join(parVal, ",") + LINEBREAK
		}
	}

	// Append http headers received
	output += LINEBREAK + "Headers:" + LINEBREAK
	for key, values := range r.Header {
		// Print all values for a given key
		for _, value := range values {
			output += key + ": " + value + LINEBREAK
		}
	}

	// Append body of request (where existing)
	content, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error parsing request body.", http.StatusInternalServerError)
	}

	output += LINEBREAK + "Body:" + LINEBREAK
	output += string(content)

	// For all options for Printf see https://yourbasic.org/golang/fmt-printf-reference-cheat-sheet/
	_, err = fmt.Fprintf(w, "%v", output)
	if err != nil {
		http.Error(w, "Error when returning output", http.StatusInternalServerError)
	}

}

func main() {
	handler.StartTime = StartTime
	// Extract PORT variable from the environment variables
	port := os.Getenv("PORT")

	// Override port with default port if not provided (e.g. local deployment)
	if port == "" {
		log.Println("$PORT has not been set. Default: 8080")
		port = "8080"
	}

	// Instantiate the router
	router := http.NewServeMux()

	// specific simple path value pattern (single element)
	router.HandleFunc(handler.INFO_PATH, handler.GeneralInfo)
	router.HandleFunc(handler.POPULATION_PATH, handler.PopulationLevel)
	router.HandleFunc(handler.STATUS_PATH, handler.Overview)

	// Start HTTP server
	log.Println("Starting server on port " + port + " ...")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
