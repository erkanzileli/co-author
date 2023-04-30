package git

import (
	"fmt"
	"github.com/erkanzileli/co-author/internal/config"
	"github.com/erkanzileli/co-author/internal/model"
	"os"
	"strings"
)

func SaveSelectedCommittersToMessage(selectedCommitters []model.Committer) error {
	commitMessageFilePath := config.AppConfig.CommitMessageFilePath
	existingMessageByteContent, err := os.ReadFile(commitMessageFilePath)
	if err != nil {
		return fmt.Errorf("could not read commit file %s: %w", commitMessageFilePath, err)
	}

	messageWithoutCoAuthors := clearAllCoAuthorsFromMessage(string(existingMessageByteContent))

	finalCommitMessage := addCoAuthorsToMessage(selectedCommitters, messageWithoutCoAuthors)

	if err = os.WriteFile(commitMessageFilePath, []byte(finalCommitMessage), 0644); err != nil {
		return fmt.Errorf("could update commit file %s: %w", commitMessageFilePath, err)
	}

	return nil
}

func addCoAuthorsToMessage(selectedCommitters []model.Committer, messageWithoutCoAuthors string) string {
	if len(selectedCommitters) == 0 {
		return messageWithoutCoAuthors
	}

	coAuthoredByLines := prepareCoAuthoredByLinesForCommitters(selectedCommitters)
	var finalCommitMessage string
	if config.AppConfig.CommitSource == "message" || len(config.AppConfig.CommitSHA1) > 0 {
		messageWithoutCoAuthorsLines := strings.Split(messageWithoutCoAuthors, "\n")
		for i, s := range messageWithoutCoAuthorsLines {
			if strings.HasPrefix(s, "#") {
				finalCommitMessage += coAuthoredByLines + "\n"
				finalCommitMessage += strings.Join(messageWithoutCoAuthorsLines[i:], "\n")
				break
			}
			finalCommitMessage += s + "\n"
			if i == len(messageWithoutCoAuthorsLines)-1 {
				finalCommitMessage += coAuthoredByLines
			}
		}
	} else {
		finalCommitMessage = fmt.Sprintf("\n\n%s%s", coAuthoredByLines, messageWithoutCoAuthors)
	}

	return finalCommitMessage
}

func eliminateDuplicatedCommitters(s []model.Committer) []model.Committer {
	m := make(map[string]struct{})
	result := make([]model.Committer, 0, len(s))
	for i := 0; i < len(s); i++ {
		elem := s[i]
		key := elem.Name + elem.Email
		if _, exist := m[key]; !exist {
			m[key] = struct{}{}
			result = append(result, elem)
		}
	}
	return result
}

func clearAllCoAuthorsFromMessage(message string) string {
	var clearedMessage string
	for _, line := range strings.Split(message, "\n") {
		if !strings.HasPrefix(line, "Co-authored-by:") {
			clearedMessage += line + "\n"
		}
	}
	return clearedMessage
}

func prepareCoAuthoredByLinesForCommitters(committers []model.Committer) (msg string) {
	lines := make([]string, 0, len(committers))
	for _, c := range committers {
		lines = append(lines, fmt.Sprintf("Co-authored-by: %s <%s>", c.Name, c.Email))
	}
	return strings.Join(lines, "\n")
}
