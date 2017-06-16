package main
import (
	"reflect"
	"testing"
)
func TestReturnInt(t *testing.T) {
	if ReturnInt() != 1 {
		t.Error("expected 1")
	}
}
func TestReturnFloat(t *testing.T) {
	if ReturnFloat() != float32(1.1) {
		t.Error("expected 1.1")
	}
}
func TestReturnIntArray(t *testing.T) {
	if ReturnIntArray() != [3]int{1, 3, 4} {
		t.Error("expected '[3]int{1, 3, 4}'")
	}
}
func TestReturnIntSlice(t *testing.T) {
	expected := []int{1, 2, 3}
	result := ReturnIntSlice()
	if !reflect.DeepEqual(result, expected) {
		t.Error("expected", expected, "have", result)
	}
}
func TestIntSliceToString(t *testing.T) {
	expected := "1723100500"
	result := IntSliceToString([]int{17, 23, 100500})
	if expected != result {
		t.Error("expected", expected, "have", result)
	}
}
func TestMergeSlices(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5}
	result := MergeSlices([]float32{1.1, 2.1, 3.1}, []int32{4, 5})
	if !reflect.DeepEqual(result, expected) {
		t.Error("expected", expected, "have", result)
	}
}

type pair struct{
    template string
    expr string
}

func TestGetMapValuesSortedByKey(t *testing.T) {
	var cases = []struct {
        expected []string
        input pair
    }{
        {
            expected: []string{
                "4",
            },
            input: pair{
                template: "$",
                expr: "asd",
            },
        },
    }

    for _, item := range cases {
        Ctemplate := C.CString(item.input.template)
        Cexpr := C.CString(item.input.expr)

        fmt.Println(C.slre_match(Ctemplate, Cexpr, 4, nil, 0, 0))
        // fmt.Println(item.input.template)
    }
	// for _, item := range cases {
	// 	result := GetMapValuesSortedByKey(item.input)
	// 	if !reflect.DeepEqual(result, item.expected) {
	// 		t.Error("expected", item.expected, "have", result)
	// 	}
	// }
}