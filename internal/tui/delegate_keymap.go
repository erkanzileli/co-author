package tui

import "github.com/charmbracelet/bubbles/key"

type DelegateKeyMap struct {
	choose   key.Binding
	complete key.Binding
	reset    key.Binding
}

func NewDelegateKeyMap() *DelegateKeyMap {
	return &DelegateKeyMap{
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

func (d *DelegateKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		d.choose,
		d.complete,
		d.reset,
	}
}

func (d *DelegateKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			d.choose,
			d.complete,
			d.reset,
		},
	}
}
