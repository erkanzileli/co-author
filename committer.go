package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
)

type committer struct {
	Name     string `yaml:"name,omitempty"`
	Email    string `yaml:"email,omitempty"`
	selected bool
}

func (c *committer) Title() string {
	if c.selected {
		return fmt.Sprintf("[X] %s", c.Name)
	}
	return fmt.Sprintf("[ ] %s", c.Name)
}

func (c *committer) Description() string {
	return c.Email
}

func (c *committer) FilterValue() string {
	return c.Name
}

func (c *committer) ToggleSelect() {
	c.selected = !c.selected
}

func (c *committer) UnSelect() {
	c.selected = false
}

func getSelectedCommitters(items []list.Item) (result []*committer) {
	for _, i := range items {
		if c, ok := i.(*committer); ok && c != nil && c.selected {
			result = append(result, c)
		}
	}
	return
}
