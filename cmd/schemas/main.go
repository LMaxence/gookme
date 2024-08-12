package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/LMaxence/gookme/packages/configuration"
	"github.com/alecthomas/jsonschema"
)

func generateGlobalSchema() (string, error) {
	schema := jsonschema.Reflect(&configuration.GookmeGlobalConfiguration{})

	// Print the JSON schema
	schemaJSON, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		return "", err
	}

	return string(schemaJSON), nil
}

func generateHookSchema() (string, error) {
	schema := jsonschema.Reflect(&configuration.HookConfiguration{})

	// Print the JSON schema
	schemaJSON, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		return "", err
	}

	return string(schemaJSON), nil
}

// Write the JSON schema files representing `GookmeGlobalConfiguration` and `HookConfiguration`
func main() {
	// Retrieve argv[1] to get the type of schema to generate
	if len(os.Args) < 2 {
		panic("Missing argument: schema type")
	}

	// Generate the JSON schema
	var schema string
	var err error
	switch os.Args[1] {
	case "global":
		schema, err = generateGlobalSchema()
	case "hooks":
		schema, err = generateHookSchema()
	default:
		panic("Unknown schema type: " + os.Args[1])
	}

	if err != nil {
		panic(err)
	}

	fmt.Println(schema)
}
