package main

const (
	max     = 5000000
	block   = 500
	bufsize = 100
)

func test() {
	done := make(chan int)
	c := make(chan int, bufsize)

	go func() {
		var count int
		for x := range c {
			count += x
		}

		//println("count:", count)
		close(done)
	}()

	for i := 0; i < max; i++ {
		c <- i
	}
	close(c)
	<-done
}

func testBlock() {
	done := make(chan int)
	c := make(chan [block]int, bufsize)

	go func() {
		var count int
		for x := range c {
			for _, d := range x {
				count += d
			}
		}

		//println("count:", count)
		close(done)
	}()

	var b [block]int
	for i := 0; i < max; i += block {
		for j := 0; j < block; j++ {
			b[j] = i + j
			if i+j == max-1 {
				break
			}
		}
		c <- b
	}

	close(c)
	<-done
}
