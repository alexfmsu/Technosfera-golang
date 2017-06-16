package main

func boolType(i interface{}) string {
	switch i.(type) {
	case bool:
		return "bool"
	default:
		return "unknown"
	}
}

func intType(i interface{}) string {
	switch i.(type) {
	case int:
		return "int"
	case int8:
		return "int8"
	case int16:
		return "int16"
	case int32:
		return "int32"
	case int64:
		return "int64"
	default:
		return "unknown"
	}
}

func uintType(i interface{}) string {
	switch i.(type) {
	case uint:
		return "uint"
	case uint8:
		return "uint8"
	case uint16:
		return "uint16"
	case uint32:
		return "uint32"
	case uint64:
		return "uint64"
	default:
		return "unknown"
	}
}

func floatType(i interface{}) string {
	switch i.(type) {
	case float32:
		return "float32"
	case float64:
		return "float64"
	default:
		return "unknown"
	}
}

func stringType(i interface{}) string {
	switch i.(type) {
	case string:
		return "string"
	default:
		return "unknown"
	}
}

func boolArrayType(i interface{}) string {
	switch i.(type) {
	case []bool:
		return "[]bool"
	default:
		return "unknown"
	}
}

func intArrayType(i interface{}) string {
	switch i.(type) {
	case []int:
		return "[]int"
	case []int8:
		return "[]int8"
	case []int16:
		return "[]int16"
	case []int32:
		return "[]int32"
	case []int64:
		return "[]int64"
	default:
		return "unknown"
	}
}

func uintArrayType(i interface{}) string {
	switch i.(type) {
	case []uint:
		return "[]uint"
	case []uint8:
		return "[]uint8"
	case []uint16:
		return "[]uint16"
	case []uint32:
		return "[]uint32"
	case []uint64:
		return "[]uint64"
	default:
		return "unknown"
	}
}

func floatArrayType(i interface{}) string {
	switch i.(type) {
	case []float32:
		return "[]float32"
	case []float64:
		return "[]float64"
	default:
		return "unknown"
	}
}

func stringArrayType(i interface{}) string {
	switch i.(type) {
	case []string:
		return "[]string"
	default:
		return "unknown"
	}
}

func scalarType(i interface{}) string {
	if t := boolType(i); t != "unknown" {
		return t
	}

	if t := intType(i); t != "unknown" {
		return t
	}

	if t := uintType(i); t != "unknown" {
		return t
	}

	if t := floatType(i); t != "unknown" {
		return t
	}

	if t := stringType(i); t != "unknown" {
		return t
	}

	return "unknown"
}

func arrayType(i interface{}) string {
	if t := boolArrayType(i); t != "unknown" {
		return t
	}

	if t := intArrayType(i); t != "unknown" {
		return t
	}

	if t := uintArrayType(i); t != "unknown" {
		return t
	}

	if t := floatArrayType(i); t != "unknown" {
		return t
	}

	if t := stringArrayType(i); t != "unknown" {
		return t
	}

	return "unknown"
}

func mapBoolType(i interface{}) string {
	switch i.(type) {
	case map[bool]bool:
		return "map[bool]bool"

	case map[bool]int:
		return "map[bool]int"
	case map[bool]int8:
		return "map[bool]int8"
	case map[bool]int16:
		return "map[bool]int16"
	case map[bool]int32:
		return "map[bool]int32"
	case map[bool]int64:
		return "map[bool]int64"

	case map[bool]uint:
		return "map[bool]uint"
	case map[bool]uint8:
		return "map[bool]uint8"
	case map[bool]uint16:
		return "map[bool]uint16"
	case map[bool]uint32:
		return "map[bool]uint32"
	case map[bool]uint64:
		return "map[bool]uint64"

	case map[bool]float32:
		return "map[bool]float32"
	case map[bool]float64:
		return "map[bool]float64"

	case map[bool]string:
		return "map[bool]string"

	default:
		return "unknown"
	}
}

func mapIntType(i interface{}) string {
	switch i.(type) {
	case map[int]bool:
		return "map[int]bool"

	case map[int]int:
		return "map[int]int"

	case map[int]uint:
		return "map[int]uint"

	case map[int]float32:
		return "map[int]float32"
	case map[int]float64:
		return "map[int]float64"

	case map[int]string:
		return "map[int]string"

	default:
		return "unknown"
	}
}

func mapUIntType(i interface{}) string {
	switch i.(type) {
	case map[uint]bool:
		return "map[uint]bool"

	case map[uint]int:
		return "map[uint]int"

	case map[uint]uint:
		return "map[uint]uint"

	case map[uint]float32:
		return "map[uint]float32"
	case map[uint]float64:
		return "map[uint]float64"

	case map[uint]string:
		return "map[uint]string"

	default:
		return "unknown"
	}
}

func mapFloatType(i interface{}) string {
	switch i.(type) {
	case map[float32]bool:
		return "map[float32]bool"
	case map[float64]bool:
		return "map[float64]bool"

	case map[float32]int:
		return "map[float32]int"
	case map[float64]int:
		return "map[float64]int"

	case map[float32]uint:
		return "map[float32]uint"
	case map[float64]uint:
		return "map[float64]uint"

	case map[float32]float32:
		return "map[float32]float32"
	case map[float64]float32:
		return "map[float64]float32"
	case map[float32]float64:
		return "map[float32]float64"

	case map[float32]string:
		return "map[float32]string"
	case map[float64]float64:
		return "map[float64]float64"

	case map[float64]string:
		return "map[float64]string"

	default:
		return "unknown"
	}
}

func mapStringType(i interface{}) string {
	switch i.(type) {
	case map[string]bool:
		return "map[string]bool"

	case map[string]int:
		return "map[string]int"

	case map[string]uint:
		return "map[string]uint"

	case map[string]float32:
		return "map[string]float32"
	case map[string]float64:
		return "map[string]float64"

	case map[string]string:
		return "map[string]string"

	default:
		return "unknown"
	}
}

func mapType(i interface{}) string {
	if t := mapBoolType(i); t != "unknown" {
		return t
	}

	if t := mapIntType(i); t != "unknown" {
		return t
	}

	if t := mapUIntType(i); t != "unknown" {
		return t
	}

	if t := mapFloatType(i); t != "unknown" {
		return t
	}

	if t := mapStringType(i); t != "unknown" {
		return t
	}

	return "unknown"
}
func showMeTheType(i interface{}) string {
	if t := scalarType(i); t != "unknown" {
		return t
	}

	if t := arrayType(i); t != "unknown" {
		return arrayType(i)
	}

	if t := mapType(i); t != "unknown" {
		return mapType(i)
	}

	return "unkwown"
}

func main() {}
