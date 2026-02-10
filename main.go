package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// Treasure represents a real-world treasure location
type Treasure struct {
	ID              int     `json:"id"`
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	Location        string  `json:"location"`
	Country         string  `json:"country"`
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	EstimatedValue  string  `json:"estimatedValue"`
	Difficulty      string  `json:"difficulty"`
	Status          string  `json:"status"`
	HistoricalInfo  string  `json:"historicalInfo"`
	DiscoveryDate   string  `json:"discoveryDate,omitempty"`
	ImageURL        string  `json:"imageURL"`
	Region          string  `json:"region"`
}

// TreasureDatabase holds all treasures
type TreasureDatabase struct {
	Treasures []Treasure
}

var db *TreasureDatabase

// Initialize treasures with Niagara Falls region data
func initTreasures() {
	db = &TreasureDatabase{
		Treasures: []Treasure{
			// World treasures
			{
				ID:          1,
				Name:        "Oak Island Money Pit",
				Description: "Legendary treasure buried on Oak Island, Nova Scotia",
				Location:    "Oak Island, Nova Scotia",
				Country:     "Canada",
				Latitude:    44.3842,
				Longitude:   -64.2843,
				EstimatedValue: "Estimated $10 million+",
				Difficulty: "extreme",
				Status:     "lost",
				HistoricalInfo: "The Money Pit has fascinated treasure hunters for over 200 years. Multiple shafts have been dug, revealing artifacts from various time periods.",
				ImageURL:   "https://images.unsplash.com/photo-1519915212116-7cfef71f0a4f",
				Region:     "Atlantic Canada",
			},
			{
				ID:          2,
				Name:        "Yamashita's Gold",
				Description: "Japanese General's treasure hidden in Southeast Asia during WWII",
				Location:    "Philippines & Southeast Asia",
				Country:     "Philippines",
				Latitude:    14.5995,
				Longitude:   120.9842,
				EstimatedValue: "$100+ billion",
				Difficulty: "hard",
				Status:     "partially recovered",
				HistoricalInfo: "General Tomoyuki Yamashita buried enormous amounts of gold and treasure throughout the Philippines before the end of WWII.",
				ImageURL:   "https://images.unsplash.com/photo-1535632066927-ab7c9ab60908",
				Region:     "Asia",
			},
			// Niagara Falls Region Treasures
			{
				ID:          10,
				Name:        "Prohibition Era Smuggler's Cache",
				Description: "Bootleg liquor and contraband hidden along the Niagara River",
				Location:    "Niagara River, near Lewiston",
				Country:     "Canada",
				Latitude:    43.1708,
				Longitude:   -79.0372,
				EstimatedValue: "$3 million",
				Difficulty: "medium",
				Status:     "lost",
				HistoricalInfo: "During the 1920s-30s Prohibition era, smugglers hid contraband and bootleg liquor along the Niagara River banks.",
				ImageURL:   "https://images.unsplash.com/photo-1516321720268-9af75e6a1427",
				Region:     "Niagara Falls, ON",
			},
			{
				ID:          11,
				Name:        "Sir Harry Oakes' Lost Fortune",
				Description: "Gold sovereigns buried by the wealthy mining tycoon",
				Location:    "Niagara-on-the-Lake Estate",
				Country:     "Canada",
				Latitude:    43.2556,
				Longitude:   -79.0629,
				EstimatedValue: "$2.8 million",
				Difficulty: "hard",
				Status:     "lost",
				HistoricalInfo: "Sir Harry Oakes, a wealthy mining magnate, is rumored to have buried gold coins on his estate in Niagara-on-the-Lake.",
				ImageURL:   "https://images.unsplash.com/photo-1515194247933-41c28a1ad136",
				Region:     "Niagara Falls, ON",
			},
			{
				ID:          12,
				Name:        "Maid of the Mist Shipwreck",
				Description: "Shipwrecked cargo from the famous passenger vessel",
				Location:    "Niagara Falls Canyon",
				Country:     "Canada",
				Latitude:    43.0896,
				Longitude:   -79.0849,
				EstimatedValue: "$1.5 million",
				Difficulty: "extreme",
				Status:     "lost",
				HistoricalInfo: "The Maid of the Mist was a historic passenger vessel. Lost cargo remains in the canyon depths.",
				ImageURL:   "https://images.unsplash.com/photo-1506905925346-21bda4d32df4",
				Region:     "Niagara Falls, ON",
			},
			{
				ID:          13,
				Name:        "Underground Railroad Conductor's Treasure",
				Description: "Gold coins and jewelry hidden to support freedom seekers",
				Location:    "Niagara-on-the-Lake Area",
				Country:     "Canada",
				Latitude:    43.2500,
				Longitude:   -79.0700,
				EstimatedValue: "$500,000+",
				Difficulty: "hard",
				Status:     "lost",
				HistoricalInfo: "Underground Railroad conductors hid gold and jewelry to fund the liberation of enslaved people.",
				ImageURL:   "https://images.unsplash.com/photo-1542571535-c46cdffc0eff",
				Region:     "Niagara Falls, ON",
			},
			{
				ID:          14,
				Name:        "Bridge Street Vault Cache",
				Description: "Safe deposit contents from a collapsed 1920s bank",
				Location:    "Bridge Street, Niagara Falls",
				Country:     "Canada",
				Latitude:    43.0861,
				Longitude:   -79.0844,
				EstimatedValue: "$800,000",
				Difficulty: "medium",
				Status:     "lost",
				HistoricalInfo: "A bank vault from the 1920s collapsed and its contents were never fully recovered.",
				ImageURL:   "https://images.unsplash.com/photo-1517604931442-7e0c6f169c02",
				Region:     "Niagara Falls, ON",
			},
			{
				ID:          15,
				Name:        "Defew Cabin Lost Artifacts",
				Description: "Colonial-era artifacts and valuables from a homestead",
				Location:    "Defewsville Area",
				Country:     "Canada",
				Latitude:    43.2400,
				Longitude:   -79.0800,
				EstimatedValue: "$1.2 million",
				Difficulty: "medium",
				Status:     "lost",
				HistoricalInfo: "The Defew family's 18th-century cabin contained valuable colonial artifacts and precious items.",
				ImageURL:   "https://images.unsplash.com/photo-1581092189519-609b473efabc",
				Region:     "Niagara Falls, ON",
			},
			{
				ID:          16,
				Name:        "SS Castle Haven Shipwreck",
				Description: "Schooner wreck in the Niagara River with merchant goods",
				Location:    "Niagara River",
				Country:     "Canada",
				Latitude:    43.2000,
				Longitude:   -79.0900,
				EstimatedValue: "$25 million",
				Difficulty: "extreme",
				Status:     "lost",
				HistoricalInfo: "The SS Castle Haven sank in the Niagara River carrying valuable merchant cargo and maritime treasures.",
				ImageURL:   "https://images.unsplash.com/photo-1500531546861-6a270d4b4c1f",
				Region:     "Niagara Falls, ON",
			},
			{
				ID:          17,
				Name:        "Tunnelton Ghost Town Treasure",
				Description: "Community wealth buried during Civil War tensions",
				Location:    "Tunnelton Historic Site",
				Country:     "Canada",
				Latitude:    43.2300,
				Longitude:   -79.1000,
				EstimatedValue: "$600,000",
				Difficulty: "hard",
				Status:     "lost",
				HistoricalInfo: "The residents of the ghost town Tunnelton buried their community's wealth for safekeeping during tumultuous times.",
				ImageURL:   "https://images.unsplash.com/photo-1518611505868-48510c2e00f7",
				Region:     "Niagara Falls, ON",
			},
		},
	}
}

