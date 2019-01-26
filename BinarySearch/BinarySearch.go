package main

import (
	"math/rand"
	"time"
	"fmt"
)

func Find(x []int,y int) int {
	n:=len(x)
	start:=0
	end:=n-1
	for start<=end {
		mid:=(start+end)/2
		if y==x[mid] {
			return mid
		} else if y<x[mid] {
			start=mid+1
		} else {
			end=mid-1
		}
	}
	return -1
}

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

func main()  {
	n:=10
	x:=make([]int,n)
	rand.Seed(time.Now().UnixNano())
	for i:=0; i<n; i++ {
		for j:=0; j<n; j++ {
			x[j]=rand.Intn(100)
		}
		QuickSort(x,0,n-1)
		for j:=0; j<n; j++ {
			fmt.Print(Find(x,x[j])," ")
		}
		fmt.Println(Find(x,-100),Find(x,100))
	}
}
