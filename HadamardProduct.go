package main

import "errors"

func HadamardProduct(t1 Tensor, t2 Tensor) (Tensor, error) {

	if t1.Dims != t2.Dims {
		return Tensor{}, errors.New("Dims attribute is different at t1 vs t2")
	}

	data := []float32{}
	// Both slices are linear vectors (we can multiply each index for another)
	for i, _ := range t1.Data {
		data = append(data, t1.Data[i]*t2.Data[i])
	}
	T := NewTensor(data, t1.Sizes)
	return T, nil
}
