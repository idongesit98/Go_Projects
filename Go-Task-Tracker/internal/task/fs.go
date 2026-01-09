package task

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

func tasksFilePath() string {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return ""
	}
	return path.Join(cwd, "tasks.json")
}

func ReadTaskFromFile() ([]Task, error) {
	//Get current working directory
	filePath := tasksFilePath()

	_, err := os.Stat(filePath)

	if os.IsNotExist(err) {
		fmt.Println("File does not exist. Creating file.....")
		file, err := os.Create(filePath)
		os.WriteFile(filePath, []byte("[]"), os.ModeAppend.Perm())

		if err != nil {
			fmt.Println("Error creating file:", err)
			return nil, err
		}

		defer file.Close()
		return []Task{}, nil
	}
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	tasks := []Task{}
	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil {
		fmt.Println("Error decoding file:", err)
		return nil, err
	}
	return tasks, nil
}

func WriteTaskToFile(tasks []Task) error {
	filepath := tasksFilePath()
	file, err := os.Create(filepath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}
	defer file.Close()
	err = json.NewEncoder(file).Encode(tasks)
	if err != nil {
		fmt.Println("Error encoding file:", err)
		return err
	}
	return nil
}
