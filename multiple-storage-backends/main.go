package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	mapRepo := NewMapRepository[string]()
	ds := NewDataStore(mapRepo)

	fmt.Println("1. Inserting a value into data store...")
	err := ds.Insert("key001", "val001")
	if err != nil {
		fmt.Printf("Expected no error, but got: %v\n", err)
		return
	}
	fmt.Println("Insertion successful!")

	fmt.Println("2. Reading a value from data store...")
	ctx, _ := context.WithTimeout(context.Background(), 100*time.Millisecond)
	val, err := ds.Read(ctx, "key001")
	if err != nil {
		fmt.Printf("Expected no error, but got: %v\n", err)
		return
	}
	fmt.Println(val)
	fmt.Println("Successfully read")

	fmt.Println("3. Removing a value from data store...")
	err = ds.Remove("key001")
	if err != nil {
		fmt.Printf("Expected no error, but got: %v\n", err)
		return
	}
	fmt.Println("Successfully removed")
}
