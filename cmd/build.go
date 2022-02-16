package main

import (
	"bytes"
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("git", "rev-list", "-1", "HEAD")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}
	version := out.String()[:len(out.String())-1]
	println("Bot version:", version)

	println("Downloading go.mod...")
	if err := exec.Command("go", "mod", "download").Run(); err != nil {
		log.Fatalln(err)
	}
	println("Building bot...")
	cmd = exec.Command("go", "build", "-ldflags", "-X main.version="+version, "-o", "build/bot", "cmd/bot/main.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}
	println("Finished building...")
}
