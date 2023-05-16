package packagejson

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/buger/jsonparser"
)

func ModifyPackageJson(packageManager string, version string) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	data, err := os.ReadFile(filepath.Join(cwd, "package.json"))
	if err != nil {
		log.Fatal(err)
	}

	packageManagerData := fmt.Sprintf("\"%v@%v\"", packageManager, version)

	data, err = jsonparser.Set(data, []byte(packageManagerData), "packageManager")
	if err != nil {
		log.Fatal(err)
	}

	os.WriteFile(filepath.Join(cwd, "package.json"), prettifyJSON(data), 0644)
}
