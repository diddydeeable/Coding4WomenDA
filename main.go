package main

import "fmt"

func MyMap(arr []string) map[string]int {
	//we initalize a map
	myMap := make(map[string]int)

	//we do a for range loop over our slice
	for _, v := range arr {

		//checking if key exists, and saving the true/false into the variable 'exists'.
		_, exists := myMap[v]

		//if the key doesn't 'v' exists in the map, we create an instance of it, and give it a value 1
		if !exists {
			myMap[v] = 1
			//if the 'v' or fruit key already exists, then we just add 1 to the value
		} else {
			myMap[v] += 1
		}
	}
	fmt.Print(myMap)
	return myMap
}

func main() {
	//f.Strawberry(99)
	//words := []string{"apple", "banana", "apple", "orange", "banana", "apple", "kiwi", "kiwi", "kiwi"}
	yadiraFruit := []string{"banana", "coconut", "coconut"}
	MyMap(yadiraFruit)
}
