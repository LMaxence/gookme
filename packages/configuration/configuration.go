package configuration

type OnModifiedFileBehaviour string

const (
	AddAndCommit OnModifiedFileBehaviour = "addAndCommit"
	Abort        OnModifiedFileBehaviour = "abort"
)

/*
Global configuration for the Gookme tool.
Behaviours defined here will be used for all hooks to run.

The Gookme global configuration is a JSON file named .gookme.json, located at the root of the repository.
*/
type GookmeGlobalConfiguration struct {
	// The behaviour to use when there are modified files at the end of the hook.
	OnModifiedFiles *OnModifiedFileBehaviour `json:"onModifiedFiles,omitempty"`
}

func DefaultGlobalConfiguration() *GookmeGlobalConfiguration {
	onModifiedFiles := Abort
	return &GookmeGlobalConfiguration{
		OnModifiedFiles: &onModifiedFiles,
	}
}
