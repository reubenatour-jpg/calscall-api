package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	apiKey  = "your-key-here"
	baseURL = "https://calscall.online"
)

type NutritionResponse struct {
	Found  bool              `json:"found"`
	Source *string           `json:"source"`
	P      *Product          `json:"p"`
	Serving *ServingValues   `json:"sv,omitempty"`
}

type Product struct {
	Barcode    string  `json:"b"`
	Name       string  `json:"n"`
	Calories   float64 `json:"c"`
	Carbs      float64 `json:"h"`
	Protein    float64 `json:"p"`
	Fat        float64 `json:"f"`
	ServingG   float64 `json:"s"`
	ImageURL   string  `json:"i"`
	Categories string  `json:"ct"`
	NutriScore string  `json:"ns"`
	NovaGroup  int     `json:"nv"`
	Allergens  string  `json:"al"`
}

type ServingValues struct {
	Calories float64 `json:"c"`
	Carbs    float64 `json:"h"`
	Protein  float64 `json:"p"`
	Fat      float64 `json:"f"`
}

func lookupBarcode(barcode string) (*NutritionResponse, error) {
	req, _ := http.NewRequest("GET", baseURL+"/v1/nutrition", nil)
	q := req.URL.Query()
	q.Add("barcode", barcode)
	req.URL.RawQuery = q.Encode()
	req.Header.Set("X-API-Key", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var data NutritionResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func main() {
	barcode := "5053990194769"
	if len(os.Args) > 1 {
		barcode = os.Args[1]
	}

	data, err := lookupBarcode(barcode)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if data.Found && data.P != nil {
		p := data.P
		fmt.Printf("✅ %s\n", p.Name)
		fmt.Printf("   Calories:  %.1f kcal/100g\n", p.Calories)
		fmt.Printf("   Protein:   %.1fg/100g\n", p.Protein)
		fmt.Printf("   Carbs:     %.1fg/100g\n", p.Carbs)
		fmt.Printf("   Fat:       %.1fg/100g\n", p.Fat)
		if data.Serving != nil {
			fmt.Printf("   Per serve: %.1f kcal\n", data.Serving.Calories)
		}
		if data.Source != nil {
			fmt.Printf("   Source:    %s\n", *data.Source)
		}
	} else {
		fmt.Printf("❌ Product not found for barcode %s\n", barcode)
	}
}
