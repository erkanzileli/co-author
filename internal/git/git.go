package git

import (
	"fmt"
	"github.com/erkanzileli/co-author/internal/config"
	"github.com/erkanzileli/co-author/internal/model"
	"os"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

var (
	ErrCommitMsgFileDoesNotSpecified = errors.New("commit msg file does not specified on args, please check your git hook")
)

func FindCommitters() ([]model.Committer, error) {
	if len(config.AppConfig.Committers) > 0 {
		return config.AppConfig.Committers, nil
	}

	gitCmd := exec.Command("git", "shortlog", "-sen", "--group=author", "--group=trailer:Co-authored-by", "--all", "--no-merges")
	cutCmd := exec.Command("cut", "-c8-")

	gitOutPipe, err := gitCmd.StdoutPipe()
	if err != nil {
		return nil, errors.Wrap(err, "failed to create git output pipe")
	}

	cutCmd.Stdin = gitOutPipe

	if gitCmd.Start() != nil {
		return nil, errors.Wrap(err, "failed to start gitCmd")
	}

	output, err := cutCmd.Output()
	if err != nil {
		return nil, errors.Wrap(err, "failed to run cutCmd")
	}

	return parseCommitters(output), nil
}

func parseCommitters(raw []byte) []model.Committer {
	rawCommitters := strings.Split(string(raw), "\n")
	committers := make([]model.Committer, 0, len(rawCommitters))
	for _, rawCommitter := range rawCommitters {
		angleIndex := strings.Index(rawCommitter, "<")
		if angleIndex == -1 {
			continue
		}
		committers = append(committers, model.Committer{
			Name:  rawCommitter[:angleIndex-1],
			Email: rawCommitter[angleIndex+1 : len(rawCommitter)-1],
		})
	}
	return committers
}

func PrepareCommitMessageForCommitters(committers []model.Committer) (msg string) {
	for _, c := range committers {
		msg += fmt.Sprintf("Co-authored-by: %s <%s>\n", c.Name, c.Email)
	}
	return
}

func UpdateCommitMessageFile(message string) error {
	commitMessageFilePath := config.AppConfig.CommitMessageFilePath
	if commitMessageFilePath == "" {
		return ErrCommitMsgFileDoesNotSpecified
	}

	content, err := os.ReadFile(commitMessageFilePath)
	if err != nil {
		return errors.Wrap(err, "could not read file "+commitMessageFilePath)
	}

	newContent := append([]byte{}, []byte("\n\n")...)
	newContent = append(newContent, []byte(message)...)
	newContent = append(newContent, content...)

	if err = os.WriteFile(commitMessageFilePath, newContent, 0644); err != nil {
		return errors.Wrap(err, "could not write file "+commitMessageFilePath)
	}

	return nil
}
