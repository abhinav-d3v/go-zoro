package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/Masterminds/sprig"
)

const (
	Event = "event"
)

func AbiTypeResolver(solidityDataType string) string {
	switch solidityDataType {
	case "uint256":
		return "big.Int"
	case "address":
		return "common.Address"
	case "bytes32":
		return "[]byte"
	}

	return solidityDataType
}
func main() {
	fileContent, err := os.ReadFile("test.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var abiDump []Abi

	err = json.Unmarshal(fileContent, &abiDump)

	if err != nil {
		log.Fatal("Not Working", err.Error())
	}

	contractEvents := []Abi{}

	for _, obj := range abiDump {
		if obj.Type == Event {
			for i := 0; i < len(obj.Inputs); i++ {
				obj.Inputs[i].Type = AbiTypeResolver(obj.Inputs[i].Type)
			}
			contractEvents = append(contractEvents, obj)
		}
	}

	tmpl, err := template.New("struct.tpl").Funcs(sprig.FuncMap()).ParseFiles("struct.tpl")
	if err != nil {
		log.Fatal("Unable to fetch template file", err.Error())
	}

	err = tmpl.Execute(os.Stdout, contractEvents)
	if err != nil {
		log.Fatal("Unable to Execute", err.Error())
	}

}
