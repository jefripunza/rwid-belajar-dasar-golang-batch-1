package main

import (
	"belajar-dasar-golang/model"
	"fmt"
	"strconv"
)

type Product struct {
	Name string `json:"name"`
	Price int `json:"price"`
}

type Merchant struct {
	Name string `json:"name"`
	Products []Product `json:"products"`
}

func main() {
	fmt.Println("Hello, World!")

	var a bool = true     // Boolean
	var b int = 5         // Integer
	var c float32 = 3.14  // Floating point number
	var d string = "Hi!"  // String

	b_string:= strconv.Itoa(b)
	fmt.Println(b_string)

	product_unggulan := Product{
		Name: "unggulan",
		Price: 123,
	}

	merchant := Merchant{
		Name: "sebelah",
		Products: []Product{product_unggulan,product_unggulan},
	}

	Tolong = "saya"

	fmt.Printf("Merchant: %+v \n", merchant)
  
	fmt.Println("Boolean: ", a)
	fmt.Println("Integer: ", b)
	fmt.Println("Float:   ", c)
	fmt.Println("String:  ", d)

	User := model.User{}

	User.CreateUser("rwid", "uyee")

}
