package main

import (
	"net/http"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const url = "https://charm.sh/"

type model struct {
	status int
	err    error
}

type statusMsg int

type errMsg struct{ err error }

// Init implements tea.Model.
func (model) Init() tea.Cmd {
	return checkServer
}

func checkServer() tea.Msg {
	c := &http.Client{Timeout: 10 * time.Second}
	res, err := c.Get(url)

	if err != nil {
		return errMsg{err}
	}
	return statusMsg(res.StatusCode)
}

// Update implements tea.Model.
func (model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	panic("unimplemented")
}

// View implements tea.Model.
func (model) View() string {
	panic("unimplemented")
}

func newModel() tea.Model {
	return model{}
}

// For messages that contain errors it's often handy to also implement the
// error interface on the message.
func (e errMsg) Error() string { return e.err.Error() }

func main() {

}
