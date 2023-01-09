package file

import (
	"crypto/md5"
	"encoding/hex"
	"fileAutomation/app/entity/connection"
	"fileAutomation/app/entity/model"
	"fileAutomation/app/entity/repository/FileRepository"
	"fmt"
	"github.com/golang-module/carbon/v2"
	"github.com/gookit/config/v2"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

/**
Support Function
*/

func parseData(fileName string) (string, int, string, error) {
	data := imgData(fileName)
	title := strings.Title(strings.ReplaceAll(data[0], "_", ""))
	price, err := strconv.Atoi(strings.ReplaceAll(data[1], "_", ""))
	description := data[2]
	desc := ""

	if description == "null" {
		desc = "Tidak ada deskripsi"
	} else {
		desc = data[2]
	}

	if err != nil {
		return "", 0, "", err
	}

	return title, price, desc, nil

}

func imgData(fileName string) []string {
	resArr := strings.SplitAfter(fileName, "_")
	return resArr
}

/**
Main Function
*/

func IterateFile() {

	currentPath := config.String("current_path")
	destinationPath := config.String("destination_path")

	dbMysql := connection.MysqlConnect()
	fileRepository := FileRepository.NewFileRepository(dbMysql)

	rootPath, err := os.Getwd()

	if err != nil {
		panic("Error")
	}

	fullPath := rootPath + currentPath

	files, err := ioutil.ReadDir(fullPath)

	if err != nil {
		panic(err)
	}

	for _, f := range files {

		h := md5.New()
		h.Write([]byte(f.Name()))
		hashedProductImg := hex.EncodeToString(h.Sum(nil)) + ".jpg"

		if err != nil {
			panic("Error on hashing")
		}

		nameProduct := strings.Replace(f.Name(), ".jpg", "", -1)

		title, price, description, errParse := parseData(nameProduct)

		if errParse != nil {
			panic("Error on parsing image name")
		}

		pathName := "assets/images/product/" + string(hashedProductImg)

		moveFilesErr := os.Rename(fullPath+"/"+f.Name(), destinationPath+hashedProductImg)

		if moveFilesErr != nil {
			panic(moveFilesErr)
		}

		var testingModel = model.ProductModels{
			CategoryId:         4,
			ProductName:        title,
			ProductDescription: description,
			ProductPrice:       int64(price),
			ProductImage:       pathName,
			CreatedAt:          carbon.Carbon{},
			UpdatedAt:          carbon.Carbon{},
		}
		msg, err := fileRepository.InsertFileInterface(&testingModel)

		if err != nil {
			panic(msg)
		}
	}

	fmt.Println("OK")

}
