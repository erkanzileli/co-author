package config

import (
	"fmt"
	"github.com/erkanzileli/co-author/internal/model"
	"os"

	"gopkg.in/yaml.v2"
)

const (
	configFilePath = ".co-author.yaml"
)

var (
	AppConfig Config

	// Version is the version of the application. It will be set by the build script.
	Version string
)

type Config struct {
	// Committers is a list of committers. It can be nil.
	Committers []model.Committer `yaml:"committers,omitempty"`

	// CommitMessageFilePath is the path of the file that will be used as the commit message.
	// It will be provided by Git when it's called as a prepare-commit-msg hook.
	CommitMessageFilePath string `yaml:"-"`
}

func Init() error {
	if !fileExists(configFilePath) {
		return nil
	}

	fileByteContent, err := os.ReadFile(configFilePath)
	if err != nil {
		return fmt.Errorf("failed to read config file %s: %w", configFilePath, err)
	}

	if err = yaml.Unmarshal(fileByteContent, &AppConfig); err != nil {
		return fmt.Errorf("failed to unmarshal config file %s: %w", configFilePath, err)
	}

	if len(os.Args) >= 3 {
		AppConfig.CommitMessageFilePath = os.Args[2]
	}

	return nil
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
