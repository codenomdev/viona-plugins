package util

import (
	"embed"
	"fmt"

	"gopkg.in/yaml.v3"
)

type Info struct {
	SlugName string `yaml:"slug_name"`
	Type     string `yaml:"type"`
	Version  string `yaml:"version"`
	Author   string `yaml:"author"`
	Link     string `yaml:"link"`
}

func (c *Info) GetInfo(info embed.FS) *Info {
	yamlFile, err := info.ReadFile("info.yaml")
	if err != nil {
		fmt.Println(err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err)
	}
	return c
}
