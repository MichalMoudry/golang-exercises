package flatten

import "fmt"

func Run() {
	input1 := []any{1, []any{2, 3, nil, 4}, nil, 5}
	input1 = Flatten(input1)
	fmt.Println(input1...)
}

func Flatten(nested interface{}) []interface{} {
	result := []any{}
	arg, ok := nested.([]any)
	if !ok {
		return nil
	}
	for _, v := range arg {
		if v == nil {
			continue
		}
		if _, isArray := v.([]any); isArray {
			result = append(result, Flatten(v)...)
		} else {
			result = append(result, v)
		}
	}
	return result
}
