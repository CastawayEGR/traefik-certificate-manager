package main

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/jpillora/opts"

	"tcm/types"
	"tcm/utils"
)

var (
	appVersion string
	buildTime string
	gitCommit string
)

func main() {
	var options types.Options
	opts.New(&options).Parse()
	if options.Version {
		fmt.Printf("Version:    %s\n", appVersion)
		fmt.Printf("Build Time: %s\n", buildTime)
		fmt.Printf("Git Commit: %s\n", gitCommit)
		return
	} else if options.File == "" {
		fmt.Println("Please provide the path to a JSON file using the --file (-f) flag")
		return
	}

	// Open and decode the JSON file
	data, err := utils.LoadJSONFile(options.File)
	if err != nil {
    		fmt.Println(err)
    		return
	}

	// Prompt for selection of domain or sans
	var selection string
	err = survey.AskOne(&survey.Select{
		Message: "Select 'domain' or 'sans':",
		Options: []string{"domain", "sans"},
	}, &selection)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Extract values based on selection
	var selectedValues []string
	var isDomain bool
	switch selection {
	case "domain":
		selectedValues, isDomain = utils.ExtractValues(data, true)
	case "sans":
		selectedValues, isDomain = utils.ExtractValues(data, false)
	}

	// Create multi-select prompt
	prompt := &survey.MultiSelect{
		Message: "Select values:",
		Options: selectedValues,
	}
	var selected []string
	err = survey.AskOne(prompt, &selected)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Delete selected values
	for _, value := range selected {
    		if isDomain {
        		utils.DeleteCertificateByDomain(&data, value)
    		} else {
        		utils.DeleteSansByDomain(&data, value)
    		}
	}

	// Save updated JSON
	err = utils.WriteJSONToFile(data, options.File)
	if err != nil {
    		fmt.Println("Error writing JSON to file:", err)
    		return
	}
	fmt.Println("Updated JSON data written to file.")

}
