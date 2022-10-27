package cmd

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestSerializeDeserialize(t *testing.T) {
	file, err := os.Create("./test_data.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	data := Records{{"Jason", "Mukim Gunung", "014534534535", "23424234"}, {"John", "Butuan City", "0495358457", "34435345"}}

	err = serialize(data, file)
	if err != nil {
		fmt.Println(err)
		return
	}

	bb := make([]byte, 2)
	n, err := file.Read(bb)
	if err == io.EOF {
		fmt.Println("EOF")
	} else if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n, "bytes read")

	file2, err := os.Open("./test_data.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	var data2 = Records{}

	err = deserialize(&data2, file2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data2)
}

func TestReadWriteJSONFile(t *testing.T) {
	records = Records{{"Jason", "Mukim Gunung", "012923485578", "2348394348"}, {"John", "Butuan City", "0934839489", "348392428394"}}
	err := writeJSONFile("./test_read_write.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	records = Records{}
	err = readJSONFile("./test_read_write.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(records)

}

func TestCreateRecord(t *testing.T) {
	record := createRecord("Jason", "Mukim Gunung", "016343478934")
	fmt.Println(*record)
}
