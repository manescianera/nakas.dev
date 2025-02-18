package main

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type CV struct {
	Name        string `yaml:"name"`
	Location    string `yaml:"location"`
	Email       string `yaml:"email"`
	Description string `yaml:"description"`
	Jobs        []Job  `yaml:"jobs"`
}

type Job struct {
	Title       string   `yaml:"title"`
	Company     string   `yaml:"company"`
	Start       string   `yaml:"start"`
	End         string   `yaml:"end"`
	Description string   `yaml:"description"`
	Skills      []string `yaml:"skills"`
}

// loads cv data from yaml into go struct
func loadData() (*CV, error) {
	file, err := os.Open("cv.yaml")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	cv := &CV{}
	if err = yaml.Unmarshal(data, cv); err != nil {
		return nil, err
	}
	return cv, nil
}
