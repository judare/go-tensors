package main

import (
	"errors"
)

func IndexSelect(t Tensor, dim int, indexes []int) (Tensor, error) {

	// Control Dim
	if dim < 0 {
		return Tensor{}, errors.New("Dims must be greater than 0")
	}
	if dim > len(t.Sizes) {
		return Tensor{}, errors.New("Dims must be less than the sizes of Tensor")
	}

	// In the case to pick by index
	if dim == 0 {
		Arr := []float32{}
		for j := 0; j < len(t.Sizes); j++ {
			// Picking indixes from data
			for _, v := range indexes {
				Val := t.Get([]int{j, v})
				Arr = append(Arr, Val)
			}
		}
		newSizes := append(t.Sizes[0:len(t.Sizes)-1], len(indexes))
		// Creating new tensor
		T := NewTensor(Arr, newSizes)
		return T, nil
	}

	dims := t.Sizes[0:1]
	dataFinal := []float32{}
	DataCount := -1
	// Split recursively de index selection
	for i := 0; i < t.Sizes[0]; i++ {
		Start := i * t.Sizes[0]
		Taux := NewTensor(t.Data[Start:Start+t.Sizes[0]], t.Sizes[1:])
		t, _ := IndexSelect(Taux, dim-1, indexes)
		// Backtracking to build the final var
		dataFinal = append(dataFinal, t.Data...)
		DataCount = Taux.Sizes[len(Taux.Sizes)-1]
	}

	T := NewTensor(dataFinal, append([]int{DataCount}, dims...))
	// Return final Tensor
	return T, nil
}
