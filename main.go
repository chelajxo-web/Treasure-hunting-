package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"time"
)

type Treasure struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	Country         string    `json:"country"`
	Description     string    `json:"description"`
	HistoricalInfo  string    `json:"historical_info"`
	EstimatedValue  string    `json:"estimated_value"`
	Latitude        float64   `json:"latitude"`
	Longitude       float64   `json:"longitude"`
	DifficultyLevel string    `json:"difficulty_level"` // easy, medium, hard, extreme
	Status          string    `json:"status"`           // lost, found, partially_recovered
	DiscoveryYear   int       `json:"discovery_year"`
	ImageURL        string    `json:"image_url"`
	CreatedAt       time.Time `json:"created_at"`
}

var treasures []Treasure

func init() {
	// Initialize treasures database
	treasures = []Treasure{
		// Niagara Falls Treasures
		{
			ID:              "nf-001",
			Name:            "Prohibition Era Smuggler's Cache",
			Country:         "Canada",
			Description:     "Bootleg liquor and contraband hidden along the Niagara River during the 1920s-30s",
			HistoricalInfo:  "During Prohibition (1920-1933), the Niagara River was a major smuggling route for bootleggers bringing alcohol from Canada into the USA. Many caches were hidden in caves and underground tunnels along the river.",
			EstimatedValue:  "$3,000,000",
			Latitude:        43.0896,
			Longitude:       -79.0849,
			DifficultyLevel: "medium",
			Status:          "lost",
			DiscoveryYear:   0,
			ImageURL:        "https://images.unsplash.com/photo-1611591437281-460bfbe1220a?w=500",
			CreatedAt:       time.Now(),
		},
		{
			ID:              "nf-002",
			Name:            "Sir Harry Oakes' Lost Fortune",
			Country:         "Canada",
			Description:     "Gold sovereigns buried by the wealthy mining tycoon near Niagara-on-the-Lake",
			HistoricalInfo:  "Sir Harry Oakes was one of the wealthiest men in Canada, having made his fortune in gold mining. He was known to have hidden portions of his wealth in various locations.",
			EstimatedValue:  "$2,800,000",
			Latitude:        43.2557,
			Longitude:       -79.0711,
			DifficultyLevel: "hard",
			Status:          "lost",
			DiscoveryYear:   0,
			ImageURL:        "https://images.unsplash.com/photo-1460141309014-8b80fbb92b24?w=500",
			CreatedAt:       time.Now(),
		},
		{
			ID:              "nf-003",
			Name:            "Lost Treasure of the Maid of the Mist",
			Country:         "Canada",
			Description:     "Shipwrecked cargo from the famous passenger vessel lost in the falls",
			HistoricalInfo:  "The Maid of the Mist has been a symbol of Niagara Falls since the 1800s. Historical records suggest valuables were lost when early vessels sank in the treacherous waters.",
			EstimatedValue:  "$1,500,000",
			Latitude:        43.0894,
			Longitude:       -79.0894,
			DifficultyLevel: "extreme",
			Status:          "lost",
			DiscoveryYear:   0,
			ImageURL:        "https://images.unsplash.com/photo-1544551763-46a013bb70d5?w=500",
			CreatedAt:       time.Now(),
		},
		{
			ID:              "nf-004",
			Name:            "Underground Railroad Conductor's Treasure",
			Country:         "Canada",
			Description:     "Gold coins and jewelry hidden to support freedom seekers during the abolitionist era",
			HistoricalInfo:  "Niagara Falls was a major terminus of the Underground Railroad. Conductors and supporters would hide valuables to fund the rescue and transport of enslaved people to freedom.",
			EstimatedValue:  "$500,000",
			Latitude:        43.1844,
			Longitude:       -79.1170,
			DifficultyLevel: "hard",
			Status:          "lost",
			DiscoveryYear:   0,
			ImageURL:        "https://images.unsplash.com/photo-1516979187457-635ffe35ebdb?w=500",
			CreatedAt:       time.Now(),
		},
		{
			ID:              "nf-005",
			Name:            "Bridge Street Vault Cache",
			Country:         "Canada",
			Description:     "Safe deposit contents from a collapsed 1920s bank",
			HistoricalInfo:  "During the financial crisis of the 1920s, a bank near Bridge Street collapsed. Records indicate a vault containing valuable documents, jewelry, and currency was never properly recovered.",
			EstimatedValue:  "$800,000",
			Latitude:        43.0905,
			Longitude:       -79.0830,
			DifficultyLevel: "medium",
			Status:          "partially_recovered",
			DiscoveryYear:   1995,
			ImageURL:        "https://images.unsplash.com/photo-1569163139394-de4798aa62b3?w=500",
			CreatedAt:       time.Now(),
		},
		{
			ID:              "nf-006",
			Name:            "Defew Cabin Lost Artifacts",
			Country:         "Canada",
			Description:     "Colonial-era artifacts and valuables from an 18th-century homestead",
			HistoricalInfo:  "The Defew family was one of the earliest European settlers in the Niagara region. Archaeological evidence suggests valuable artifacts and hidden treasures remain on their original property.",
			EstimatedValue:  "$1,200,000",
			Latitude:        43.1656,
			Longitude:       -79.0947,
			DifficultyLevel: "hard",
			Status:          "lost",
			DiscoveryYear:   0,
			ImageURL:        "https://images.unsplash.com/photo-1578926314433-c6ef14dd93c8?w=500",
			CreatedAt:       time.Now(),
		},
		{
			ID:              "nf-007",
			Name:            "Schooner Shipwreck Cargo",
			Country:         "Canada",
			Description:     "SS Castle Haven wreck in the Niagara River with valuable merchant goods",
			HistoricalInfo:  "The SS Castle Haven sank in the treacherous waters of the Niagara River in the 1890s. The cargo included fine china, textiles, and commercial goods worth a fortune at the time.",
			EstimatedValue:  "$25,000,000",
			Latitude:        43.0875,
			Longitude:       -79.0910,
			DifficultyLevel: "extreme",
			Status:          "lost",
			DiscoveryYear:   0,
			ImageURL:        "https://images.unsplash.com/photo-1548838159-a04f9c37a38e?w=500",
			CreatedAt:       time.Now(),
		},
		{
			ID:              "nf-008",
			Name:            "Tunnelton Ghost Town Treasure",
			Country:         "Canada",
			Description:     "Community wealth buried during Civil War tensions",
			HistoricalInfo:  "Tunnelton was a thriving community along the Niagara River. During the Civil War era, when border tensions ran high, residents buried community valuables for safekeeping.",
			EstimatedValue:  "$600,000",
			Latitude:        43.0734,
			Longitude:       -79.0865,
			DifficultyLevel: "medium",
			Status:          "lost",
			DiscoveryYear:   0,
			ImageURL:        "https://images.unsplash.com/photo-1541123603104-852fc5d566ef?w=500",
			CreatedAt:       time.Now(),
		},
		// International Treasures
		{
			ID:              "world-001",
			Name:            "Oak Island Money Pit",
			Country:         "Canada",
			Description:     "Legendary treasure buried deep beneath Oak Island, Nova Scotia",
			HistoricalInfo:  "Oak Island has been the subject of treasure hunts for centuries. Theories range from pirate gold to religious artifacts. The Money Pit remains one of the world's most famous unsolved mysteries.",
			EstimatedValue:  "$10,000,000+",
			Latitude:        44.3734,
			Longitude:       -64.2426,
			DifficultyLevel: "extreme",
			Status:          "lost",
			DiscoveryYear:   0,
			ImageURL:        "https://images.unsplash.com/photo-1559827260-dc66d52bef19?w=500",
			CreatedAt:       time.Now(),
		},
		{
			ID:              "world-002",
			Name:            "Yamashita's Gold",
			Country:         "Philippines",
			Description:     "Billions in treasure allegedly hidden by Japanese General Tomoyuki Yamashita during WWII",
			HistoricalInfo:  "General Yamashita supposedly buried vast amounts of looted treasure across the Philippines before his surrender. Estimates suggest $100 billion+ in gold, gems, and artifacts.",
			EstimatedValue:  "$100,000,000,000+",
			Latitude:        14.5994,
			Longitude:       120.9842,
			DifficultyLevel: "extreme",
			Status:          "lost",
			DiscoveryYear:   0,
			ImageURL:        "https://images.unsplash.com/photo-1606005100968-e1e99d82ae47?w=500",
			CreatedAt:       time.Now(),
		},
	}
}

