package pipeline

import "sync"

type job func(in, out chan interface{})

func do_job(in, out chan interface{}, job func(in, out chan interface{}), wg *sync.WaitGroup) {
	job(in, out)

	wg.Done()

	close(out)
}

func Pipe(funcs ...job) {
	var wg sync.WaitGroup

	out := make(chan interface{})

	for i := len(funcs) - 1; i >= 0; i-- {
		in := make(chan interface{})

		wg.Add(1)

		go do_job(in, out, funcs[i], &wg)

		out = in
	}

	close(out)

	wg.Wait()
}
