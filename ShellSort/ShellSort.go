package main

import (
	"math/rand"
	"time"
	"fmt"
)

func ShellSort(x []int)  {
	n:=len(x)
	for i:=n/2; i>0; i/=2 {
		for j:=i; j<n; j++ {
			tmp:=x[j]
			for k:=j-i; k>=0; k-=i {
				if tmp>x[k] {
					x[k],x[k+i]=x[k+i],x[k]
				}
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
		ShellSort(x)
		fmt.Println(IsDecrease(x))
	}
}
