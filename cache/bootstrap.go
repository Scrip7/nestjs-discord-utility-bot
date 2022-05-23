package cache

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
)

func BootstrapContent() {
	files, err := ioutil.ReadDir("./content/")
	if err != nil {
		logrus.Fatal(err)
	}

	for _, f := range files {
		// Skip directories
		if f.IsDir() {
			continue
		}

		// Extract file extension
		ext := filepath.Ext(f.Name())
		// Ignore non .md files
		if ext != ".md" {
			continue
		}

		// TODO: refactor this

		// extract file name without the extension
		name := f.Name()
		name = name[:len(name)-len(ext)]
		// replace spaces with dash
		name = strings.Replace(name, " ", "-", -1)
		// convert to lowercase
		name = strings.ToLower(name)

		// read file content
		fileContent, err := ioutil.ReadFile("./content/" + f.Name())
		if err != nil {
			logrus.Fatalf("error reading file: %v", err)
		}

		// convert content to string
		content := string(fileContent)
		// trim content
		content = strings.TrimSpace(content)

		// cache file content with a cost of 1
		Driver.Set(name, content, 1)
		if err != nil {
			logrus.Fatalf("error caching file: %v", err)
		}

		// wait for value to pass through buffers
		Driver.Wait()

		logrus.WithField("cmd", name).Info("Content cached")
	}
}
