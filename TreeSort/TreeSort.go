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

func Create(data int) *BTreeNode {
	return &BTreeNode{Data:data}
}

func (t *BTreeNode) Insert(data int)  {
	if t==nil {
		return
	}
	if data<t.Data {
		if t.LeftChild==nil {
			t.LeftChild=Create(data)
		} else {
			t.LeftChild.Insert(data)
		}
	} else {
		if t.RightChild==nil {
			t.RightChild=Create(data)
		} else {
			t.RightChild.Insert(data)
		}
	}
}

func (t *BTreeNode) MidTravel() []int {
	if t==nil {
		return []int{}
	}
	x:=t.LeftChild.MidTravel()
	x=append(x,t.Data)
	y:=t.RightChild.MidTravel()
	x=append(x,y...)
	return x
}

func IsIncrease(x []int) bool {
	n:=len(x)
	for i:=0; i<n-1; i++ {
		if x[i]>x[i+1] {
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
		tree:=Create(x[0])
		for j:=1; j<n; j++ {
			tree.Insert(x[j])
		}
		x=tree.MidTravel()
		fmt.Println(IsIncrease(x))
	}
}
