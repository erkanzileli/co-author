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

	ErrCommitMsgFileDoesNotSpecified = fmt.Errorf("commit msg file does not specified on args, please check your git hook")
)

type Config struct {
	// Committers is a list of committers. It can be nil.
	Committers []model.Committer `yaml:"committers,omitempty"`

	// CommitMessageFilePath is the path of the file that will be used as the commit message.
	// It will be provided by Git when it's called as a prepare-commit-msg hook.
	CommitMessageFilePath string `yaml:"-"`

	// CommitSource specifies the source of the commit message.
	// It will be provided by Git when it's called as a prepare-commit-msg hook.
	// It can be one of the following values:
	// - "message": The commit message is provided by the user.
	// - "template": The commit message is provided by Git when it's called as a prepare-commit-msg hook.
	// - "merge": The commit message is provided by Git when it's called as a prepare-commit-msg hook.
	// - "squash": The commit message is provided by Git when it's called as a prepare-commit-msg hook.
	CommitSource string `yaml:"-"`

	// CommitSHA1 is the SHA1 of the commit.
	CommitSHA1 string `yaml:"-"`
}

func Init() error {
	if len(os.Args) >= 3 {
		AppConfig.CommitMessageFilePath = os.Args[2]
	}
	if len(os.Args) >= 4 {
		AppConfig.CommitSource = os.Args[3]
	}
	if len(os.Args) >= 5 {
		AppConfig.CommitSHA1 = os.Args[4]
	}

	if AppConfig.CommitMessageFilePath == "" {
		return ErrCommitMsgFileDoesNotSpecified
	}

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

	return nil
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
