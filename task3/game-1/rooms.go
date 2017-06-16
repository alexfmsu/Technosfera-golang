package main

type Room struct {
	name         string
	neighbour    []*Room
	environment  string
	items        []*Item
	move_print   string
	in_condition func() (bool, string)
}

func (room Room) has_item(item_name string) (*Item, bool) {
	for i := range room.items {
		if room.items[i].name == item_name {
			return room.items[i], true
		}

		for j := range room.items[i].items {
			if room.items[i].items[j].name == item_name {
				return room.items[i].items[j], true
			}
		}
	}

	return nil, false
}

func (room *Room) remove_item(item *Item) {
	for i := range room.items {
		if room.items[i] == item {
			room.items = cut(room.items, i)

			return
		}

		for j := range room.items[i].items {
			if room.items[i].items[j] == item {
				room.items[i].items = cut(room.items[i].items, j)

				return
			}
		}
	}
}

func (p *Player) get_neighbour(neighbour_name string) (*Room, bool) {
	for i := range p.room.neighbour {
		if p.room.neighbour[i].name == neighbour_name {
			return p.room.neighbour[i], true
		}
	}

	return nil, false
}

func (room *Room) has_neighbour() bool {
	return len(room.neighbour) > 0
}

func init_rooms() {
	ROOMS = []*Room{&KITCHEN, &FLAT, &HALL, &STREET}

	KITCHEN = Room{
		neighbour:   []*Room{&HALL},
		name:        "кухня",
		environment: "ты находишься на кухне",
		items: []*Item{
			&Item{
				name:     "стол",
				to_print: "на столе",

				items: []*Item{
					&Item{
						name:  "чай",
						items: nil,
					},
				},
			},
		},
		move_print:   "кухня, ничего интересного",
		in_condition: func() (bool, string) { return true, "" },
	}

	FLAT = Room{
		neighbour:   []*Room{&HALL},
		name:        "комната",
		environment: "ты в своей комнате",
		items: []*Item{
			&Item{
				name:     "стол",
				to_print: "на столе",

				items: []*Item{
					&Item{
						name:  "ключи",
						items: nil,
					},
					&Item{
						name:  "конспекты",
						items: nil,
					},
				},
			},

			&Item{
				name:     "стул",
				to_print: "на стуле -",

				items: []*Item{
					&BACKPACK,
				},
			},
		},
		move_print:   "ты в своей комнате",
		in_condition: func() (bool, string) { return true, "" },
	}

	HALL = Room{
		neighbour:    []*Room{&KITCHEN, &FLAT, &STREET},
		name:         "коридор",
		environment:  "ничего интересного",
		move_print:   "ничего интересного",
		in_condition: func() (bool, string) { return true, "" },
	}

	STREET = Room{
		neighbour:    []*Room{&HOME},
		name:         "улица",
		environment:  "на улице весна",
		move_print:   "на улице весна",
		in_condition: func() (bool, string) { return DOOR.status == "открыт", "дверь закрыта" },
	}

	HOME = Room{
		name: "домой",
	}
}
