package main

import (
	"math/rand"
	"time"
	"fmt"
)

func BubbleSort(x []int)  {
	n:=len(x)
	for i:=1; i<n; i++ {
		for j:=0; j<n-i; j++ {
			if x[j]<x[j+1] {
				x[j],x[j+1]=x[j+1],x[j]
			}
		}
	}
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
		BubbleSort(x)
		fmt.Println(IsDecrease(x))
	}
}
