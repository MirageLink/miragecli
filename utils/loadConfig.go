package utils

import (
	"fmt"
	"main/lang"
	"os"
	"path/filepath"
)

var ConfigFileContent string

func LoadConfig() {
	if !folderexists() {
		fmt.Println(lang.GetMessage("warn_no_conf_dir", Getfromconfig("lang")))
		createfolder()
	}
	loadfile()

}

func loadfile() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	filePath := filepath.Join(homeDir, ".mirage", "conf.uwu")
	_, err = os.Stat(filePath)
	if os.IsNotExist(err) {
		err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
		if err != nil {
			return err
		}

		err = os.WriteFile(filePath, []byte("lang uwu"), os.ModePerm)
		if err != nil {
			return err
		}

		ConfigFileContent = "test"
	} else {
		content, err := os.ReadFile(filePath)
		if err != nil {
			return err
		}
		if len(content) == 0 {
			ConfigFileContent = "test"

			err := os.WriteFile(filePath, []byte(ConfigFileContent), os.ModePerm)
			if err != nil {
				return err
			}
		} else {
			ConfigFileContent = string(content)
		}
	}

	return nil
}

func createfolder() {
	dirname, _ := os.UserHomeDir()
	os.Mkdir(dirname+"/.mirage", os.ModePerm)
}

func folderexists() bool {
	dirname, _ := os.UserHomeDir()
	info, _ := os.Stat(dirname + "/.mirage")
	if info == nil {
		return false
	}
	return true
}
