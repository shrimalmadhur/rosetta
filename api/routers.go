/*
 * Rosetta
 *
 * <h2>Backstory</h2> Writing reliable blockchain integrations is complicated and time-consuming. The process requires careful analysis of the unique aspects of each blockchain and extensive communication with its developers to understand the best strategies to deploy nodes, recognize deposits, broadcast transactions, etc. Even a minor misunderstanding can lead to downtime, or even worse, incorrect fund attribution. Not to mention, this integration must be continuously modified and tested each time a blockchain team releases new software.  Instead of spending time working on their blockchain, project developers spend countless hours answering similar support questions for each team integrating their blockchain. With their questions answered, each integrating team then writes similar code to interface with the blockchain instead of spending their engineering resources adding support for more blockchain projects or working on unique products and applications.  <h2>Standard for Blockchain Interaction</h2> Rosetta is a new project from Coinbase to standardize the process of deploying and interacting with blockchains. With an explicit specification to adhere to, all parties involved in blockchain development can spend less time figuring out how to integrate with each other and more time working on the novel advances that will push the blockchain ecosystem forward. In practice, this means that any blockchain project that implements the requirements outlined in this specification will enable exchanges, block explorers, and wallets to integrate with much less communication overhead and network-specific work.  <h5>© 2020 Coinbase</h5>
 *
 * API version: 1.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

// A Route defines the parameters for an api endpoint
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes are a collection of defined api endpoints
type Routes []Route

// Router defines the required methods for retrieving api routes
type Router interface {
	Routes() Routes
}

// NewRouter creates a new router for any number of api routers
func NewRouter(routers ...Router) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, api := range routers {
		for _, route := range api.Routes() {
			var handler http.Handler
			handler = route.HandlerFunc

			router.
				Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				Handler(handler)
		}
	}

	return router
}

// EncodeJSONResponse uses the json encoder to write an interface to the http response with an optional status code
func EncodeJSONResponse(i interface{}, status *int, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if status != nil {
		w.WriteHeader(*status)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	return json.NewEncoder(w).Encode(i)
}

// ReadFormFileToTempFile reads file data from a request form and writes it to a temporary file
func ReadFormFileToTempFile(r *http.Request, key string) (*os.File, error) {
	r.ParseForm()
	formFile, _, err := r.FormFile(key)
	if err != nil {
		return nil, err
	}

	defer formFile.Close()
	file, err := ioutil.TempFile("tmp", key)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	fileBytes, err := ioutil.ReadAll(formFile)
	if err != nil {
		return nil, err
	}

	file.Write(fileBytes)
	return file, nil
}

// parseIntParameter parses a sting parameter to an int64
func parseIntParameter(param string) (int64, error) {
	return strconv.ParseInt(param, 10, 64)
}
