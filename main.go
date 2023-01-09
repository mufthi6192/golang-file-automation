package main

import (
	"fileAutomation/app/helper"
	"fileAutomation/app/services/file"
)

func main() {
	helper.YamlConverterHelper()
	file.IterateFile()
}
