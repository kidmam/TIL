package main

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
)

// https://preslav.me/2020/08/12/lessons-from-golang-keep-calm-and-use-the-byte-array/

func main() {
	/*f, err := os.Open("./filename")
	defer f.Close()
	if err != nil {
		// handle the error
	}

	buf := bytes.NewBuffer(make([]byte, 0))
	buf.ReadFrom(f)
	fmt.Printf(string(buf.Bytes())) */

	jsonString := `{"Foo": "Bar"}`

	// Parse the incoming JSON string
	type S struct{ Foo string }
	var s S
	json.Unmarshal([]byte(jsonString), &s)

	// Turn into XML
	xmlResult, _ := xml.MarshalIndent(&s, "", "")

	// Write to a file
	ioutil.WriteFile("filename", xmlResult, 0611)
}
