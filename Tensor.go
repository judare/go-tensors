package main

import (
	"fmt"
)

type Tensor struct {
	Dims    int
	Sizes   []int
	Weights []int
	Data    []float32
}

func NewTensor(data []float32, sizes []int) Tensor {
	obj := Tensor{
		Sizes:   sizes,
		Dims:    len(sizes),
		Weights: []int{},
		Data:    data,
	}
	obj.Process()
	return obj
}

func (t *Tensor) Process() {
	sizes := t.Sizes
	weights := make([]int, len(sizes))
	weights[0] = 1
	for i, _ := range sizes {
		if i > 0 {
			weights[i] = weights[i-1] * sizes[i-1]
		}
	}
	t.Weights = weights
}

func (t *Tensor) Reshape(sizes []int) {
	t.Sizes = sizes
	t.Process()
}

func (t Tensor) BuildIndex(keys []int) int {
	index := 0
	for i, v := range t.Weights {
		index += keys[len(keys)-1-i] * v
	}
	return index
}

func (t Tensor) Get(keys []int) float32 {
	return t.Data[t.BuildIndex(keys)]
}

func (t Tensor) Set(keys []int, value float32) {
	t.Data[t.BuildIndex(keys)] = value
}

func (t Tensor) GetIndex(keys []int) float32 {
	return t.Data[t.BuildIndex(keys)]
}

func (t Tensor) printAux(indexes []int) {
	if len(indexes) == len(t.Sizes) {
		fmt.Printf(" %d ", int64(t.Get(indexes)))
		return
	}

	last := t.Sizes[len(t.Sizes)-1-len(indexes)]
	fmt.Print("[")
	for i := 0; i < last; i++ {
		sizes := append(indexes, i)
		t.printAux(sizes)
	}
	fmt.Print("]")
}

func (t Tensor) print() {
	fmt.Print("Tensor:\n")
	t.printAux([]int{})
	fmt.Print("\n")

}
