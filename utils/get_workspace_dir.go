package utils

import (
	"fmt"
	"os"
)

func GetWorkspaceDir() string {
	home, err := os.UserHomeDir()

	if err != nil {
		panic(err)
	}

	workspaceDir := fmt.Sprintf(`%s\OneDrive\Área de Trabalho`, home)

	_, err = os.Stat(workspaceDir)

	if err == nil {
		return workspaceDir
	}

	workspaceDir = fmt.Sprintf(`%s\Área de Trabalho`, home)

	_, err = os.Stat(workspaceDir)

	if err == nil {
		return workspaceDir
	}

	return "."
}
