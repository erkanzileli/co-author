package file

import (
	"github.com/erkanzileli/co-author/internal/model"
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

const (
	coAuthorsFilePath = ".git-co-authors.yaml"
)

type config struct {
	Committers []model.Committer `yaml:"committers,omitempty"`
}

func fileOpen(filename string) ([]byte, error) {
	buf, err := os.ReadFile(filename)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return buf, nil
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func configload(filename string) (*config, error) {
	buf, err := fileOpen(filename)
	if err != nil {
		return nil, err
	}
	cfg := &config{}
	err = yaml.Unmarshal(buf, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
