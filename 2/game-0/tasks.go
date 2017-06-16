// < DONE
package main

type task struct {
	name    string
	is_done func() bool
}

// > DONE

func get_tasks() string {
	var ret string

	for i := range TASKS {
		ret += TASKS[i].name

		if i == len(TASKS)-2 {
			ret += " и "
		} else if i != len(TASKS)-1 {
			ret += ", "
		}
	}

	return ret
}

func remove_task(tasks []*task, k int) []*task {
	tasks = tasks[:k+copy(tasks[k:], tasks[k+1:])]

	return tasks
}

func init_tasks() {
	TASKS = []*task{
		&task{
			name: "собрать рюкзак",
			is_done: func() bool {
				_, ok := player.get_item(&BACKPACK)
				return ok
			},
		},

		&task{
			name: "идти в универ",
			is_done: func() bool {
				return player.room == &STREET
			},
		},
	}
}
