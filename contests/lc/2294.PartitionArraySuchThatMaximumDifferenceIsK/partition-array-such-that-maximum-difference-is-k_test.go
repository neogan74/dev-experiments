package _2294_partition_array_such_that_maximum_difference_is_k

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func Test_partitionArray(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				nums: []int{3, 6, 1, 2, 5},
				k:    2,
			},
			want: 2,
		},
		{
			name: "Test 2",
			args: args{
				nums: []int{1, 2, 3},
				k:    1,
			},
			want: 2,
		},
		{
			name: "Test 3",
			args: args{
				nums: []int{2, 2, 4, 5},
				k:    0,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partitionArray(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("partitionArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Генерация случайного массива для тестирования
func generateRandomArray(size, maxVal int) []int {
	// Создаем новый источник случайных чисел
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	arr := make([]int, size)
	for i := 0; i < size; i++ {
		// Генерируем случайное число от 0 до maxVal-1
		arr[i] = r.Intn(maxVal)
	}
	return arr
}

// Benchmark
func BenchmarkPartitionArray(b *testing.B) {
	sizes := []int{100, 1000, 10000, 100000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("Size-%d", size), func(b *testing.B) {
			arr := generateRandomArray(size, 1000000)
			k := 1000

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				partitionArray(append([]int{}, arr...), k)
			}
		})
	}
}
