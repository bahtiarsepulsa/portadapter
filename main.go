package main

import (
	"encoding/json"
	"fmt"

	operatorRepositoryMock "portadapter/repository/mock/operator"
	operatorService "portadapter/service/operator"
	structService "portadapter/struct/service"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// Config DB

	// init repository
	// db, _ := gorm.Open("mysql", "root:@tcp(localhost:3306)/portadapter?charset=utf8&parseTime=True&loc=Local")
	// repository := operatorRepository.New(db)
	repositoryMock := operatorRepositoryMock.New()

	// init Service
	// operatorService := operatorService.New(repository)
	operatorService := operatorService.New(repositoryMock)

	// Run Method
	err := operatorService.CreateData(structService.SaveOperator{
		Code:  "label12",
		Label: "name12",
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	readData, err := operatorService.ReadData("1")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		b, _ := json.Marshal(readData)
		fmt.Println(string(b))
	}

	err = operatorService.UpdateData("1", structService.SaveOperator{
		Code:  "label12",
		Label: "name12",
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	listData, err := operatorService.ListData("name12")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		b, _ := json.Marshal(listData)
		fmt.Println(string(b))
	}
}
