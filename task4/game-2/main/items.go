package main

// < NOT DONE
type Item struct {
	name string

	items     []*Item
	room      *Room
	init_room *Room

	to_print  string
	positions []string
	status    string
	apply_to  []*Item

	can_take bool
	can_wear bool

	postfix string
}

func (game *Game) item_exists(name string) bool {
	if _, ok := game.items[name]; ok {
		return true
	}

	return false
}

func (item *Item) has_position(position_name string) bool {
	for i := range item.positions {
		if item.positions[i] == position_name {
			return true
		}
	}

	return false
}

func (game *Game) init_items() {
	game.items = map[string]*Item{
		"дверь": &Door,
		"шкаф":  &Wardrobe,
		"ключи": &Key,
	}
}

func (player *Player) get_items() string {
	var ret string

	items := player.room.items

	var str []string

	for i := range items {
		var _str string

		if len(items[i].items) == 0 {
			continue
		}

		_str += items[i].to_print

		if len(items) > 1 && items[i].name == "стол" {
			_str += ":"
		}

		_str += " "

		ss := items[i]

		for j := range ss.items {
			if obj := ss.items[j]; obj != nil {
				_str += obj.name
			}

			if j != len(ss.items)-1 {
				_str += ", "
			}
		}

		str = append(str, _str)
	}

	ret += join(", ", str)
	// for i := range str {
	// 	if i == len(str)-1 {
	// 		ret += str[i]
	// 	} else {
	// 		ret += str[i] + ", "
	// 	}
	// }

	room := player.room

	if ret == "" {
		if room.name == "комната" {
			ret = "пустая комната"
		} else {
			ret = "ничего интересного"
		}
	}

	return ret
}

func cut(arr []*Item, index int) []*Item {
	arr = arr[:index+copy(arr[index:], arr[index+1:])]

	return arr
}
