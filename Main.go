package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

var app = cli.NewApp()

func main() {
	info()
	commands()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func info() {
	app.Name = "GO Tensors operations"
	app.Usage = ""
	app.Authors = []*cli.Author{
		&cli.Author{
			Name:  "Juan David Restrepo Montoya",
			Email: "fosebadgame@gmail.com",
		},
	}
	app.Version = "1.0.0"
}

func commands() {
	app.Commands = []*cli.Command{
		&cli.Command{
			Name:    "reshape",
			Aliases: []string{"r"},
			Usage:   "",
			Action: func(cCtx *cli.Context) error {
				// Reshape
				fmt.Println(string("\033[33m"), "Reshape")
				fmt.Println(string("\033[36m  Original \033[37m"))

				t := NewTensor([]float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, []int{4, 4})

				t.print()

				fmt.Println(string("\033[36m  Reshaped \033[37m"))
				Reshape(t, []int{2, 2, 4})

				return nil
			},
		},
		&cli.Command{
			Name:    "hadamardproduct",
			Aliases: []string{"hp"},
			Usage:   "",
			Action: func(cCtx *cli.Context) error {
				// HadamardProduct
				fmt.Println("\033[33m", "HadamardProduct")

				fmt.Println("\033[36m", "Tensor 1 \033[37m")
				t1 := NewTensor([]float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, []int{4, 4})
				t1.print()

				fmt.Println(string("\033[36m  Tensor 2 \033[37m"))
				t2 := NewTensor([]float32{16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{4, 4})
				t2.print()

				t3, errorMessage := HadamardProduct(t1, t2)
				if errorMessage != nil {
					fmt.Println(errorMessage)
				} else {
					fmt.Println(string("\033[36m  Result \033[37m"))
					t3.print()
				}

				return nil
			},
		},
		&cli.Command{
			Name:    "indexselect",
			Aliases: []string{"is"},
			Usage:   "",
			Action: func(cCtx *cli.Context) error {

				t := NewTensor([]float32{1, 2, 3, 4}, []int{2, 2})
				t, _ = IndexSelect(t, 1, []int{0})
				fmt.Println("=> IndexSelect([[1, 2], [3, 4]], 1, [0])")
				t.print()

				t = NewTensor([]float32{1, 2, 3, 4}, []int{2, 2})
				t, _ = IndexSelect(t, 1, []int{0, 0})
				fmt.Println("=> IndexSelect([[1, 2], [3, 4]], 1, [0,0])")
				t.print()

				t = NewTensor([]float32{1, 2, 3, 4}, []int{2, 2})
				t, _ = IndexSelect(t, 1, []int{0, 0, 1, 1})
				fmt.Println("=> IndexSelect([[1, 2], [3, 4]], 1, [0,0,1,1])")
				t.print()

				return nil
			},
		},
	}
}
