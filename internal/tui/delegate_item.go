package tui

import (
	"fmt"
	"github.com/erkanzileli/co-author/internal/git"
	"github.com/erkanzileli/co-author/internal/model"
	"log"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	itemSelectMessageFormat   = "Selected %s"
	itemUnSelectMessageFormat = "Unselected %s"
)

func NewItemDelegate(keys *DelegateKeyMap) list.DefaultDelegate {
	d := list.NewDefaultDelegate()

	d.UpdateFunc = func(msg tea.Msg, m *list.Model) tea.Cmd {
		item, ok := m.SelectedItem().(*model.Committer)
		if !ok {
			return nil
		}

		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, keys.choose):
				item.ToggleSelect()

				var message string
				if item.IsSelected() {
					message = fmt.Sprintf(itemSelectMessageFormat, item.Name)
				} else {
					message = fmt.Sprintf(itemUnSelectMessageFormat, item.Name)
				}

				return m.NewStatusMessage(StatusMessageStyle(message))

			case key.Matches(msg, keys.complete):
				selectedCommitters := getSelectedCommitters(m.Items())
				if err := git.SaveSelectedCommittersToMessage(selectedCommitters); err != nil {
					log.Fatal(err)
				}
				return tea.Quit

			case key.Matches(msg, keys.reset):
				for _, i := range m.Items() {
					i.(*model.Committer).UnSelect()
				}
				return m.NewStatusMessage(StatusMessageStyle("Cleaned all selections"))
			}
		}

		return nil
	}

	help := []key.Binding{keys.choose, keys.complete, keys.reset}

	d.ShortHelpFunc = func() []key.Binding {
		return help
	}

	d.FullHelpFunc = func() [][]key.Binding {
		return [][]key.Binding{help}
	}

	return d
}
