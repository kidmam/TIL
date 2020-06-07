package main

// https://bitfieldconsulting.com/golang/map-string-interface
func main() {
	/* foods := map[string]interface{}{
		"bacon": "delicious",
		"eggs": struct {
			source string
			price  float64
		}{"chicken", 1.75},
		"steak": true,
	} */

	type Person struct {
		Name    string
		Age     int
		Hobbies []string
	}

	/*data := `{
			    "name":"John",
				"age":29,
				"hobbies":[
					"martial arts",
					"breakfast foods",
					"piano"
	            ]
			}` */

	//var p map[string]interface{}
	//_ = json.Unmarshal(data, &p)

}
