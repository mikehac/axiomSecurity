package manifestlib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

// func (a *Accounts) LoadData() {
// 	file, err := ioutil.ReadFile("./Data/values/accounts.json")
// 	if err != nil {
// 		log.Fatal("Error when opening file: ", err)
// 	}
// 	err = json.Unmarshal([]byte(file), &a)
// 	if err != nil {
// 		log.Fatal("Error when parsing to the object", err)
// 	}
// }

//	func (d *DatabaseInstances) LoadData() {
//		file, err := ioutil.ReadFile("./Data/values/databaseInstances.json")
//		if err != nil {
//			log.Fatal("Error when opening file: ", err)
//		}
//		err = json.Unmarshal([]byte(file), &d)
//		if err != nil {
//			log.Fatal("Error when parsing to the object", err)
//		}
//	}
func LoadSourceValues() {
	var s SourceValues
	items, err := ioutil.ReadDir("./Data/values")
	if err != nil {
		log.Fatal("Error when reading the the ManifestFolder", err)
	}
	for _, item := range items {
		fileName := item.Name()
		file, err := ioutil.ReadFile("./Data/Form/" + fileName)
		if err != nil {
			log.Fatal("Error when opening file: ", err)
		}
		err = json.Unmarshal([]byte(file), &s)
		if err != nil {
			log.Fatal("Error when parsing to the object", err)
		}
		fileNameNoExtantion := strings.Split(fileName, ".")[0]
		SourceValueMap[fileNameNoExtantion] = s
	}
}

func LoadManifest() {
	var m Manifest
	items, err := ioutil.ReadDir("./Data/Form")
	if err != nil {
		log.Fatal("Error when reading the the ManifestFolder", err)
	}
	for _, item := range items {
		fileName := item.Name()
		file, err := ioutil.ReadFile("./Data/Form/" + fileName)
		if err != nil {
			log.Fatal("Error when opening file: ", err)
		}
		err = json.Unmarshal([]byte(file), &m)
		if err != nil {
			log.Fatal("Error when parsing to the object", err)
		}
		fileNameNoExtantion := strings.Split(fileName, ".")[0]
		ManifestsMap[fileNameNoExtantion] = m
	}
}
