package main

import (
	"math/rand"
	"time"
	"fmt"
)

func HeapSort(x []int,s int,m int)  {
	l:=s*2+1
	if l>=m {
		return
	}
	r:=l+1
	var min int
	if x[l]<x[r] {
		min=l
	} else {
		min=r
	}
	if x[s]<=x[min] {
		return
	}
	x[s],x[min]=x[min],x[s]
	HeapSort(x,min,m)
}

func HeapInit(x []int)  {
	n:=len(x)
	for i:=n/2; i>=0; i-- {
		HeapSort(x,i,n-1)
	}
	for i:=n-1; i>0; i-- {
		if x[0]>=x[i] {
			break
		}
		x[0],x[i]=x[i],x[0]
		HeapSort(x,0,i-1)
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
	n:=9
	x:=make([]int,n)
	rand.Seed(time.Now().UnixNano())
	for i:=0; i<n; i++ {
		for j:=0; j<n; j++ {
			x[j]=rand.Intn(100)
		}
		HeapInit(x)
		fmt.Println(IsDecrease(x))
	}
}
