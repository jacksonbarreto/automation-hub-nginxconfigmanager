package entities

import (
	"fmt"
)

type Automation struct {
	Name    string `json:"name"`
	URLPath string `json:"urlPath"`
	Host    string `json:"hostname"`
	Port    int    `json:"port"`
}

func (a *Automation) Validate() error {
	if a.Name == "" {
		return fmt.Errorf("name is required")
	}
	if a.Host == "" {
		return fmt.Errorf("hostname is required")
	}
	if a.Port <= 0 || a.Port > 65535 {
		return fmt.Errorf("error: Port %d is not valid", a.Port)
	}
	return nil
}
