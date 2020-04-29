package main

import (
	"fmt"
	"io/ioutil"
	"log"
	complex "protobuf/src/complex"
	enumpb "protobuf/src/enum_example"
	example_simple "protobuf/src/simple"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func main() {

	m := doSimple()
	err := writeToFile("SimpleFile.bin", m)
	if err != nil {
		log.Fatalln("Error Occured While Creating File", err)
	}
	nsm := &example_simple.SimpleMessage{}
	err2 := readFromFile("SimpleFile.bin", nsm)
	if err2 != nil {
		log.Fatalln("Error Occured While Creating File", err2)
	}
	nsmString := toJSON(nsm)

	newpb := &example_simple.SimpleMessage{}
	fromJSON(nsmString, newpb)
	// doEnum()
	doComplex()
}
func doComplex() {
	cd := []*complex.DummyMessage{&complex.DummyMessage{Id: 1, Name: "sdcsdc"}, &complex.DummyMessage{Id: 2, Name: "vdvsdv5436"}, &complex.DummyMessage{Id: 345, Name: "34567bds"}}

	cm := complex.ComplexMessage{
		Id:            1,
		MultipleDummy: cd,
	}

	fmt.Println(cm, "=====================.")
}
func doEnum() {
	em := enumpb.EnumMessage{Id: 42, DayofWeek: enumpb.DayofWeek_MONDAY}
	fmt.Println(em)
}
func fromJSON(in string, pb proto.Message) {
	err := jsonpb.UnmarshalString(in, pb)
	if err != nil {
		log.Fatalln("Could not Marshal into protocol buffer", err)

	}

}
func doSimple() *example_simple.SimpleMessage {
	sm := example_simple.SimpleMessage{
		Id: 12345, IsSimple: true, Name: "My Simple Message", SampleList: []int32{1, 2, 3, 4, 5, 6, 7},
	}
	sm.Name = "Hey Theri"
	fmt.Println(sm.GetId())
	return &sm
}

func writeToFile(fname string, pb proto.Message) error {
	bs, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Error Can't Serialize to Bytres", err)
		return err
	}
	if err := ioutil.WriteFile(fname, bs, 0644); err != nil {
		log.Fatalln("Error Can;t Write to File", err)
		return err
	}
	fmt.Println("File Written ")
	return nil

}
func readFromFile(fname string, pb proto.Message) error {
	f, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Error Occured While Reading File", err)
		return err
	}
	err2 := proto.Unmarshal(f, pb)
	if err2 != nil {
		log.Fatalln("Con't Pass Bytes", err)
		return err2
	}
	return nil
}

func toJSON(pb proto.Message) string {
	marsheller := jsonpb.Marshaler{}
	out, err := marsheller.MarshalToString(pb)
	if err != nil {
		log.Fatalln("Error While Converting to Json", err)
		return ""
	}
	return out

}
