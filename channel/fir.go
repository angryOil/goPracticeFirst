package main

/*func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	go square(&wg, ch)
	ch <- 9
	wg.Wait()
}

func square(w *sync.WaitGroup, ch chan int) {
	n := <-ch
	time.Sleep(time.Second)
	fmt.Printf("Square():%d", n*n)
	w.Done()
}
*/
