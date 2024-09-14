package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"

    "gopkg.in/yaml.v2"
    "github.com/spf13/cobra"
    "github.com/ljtian/j2y/pkg/help"
)

func main() {

    // Get language setting from environment variable
    helpLang := os.Getenv("J2Y_LANG")
    if helpLang == "" {
        helpLang = os.Getenv("LANG")
    }
    if _, exists := help.HelpData[helpLang]; !exists {
        helpLang = "en" // Default language
    }

    var inputFile string
    var jsonData string
    var outputFile string

    var rootCmd = &cobra.Command{
        Use:   help.HelpData[helpLang].Usage,
        Short: help.HelpData[helpLang].ShortDesc,
        Long:  help.HelpData[helpLang].LongDesc,
        Args:  cobra.MaximumNArgs(1), // Maximum one output file argument
        Run: func(cmd *cobra.Command, args []string) {
            run(inputFile, jsonData, args)
        },
    }

    // Add command line arguments
    rootCmd.Flags().StringVarP(&inputFile, "file", "f", "", "Input JSON file path")
    rootCmd.Flags().StringVarP(&jsonData, "input", "i", "", "Input JSON data")
    rootCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Output YAML file (optional)")

    // Run command
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

// Read JSON data from file
func readJSONFromFile(filePath string) (interface{}, error) {
    jsonFile, err := ioutil.ReadFile(filePath)
    if err != nil {
        return nil, err
    }

    var data interface{}
    if err := json.Unmarshal(jsonFile, &data); err != nil {
        return nil, err
    }

    return data, nil
}

// Parse JSON data
func parseJSON(data string) (interface{}, error) {
    var result interface{}
    if err := json.Unmarshal([]byte(data), &result); err != nil {
        return nil, err
    }
    return result, nil
}

// Convert to YAML format
func convertToYAML(data interface{}) ([]byte, error) {
    return yaml.Marshal(data)
}

// Handle command line input and conversion logic
func run(inputFile, jsonData string, args []string) {
    var data interface{}
    var err error

    if isInputFromPipe() {
        jsonData, err = readFromPipe()
        if err != nil {
            fmt.Printf("Failed to read from standard input: %v\n", err)
            return
        }
    }

    if inputFile != "" {
        data, err = readJSONFromFile(inputFile)
        if err != nil {
            fmt.Printf("Failed to read file: %v\n", err)
            return
        }
    } else if jsonData != "" {
        data, err = parseJSON(jsonData)
        if err != nil {
            fmt.Printf("Failed to parse JSON: %v\n", err)
            return
        }
    } else {
        fmt.Println("Please provide input JSON data or file path.")
        return
    }

    yamlData, err := convertToYAML(data)
    if err != nil {
        fmt.Printf("Failed to convert to YAML: %v\n", err)
        return
    }

    outputFile := ""
    if len(args) > 0 {
        outputFile = args[0]
    }

    if outputFile == "" {
        fmt.Printf("%v\n", string(yamlData))
        return
    }

    if err := ioutil.WriteFile(outputFile, yamlData, 0644); err != nil {
        fmt.Printf("Failed to write file: %v\n", err)
        return
    }

    fmt.Println("Conversion successful!")
}

// Check if there is data from standard input
func isInputFromPipe() bool {
    fi, err := os.Stdin.Stat()
    if err != nil {
        return false
    }
    return fi.Mode()&os.ModeNamedPipe != 0
}

// Read JSON data from standard input
func readFromPipe() (string, error) {
    bytes, err := ioutil.ReadAll(os.Stdin)
    if err != nil {
        return "", err
    }
    return string(bytes), nil
}
