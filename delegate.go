package main

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	itemSelectMessageFormat   = "Selected %s"
	itemUnSelectMessageFormat = "Unselected %s"
)

func newItemDelegate(keys *delegateKeyMap) list.DefaultDelegate {
	d := list.NewDefaultDelegate()

	d.UpdateFunc = func(msg tea.Msg, m *list.Model) tea.Cmd {
		item, ok := m.SelectedItem().(*committer)
		if !ok {
			return nil
		}

		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, keys.choose):
				item.ToggleSelect()

				var message string
				if item.selected {
					message = fmt.Sprintf(itemSelectMessageFormat, item.Name)
				} else {
					message = fmt.Sprintf(itemUnSelectMessageFormat, item.Name)
				}

				return m.NewStatusMessage(statusMessageStyle(message))

			case key.Matches(msg, keys.complete):
				selectedCommitters := getSelectedCommitters(m.Items())
				if err := writeToCommitMsgFile(prepareCommitMsg(selectedCommitters)); err != nil {
					log.Fatal(err)
				}
				return tea.Quit

			case key.Matches(msg, keys.reset):
				for _, i := range m.Items() {
					i.(*committer).UnSelect()
				}
				return m.NewStatusMessage(statusMessageStyle("Cleaned all selections"))
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

type delegateKeyMap struct {
	choose   key.Binding
	complete key.Binding
	reset    key.Binding
}

func (d *delegateKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		d.choose,
		d.complete,
		d.reset,
	}
}

func (d *delegateKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			d.choose,
			d.complete,
			d.reset,
		},
	}
}

func newDelegateKeyMap() *delegateKeyMap {
	return &delegateKeyMap{
		choose: key.NewBinding(
			key.WithKeys(" "),
			key.WithHelp("space", "choose"),
		),
		complete: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "complete"),
		),
		reset: key.NewBinding(
			key.WithKeys("r"),
			key.WithHelp("r", "reset"),
		),
	}
}
