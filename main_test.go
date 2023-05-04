package main

import (
	"fmt"
	"math"
	"testing"
)

func TestNM(t *testing.T) {
	n := 0
	m := 0
	fmt.Scan(&n, &m)
	ns := make([]int, 0, n)
	ms := make([]int, 0, m)
	for j := 0; j < n; j++ {
		x := 0
		fmt.Scan(&x)
		ns = append(ns, x)
	}
	for j := 0; j < m; j++ {
		x := 0
		fmt.Scan(&x)
		ms = append(ms, x)
	}
	nSum := 0
	for i := range ns {
		nSum += ns[i]
	}
	stp := make(map[int][]int)
	var mSum int
	for i := range ms {
		mSum = 0
		for j := i; j < len(ms); j++ {
			mSum += ms[j]
			z := int(math.Abs(float64(mSum - nSum)))
			stp[z] = append(stp[z], ms[j])
		}
	}
	min := ms[0]
	arr := []int{ms[0]}
	for k, v := range stp {
		if k < min {
			min = k
			arr = v
		}
	}
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], "")
	}
}

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}

func TestTrace(t *testing.T) {
	a()
}
