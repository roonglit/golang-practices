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

type Person struct {
	Name  string `validate:"required"`
	Age   int    `validate:"gte=18"`
	Email string `validate:"required,email"`
}

func main() {
	user := User{
		ID:    1,
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}

	fmt.Println("user", user) // Output: user {1 John Doe john.doe@example.com}
	jsonData, _ := json.Marshal(user)
	fmt.Println(string(jsonData)) // Output: {"id":1,"name":"Alice"}

	// Example of struct tags for validation
	// validate := validator.New()

	// person := Person{
	// 	Name:  "",
	// 	Age:   16,
	// 	Email: "invalid-email",
	// }

	// err := validate.Struct(person)
	// if err != nil {
	// 	fmt.Println("Validation errors:")
	// 	for _, err := range err.(validator.ValidationErrors) {
	// 		fmt.Printf("- Field '%s': %s\n", err.Field(), err.Tag())
	// 	}
	// }
}
