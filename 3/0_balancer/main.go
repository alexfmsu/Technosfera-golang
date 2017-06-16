package main

type RoundRobinBalancer struct {
	stat         []int
	node_num_chl chan int
	do_job_chl   chan bool
}

func (balancer *RoundRobinBalancer) Init(n int) {
	balancer.stat = make([]int, n)
	balancer.node_num_chl = make(chan int)
	balancer.do_job_chl = make(chan bool)

	node := 0

	go func() {
		do := balancer.do_job_chl

		for {
			if <-do {
				balancer.stat[node]++

				node = (node + 1) % n

				balancer.node_num_chl <- node
			}
		}
	}()
}

func (balancer *RoundRobinBalancer) GiveStat() []int {
	return balancer.stat
}

func (balancer *RoundRobinBalancer) GiveNode() int {
	balancer.do_job_chl <- true

	return <-balancer.node_num_chl
}
