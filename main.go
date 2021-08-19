package main

import (
	"encoding/json"
	"fmt"

	"github.com/jinzhu/gorm"
	// operatorRepositoryDummy "portadapter/repository/dummy/operator"
	operatorRepository "portadapter/repository/mysql/operator"
	operatorService "portadapter/service/operator"
	structService "portadapter/struct/service"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// Config DB

	// init repository MySql
	db, _ := gorm.Open("mysql", "root:@tcp(localhost:3306)/portadapter?charset=utf8&parseTime=True&loc=Local")
	repository := operatorRepository.New(db)
	// init Service use repository mysql
	operatorService := operatorService.New(repository)

	// init repository Dummy
	// repositoryDummy := operatorRepositoryDummy.New()
	// init Service use repository dummy
	// operatorService := operatorService.New(repositoryDummy)

	// Run Service Method
	// CreateData
	err := operatorService.CreateData(structService.SaveOperator{
		Code:  "label12",
		Label: "name12",
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	// ReadData
	readData, err := operatorService.ReadData("1")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		b, _ := json.Marshal(readData)
		fmt.Println(string(b))
	}

	// UpdateData
	err = operatorService.UpdateData("1", structService.SaveOperator{
		Code:  "label12",
		Label: "name12",
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	// ListData
	listData, err := operatorService.ListData("name12")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		b, _ := json.Marshal(listData)
		fmt.Println(string(b))
	}
}
