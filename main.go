package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/waynefoux/asset-manager/asset"
)

type Asset struct {
	Item  string `json:"item"`
	ID   int    `json:"id"`
	Description string `json:"description"`
}

var db asset.AssetDB = asset.AssetDB{}

func main() {

	// Add assets to the database
	db.AddAsset(asset.Asset{ID: 1, Item: "IMSAI 8080", Description: "Vintage Computer of the 1970s"})
	db.AddAsset(asset.Asset{ID: 2, Item: "Apple II", Description: "Vintage Computer of the 1970s"})

	http.HandleFunc("/", assetHandler)

	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func assetHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		// Handler for GET /users
		//fmt.Fprintf(w, "GET /users")
	} else if r.Method == http.MethodPost {
		// Handler for POST /users
		//fmt.Fprintf(w, "POST /users")
	} else if r.Method == http.MethodPut {
		// Handler for PUT /users
		//fmt.Fprintf(w, "PUT /users")
	} else {
		// Method not supported
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	// Set the response content type to text/plain
	w.Header().Set("Content-Type", "json")

	assets := db.GetAssetList()

	jsonData, err := json.Marshal(assets)
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write(jsonData)
}
