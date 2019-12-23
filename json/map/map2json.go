package main 

import (

	"fmt"
	"encoding/json"
)

func main(){
	m := make(map[string]interface{} , 4)
	m["a"] = "America"
	m["b"] = "Great Britain"
	m["c"] = []string{"Canada", "china", "Chicago"}
	m["d"] = false

	result ,error := json.MarshalIndent(m , "" , "	")

	if error != nil {
		fmt.Println(" 	error is " , error)
		return 
	}
	fmt.Printf(" result	 is %v \n " , string(result))


}
