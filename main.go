package main

import (
	"sync"
	"syscall/js"
)

func main() {
	var wg sync.WaitGroup
	window := js.Global()
	window.Set("wasmGroup", make(map[string]interface{}))
	wasmGroup := window.Get("wasmGroup")

	loop := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		for i := 0; i < 10000; i++ {
		}
		return nil
	})

	loop2 := func() {
		for i := 0; i < 10000; i++ {
		}
	}

	fibonacci := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var fib func(int) int
		fib = func(num int) int {
			if num < 1 {
				return 0
			}
			if num < 3 {
				return 1
			}
			return fib(num-1) + fib(num-2)
		}

		fib(20)
		return nil
	})

	fibonacci2 := func() {
		var fib func(int) int
		fib = func(num int) int {
			if num < 1 {
				return 0
			}
			if num < 3 {
				return 1
			}
			return fib(num-1) + fib(num-2)
		}

		fib(20)
	}

	selectionSort := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		s := func(arr []int) []int {
			for i := 0; i < len(arr); i++ {
				min := i
				for j := i + 1; j < len(arr); j++ {
					if arr[j] < arr[min] {
						min = j
					}
				}
				if min != i {
					arr[i], arr[min] = arr[min], arr[i]
				}
			}
			return arr
		}

		arr := []int{12, 4124, 3, 24, 5, 43, 34, 654, 7, 54, 23, 643, 34, 76, 32, 125, 43256, 4, 6546, 45, 7, 45, 3, 423, 5, 43, 65, 76, 554, 745, 7}
		s(arr)
		return nil
	})

	selectionSort2 := func() {
		s := func(arr []int) []int {
			for i := 0; i < len(arr); i++ {
				min := i
				for j := i + 1; j < len(arr); j++ {
					if arr[j] < arr[min] {
						min = j
					}
				}
				if min != i {
					arr[i], arr[min] = arr[min], arr[i]
				}
			}
			return arr
		}

		arr := []int{12, 4124, 3, 24, 5, 43, 34, 654, 7, 54, 23, 643, 34, 76, 32, 125, 43256, 4, 6546, 45, 7, 45, 3, 423, 5, 43, 65, 76, 554, 745, 7}
		s(arr)
	}

	createElement := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		document := js.Global().Get("document")
		document.Call("createElement", "div")
		document.Call("createElement", "div")
		document.Call("createElement", "div")
		document.Call("createElement", "div")
		document.Call("createElement", "div")
		document.Call("createElement", "div")
		document.Call("createElement", "div")
		document.Call("createElement", "div")
		document.Call("createElement", "div")
		document.Call("createElement", "div")
		return nil
	})

	createElement2 := func() {
		document := js.Global().Get("document")
		document.Call("createElement", "div")
		document.Call("createElement", "div")
		document.Call("createElement", "div")
		document.Call("createElement", "div")
		document.Call("createElement", "div")
		document.Call("createElement", "div")
		document.Call("createElement", "div")
		document.Call("createElement", "div")
		document.Call("createElement", "div")
		document.Call("createElement", "div")
	}

	appendChild := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		document := js.Global().Get("document")
		a := document.Call("createElement", "div")
		b := document.Call("createElement", "div")
		c := document.Call("createElement", "div")
		d := document.Call("createElement", "div")
		a.Call("appendChild", b.Call("appendChild", c.Call("appendChild", d)))
		return nil
	})

	appendChild2 := func() {
		document := js.Global().Get("document")
		a := document.Call("createElement", "div")
		b := document.Call("createElement", "div")
		c := document.Call("createElement", "div")
		d := document.Call("createElement", "div")
		a.Call("appendChild", b.Call("appendChild", c.Call("appendChild", d)))
	}

	complex := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		loop2()
		fibonacci2()
		createElement2()
		selectionSort2()
		appendChild2()
		loop2()
		fibonacci2()
		createElement2()
		selectionSort2()
		appendChild2()
		return nil
	})

	wasmGroup.Set("loop", loop)
	wasmGroup.Set("fibonacci", fibonacci)
	wasmGroup.Set("selectionSort", selectionSort)
	wasmGroup.Set("createElement", createElement)
	wasmGroup.Set("appendChild", appendChild)
	wasmGroup.Set("complex", complex)

	window.Call("main")
	wg.Wait()
}
