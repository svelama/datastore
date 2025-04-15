package main

import (
	"fmt"
)

func main() {

	ds := NewDataStore[string]()

	fmt.Println("1. Inserting a value into data store...")
	err := ds.Insert("key001", "val001")
	if err != nil {
		fmt.Printf("Expected no error, but got: %v\n", err)
		return
	}
	fmt.Printf("Insertion successful!\n\n")

	fmt.Println("2. Reading a value from data store...")
	val, err := ds.Read("key001")
	if err != nil {
		fmt.Printf("Expected no error, but got: %v\n", err)
		return
	}
	fmt.Println(val)
	fmt.Printf("Successfully read\n\n")

	fmt.Println("3. Removing a value from data store...")
	err = ds.Remove("key001")
	if err != nil {
		fmt.Printf("Expected no error, but got: %v\n", err)
		return
	}
	fmt.Println("Successfully removed")
}
