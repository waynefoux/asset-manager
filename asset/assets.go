package asset

type Asset struct {
	Item  string `json:"item"`
	ID   int    `json:"id"`
	Description string `json:"description"`
}

type AssetDB struct {
	Assets []Asset
}

func (db *AssetDB) AddAsset(asset Asset) {
	db.Assets = append(db.Assets, asset)
}

func (db *AssetDB) DeleteAsset(asset Asset) {
	for i, a := range db.Assets {
		if a.ID == asset.ID {
			db.Assets = append(db.Assets[:i], db.Assets[i+1:]...)
			break
		}
	}
}

func (db *AssetDB) GetAssetList() []Asset {
	return db.Assets
}


// assetList := []Asset
// }













// addAsset := func(w http.ResponseWriter, r *http.Request) {
// 	// Set the response content type to text/plain
// 	w.Header().Set("Content-Type", "json")

// 	jsonData, err := json.Marshal(assetList)
// 	if err != nil {
// 		log.Fatal(err)
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}

// 	w.Write(jsonData)
// }

// deleteAsset := func(w http.ResponseWriter, r *http.Request) {	
// 	// Set the response content type to text/plain
// 	w.Header().Set("Content-Type", "json")

// 	jsonData, err := json.Marshal(assetList)
// 	if err != nil {
// 		log.Fatal(err)
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}

// 	w.Write(jsonData)
// }




// getAssetList := func(w http.ResponseWriter, r *http.Request) {
// 	// Set the response content type to text/plain
// 	w.Header().Set("Content-Type", "json")

// 	jsonData, err := json.Marshal(assetList)
// 	if err != nil {
// 		log.Fatal(err)
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}

// 	w.Write(jsonData)
// }

// setAsset := func(w http.ResponseWriter, r *http.Request) {
// 	// Set the response content type to text/plain
// 	w.Header().Set("Content-Type", "json")

// 	jsonData, err := json.Marshal(assetList)
// 	if err != nil {
// 		log.Fatal(err)
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}

// 	w.Write(jsonData)
// }