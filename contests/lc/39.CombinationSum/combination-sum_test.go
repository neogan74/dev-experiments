package _39_CombinationSum

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"reflect"
	"testing"
)

func Test_combinationSum(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]int
	}{
		{
			name: "Test 1",
			args: args{
				candidates: []int{2, 3, 6, 7},
				target:     7,
			},
			wantAns: [][]int{{2, 2, 3}, {7}},
		},
		{
			name: "Test 2",
			args: args{
				candidates: []int{2, 3, 5},
				target:     8,
			},
			wantAns: [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			name: "Test 3",
			args: args{
				candidates: []int{2},
				target:     1,
			},
			wantAns: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := combinationSum(tt.args.candidates, tt.args.target); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("combinationSum() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

// Генерация случайного массива для тестирования
func generateRandomArray(size, maxVal int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		// Read 8 bytes for a uint64
		b := make([]byte, 8)
		_, err := rand.Read(b)
		if err != nil {
			return nil
		}
		// Convert bytes to uint64 and then to int, taking modulo maxVal
		val := int(binary.BigEndian.Uint64(b) % uint64(maxVal))
		arr[i] = val
	}
	return arr
}

// Benchmark
func BenchmarkPartitionArray(b *testing.B) {
	sizes := []int{100, 1000, 10000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("Size-%d", size), func(b *testing.B) {
			arr := generateRandomArray(size, 1000000)
			k := 1000

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				combinationSum(append([]int{}, arr...), k)
			}
		})
	}
}
