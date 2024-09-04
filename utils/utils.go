package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func WriteToJson[T any](data T, filePath string) {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println("Failed to Write Json for Marshal error: ", err)
	} else {
		file, err := os.Create(path + filePath)
		if err != nil {
			log.Println("Error in file creation: ", err)
		} else {
			_, err := file.Write(jsonData)
			if err != nil {
				log.Println("Error in writing to file: ", err)
			}
		}
	}
}
