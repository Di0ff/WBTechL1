package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	var examples = []interface{}{
		true,                           // bool
		42,                             // int
		3.14,                           // float64
		"Hello, World!",                // string
		[]int{1, 2, 3},                 // slice
		map[string]int{"one": 1},       // map
		Person{Name: "Alice", Age: 30}, // struct
		&Person{Name: "Bob", Age: 25},  // pointer to struct
		make(chan int),                 // channel
		func() {},                      // function
	}

	for _, example := range examples {
		checkType(example)
		fmt.Println()
	}

	for _, example := range examples {
		checkTypeReflect(example)
		fmt.Println()
	}
}

func checkType(i interface{}) { // Пример решения задачи без бибилиотеки reflect
	switch i.(type) {
	case bool:
		fmt.Println("bool")
	case int, int8, int16, int32, int64:
		fmt.Println("int")
	case uint, uint8, uint16, uint32, uint64, uintptr:
		fmt.Println("uint")
	case float32, float64:
		fmt.Println("float")
	case complex64, complex128:
		fmt.Println("complex number")
	case string:
		fmt.Println("string")
	case []int:
		fmt.Printf("slice")
		// и т.д.
	}
}

func checkTypeReflect(i interface{}) { // Пример решения задачи с использованием бибилиотеки reflect
	t := reflect.TypeOf(i)

	switch t.Kind() {
	case reflect.Bool:
		fmt.Println("bool")
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Println("int")
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Println("uint")
	case reflect.Float32, reflect.Float64:
		fmt.Println("float")
	case reflect.Complex64, reflect.Complex128:
		fmt.Println("complex number")
	case reflect.String:
		fmt.Println("string")
	case reflect.Array:
		fmt.Printf("array of %s\n", t.Elem().Kind())
	case reflect.Slice:
		fmt.Printf("slice of %s\n", t.Elem().Kind())
	case reflect.Map:
		fmt.Printf("map with keys of %s and values of %s\n", t.Key().Kind(), t.Elem().Kind())
	case reflect.Struct:
		fmt.Println("struct")
	case reflect.Ptr:
		fmt.Printf("pointer to %s\n", t.Elem().Kind())
	case reflect.Interface:
		fmt.Println("interface")
	case reflect.Chan:
		fmt.Printf("channel of %s\n", t.Elem().Kind())
	case reflect.Func:
		fmt.Println("function")
	default:
		fmt.Printf("unknown type: %s\n", t)
	}
}
