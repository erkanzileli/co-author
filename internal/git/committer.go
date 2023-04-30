package git

import (
	"fmt"
	"github.com/erkanzileli/co-author/internal/config"
	"github.com/erkanzileli/co-author/internal/model"
	"github.com/pkg/errors"
	"os"
	"os/exec"
	"strings"
)

func FindCommitters() ([]model.Committer, error) {
	var committers []model.Committer

	committersFromCommitMessage, err := findCommittersFromCommitMessage()
	if err != nil {
		return nil, fmt.Errorf("failed to find committers from commit message: %w", err)
	}

	committers = append(committers, committersFromCommitMessage...)

	if len(config.AppConfig.Committers) > 0 {
		committers = append(committers, config.AppConfig.Committers...)
		return committers, nil
	}

	committersFromHistory, err := findCommittersFromHistory()
	if err != nil {
		return nil, fmt.Errorf("failed to find committers from history: %w", err)
	}

	committers = append(committers, committersFromHistory...)

	return eliminateDuplicatedCommitters(committers), nil
}

func findCommittersFromHistory() ([]model.Committer, error) {
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

	return parseCommitters(strings.Split(string(output), "\n")), nil
}

func findCommittersFromCommitMessage() ([]model.Committer, error) {
	commitMessageFilePath := config.AppConfig.CommitMessageFilePath
	existingMessage, err := os.ReadFile(commitMessageFilePath)
	if err != nil {
		return nil, fmt.Errorf("could not read commit file %s: %w", commitMessageFilePath, err)
	}

	var possibleCommitters []string
	for _, s := range strings.Split(string(existingMessage), "\n") {
		if strings.HasPrefix(s, "Co-authored-by:") {
			possibleCommitters = append(possibleCommitters, s)
		}
	}

	committers := parseCommitters(possibleCommitters)
	for i := 0; i < len(committers); i++ {
		(&committers[i]).ToggleSelect()
	}

	return committers, nil
}

func parseCommitters(committerLines []string) []model.Committer {
	committers := make([]model.Committer, 0, len(committerLines))
	for _, rawCommitter := range committerLines {
		rawCommitter = strings.TrimPrefix(rawCommitter, "Co-authored-by:")
		rawCommitter = strings.TrimSpace(rawCommitter)
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
