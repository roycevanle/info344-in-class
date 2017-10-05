package models

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

//Declare types so you have control over it later on.
type Zip struct {
	Code  string
	City  string
	State string
}

type ZipSlice []*Zip

type ZipIndex map[string]ZipSlice

//File Path
func LoadZips(fileName string) (ZipSlice, error) {
	f, err := os.Open(fileName)
	if err != nil { //make a call that might fail so standard notation
		return nil, fmt.Errorf("error opening file: %v", err)
	}

	//Reader allows us to process 1 line at a time
	reader := csv.NewReader(f)
	//The underscore means we don't care about the value
	_, err = reader.Read()
	if err != nil {
		return nil, fmt.Errorf("error reading header row: %v", err)
	}

	zips := make(ZipSlice, 0, 43000) //Pre allocating array to 43k
	for {
		//io.eof is the error that signifies end of iteration
		fields, err := reader.Read()
		if err == io.EOF {
			return zips, nil
		}
		if err != nil {
			return nil, fmt.Errorf("error reading record: %v", err)
		}
		z := &Zip{
			Code:  fields[0],
			City:  fields[3],
			State: fields[6],
		}
		zips = append(zips, z)
	}
}
