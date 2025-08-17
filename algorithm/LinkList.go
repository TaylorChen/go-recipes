package algorithm

import (
	"fmt"
)

type Element interface{}

type linkNode struct {
	Data Element
	next *linkNode
}

func NewLinkNode() *linkNode {
	return &linkNode{Data: "head"}
}

func (head *linkNode) Add(data Element) {
	point := head

	for point.next != nil {
		point = point.next
	}

	newNode := linkNode{Data: data}
	point.next = &newNode
}

func (head *linkNode) Delete(index int) Element {
	if index < 0 || index > head.GetLength() {
		fmt.Println("please check index")
		return fmt.Errorf("check index error")
	}

	point := head

	for i := 0; i < index; i++ {
		point = point.next
	}

	point.next = point.next.next
	data := point.next.Data
	return data
}

func (head *linkNode) Insert(index int, data Element) {
	if index < 0 || index > head.GetLength() {
		fmt.Println("please check index")
	} else {
		point := head
		for i := 0; i < index-1; i++ {

			point = point.next
		}
		newNode := &linkNode{Data: data}
		newNode.next = point.next
		point.next = newNode
	}
}

func (head *linkNode) Search(data Element) int {

	point := head

	index := 0

	for point.next != nil {
		if point.Data == data {
			return index
		} else {
			index++
			point = point.next
			if index > head.GetLength()-1 {
				fmt.Println(data, "no data exist")
				break
			}
			continue
		}
	}

	if point.Data == data {
		return index
	}
	return -1
}

func (head *linkNode) GetData(index int) Element {

	point := head

	if index < 0 || index > head.GetLength() {
		return fmt.Errorf("please check index error")
	}

	for i := 0; i < index-1; i++ {
		point = point.next
	}

	return point.Data
}

func (head *linkNode) GetAll() []Element {
	dataList := make([]Element, 0, head.GetLength())
	point := head
	for point.next != nil {
		dataList = append(dataList, point.Data)
		point = point.next
	}
	dataList = append(dataList, point.Data)
	return dataList
}

func (head *linkNode) GetLength() int {
	point := head
	var length int
	for point.next != nil {
		length++
		point = point.next
	}
	length++
	return length
}

// func main() {}
