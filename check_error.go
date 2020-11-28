package main

import "fmt"

//check error
func checkError(err error) {
	if err != nil {
		fmt.Printf("Error in unmarshal: %s", err.Error())
	}

}
