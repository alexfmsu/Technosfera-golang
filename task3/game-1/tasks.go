// < DONE
package main

type task struct {
	name    string
	is_done func() bool
}

// > DONE

func (p *Player) get_tasks() string {
	var ret string

	for i := range p.Tasks {
		ret += p.Tasks[i].name

		if i == len(p.Tasks)-2 {
			ret += " и "
		} else if i != len(p.Tasks)-1 {
			ret += ", "
		}
	}

	return ret
}

func remove_task(tasks []*task, k int) []*task {
	tasks = tasks[:k+copy(tasks[k:], tasks[k+1:])]

	return tasks
}

func (p *Player) init_tasks() {
	TASKS = []*task{
		&task{
			name: "собрать рюкзак",
			is_done: func() bool {
				_, ok := p.get_item(&BACKPACK)
				return ok
			},
		},

		&task{
			name: "идти в универ",
			is_done: func() bool {
				return p.room == &STREET
			},
		},
	}
}
