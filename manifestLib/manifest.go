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

func LoadSourceValues() {
	items, err := ioutil.ReadDir("./Data/values")
	if err != nil {
		log.Fatal("Error when reading the the ManifestFolder", err)
	}
	for _, item := range items {
		var s SourceValues
		fileName := item.Name()
		file, err := ioutil.ReadFile("./Data/values/" + fileName)
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
	items, err := ioutil.ReadDir("./Data/Form")
	if err != nil {
		log.Fatal("Error when reading the the ManifestFolder", err)
	}
	for _, item := range items {
		var m Manifest
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
