package tui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/erkanzileli/co-author/internal/model"
)

func getSelectedCommitters(items []list.Item) (result []model.Committer) {
	for _, i := range items {
		if c, ok := i.(*model.Committer); ok && c != nil && c.IsSelected() {
			result = append(result, *c)
		}
	}
	return
}

func ConvertCommittersToListItems(committers []model.Committer) (items []list.Item) {
	for i := 0; i < len(committers); i++ {
		items = append(items, &committers[i])
	}
	return
}
