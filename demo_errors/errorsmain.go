package main

import (
	"fmt"
	"io/ioutil"

	"github.com/pkg/errors"
)

func getdata(filename string) ([]byte, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		err = errors.Wrap(err, "read failed")
	}
	return data, err
}

func calculate() ([]byte, error) {
	data, err := getdata("/tmp/frobniz")
	if err != nil {
		err = errors.Wrap(err, "calc failed")
	}
	return data, err
}

func main() {
	data, err := calculate()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Data:", data)

}
