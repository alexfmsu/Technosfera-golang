// < DONE
package main

// > DONE

// < NOT DONE
type Item struct {
	name      string
	items     []*Item
	to_print  string
	positions []string
	status    string
	apply_to  []*Item
	postfix   string
}

func item_exists(name string) bool {
	if _, ok := ITEMS[name]; ok {
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

func init_items() {
	ITEMS = map[string]*Item{
		"дверь": &DOOR,
		"шкаф":  &WARDROBE,
		"ключи": &KEY,
	}
}

func get_items() string {
	var ret string

	st := player.room.items

	var str []string

	for i := range st {
		var _str string

		if len(st[i].items) == 0 {
			continue
		}

		_str += st[i].to_print

		if len(st) > 1 && st[i].name == "стол" {
			_str += ":"
		}

		_str += " "

		ss := st[i]

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

	for i := range str {
		if i == len(str)-1 {
			ret += str[i]
		} else {
			ret += str[i] + ", "
		}
	}

	if ret == "" {
		if player.room.name == "комната" {
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
