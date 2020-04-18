package main

import (
	"fmt"
	"github.com/go-yaml/yaml"
	jsoniter "github.com/json-iterator/go"
	"github.com/tidwall/pretty"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var inputFile string
	var err error
	if len(os.Args) < 2 {
		fmt.Println("please provide a yaml file")
		return
	}

	inputFile = os.Args[1]
	YamlFile, err := os.Open(inputFile)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("File： %s does not exists\n", inputFile)
		} else {
			fmt.Println(err.Error())
		}
		return
	}

	var inputYaml []byte
	inputYaml, _ = ioutil.ReadAll(YamlFile)
	outputJosn := make(map[interface{}]interface{})
	err = yaml.Unmarshal(inputYaml, &outputJosn)
	if err != nil {
		log.Fatalln(err.Error())
	}

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	jsonString, err := json.MarshalIndent(&outputJosn, "", "")
	if err != nil {
		log.Fatalln("encode data to json error")
	}
	jsonString = pretty.Pretty(jsonString)
	fmt.Println(string(jsonString))
}
