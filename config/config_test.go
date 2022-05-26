package config

import (
	"fmt"
	"os"
	"testing"
)

func createFile(path string) {
	// check if file exists
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return
		}
		defer file.Close()
	}

	fmt.Println("File Created Successfully", path)
}

func writeFile(path string) {
	// Open file using READ & WRITE permission.
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// Write some text line-by-line to file.
	_, err = file.WriteString("SAMPLE=data")
	if isError(err) {
		return
	}

	// Save file changes.
	err = file.Sync()
	if isError(err) {
		return
	}

	fmt.Println("File Updated Successfully.")
}

func deleteFile(path string) {
	// delete file
	var err = os.Remove(path)
	if isError(err) {
		return
	}

	fmt.Println("File Deleted")
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func TestLoadConfig(t *testing.T) {

	type envData struct {
		Sample string
	}
	wrongPaths := []string{"."}
	_, err := LoadConfig(wrongPaths, "app", envData{})
	if err == nil {
		t.Errorf("error while asdasdasd file %v", err)
	}
	createFile("default.env")
	writeFile("default.env")

	defer deleteFile("default.env")
	paths := []string{"./config", "/app/config", "."}
	configInterface, err := LoadConfig(paths, "default", envData{})
	if err != nil {
		fmt.Println("into correct 1")
		t.Errorf("error while reading file %v", err)
	}

	config := configInterface.(envData)
	if config.Sample != "data" {
		fmt.Println("into correct 2")
		t.Errorf("error while reading file %v", err)
	}

}
