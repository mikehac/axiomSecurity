package manifestlib

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

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

func BuildManifests() {
	manifestsFolders, err := os.ReadDir("./Data/Manifests/Components")
	if err != nil {
		log.Fatal("Error when reading the the ManifestFolder", err)
	}
	for _, mf := range manifestsFolders {
		var m Manifest
		componetnFolderName := mf.Name()
		currentComponents, err := os.ReadDir("./Data/Manifests/Components/" + componetnFolderName)
		if err != nil {
			log.Fatal("Error when reading Component's folder:", err)
		}
		for _, component := range currentComponents {
			var c Component
			fileName := component.Name()
			file, err := os.ReadFile("./Data/Manifests/Components/" + componetnFolderName + "/" + fileName)
			if err != nil {
				log.Fatal("Error when opening file: ", err)
			}
			err = json.Unmarshal([]byte(file), &c)
			if err != nil {
				log.Fatal("Error when parsing to the object", err)
			}
			m.Components = append(m.Components, c)
		}
		ManifestsMap[componetnFolderName] = m
	}
}
