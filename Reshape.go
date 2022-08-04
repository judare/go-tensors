package main

func Reshape(t Tensor, dim []int) {

	// Call Reshape fn, only changes the sizes slice
	t.Reshape(dim)
	t.print()
}
