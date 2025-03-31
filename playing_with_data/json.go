package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// 1. Encoding (Marshaling) Go struct to JSON
	user := User{ID: 1, Name: "John Doe", Email: "john@example.com"}
	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	fmt.Println("Encoded JSON:", string(jsonData))

	// 2. Decoding (Unmarshaling) JSON to Go struct
	jsonString := `{"id":2, "name":"Jane Doe", "email":"jane@example.com"}`
	var user2 User
	if err := json.Unmarshal([]byte(jsonString), &user2); err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}
	fmt.Println("Decoded Struct:", user2)

	// 3. Working with JSON as a map (Dynamic JSON Parsing)
	jsonMap := make(map[string]interface{})
	if err := json.Unmarshal([]byte(jsonString), &jsonMap); err != nil {
		fmt.Println("Error unmarshaling into map:", err)
		return
	}
	fmt.Println("JSON as Map:", jsonMap)

	// 4. Modifying JSON Data in a map
	jsonMap["name"] = "Jane Smith"
	updatedJSON, err := json.Marshal(jsonMap)
	if err != nil {
		fmt.Println("Error marshaling updated JSON:", err)
		return
	}
	fmt.Println("Updated JSON:", string(updatedJSON))

	// 5. Pretty Print JSON
	prettyJSON, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling pretty JSON:", err)
		return
	}
	fmt.Println("Pretty JSON:")
	fmt.Println(string(prettyJSON))
}
