package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
)

type committer struct {
	name     string
	email    string
	selected bool
}

func (c *committer) Title() string {
	if c.selected {
		return fmt.Sprintf("[X] %s", c.name)
	}
	return fmt.Sprintf("[ ] %s", c.name)
}

func (c *committer) Description() string {
	return c.email
}

func (c *committer) FilterValue() string {
	return c.name
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
