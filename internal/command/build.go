package command

import "gbend/internal/configuration"

// Build represents build command. The core command of gbend.
// It builds whole backend structure.
type Build struct {
	outputDirectory string                // outputDirectory is a directory when generated code will be stored in.
	cfg             *configuration.Config // cfg represents the configuration of the backend
}

func NewBuildCommand(outputDirectory string, cfg *configuration.Config) *Build {
	return &Build{
		outputDirectory: outputDirectory,
		cfg:             cfg,
	}
}

func (b *Build) Execute() error {
	return nil
}
