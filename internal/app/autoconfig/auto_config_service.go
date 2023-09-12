package autoconfig

import (
	"automation-hub-nginxconfigmanager/internal/app/config"
	"automation-hub-nginxconfigmanager/internal/app/dto"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
)

type ConfigAction int

const (
	Add ConfigAction = iota
	Remove
	Update
)

func manageConfig(action ConfigAction, auto dto.Automation) error {
	var err error

	switch action {
	case Add:
		err = addConfig(auto)
	case Remove:
		err = removeConfig(auto.Name)
	case Update:
		err = updateConfig(auto)
	default:
		return fmt.Errorf("invalid action")
	}

	if err != nil {
		return err
	}

	return reloadNginx()
}

func addConfig(auto dto.Automation) error {
	filePath := filepath.Join(config.Config.ConfigDir, auto.Name+".conf")
	tmpl, err := template.New("config").Parse(configTemplate)
	if err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	return tmpl.Execute(file, auto)
}

func removeConfig(name string) error {
	filePath := filepath.Join(config.Config.ConfigDir, name+".conf")
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return fmt.Errorf("config %s not found", name)
	}
	return os.Remove(filePath)
}

func updateConfig(auto dto.Automation) error {
	if err := removeConfig(auto.Name); err != nil {
		return err
	}
	return addConfig(auto)
}

func reloadNginx() error {
	cmd := exec.Command("docker", "exec", config.Config.NginxContainer, "nginx", "-s", "reload")
	return cmd.Run()
}
