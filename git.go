package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	"github.com/pkg/errors"
	"os"
	"os/exec"
	"strings"
)

var (
	errCommitMsgFileDoesNotSpecified = errors.New("commit msg file does not specified on args, please check your git hook")
)

func findCommitters() ([]list.Item, error) {
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

func parseCommitters(raw []byte) []list.Item {
	rawCommitters := strings.Split(string(raw), "\n")
	committers := make([]list.Item, 0, len(rawCommitters))
	for _, rawCommitter := range rawCommitters {
		angleIndex := strings.Index(rawCommitter, "<")
		if angleIndex == -1 {
			continue
		}
		committers = append(committers, &committer{
			name:  rawCommitter[:angleIndex-1],
			email: rawCommitter[angleIndex+1 : len(rawCommitter)-1],
		})
	}
	return committers
}

func prepareCommitMsg(committers []*committer) (msg string) {
	for _, c := range committers {
		msg += fmt.Sprintf("Co-authored-by: %s <%s>\n", c.name, c.email)
	}
	return
}

func writeToCommitMsgFile(text string) error {
	if len(os.Args) < 3 {
		return errCommitMsgFileDoesNotSpecified
	}

	commitMsgFilePath := os.Args[2]

	content, err := os.ReadFile(commitMsgFilePath)
	if err != nil {
		return errors.Wrap(err, "could not read file "+commitMsgFilePath)
	}

	newContent := append([]byte{}, []byte("\n\n")...)
	newContent = append(newContent, []byte(text)...)
	newContent = append(newContent, content...)

	err = os.WriteFile(commitMsgFilePath, newContent, 0644)
	if err != nil {
		return errors.Wrap(err, "could not write file "+commitMsgFilePath)
	}

	return nil
}
