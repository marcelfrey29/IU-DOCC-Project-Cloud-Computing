package main

// A Travel Guide is a collection of activities and attractions at a specific location.
type TravelGuide struct {
	Id          string   `json:"id" dynamodbav:"id"`
	Name        string   `json:"name" dynamodbav:"name"`
	Private     bool     `json:"isPrivate" dynamodbav:"private"`
	Secret      string   `json:"secret" dynamodbav:"secret"`
	Description string   `json:"description" dynamodbav:"description"`
	Location    Location `json:"location" dynamodbav:"location"`
	Category    Category `json:"category" dynamodbav:"category"`
}

// A Category is represented by a number (ID).
type Category int

const (
	TG_CATEGORY_CULTURE Category = iota
	TG_CATEGORY_ACTION  Category = iota
	TG_CATEGORY_RELAX   Category = iota
)

// Location represents a geographic location where an attraction is located.
type Location struct {
	Street  string `json:"street" dynamodbav:"street"`
	Zip     string `json:"zip" dynamodbav:"zip"`
	City    string `json:"city" dynamodbav:"city"`
	State   string `json:"state" dynamodbav:"state"`
	Country string `json:"country" dynamodbav:"country"`
}
