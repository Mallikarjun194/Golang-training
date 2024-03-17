package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type sample struct {
	Hey string `json:"hey"`
	Age int    `json:"age"`
}

func main() {
	fileContent := readFromJson()
	//fmt.Println(fileContent)
	//	var s sample
	var s1 sample
	if data, ok := fileContent["print"]; ok {
		jsonArray, _ := json.Marshal(data)
		json.Unmarshal(jsonArray, &s1)
		res := data.(sample)
		fmt.Println(s1.Hey)
		fmt.Println(s1)
		fmt.Println(res)
		fmt.Println(data)
	}
}

// read From json file
func readFromJson() map[string]interface{} {
	directory, err := os.Getwd()
	if err != nil {
		fmt.Println("Unable to fetch directory", err)
		return nil
	}
	fmt.Println("Currently working direcotry: ", directory)

	fileContent, err := os.Open(directory + "/sample.json")
	if err != nil {
		fmt.Println("Unable to read a file", err)
		return nil
	}
	var res map[string]interface{}
	defer fileContent.Close()
	byteArray, _ := ioutil.ReadAll(fileContent)

	err = json.Unmarshal(byteArray, &res)
	if err != nil {
		fmt.Println("Errror while unmarshelling")
		return nil
	}
	return res
}
