// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package treesort_test

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"

	"gopl.io/ch4/treesort"
)

func TestSort(t *testing.T) {
	data := createInts(500)
	//treesort.Sort(data)
	sort.Ints(data)
	if !sort.IntsAreSorted(data) {
		t.Errorf("not sorted: %v", data)
	}
}

func TestMySort(t *testing.T) {
	data := createInts(500)
	//treesort.Sort(data)
	mySort(data)
	if !sort.IntsAreSorted(data) {
		t.Errorf("not sorted: %v", data)
	}
}

//!+bench
func BenchmarkNonempty0(b *testing.B) {
	data := createInts(5000)
	fmt.Println(b.N)
	for i := 0; i < b.N; i++ {
		treesort.Sort(data)
	}
}

//!+bench
func BenchmarkNonempty1(b *testing.B) {
	data := createInts(5000)
	fmt.Println(b.N)
	for i := 0; i < b.N; i++ {
		sort.Ints(data)
	}
}

//!+bench
func BenchmarkNonempty3(b *testing.B) {
	data := createInts(5000)
	fmt.Println(b.N)
	for i := 0; i < b.N; i++ {
		mySort(data)
	}
}

func createInts(n int) []int {
	data := make([]int, n)
	for i := range data {
		data[i] = rand.Int() % n
	}
	return data
}

func mySort(data []int) {
	cnt := make([]int, 5000)
	for _, num := range data {
		cnt[num]++
	}
	idata := 0
	for ic, c := range cnt {
		if c == 0 {
			continue
		}
		for i := 0; i < c; i++ {
			data[idata] = ic
			idata++
		}
	}
}
