package main

import (
	"fmt"
	"reflect"
)

// An UnmarshalTypeError describes a JSON value that was not appropriate for
// a value of a specific Go type.
// Naming convention: The word "Error" ends at the name of the type.
type UnmarshalTypeError struct {
	Value string       // description of JSON value
	Type  reflect.Type // type of Go value it could not be assigned to
}

// UnmarshalTypeError implements the error interface.
// We are using pointer semantic.
// In the implementation, we are validating all the fields are being used in the error message. If
// not, we have a problem. Because why would you add a field to the custom error type and not
// displaying on your log when this method would call. We only do this when we really need it.
func (e *UnmarshalTypeError) Error() string {
	return "json: cannot unmarshal " + e.Value + " into Go value of type " + e.Type.String()
}

// An InvalidUnmarshalError describes an invalid argument passed to Unmarshal.
// (The argument to Unmarshal must be a non-nil pointer.)
// This concrete type is used when we don't pass the address of a value into Unmarshal function.
type InvalidUnmarshalError struct {
	Type reflect.Type
}

// InvalidUnmarshalError implements the error interface.
func (e *InvalidUnmarshalError) Error() string {
	if e.Type == nil {
		return "json: Unmarshal(nil)"
	}

	if e.Type.Kind() != reflect.Ptr {
		return "json: Unmarshal(non-pointer " + e.Type.String() + ")"
	}
	return "json: Unmarshal(nil " + e.Type.String() + ")"
}

// user is a type for use in the Unmarshal call.
type user2 struct {
	Name int
}

func main() {
	var u user2
	err := Unmarshal([]byte(`{"name":"bill"}`), u) // Run with a value and pointer.
	if err != nil {
		// This is a special type assertion that only works on the switch.
		switch e := err.(type) {
		case *UnmarshalTypeError:
			fmt.Printf("UnmarshalTypeError: Value[%s] Type[%v]\n", e.Value, e.Type)
		case *InvalidUnmarshalError:
			fmt.Printf("InvalidUnmarshalError: Type[%v]\n", e.Type)
		default:
			fmt.Println(err)
		}
		return
	}

	fmt.Println("Name:", u.Name)
}

// Unmarshal simulates an unmarshal call that always fails.
// Notice the parameters here: The first one is a slice of byte and the second one is an empty
// interface. The empty interface basically says nothing, which means any value can be passed into
// this function.
// We are going to reflect on the concrete type that is stored inside this interface and we are
// going to validate that if it is a pointer or not nil. We then return different error types
// depending on these.
func Unmarshal(data []byte, v interface{}) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return &InvalidUnmarshalError{reflect.TypeOf(v)}
	}

	return &UnmarshalTypeError{"string", reflect.TypeOf(v)}
}
