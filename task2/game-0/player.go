package main

type Player struct {
	room  *Room
	items []*Item
}

func (player *Player) get_item(item *Item) (*Item, bool) {
	for i := range player.items {
		if player.items[i].name == item.name {
			return player.items[i], true
		}
	}

	return nil, false
}

func (player *Player) wear(item_name string) string {
	var ret string

	if item, ok := player.room.has_item(item_name); ok {
		player.add_item(item)
		player.room.remove_item(item)

		for i := range TASKS {
			if TASKS[i].is_done() {
				TASKS = remove_task(TASKS, i)

				break
			}
		}

		ret = "вы одели: " + item_name
	} else {
		ret = "нечего одевать"
	}

	return ret
}

func (player *Player) add_item(item *Item) {
	player.items = append(player.items, item)
}
