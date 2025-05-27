package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type Data struct {
	MetaData `json:"metadata"`
	Tasks    map[int]Task `json:"tasks"`
}

type MetaData struct {
	TaskCount int `json:"task_count"`
}

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func getData() (Data, error) {
	jsonPath := os.Getenv("FILE_PATH")
	if jsonPath == "" {
		log.Fatalf("Json file path not set!")
	}
	jsonFile, err := os.Open(jsonPath)
	if err != nil {
		return Data{}, fmt.Errorf("error opening json file : %w", err)
	}
	defer jsonFile.Close()

	tasks := Data{
		Tasks: make(map[int]Task),
	}
	err = json.NewDecoder(jsonFile).Decode(&tasks)
	if err != nil {
		return Data{}, fmt.Errorf("error decodong json file : %w", err)
	}

	return tasks, nil
}

func saveJSON(data Data) error {
	jsonPath := os.Getenv("FILE_PATH")
	if jsonPath == "" {
		log.Fatalf("Json file path not set!")
	}
	jsonFile, err := os.Create(jsonPath)
	if err != nil {
		return fmt.Errorf("error opening json file : %w", err)
	}
	defer jsonFile.Close()

	err = json.NewEncoder(jsonFile).Encode(&data)
	if err != nil {
		return fmt.Errorf("error saving json file : %w", err)
	}

	return nil

}
