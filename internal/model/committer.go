package model

import (
	"fmt"
)

type Committer struct {
	Name     string `yaml:"name,omitempty"`
	Email    string `yaml:"email,omitempty"`
	selected bool
}

func (c *Committer) Title() string {
	if c.selected {
		return fmt.Sprintf("[X] %s", c.Name)
	}
	return fmt.Sprintf("[ ] %s", c.Name)
}

func (c *Committer) Description() string {
	return c.Email
}

func (c *Committer) FilterValue() string {
	return c.Name
}

func (c *Committer) ToggleSelect() {
	c.selected = !c.selected
}

func (c *Committer) UnSelect() {
	c.selected = false
}

func (c *Committer) IsSelected() bool {
	return c.selected
}
