package helper

import "encoding/json"

func ConvertMapToStruct(structType any, data any) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, &structType)
	if err != nil {
		return err
	}

	return nil

}

// Return Value Append Slice
func AppendSlice(slice []string, item string) []string {
	slice = append(slice, item)
	return slice
}

// Append Slice By Pointer
func AppendSlice2(slice *[]string, item string) {
	*slice = append(*slice, item)
}
