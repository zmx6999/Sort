package main

import (
	"math/rand"
	"time"
	"fmt"
)

type BTreeNode struct {
	Data int
	LeftChild *BTreeNode
	RightChild *BTreeNode
}

func CreateBTree(data int) *BTreeNode {
	return &BTreeNode{Data:data}
}

func (t *BTreeNode) Insert(data int)  {
	if t==nil {
		return
	}
	if data>t.Data {
		if t.LeftChild==nil {
			t.LeftChild=CreateBTree(data)
		} else {
			t.LeftChild.Insert(data)
		}
	} else {
		if t.RightChild==nil {
			t.RightChild=CreateBTree(data)
		} else {
			t.RightChild.Insert(data)
		}
	}
}

func (t *BTreeNode) TreeSort() []int {
	if t==nil {
		return []int{}
	}
	x:=t.LeftChild.TreeSort()
	x=append(x,t.Data)
	y:=t.RightChild.TreeSort()
	return append(x,y...)
}

func IsDecrease(x []int) bool {
	n:=len(x)
	for i:=0; i<n-1; i++ {
		if x[i]<x[i+1] {
			return false
		}
	}
	return true
}

func main()  {
	n:=10
	x:=make([]int,n)
	rand.Seed(time.Now().UnixNano())
	for i:=0; i<n; i++ {
		for j:=0; j<n; j++ {
			x[j]=rand.Intn(100)
		}
		tree:=CreateBTree(x[0])
		for j:=1; j<n; j++ {
			tree.Insert(x[j])
		}
		x=tree.TreeSort()
		fmt.Println(IsDecrease(x))
	}
}
