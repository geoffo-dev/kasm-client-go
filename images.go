package kasm

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetImages - Returns list of coffees
func (c *Client) GetImages(authToken *string) ([]Image, error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/public/get_images", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	images := []Image{}
	err = json.Unmarshal(body, &images)
	if err != nil {
		return nil, err
	}

	return images, nil
}

// // CreateCoffee - Create new coffee
// func (c *Client) CreateCoffee(coffee Coffee, authToken *string) (*Coffee, error) {
// 	rb, err := json.Marshal(coffee)
// 	if err != nil {
// 		return nil, err
// 	}

// 	req, err := http.NewRequest("POST", fmt.Sprintf("%s/coffees", c.HostURL), strings.NewReader(string(rb)))
// 	if err != nil {
// 		return nil, err
// 	}

// 	body, err := c.doRequest(req, authToken)
// 	if err != nil {
// 		return nil, err
// 	}

// 	newImage := Image{}
// 	err = json.Unmarshal(body, &newImage)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &newImage, nil
// }

// // GetImage - Returns specific Image (no auth required)
// func (c *Client) GetImage(coffeeID string) ([]Image, error) {
// 	req, err := http.NewRequest("GET", fmt.Sprintf("%s/coffees/%s", c.HostURL, coffeeID), nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	body, err := c.doRequest(req, nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	coffees := []Image{}
// 	err = json.Unmarshal(body, &coffees)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return coffees, nil
// }

// // GetImageIngredients - Returns list of coffee ingredients (no auth required)
// func (c *Client) GetImageIngredients(coffeeID string) ([]Ingredient, error) {
// 	req, err := http.NewRequest("GET", fmt.Sprintf("%s/coffees/%s/ingredients", c.HostURL, coffeeID), nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	body, err := c.doRequest(req, nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	ingredients := []Ingredient{}
// 	err = json.Unmarshal(body, &ingredients)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return ingredients, nil
// }

// // CreateCoffeeIngredient - Create new coffee ingredient
// func (c *Client) CreateCoffeeIngredient(coffee Coffee, ingredient Ingredient, authToken *string) (*Ingredient, error) {
// 	reqBody := struct {
// 		CoffeeID     int    `json:"coffee_id"`
// 		IngredientID int    `json:"ingredient_id"`
// 		Quantity     int    `json:"quantity"`
// 		Unit         string `json:"unit"`
// 	}{
// 		CoffeeID:     coffee.ID,
// 		IngredientID: ingredient.ID,
// 		Quantity:     ingredient.Quantity,
// 		Unit:         ingredient.Unit,
// 	}
// 	rb, err := json.Marshal(reqBody)
// 	if err != nil {
// 		return nil, err
// 	}

// 	req, err := http.NewRequest("POST", fmt.Sprintf("%s/coffees/%d/ingredients", c.HostURL, coffee.ID), strings.NewReader(string(rb)))
// 	if err != nil {
// 		return nil, err
// 	}

// 	body, err := c.doRequest(req, authToken)
// 	if err != nil {
// 		return nil, err
// 	}

// 	newIngredient := Ingredient{}
// 	err = json.Unmarshal(body, &newIngredient)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &newIngredient, nil
// }
