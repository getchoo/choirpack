package corepack

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type CorepackVersion struct {
	Npm  string
	Pnpm string
	Yarn string
}

func GetCorepackVersions() CorepackVersion {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		log.Fatal(err)
	}

	fileLoc := filepath.Join(homeDir, ".cache", "node", "corepack", "lastKnownGood.json")
	file, err := os.ReadFile(fileLoc)

	if err != nil {
		log.Fatal(err)
	}

	var data CorepackVersion
	err = json.Unmarshal(file, &data)

	if err != nil {
		log.Fatal(err)
	}

	return data
}
