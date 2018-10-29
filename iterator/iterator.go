package iterator

type Ints []int

func (i Ints) Iteracor() <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range i {
			c <- v
		}
		close(c)
	}()
	return c
}
