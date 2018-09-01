package slice

type Slice interface {
	// 移除切片中的值
	RemoveValue(ids []int64, value int64) []int64
}

func NewSlice() Slice {
	return new(slice)
}

type slice struct{}

func (t slice) RemoveValue(ids []int64, value int64) []int64 {
	for index, id := range ids {
		if id == value {
			ids = append(ids[:index], ids[index+1:]...)
		}
	}
	return ids
}
