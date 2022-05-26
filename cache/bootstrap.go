package cache

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/Scrip7/nestjs-discord-utility-bot/commands"
	"github.com/bwmarrin/discordgo"
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

		// read file content
		fileContent, err := ioutil.ReadFile("./content/" + f.Name())
		if err != nil {
			logrus.Fatalf("error reading file: %v", err)
		}

		// convert content to string
		content := string(fileContent)
		// trim content
		content = strings.TrimSpace(content)

		name := f.Name()
		// extract file name without the extension
		name = name[:len(name)-len(ext)]

		// cache file content with a cost of 1
		Driver.Set(name, content, 1)
		if err != nil {
			logrus.Fatalf("error caching file: %v", err)
		}

		// wait for value to pass through buffers
		Driver.Wait()

		// trim content variable by max 95 characters
		if len(content) > 95 {
			// when the text is long, add "..." at the end of the string
			content = content[:95] + "..."
		}

		commands.Commands = append(commands.Commands, &discordgo.ApplicationCommand{
			Name: name,
			// set description to the content of the file trim by 100 characters
			Description: content,
		})

		logrus.WithField("cmd", name).Info("Content cached")
	}
}
