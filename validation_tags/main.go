package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type FullExample struct {
    // Basic string validations
    Name           string   `validate:"required,min=2,max=30"`       // required, 2-30 chars
    Email          string   `validate:"required,email"`             // required, valid email
    Username       string   `validate:"required,alphanum"`          // letters & numbers only
    Phone          string   `validate:"required,regexp=^\\+88[0-9]{11}$"` // Bangladesh phone format
    Nickname       string   `validate:"omitempty,min=2,max=10"`     // optional, 2-10 chars if set

    // Number validations
    Age            int      `validate:"gte=18,lte=100"`              // 18-100
    Price          float64  `validate:"required,gt=0"`               // must exist & >0
    Discount       int      `validate:"lt=50"`                        // <50
    Quantity       int      `validate:"eq=10"`                        // must equal 10

    // Enum / specific values
    Color          string   `validate:"oneof=red green blue"`         // must be one of these
    Size           string   `validate:"oneof=S M L XL"`               // must match S/M/L/XL

    // Conditional validations
    Password       string   `validate:"required,min=8"`              // required, min 8 chars
    ConfirmPassword string  `validate:"required_with=Password,eqfield=Password"` // matches Password if set

    // UUID / URL
    ID             string   `validate:"required,uuid"`               // valid UUID
    Website        string   `validate:"required,url"`                // valid URL

    // Slice / array
    Tags           []string `validate:"required,dive,required,min=2"` // each element required, min 2 chars
}

func main() {
    validate := validator.New()

    example := FullExample{
        Name:            "John",
        Email:           "john@example.com",
        Username:        "john123",
        Phone:           "+8801712345678",
        Nickname:        "JD",
        Age:             25,
        Price:           100.5,
        Discount:        20,
        Quantity:        10,
        Color:           "red",
        Size:            "M",
        Password:        "mypassword",
        ConfirmPassword: "mypassword",
        ID:              "550e8400-e29b-41d4-a716-446655440000",
        Website:         "https://example.com",
        Tags:            []string{"go", "api"},
    }

    err := validate.Struct(example)
    if err != nil {
        fmt.Println("Validation errors:", err)
    } else {
        fmt.Println("All validations passed!")
    }
}
