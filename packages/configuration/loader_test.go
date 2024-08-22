package configuration

import (
	"path/filepath"
	"testing"

	helpers "github.com/LMaxence/gookme/packages/test-helpers"
	"github.com/stretchr/testify/assert"
)

func TestLoadGlobalConfiguration(t *testing.T) {
	directory, err := helpers.SetupTmpGit()
	assert.NoError(t, err)

	// Create a configuration file
	err = helpers.WriteFile(directory, ".gookme.json", `{"onModifiedFiles": "addAndCommit"}`)
	assert.NoError(t, err)

	// Call the function
	config, err := LoadGlobalConfiguration(directory)
	assert.NoError(t, err)
	assert.NotNil(t, config)
	assert.Equal(t, *config.OnModifiedFiles, AddAndCommit)
}

func TestLoadDefaultGlobalConfigurationIfNotFound(t *testing.T) {
	directory, err := helpers.SetupTmpGit()
	assert.NoError(t, err)

	// Call the function
	config, err := LoadGlobalConfiguration(directory)
	assert.NoError(t, err)
	assert.NotNil(t, config)
	assert.Equal(t, *config.OnModifiedFiles, Abort)
}

func TestLoadHooksConfiguration(t *testing.T) {
	directory, err := helpers.SetupTmpGit()
	assert.NoError(t, err)

	// Create a configuration file
	err = helpers.WriteFile(directory, "pre-commit.json", `{"steps": [{"name": "step1", "command": "echo 'Hello World'"}]}`)
	assert.NoError(t, err)

	// Call the function
	hooks, err := LoadHooksConfiguration(directory, PreCommit)
	assert.NoError(t, err)

	assert.Len(t, hooks, 1)
	assert.Equal(t, hooks[0].Path, directory)
	assert.Len(t, hooks[0].Steps, 1)
	assert.Equal(t, hooks[0].Steps[0].Name, "step1")
	assert.Equal(t, hooks[0].Steps[0].Command, "echo 'Hello World'")
}

func TestLoadNestedHooksConfiguration(t *testing.T) {
	directory, err := helpers.SetupTmpGit()
	assert.NoError(t, err)

	// Create a configuration file
	err = helpers.WriteFile(filepath.Join(directory, "packages/a"), "pre-commit.json", `{"steps": [{"name": "step1", "command": "echo 'Hello World'"}]}`)
	assert.NoError(t, err)

	// Call the function
	hooks, err := LoadHooksConfiguration(directory, PreCommit)
	assert.NoError(t, err)

	assert.Len(t, hooks, 1)
	assert.Equal(t, hooks[0].Path, filepath.Join(directory, "packages/a"))
	assert.Len(t, hooks[0].Steps, 1)
	assert.Equal(t, hooks[0].Steps[0].Name, "step1")
	assert.Equal(t, hooks[0].Steps[0].Command, "echo 'Hello World'")
}

func TestLoadHooksConfigurationFromEmptyDirectory(t *testing.T) {
	directory, err := helpers.SetupTmpGit()
	assert.NoError(t, err)

	hooks, err := LoadHooksConfiguration(directory, PreCommit)
	assert.NoError(t, err)
	assert.Equal(t, []Hook{}, hooks)
}

func TestLoadConfiguration(t *testing.T) {
	directory, err := helpers.SetupTmpGit()
	assert.NoError(t, err)

	// Create a configuration file
	err = helpers.WriteFile(directory, ".gookme.json", `{"onModifiedFiles": "addAndCommit"}`)
	assert.NoError(t, err)

	err = helpers.WriteFile(directory, "pre-commit.json", `{"steps": [{"name": "step1", "command": "echo 'Hello World'"}]}`)
	assert.NoError(t, err)

	config, err := LoadConfiguration(directory, PreCommit)

	assert.NoError(t, err)
	assert.NotNil(t, config)
	assert.NotNil(t, config.Global)
	assert.Equal(t, *config.Global.OnModifiedFiles, AddAndCommit)
	assert.Len(t, config.Hooks, 1)
	assert.Equal(t, config.Hooks[0].Path, directory)
	assert.Len(t, config.Hooks[0].Steps, 1)
	assert.Equal(t, config.Hooks[0].Steps[0].Name, "step1")
	assert.Equal(t, config.Hooks[0].Steps[0].Command, "echo 'Hello World'")
}
