package main

import (
	"errors"
	"fmt"
	"strings"
)

type ArgValidator struct {
}

func (a *ArgValidator) ModulePath(path string) error {
	if path == "" {
		return errors.New("please provide absolute path to the module e.g. /var/www/magento/web/vendor/magento/module-catalog/")
	}

	if !DirExists(path) {
		return fmt.Errorf("double check path is correct? %s", path)
	}

	if !FileExists(path + "etc/module.xml") {
		return errors.New("could not find valid module.xml")
	}

	return nil
}

func (a *ArgValidator) OutputFile(path string) error {
	if path == "" {
		return errors.New("please provide an output file name e.g. MAGE_DOCS_README.md")
	}

	if !strings.HasSuffix(path, ".md") {
		return errors.New("output file extension should be .md")
	}

	return nil
}
