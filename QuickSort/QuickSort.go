package main

import (
	"time"
	"math/rand"
	"fmt"
)

func QuickSort(x []int,s int,t int)  {
	if s>=t {
		return
	}
	j:=s
	k:=t
	m:=x[s]
	for j<k {
		for j<k && x[k]<=m {
			k--
		}
		if j<k {
			x[j],x[k]=x[k],x[j]
		}
		for j<k && x[j]>=m {
			j++
		}
		if j<k {
			x[j],x[k]=x[k],x[j]
		}
	}
	QuickSort(x,s,k-1)
	QuickSort(x,k+1,t)
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
		QuickSort(x,0,n-1)
		fmt.Println(IsDecrease(x))
	}
}
