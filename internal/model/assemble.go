package model

import (
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

func AssembleSystem(opts ...AssemblyOption) (*System, error) {
	collector := &assemblyOptionsCollector{}
	for _, opt := range opts {
		opt(collector)
	}

	f := collector.systemFile
	if f == "" {
		dir := collector.directory
		if dir == "" {
			var err error
			dir, err = os.Getwd()
			if err != nil {
				return nil, err
			}
		}
		f = filepath.Join(dir, "system.yaml")
	}

	sys, err := readSystemYaml(f)
	if err != nil {
		return nil, err
	}

	dir := filepath.Dir(f)

	apps, err := traverseAndReadApps(dir)
	if err != nil {
		return nil, err
	}

	sys.Apps = apps
	return sys, nil
}

type assemblyOptionsCollector struct {
	directory  string
	systemFile string
}

type AssemblyOption func(*assemblyOptionsCollector)

func FromDirectory(directory string) AssemblyOption {
	return func(c *assemblyOptionsCollector) {
		c.directory = directory
	}
}

func WithSystemFile(systemFile string) AssemblyOption {
	return func(c *assemblyOptionsCollector) {
		c.systemFile = systemFile
	}
}

func readSystemYaml(systemFile string) (*System, error) {
	data, err := os.ReadFile(systemFile)
	if err != nil {
		return nil, err
	}

	var system System
	err = yaml.Unmarshal(data, &system)
	if err != nil {
		return nil, err
	}

	return &system, nil
}

func traverseAndReadApps(dir string) ([]App, error) {
	var apps []App
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && info.Name() == "app.yaml" {
			app, err := readAppYaml(path)
			if err != nil {
				return err
			}
			app.Directory = filepath.Dir(path)
			apps = append(apps, *app)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return apps, nil
}

func readAppYaml(path string) (*App, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var app App
	err = yaml.Unmarshal(data, &app)
	if err != nil {
		return nil, err
	}

	return &app, nil
}
