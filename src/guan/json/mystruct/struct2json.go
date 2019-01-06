package mystruct

import (
	"encoding/json"
	"fmt"
)

/*
JsonText is ...
*/
var JsonText = ` {
	"name": "jacky",
	"age": 12,
	"country": {
			"ccode": "America",
			"cname": "usa"
	}
}`

type Person struct {
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Country Country `json:"country"`
}

type Country struct {
	Countrycode string `json:"ccode"`
	Countryname string `json:"cname"`
}

func main() {

	p := Person{"jacky", 12, Country{"America", "usa"}}
	// p := Person{ "jacky" , 12}

	result, error := json.MarshalIndent(p, "", "	")
	if error != nil {

		fmt.Println(" error : ", error)
	}

	fmt.Printf(" result	 is %v \n ", string(result))

	p2 := Person{}
	error2 := json.Unmarshal([]byte(JsonText), &p2)

	if error2 != nil {
		fmt.Println("error2 : ", error2)
	}
	fmt.Println(" result2 = ", p2)

}
