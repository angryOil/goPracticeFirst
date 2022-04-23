package main

import "fmt"

type DB interface {
	Get()
	Set()
}

type CDB struct {
}

func (c CDB) GetData() {

}
func (c CDB) SetData() {

}

type wrapper struct {
	cdb CDB
}

func (w wrapper) Get() {
	w.cdb.GetData()
}
func (w wrapper) Set() {
	w.cdb.SetData()
}

func main() {
	var cDatabase CDB
	var database DB
	database = wrapper{cDatabase}
	fmt.Println(database)
}
