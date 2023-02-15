package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	f, err := os.Open("./websitedata.json")
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	bytes, _ := ioutil.ReadAll(f)

	json.Unmarshal(bytes, &data)

	fmt.Printf("%+v\n", data)

}
