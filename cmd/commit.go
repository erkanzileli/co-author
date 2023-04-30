package cmd

import (
	"fmt"
	"github.com/erkanzileli/co-author/internal/git"
	"github.com/erkanzileli/co-author/internal/tui"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var (
	commitCmd = &cobra.Command{
		Use:     "commit",
		Short:   "start selecting committer from your repository and produce commit message",
		Example: "co-author commit $COMMIT_MSG_FILE",
		Run: func(cmd *cobra.Command, args []string) {
			if err := tea.NewProgram(newModel()).Start(); err != nil {
				fmt.Println("Error running program:", err)
				os.Exit(1)
			}
		},
	}
)

type listKeyMap struct {
	toggleSpinner  key.Binding
	toggleHelpMenu key.Binding
}

func newListKeyMap() *listKeyMap {
	return &listKeyMap{
		toggleSpinner: key.NewBinding(
			key.WithKeys("s"),
			key.WithHelp("s", "toggle spinner"),
		),
		toggleHelpMenu: key.NewBinding(
			key.WithKeys("H"),
			key.WithHelp("H", "toggle help"),
		),
	}
}

type model struct {
	list         list.Model
	keys         *listKeyMap
	delegateKeys *tui.DelegateKeyMap
}

func newModel() *model {
	var (
		delegateKeys = tui.NewDelegateKeyMap()
		listKeys     = newListKeyMap()
	)

	// Make initial list of items
	committers, err := git.FindCommitters()
	if err != nil {
		log.Fatal(err)
	}

	// Setup list
	delegate := tui.NewItemDelegate(delegateKeys)
	committerList := list.New(tui.ConvertCommittersToListItems(committers), delegate, 0, 0)
	committerList.Title = "Committers"
	committerList.Styles.Title = tui.TitleStyle
	committerList.AdditionalFullHelpKeys = func() []key.Binding {
		return []key.Binding{
			listKeys.toggleSpinner,
			listKeys.toggleHelpMenu,
		}
	}
	return &model{
		list:         committerList,
		keys:         listKeys,
		delegateKeys: delegateKeys,
	}
}

func (m *model) Init() tea.Cmd {
	return tea.EnterAltScreen
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		topGap, rightGap, bottomGap, leftGap := tui.AppStyle.GetPadding()
		m.list.SetSize(msg.Width-leftGap-rightGap, msg.Height-topGap-bottomGap)

	case tea.KeyMsg:
		// Don't match any of the keys below if we're actively filtering.
		if m.list.FilterState() == list.Filtering {
			break
		}

		switch {
		case key.Matches(msg, m.keys.toggleSpinner):
			cmd := m.list.ToggleSpinner()
			return m, cmd

		case key.Matches(msg, m.keys.toggleHelpMenu):
			m.list.SetShowHelp(!m.list.ShowHelp())
			return m, nil
		}
	}

	// This will also call our delegate's update function.
	newListModel, cmd := m.list.Update(msg)
	m.list = newListModel
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m *model) View() string {
	return tui.AppStyle.Render(m.list.View())
}
