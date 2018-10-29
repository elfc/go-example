package container

type IntHeap []int

func (t IntHeap) Len() int           { return len(t) }
func (t IntHeap) Less(i, j int) bool { return t[i] < t[i] }
func (t IntHeap) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }

func (t *IntHeap) Push(x interface{}) {
	*t = append(*t, x.(int))
}

func (t *IntHeap) Pop() interface{} {
	old := *t
	n := len(old)
	x := old[n-1]
	*t = old[0 : n-1]
	return x
}
