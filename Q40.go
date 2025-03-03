package main

import "fmt"

func main(){
	type parent struct{
		value int
	}

	func(c parent) child() string {
		return fmt.Printf("value:%d",c.value)
	}

	parent:={
		value:10
	}

	fmt.Println(parent.value)
}