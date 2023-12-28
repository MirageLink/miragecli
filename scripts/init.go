package scripts

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var Scripts []string

var homeDir, _ = os.UserHomeDir()

func InitScripts() {
	getScriptList()
}

func ExecuteScript(name string) {
	content := getScriptContent(name)
	shell := strings.Split(content, "\n")[0]

	cmd := exec.Command("pwsh", "-c", strings.Trim(content, shell))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Start()
	if err != nil {
		fmt.Println("Error starting command:", err)
		return
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Println("Error waiting for command:", err)
		return
	}
}

func getScriptContent(name string) string {
	scriptLocation := filepath.Join(homeDir, ".mirage", "scripts", name+".uwuscript")
	content, _ := os.ReadFile(scriptLocation)
	return string(content)
}

func getScriptList() {

	filePath := filepath.Join(homeDir, ".mirage", "scripts")
	dir, _ := os.Open(filePath)
	if dir == nil {
		os.Mkdir(filePath, os.ModePerm)
	}
	fileScripts, _ := dir.ReadDir(-1)
	for _, i := range fileScripts {
		if strings.HasSuffix(i.Name(), ".uwuscript") && !i.IsDir() {
			Scripts = append(Scripts, strings.TrimSuffix(i.Name(), ".uwuscript"))
		}
	}
}
