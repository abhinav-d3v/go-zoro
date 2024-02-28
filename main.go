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
	Event        = "event"
	FetchLogData = `decodeLog, err := contractABI.Unpack(event.Name, vLog.Data)
		if err != nil {
			log.Fatal("unable to decode log")
		}`
)

func AbiTypeResolver(solidityDataType string) string {
	switch solidityDataType {
	case "uint256":
		return "*big.Int"
	case "address":
		return "common.Address"
	case "bytes32":
		return "[]byte"
	}

	return solidityDataType
}
func CreateIndxedType(golangDataType string) string {
	switch golangDataType {
	case "common.Address":
		return "common.HexToAddress"
	case "*big.Int":
		return "new(big.Int).SetBytes"
	}
	return ""

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
	/*

	  {
	    "anonymous": false,
	    "inputs": [
	      {
	        "indexed": true,
	        "internalType": "bytes32",
	        "name": "role",
	        "type": "bytes32"
	      },
	      {
	        "indexed": true,
	        "internalType": "address",
	        "name": "account",
	        "type": "address"
	      },
	      {
	        "indexed": true,
	        "internalType": "address",
	        "name": "sender",
	        "type": "address"
	      }
	    ],
	    "name": "RoleGranted",
	    "type": "event"
	  },
	*/
	renderEvent := []RenderEvent{}
	for _, obj := range abiDump {
		if obj.Type == Event {
			logDataIndex := 0
			topicIndex := 1
			event := RenderEvent{Name: obj.Name, Inputs: []RenderInput{}}
			for _, contractEvent := range obj.Inputs {
				eventInput := RenderInput{
					Name: contractEvent.Name,
					Type: AbiTypeResolver(contractEvent.Type),
				}
				if contractEvent.Indexed {
					createArg := CreateIndxedType(eventInput.Type)
					eventInput.FetchFrom = fmt.Sprintf("vLog.Topics[%d]", topicIndex)
					topicIndex++
					eventInput.InitValue = fmt.Sprintf("%s(%s)", createArg, eventInput.FetchFrom)
				} else {

					createArg := fmt.Sprintf("decodeLog[%d].(%s)", logDataIndex, eventInput.Type)
					logDataIndex++
					eventInput.InitValue = createArg
				}
				event.Inputs = append(event.Inputs, eventInput)
			}
			if logDataIndex > 0 {
				event.IsFetchLogData = true
			}
			renderEvent = append(renderEvent, event)
		}
	}
	tmpl, err := template.New("struct.tpl").Funcs(sprig.FuncMap()).ParseFiles("struct.tpl")

	if err != nil {
		log.Fatal("Unable to fetch template file", err.Error())
	}

	err = tmpl.Execute(os.Stdout, renderEvent)

	if err != nil {
		log.Fatal("Unable to Execute", err.Error())
	}
}
