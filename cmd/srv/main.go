package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"time"

    godotenv "github.com/joho/godotenv"
)

type Data struct {
	Script map[string]string `json:"Script"`
}

func main() {
	// Load environment variables (assuming .env in the same directory)
    godotenv.Load(.env)

	mainDir := os.Getenv("MAIN_DIR")
	schemaDir := filepath.Join(mainDir, os.Getenv("SCHEMA_DIR"))
	schemaFile := filepath.Join(schemaDir, "script.json")
	venvDir := filepath.Join(mainDir, os.Getenv("PY_ENV")) // Not used directly in this Go script

	// Read the JSON file
	file, err := ioutil.ReadFile(schemaFile)
	if err != nil {
		panic(err)
	}

	var data Data
	err = json.Unmarshal(file, &data)
	if err != nil {
		panic(err)
	}

	batch_size := 8 // Max 16, Ideal 4, Best 8
	total_execution_time := 0.0

	for batch_number, i := 0, 0; i < len(data.Script); batch_number, i = batch_number+1, i+batch_size {
		batch_end := i + batch_size
		if batch_end > len(data.Script) {
			batch_end = len(data.Script)
		}

		batch := make(map[string]string)
		script_number := 0
		for k, v := range data.Script {
			if script_number >= i && script_number < batch_end {
				batch[k] = v
			}
			script_number++
		}

		// Run scripts in parallel
		for script_number, script_name := range batch {
			fmt.Printf("Running script %s: %s\n", script_number, script_name)
			script_start_time := time.Now()

			os.Setenv("BATCH_NUMBER", strconv.Itoa(batch_number+1))

			regionTrim := regexp.MustCompile(`_(.+?)/`).FindStringSubmatch(script_name)[1]
			numberDecimalTrim := regexp.MustCompile(`(\d+)_`).FindStringSubmatch(script_name)[1]
			numberTrim, _ := strconv.Atoi(regexp.MustCompile(`(\d+)_`).FindStringSubmatch(script_name)[1])

			os.Setenv("MY_REGION", regionTrim)
			os.Setenv("SCRIPT_NUMBER_DECIMAL", numberDecimalTrim)
			os.Setenv("SCRIPT_NUMBER", strconv.Itoa(numberTrim))
			os.Setenv("SCRIPT_NAME", script_name)

			// Start the script as a subprocess
			cmd := exec.Command("python", script_name)
			cmd.Start()
			cmd.Wait()

			fmt.Printf("Script %s executed!\n", script_number)

			script_end_time := time.Now()
			script_execution_time := script_end_time.Sub(script_start_time).Seconds()
			fmt.Printf("Script %s executed in %f seconds.\n", script_number, script_execution_time)

			total_execution_time += script_execution_time
		}
	}
	fmt.Printf("All scripts executed in %f seconds.\n", total_execution_time)
}

func loadEnv(filePath string) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	for _, line := range regexp.MustCompile(`\r?\n`).Split(string(file), -1) {
		if len(line) == 0 {
			continue
		}
		parts := regexp.MustCompile(`=`).Split(line, 2)
		if len(parts) != 2 {
			continue
		}
		os.Setenv(parts[0], parts[1])
	}
}
