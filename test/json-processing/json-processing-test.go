package json_processing

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"testing"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

type BoxType struct {
	Items []ItemType `json:"items"`
}

type ItemType struct {
	ID int `json:"id"`
}

var data = []byte(`{"items":[
  {"id": 1},
  {"id": 2},
  {"id": 3},
  {"id": 4},
  {"id": 5}
]}`)

func TestUnmarshal(t *testing.T) {
	body := ioutil.NopCloser(bytes.NewBuffer(data))

	bodyBytes, err := ioutil.ReadAll(body)
	must(err)

	box := BoxType{}
	err = json.Unmarshal(bodyBytes, &box)
	must(err)

	if len(box.Items) != 5 {
		t.Fatalf("box wrong size %+v", box)
	}
}

func BenchmarkUnmarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		body := ioutil.NopCloser(bytes.NewBuffer(data))

		bodyBytes, err := ioutil.ReadAll(body)
		must(err)

		box := BoxType{}
		err = json.Unmarshal(bodyBytes, &box)
		must(err)

		if len(box.Items) != 5 {
			panic("box wrong length")
		}
	}
}
