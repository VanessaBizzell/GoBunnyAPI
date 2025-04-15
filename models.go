package main

type Bunny struct {
	ID              int    `json:"id"`              // Unique identifier for the bunny
	Name            string `json:"name"`            // Name of the bunny	//
	Breed           string `json:"breed"`           // Breed of the bunny
	Age             int    `json:"age"`             // Age of the bunny in years
	Description     string `json:"description"`     // Description of the bunny
	Characteristics string `json:"characteristics"` // Characteristics of the bunny
}

// uses slice of pointers to Bunny struct. This allows for more efficient memory usage and easier manipulation of the data.
// The slice is initialized with some sample data for demonstration purposes.
var bunnies = []*Bunny{
	{ID: 1, Name: "Chiggles", Breed: "Lionhead", Age: 3, Description: "A fluffy lionhead bunny. White with ginger spots.", Characteristics: "Fluffy, friendly, eater of cables."},
	{ID: 2, Name: "Jasmine", Breed: "New Zealand White", Age: 2, Description: "A white rabbit with red eyes", Characteristics: "big ears, loves to hop around. Demands attention."},
	{ID: 3, Name: "Thumper", Breed: "Holland Lop", Age: 4, Description: "A small bunny with floppy ears.", Characteristics: "Loves to dig and chew."},
}
