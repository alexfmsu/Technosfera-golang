// < DONE
package main

// > DONE

type task struct {
	name    string
	is_done func() bool
}

func (player *Player) get_tasks() string {
	var ret string

	tasks := player.tasks

	for i := range tasks {
		ret += tasks[i].name

		if i == len(tasks)-2 {
			ret += " и "
		} else if i != len(tasks)-1 {
			ret += ", "
		}
	}

	return ret
}

func remove_task(tasks []*task, k int) []*task {
	tasks = tasks[:k+copy(tasks[k:], tasks[k+1:])]

	return tasks
}

// func (p *Player) init_tasks() {
// 	TASKS = []*task{
// 		&task{
// 			name: "собрать рюкзак",
// 			is_done: func() bool {
// 				_, ok := p.get_item(&BACKPACK)
// 				return ok
// 			},
// 		},

// 		&task{
// 			name: "идти в универ",
// 			is_done: func() bool {
// 				return p.room == &STREET
// 			},
// 		},
// 	}
// }
