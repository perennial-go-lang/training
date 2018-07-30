package main

import (
	"github.com/perennial-go-training/training/proto/practice/proto/simple"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"github.com/golang/protobuf/jsonpb"
	"log"
	"fmt"
	"github.com/perennial-go-training/training/proto/practice/proto/enum"
	"github.com/perennial-go-training/training/proto/practice/proto/complex"
)

func main() {
	var sm, sm1, sm2 simplepb.SimpleMessage
	sm = doSimple()

	readWriteBin(sm, sm1, "simpleMsg")

	readWriteJson(sm, sm2, "simpleMsg")

	em := doEnum()
	fmt.Println("Created Enum")
	fmt.Println(em)

	cm := doComplex()
	fmt.Println("Created Complex Message")
	fmt.Println(cm)
}

func doSimple() simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id: 12345,
		IsSimple: true,
		Name: "Hello, World!",
		SampleList: []int32{1,3,5,7,9},
	}

	return sm
}

func writeToFile(fName string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fName, out, 0644)
	if err != nil {
		return err
	}

	return nil
}

func readFromFile(fName string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fName)
	if err != nil {
		return err
	}

	err = proto.Unmarshal(in, pb)
	if err != nil {
		return err
	}

	return nil
}

func readWriteBin(sm, sm1 simplepb.SimpleMessage, fileName string) {
	fmt.Println("Created Message", sm)
	writeToFile(fileName + ".bin", &sm)
	fmt.Println("Data written!")
	readFromFile(fileName + ".bin", &sm1)
	fmt.Println("Read from File", sm1)
}

func toJson(fileName string, pb proto.Message) {
	marshaler := jsonpb.Marshaler{}
	json, err := marshaler.MarshalToString(pb)
	if err != nil {
		log.Fatal(err)
	}

	ioutil.WriteFile(fileName, []byte(json), 0644)
}

func fromJson(fileName string, pb proto.Message) {
	json, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	jsonpb.UnmarshalString(string(json), pb)
}

func readWriteJson(sm, sm2 simplepb.SimpleMessage, fileName string) {
	fmt.Println("Created Message", sm)
	toJson(fileName + ".json", &sm)
	fmt.Println("Data written!")
	fromJson(fileName + ".json", &sm2)
	fmt.Println("Read from File", sm2)
}

func doEnum() enumpb.EnumMessage {
	em := enumpb.EnumMessage{
		Id: 30,
		DayOfTheWeek:enumpb.DayOfTheWeek_THURSDAY,
	}

	return em
}

func doComplex() complexpb.ComplexMessage {
	cm := complexpb.ComplexMessage{
		OneDummy: &complexpb.DummyMessage{
			Id: 1,
			Name: "First Message",
		},
		MultipleDummy:[]*complexpb.DummyMessage{
			&complexpb.DummyMessage{
				Id: 1,
				Name: "First Message",
			},
			&complexpb.DummyMessage{
				Id: 1,
				Name: "First Message",
			},
		},
	}

	return cm
}