package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fxivan/data-structure-go/hashmap"
)

// Estructura del cliente
type Customer struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type Customers struct {
	Customers []Customer `json:"customers"`
}

// Estructura de la cuenta bancaria
type Account struct {
	ID         int     `json:"id"`
	CustomerID int     `json:"customer_id"`
	Type       string  `json:"type"`
	Balance    float64 `json:"balance"`
}

// Estructura de la transacción
type Transaction struct {
	ID          int     `json:"id"`
	FromAccount int     `json:"from_account"`
	ToAccount   int     `json:"to_account"`
	Amount      float64 `json:"amount"`
	Date        string  `json:"date"`
}

func main() {
	// Create a new hashmap with size 10
	myHashMap := hashmap.NewHashMap(10)

	jsonFile, err := os.Open("bank_data.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var dataBank Customers
	json.Unmarshal(byteValue, &dataBank)

	defer jsonFile.Close()
	for _, data := range dataBank.Customers {
		myHashMap.Insert(data.ID, data.Name)
	}

	// Get and print values
	value := myHashMap.Get(10000)
	fmt.Println("Value for key:", value)

	// Delete a key
	//  myHashMap.Delete(10000)
	/* If we try to get the value for key "foo" we will get an empty string. (You can return a
	proper error or a flag in your get method) */
}
