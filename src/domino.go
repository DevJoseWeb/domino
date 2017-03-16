package domino

import (
	"fmt"
)

const VERSION = "0.1"

func NewDomino() Learner {
	var L Learner

	fmt.Println("Domino! version", VERSION)

	return L
}

func (l Learner) SetModel(m interface{}) {
	l.Type = m
}

func (l Learner) GetModel() interface{} {
	return l.Type
}

func (l Learner) LoadData() {

}

func (l Learner) SplitData() {

}
