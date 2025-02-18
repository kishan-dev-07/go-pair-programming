package main

import "fmt"

func main(){
	var map1 = make(map[string]int)
	map1["one"] = 1
	map1["two"] = 2
	map1["three"] = 3
	var map2 = make(map[int]string)

	for key,value:= range(map1){
		map2[value] = key
	}
	
	for key,value:= range(map2){
		fmt.Println(key,":",value)
	}
}