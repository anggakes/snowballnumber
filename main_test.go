package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type detail struct {
	N int
	R [][]int
}
var testTable = []detail{
	{
		N: 1,
		R: [][]int{{1}},
	},
	{
		N: 2,
		R: [][]int{{2}},
	},
	{
		N: 6,
		R: [][]int{{1, 2, 3}, {6}},
	},
	{
		N: 9,
		R: [][]int{{2, 3, 4}, {4, 5}, {9}},
	}, {
		N: 10,
		R: [][]int{{1,2,3,4}, {10}},
	},
	{
		N: 100,
		R: [][]int{{9,10,11,12,13,14,15,16},{18,19,20,21,22},{100}},
	},
	{
		N: 999,
		R: [][]int{{9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45},{24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50},{47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,64},{107,108,109,110,111,112,113,114,115},{164,165,166,167,168,169},{332,333,334},{499,500},{999}},
	},
}

func TestSnowballNumberSimple(t *testing.T) {
	for _, v := range testTable {
		t.Run(fmt.Sprintf("N = %v", v.N),func(t *testing.T) {
			Rs := SnowballNumberSimple(v.N)
			assert.Equal(t, Rs, v.R)
		})
	}
}

func BenchmarkSnowballNumberSimple(b *testing.B) {
	for _, v := range testTable {
		b.Run(fmt.Sprintf("N = %d", v.N), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SnowballNumberSimple(v.N)
			}
		})
	}
}


func TestSnowballNumberSnowballNumberReduceIteration(t *testing.T) {
	for _, v := range testTable {
		t.Run(fmt.Sprintf("N = %v", v.N),func(t *testing.T) {
			Rs := SnowballNumberReduceIteration(v.N)
			assert.Equal(t, Rs, v.R)
		})
	}
}

func BenchmarkSnowballNumberReduceIteration(b *testing.B) {
	for _, v := range testTable {
		b.Run(fmt.Sprintf("N = %d", v.N), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SnowballNumberReduceIteration(v.N)
			}
		})
	}
}


func TestSnowballNumberArithmetic(t *testing.T) {
	for _, v := range testTable {
		t.Run(fmt.Sprintf("N = %v", v.N),func(t *testing.T) {
			Rs := SnowballNumberArithmetic(v.N)
			assert.Equal(t, Rs, v.R)
		})
	}
}

func BenchmarkAwesomeNumbeArithmetic(b *testing.B) {
	for _, v := range testTable {
		b.Run(fmt.Sprintf("N = %d", v.N), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SnowballNumberArithmetic(v.N)
			}
		})
	}
}
//
//func BenchmarkAwesomeReduceIteration(b *testing.B) {
//	for _, v := range testTable {
//		b.Run(fmt.Sprintf("input_size_%d", v), func(b *testing.B) {
//			for i := 0; i < b.N; i++ {
//				SnowballNumberReduceIteration(v)
//			}
//		})
//	}
//}
//
//func BenchmarkSnowballNumberWithAritmatik(b *testing.B) {
//	for _, v := range testTable {
//		b.Run(fmt.Sprintf("input_size_%d", v), func(b *testing.B) {
//			for i := 0; i < b.N; i++ {
//				SnowballNumberWithAritmatik(v)
//			}
//		})
//	}
//}
//
//func BenchmarkSnowballNumberWithAritmatikChannel(b *testing.B) {
//	for _, v := range testTable {
//		b.Run(fmt.Sprintf("input_size_%d", v), func(b *testing.B) {
//			for i := 0; i < b.N; i++ {
//				SnowballNumberWithAritmatikChannel(v)
//			}
//		})
//	}
//}

//func TestPersamaanKuadrat(t *testing.T) {
//	b, c := getBandC(2, 9)
//	fmt.Println(c,b)
//	fmt.Println(persamaanKuadrat(1, float64(b), float64(c)))
//	fmt.Println(persamaanKuadrat(1, 4, -60))
//}
