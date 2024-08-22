package hooksscripts

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/LMaxence/gookme/packages/configuration"
	"github.com/LMaxence/gookme/packages/logging"
)

var logger = logging.NewLogger("hooks-scripts")

var (
	ErrorScriptAlreadyExists   = errors.New("script already exists")
	ErrorScriptVersionNotFound = errors.New("script version not found")
)

type ScriptVariables struct {
	HookType configuration.HookType
	Version  string
}

func GenerateScript(variables *ScriptVariables, template string) string {
	script := template

	script = strings.ReplaceAll(script, "{{.Version}}", variables.Version)
	script = strings.ReplaceAll(script, "{{.HookType}}", string(variables.HookType))

	return script
}

// Look for a script in the provided content, and parse it to get the version it has been generated with
func ParseScriptVersion(content string) (string, error) {
	// Look for the version of the script
	version := ""
	prefix := "# GOOKME_CLI_VERSION: "

	lines := strings.Split(content, "\n")

	for _, line := range lines {
		if strings.HasPrefix(line, prefix) {
			version = strings.TrimSpace(strings.TrimPrefix(line, prefix))
			logger.Debugf("Found script version: %s", version)
			return version, nil
		}
	}

	return version, ErrorScriptVersionNotFound
}

// Remove the gookme script from the provided content
func RemoveExistingGookmeScript(
	content string,
) string {
	lines := strings.Split(content, "\n")

	startLine := "# Start of automatically generated script"
	endLine := "# End of automatically generated script"

	// Remove everything between the start and end line
	startIndex := -1
	endIndex := -1

	for i, line := range lines {
		if strings.HasPrefix(line, startLine) {
			startIndex = i
		}
		if strings.HasPrefix(line, endLine) {
			endIndex = i
		}
	}

	logger.Debugf("Found start line %d and end line %d", startIndex, endIndex)
	if startIndex == -1 {
		logger.Debugf("No start line found in script")
		return content
	}
	if endIndex == -1 {
		logger.Debugf("No end line found in script")
		return content
	}

	return strings.Join(append(lines[:startIndex], lines[endIndex+1:]...), "\n")
}

func AddGookmeScript(
	content string,
	scriptscriptVariables *ScriptVariables,
) (string, error) {
	// Check if the script already exists, error if it does
	_, err := ParseScriptVersion(content)
	if err != nil && !errors.Is(err, ErrorScriptVersionNotFound) {
		return "", err
	}

	// Load the template file and read its content, using the directory of this file's path
	_, filePath, _, ok := runtime.Caller(0)
	if !ok {
		return content, errors.New("failed to load script template")
	}
	dir := filepath.Dir(filePath)

	// Read the template file
	templatePath := filepath.Join(dir, "template.sh")
	templateContent, err := os.ReadFile(templatePath)
	if err != nil {
		return content, err
	}

	// Generate the script
	script := GenerateScript(scriptscriptVariables, string(templateContent))

	if strings.HasSuffix(content, "\n\n") {
		return content + script, nil
	} else if strings.HasSuffix(content, "\n") {
		return content + "\n" + script, nil
	} else {
		return content + "\n\n" + script, nil
	}
}
