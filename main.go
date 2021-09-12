package main

import (
	"bitbucket.org/meete/genesis-core/component/validation/goplayground"
	"fmt"
)

func main() {
	config := goplayground.Configuration{}
	validation := goplayground.New(config)

	//<nil>---PASSED
	fmt.Println("---------------------------------------")
	fmt.Println("TEST1")

	dataValidate := struct {
		Phone string `json:"phone" validate:"required,phone"`
	}{"0987654321"}

	err := validation.Struct(dataValidate)

	fmt.Println(err)

	//Key: 'Phone' Error:Field validation for 'Phone' failed on the 'required' tag---PASSED
	fmt.Println("\n\n---------------------------------------")
	fmt.Println("TEST2")

	dataValidate = struct {
		Phone string `json:"phone" validate:"required,phone"`
	}{"9876543210"}

	err = validation.Struct(dataValidate)

	fmt.Println(err)

	//Key: 'Phone' Error:Field validation for 'Phone' failed on the 'required' tag---PASSED
	fmt.Println("\n\n---------------------------------------")
	fmt.Println("TEST3")

	dataValidate = struct {
		Phone string `json:"phone" validate:"required,phone"`
	}{"098765432199"}

	err = validation.Struct(dataValidate)

	fmt.Println(err)

	//Key: 'Phone' Error:Field validation for 'Phone' failed on the 'required' tag---PASSED
	fmt.Println("\n\n---------------------------------------")
	fmt.Println("TEST4")

	dataValidate = struct {
		Phone string `json:"phone" validate:"required,phone"`
	}{"098765"}

	err = validation.Struct(dataValidate)

	fmt.Println(err)

	//Key: 'Phone' Error:Field validation for 'Phone' failed on the 'required' tag---PASSED
	fmt.Println("\n\n---------------------------------------")
	fmt.Println("TEST5")

	dataValidate = struct {
		Phone string `json:"phone" validate:"required,phone"`
	}{"98765"}

	err = validation.Struct(dataValidate)

	fmt.Println(err)

	//Key: 'Phone' Error:Field validation for 'Phone' failed on the 'required' tag---PASSED
	fmt.Println("\n\n---------------------------------------")
	fmt.Println("TEST6")

	dataValidate = struct {
		Phone string `json:"phone" validate:"required,phone"`
	}{"098765432A"}

	err = validation.Struct(dataValidate)

	fmt.Println(err)

	//Key: 'Phone' Error:Field validation for 'Phone' failed on the 'required' tag---PASSED
	fmt.Println("\n\n---------------------------------------")
	fmt.Println("TEST7")

	dataValidate = struct {
		Phone string `json:"phone" validate:"required,phone"`
	}{"09B76A4321"}

	err = validation.Struct(dataValidate)

	fmt.Println(err)

	//Key: 'Phone' Error:Field validation for 'Phone' failed on the 'required' tag---PASSED
	fmt.Println("\n\n---------------------------------------")
	fmt.Println("TEST8")

	dataValidate = struct {
		Phone string `json:"phone" validate:"required,phone"`
	}{"098765432-"}

	err = validation.Struct(dataValidate)

	fmt.Println(err)

	//Key: 'Phone' Error:Field validation for 'Phone' failed on the 'required' tag---PASSED
	fmt.Println("\n\n---------------------------------------")
	fmt.Println("TEST9")

	dataValidate = struct {
		Phone string `json:"phone" validate:"required,phone"`
	}{"0987-54321"}

	err = validation.Struct(dataValidate)

	fmt.Println(err)

	//Key: 'Phone' Error:Field validation for 'Phone' failed on the 'phone' tag---FAILED
	fmt.Println("\n\n---------------------------------------")
	fmt.Println("TEST9")

	dataValidate1 := struct {
		Phone string `json:"phone" validate:"phone"`
	}{}

	err = validation.Struct(dataValidate1)

	fmt.Println(err)

}
