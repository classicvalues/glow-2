package ui

import (
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
)

type editorFinishedMsg struct{ err error }

func openEditor(path string) tea.Cmd {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vim"
	}
	cmd := exec.Command(editor, path)
	cb := func(err error) tea.Msg {
		return editorFinishedMsg{err}
	}
	return tea.ExecProcess(cmd, cb)
}
