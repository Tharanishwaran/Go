package main

import "fmt"

func main()  { //map is faster than array and instantly adding or remove elements then array the only problems in map is order

                //key   //value
   var mymap map[string]int  = map[string]int {


   "apples" :10,
   "orange" : 5,
   "mango" : 18,

    
}

fmt.Println(mymap)


fmt.Println(mymap["mango"])

// change value
mymap["orange"] = 890

// add new key value

mymap["bannana"] = 70


// to delete key value 
delete(mymap,"bannana")

// to check whether the key exist or not 
val, ok := mymap["apples"]
fmt.Println(val, ok)

fmt.Println(len(mymap))


}








// func main() {
//     // Create a map using the make function
//     myMap := make(map[string]int)

//     // Add key-value pairs to the map
//     myMap["apple"] = 5
//     myMap["banana"] = 3

//     // Access and print a value from the map
//     fmt.Println("Apple count:", myMap["apple"]) // Output: 5

//     // Check if a key exists in the map
//     value, exists := myMap["apple"]
//     if exists {
//         fmt.Println("Apple count:", value) // Output: 5
//     } else {
//         fmt.Println("Apple not found")
//     }

//     // Delete a key-value pair from the map
//     delete(myMap, "apple")

//     // Check if the key was deleted
//     _, exists = myMap["apple"]
//     if !exists {
//         fmt.Println("Apple has been deleted")
//     }

//     // Create and initialize a map using a map literal
//     anotherMap := map[string]int{
//         "one": 1,
//         "two": 2,
//         "three": 3,
//     }

//     // Iterate over the map and print each key-value pair
//     for key, value := range anotherMap {
//         fmt.Printf("Key: %s, Value: %d\n", key, value)
//     }
// }