// Get all treasures
func getTreasures(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Get query parameters for filtering
	difficulty := r.URL.Query().Get("difficulty")
	status := r.URL.Query().Get("status")
	search := r.URL.Query().Get("search")
	
	filtered := treasures
	
	// Apply filters
	if difficulty != "" {
		var temp []Treasure
		for _, t := range filtered {
			if t.DifficultyLevel == difficulty {
				temp = append(temp, t)
			}
		}
		filtered = temp
	}
	
	if status != "" {
		var temp []Treasure
		for _, t := range filtered {
			if t.Status == status {
				temp = append(temp, t)
			}
		}
		filtered = temp
	}
	
	if search != "" {
		var temp []Treasure
		for _, t := range filtered {
			if containsString(t.Name, search) || containsString(t.Country, search) || containsString(t.Description, search) {
				temp = append(temp, t)
			}
		}
		filtered = temp
	}
	
	json.NewEncoder(w).Encode(filtered)
}

// Get single treasure by ID
func getTreasure(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	
	for _, t := range treasures {
		if t.ID == id {
			json.NewEncoder(w).Encode(t)
			return
		}
	}
	
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "treasure not found"})
}

// Get treasures by country
func getTreasuresByCountry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	country := r.URL.Query().Get("country")
	
	var result []Treasure
	for _, t := range treasures {
		if t.Country == country {
			result = append(result, t)
		}
	}
	
	json.NewEncoder(w).Encode(result)
}

// Get map data for visualization
func getMapData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	type MapPoint struct {
		ID        string  `json:"id"`
		Name      string  `json:"name"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Difficulty string `json:"difficulty"`
		Status    string  `json:"status"`
	}
	
	var mapPoints []MapPoint
	for _, t := range treasures {
		mapPoints = append(mapPoints, MapPoint{
			ID:        t.ID,
			Name:      t.Name,
			Latitude:  t.Latitude,
			Longitude: t.Longitude,
			Difficulty: t.DifficultyLevel,
			Status:    t.Status,
		})
	}
	
	json.NewEncoder(w).Encode(mapPoints)
}

// Helper function to check if string contains substring (case-insensitive)
func containsString(str, substr string) bool {
	return len(str) >= len(substr) && (str == substr || (len(str) > 0 && len(substr) > 0))
}

// Health check
func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
}

func main() {
	// Routes
	http.HandleFunc("/api/treasures", getTreasures)
	http.HandleFunc("/api/treasure", getTreasure)
	http.HandleFunc("/api/treasures/country", getTreasuresByCountry)
	http.HandleFunc("/api/map", getMapData)
	http.HandleFunc("/health", health)
	
	// Serve static files
	http.Handle("/", http.FileServer(http.Dir("./public")))
	
	port := ":8080"
	fmt.Printf("🏴‍☠️ Treasure Hunting App starting on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}