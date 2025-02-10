package ruleset

import (
	"fmt"
	"net/url"
	"path/filepath"

	"github.com/xeipuuv/gojsonschema"
)

func ValidateJSON(jsonData interface{}, schemaPath string) error {
	absSchemaPath, err := filepath.Abs(schemaPath)
	if err != nil {
		return fmt.Errorf("error getting absolute path: %v", err)
	}

	fileURL := url.PathEscape(absSchemaPath)
	if filepath.Separator == '\\' {
		fileURL = "file:///" + filepath.ToSlash(absSchemaPath)
	} else {
		fileURL = "file://" + absSchemaPath
	}

	schemaLoader := gojsonschema.NewReferenceLoader(fileURL)

	jsonLoader := gojsonschema.NewGoLoader(jsonData)

	result, err := gojsonschema.Validate(schemaLoader, jsonLoader)
	if err != nil {
		return fmt.Errorf("error validating JSON: %v", err)
	}

	if !result.Valid() {
		errMsg := "JSON validation failed:\n"
		for _, desc := range result.Errors() {
			errMsg += fmt.Sprintf("- %s\n", desc)
		}
		return fmt.Errorf(errMsg)
	}

	return nil
}
