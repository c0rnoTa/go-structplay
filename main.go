package main

import (
"fmt"
)

type myChan struct {
	id string
	value string
}

type myObject struct {
	channelsInfo []myChan
}

func (self *myObject) myFind(searchid string) *myChan {
	for i, savedChannel := range self.channelsInfo {
		if savedChannel.id == searchid {
			return &self.channelsInfo[i]
		}
	}
	return nil
}

func (self *myObject) myChange(searchid string, newvalue string){
	foundChan := self.myFind(searchid)
	if foundChan != nil {
		foundChan.value = newvalue
	} else {
		self.myAppend(searchid, newvalue)
	}
}

func (self *myObject) printAll() {
	for num, value := range self.channelsInfo {
		fmt.Println(num, value.id, value.value)
	}
	return
}

func (self *myObject) myAppend(id string, value string) {
	newChannel := myChan{id, value}
	self.channelsInfo = append(self.channelsInfo, newChannel)
}

func main() {
	testObject := new(myObject)
	fmt.Println("adding values")
	testObject.myAppend("a","anton")
	testObject.myAppend("b","burbon")
	testObject.myAppend("c","c0rn")
	testObject.printAll()
	fmt.Println("changing values")
	testObject.myChange("","c0rnoTa")
	testObject.printAll()
}