// GetTreasures returns all treasures or filtered results
func GetTreasures(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	difficulty := r.URL.Query().Get("difficulty")
	status := r.URL.Query().Get("status")
	region := r.URL.Query().Get("region")
	search := strings.ToLower(r.URL.Query().Get("search"))

	var results []Treasure

	for _, treasure := range db.Treasures {
		match := true

		if difficulty != "" && treasure.Difficulty != difficulty {
			match = false
		}
		if status != "" && treasure.Status != status {
			match = false
		}
		if region != "" && treasure.Region != region {
			match = false
		}
		if search != "" && !strings.Contains(strings.ToLower(treasure.Name), search) &&
			!strings.Contains(strings.ToLower(treasure.Location), search) &&
			!strings.Contains(strings.ToLower(treasure.Country), search) {
			match = false
		}

		if match {
			results = append(results, treasure)
		}
	}

	json.NewEncoder(w).Encode(results)
}

// GetTreasure returns a single treasure by ID
func GetTreasure(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for _, treasure := range db.Treasures {
		if treasure.ID == id {
			json.NewEncoder(w).Encode(treasure)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Treasure not found"})
}

// GetStats returns statistics about treasures
func GetStats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	stats := map[string]interface{}{
		"total_treasures": len(db.Treasures),
		"by_difficulty": map[string]int{
			"easy":     len(filterByDifficulty("easy")),
			"medium":   len(filterByDifficulty("medium")),
			"hard":     len(filterByDifficulty("hard")),
			"extreme":  len(filterByDifficulty("extreme")),
		},
		"by_status": map[string]int{
			"lost":                len(filterByStatus("lost")),
			"found":               len(filterByStatus("found")),
			"partially_recovered": len(filterByStatus("partially recovered")),
		},
	}

	json.NewEncoder(w).Encode(stats)
}

func filterByDifficulty(difficulty string) []Treasure {
	var result []Treasure
	for _, t := range db.Treasures {
		if t.Difficulty == difficulty {
			result = append(result, t)
		}
	}
	return result
}

func filterByStatus(status string) []Treasure {
	var result []Treasure
	for _, t := range db.Treasures {
		if t.Status == status {
			result = append(result, t)
		}
	}
	return result
}

func main() {
	initTreasures()

	router := mux.NewRouter()

	// API endpoints
	router.HandleFunc("/api/treasures", GetTreasures).Methods("GET")
	router.HandleFunc("/api/treasures/{id}", GetTreasure).Methods("GET")
	router.HandleFunc("/api/stats", GetStats).Methods("GET")

	// Serve static files
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public")))

	fmt.Println("Treasure Hunting App Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}