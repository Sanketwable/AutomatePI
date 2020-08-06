package getauto

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// GetModels is a func
func GetModels() {
	file, err := ioutil.ReadFile("json.auto")
	if err != nil {
		log.Fatal("error occured wile reading file ", err)
	}

	fileString := string(file)
	NumberOfModels := count(file)
	// fmt.Println(NumberOfModels)
	// fmt.Println(fileString)

	Models := make([][]string, NumberOfModels)

	ModelsArray := strings.Fields(fileString)
	ModelsEndpoint := ModelsArray[1]

	fmt.Println("Endpoint : ", ModelsEndpoint)

	modelstart := modelStart(ModelsArray)

	for i := 0; i < NumberOfModels; i++ {
		Models[i] = append(Models[i], string(ModelsArray[3*i+modelstart]))
		Models[i] = append(Models[i], string(ModelsArray[3*i+1+modelstart]))
		Models[i] = append(Models[i], string(ModelsArray[3*i+2+modelstart]))
	}
	fmt.Println("models are: ", Models)

}

func count(file []byte) (no int) {
	no = 0
	for _, i := range file {
		if string(i) == ";" {
			no++
		}
	}
	return
}

func modelStart(file []string) int {

	for i, str := range file {
		if string(str) == "Models:" {
			return i
		}
	}
	return 0
}
