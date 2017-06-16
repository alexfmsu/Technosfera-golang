package main

// < NOT DONE
type Room struct {
	name         string
	neighbour    []*Room
	environment  string
	items        []*Item
	move_print   string
	in_condition func() (bool, string)
}

// > NOT DONE

// < NOT DONE second item
func (room Room) get_item(item_name string) (*Item, bool) {
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

// > NOT DONE second item

// < NOT DONE second item, cut
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

// > NOT DONE second item, cut

// < DONE
func (player *Player) get_neighbour(neighbour_name string) (*Room, bool) {
	for i := range player.room.neighbour {
		if player.room.neighbour[i].name == neighbour_name {
			return player.room.neighbour[i], true
		}
	}

	return nil, false
}

// > DONE

func (room *Room) has_neighbour() bool {
	return len(room.neighbour) > 0
}

// func (room *Room) get_players() []*Player {
//  var pl []*Player

//  for i := range PLAYERS {
//      if PLAYERS[i].room == room {
//          pl = append(pl, PLAYERS[i])
//      }
//  }

//  return pl
// }

// < NOT DONE
func (game *Game) init_rooms() {
	game.rooms = []*Room{&Kitchen, &Flat, &Hall, &Street}

	Kitchen = Room{
		neighbour:   []*Room{&Hall},
		name:        "кухня",
		environment: "ты находишься на кухне",
		items: []*Item{
			&Item{
				name:     "стол",
				to_print: "на столе",

				can_take: false,
				can_wear: false,

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

	Flat = Room{
		neighbour:   []*Room{&Hall},
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

						init_room: &Flat,

						can_take: true,
						can_wear: false,
					},
					&Item{
						name:  "конспекты",
						items: nil,

						init_room: &Flat,

						can_take: true,
						can_wear: false,
					},
				},
			},

			&Item{
				name:     "стул",
				to_print: "на стуле -",

				items: []*Item{
					&Backpack,
				},
			},
		},
		move_print:   "ты в своей комнате",
		in_condition: func() (bool, string) { return true, "" },
	}

	Hall = Room{
		neighbour:    []*Room{&Kitchen, &Flat, &Street},
		name:         "коридор",
		environment:  "ничего интересного",
		move_print:   "ничего интересного",
		in_condition: func() (bool, string) { return true, "" },
	}

	Street = Room{
		neighbour:    []*Room{&HOME},
		name:         "улица",
		environment:  "на улице весна",
		move_print:   "на улице весна",
		in_condition: func() (bool, string) { return Door.status == "открыт", "дверь закрыта" },
	}

	HOME = Room{
		name: "домой",
	}
}

// > NOT DONE
