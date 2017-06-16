// < DONE
package main

// > DONE

import (
	"strings"
)

// Player
var player Player

// Player

// Rooms
var KITCHEN Room
var FLAT Room
var HALL Room
var STREET Room
var HOME Room

var ROOMS []*Room

var INIT_ROOM *Room

// Rooms

// Tasks
var TASKS []*task

// Tasks

// Items
var BACKPACK = Item{
	name: "рюкзак",
}

var ITEMS map[string]*Item

// > NOT DONE

var DOOR = Item{
	name:      "дверь",
	positions: []string{"открыт", "закрыт"},
	status:    "закрыт",
	postfix:   "а",
}

var WARDROBE = Item{
	name:      "шкаф",
	positions: []string{"открыт", "закрыт"},
	status:    "закрыт",
	postfix:   "",
}

var KEY = Item{
	apply_to: []*Item{&DOOR},
}

// Items

func look_around() string {
	var ret string

	if player.room == INIT_ROOM {
		ret += player.room.environment + ", " + get_items() + ", " + "надо " + get_tasks() + ". "
	} else {
		ret += get_items() + ". "
	}

	ret += can_move()

	return ret
}

func can_move() string {
	var ret string

	if player.room.has_neighbour() {
		var str []string

		for i := range player.room.neighbour {
			str = append(str, (*player.room.neighbour[i]).name)
		}

		ret = "можно пройти - " + strings.Join(str, ", ")
	} else {
		ret = "нельзя никуда пройти"
	}

	return ret
}

func move(room_name string) string {
	var ret string

	if room, ok := player.get_neighbour(room_name); ok {
		if ok, fault_msg := room.in_condition(); !ok {
			return fault_msg
		}

		player.room = room

		ret = player.room.move_print + ". " + can_move()
	} else {
		ret = "нет пути в " + room_name
	}

	return ret
}

func take(item_name string) string {
	var ret string

	if _, ok := player.get_item(&BACKPACK); ok {
		if item, ok := player.room.has_item(item_name); ok {
			player.add_item(item)
			player.room.remove_item(item)

			ret = "предмет добавлен в инвентарь: " + item.name
		} else {
			ret = "нет такого"
		}
	} else {
		ret = "некуда класть"
	}

	return ret
}

func apply(subj Item, obj Item) string {
	var ret string

	if _, ok := player.get_item(&subj); !ok {
		ret = "нет предмета в инвентаре - " + subj.name

		return ret
	}

	if item_exists(subj.name) {
		objX := ITEMS[subj.name]

		for i := range objX.apply_to {
			if objX.apply_to[i].name == obj.name {
				objY := ITEMS[obj.name]

				if objY.status == "закрыт" && objY.has_position("открыт") {
					ITEMS[obj.name].status = "открыт"

					ret = ITEMS[obj.name].name + " " + ITEMS[obj.name].status + ITEMS[obj.name].postfix

					return ret
				}
			}
		}
	}

	return "не к чему применить"
}

func handleCommand(s string) string {
	params := strings.Split(s, " ")

	switch params[0] {
	case "осмотреться":
		return look_around()
	case "идти":
		if len(params) < 2 {
			return "неверный формат команды"
		}

		return move(params[1])
	case "взять":
		if len(params) < 2 {
			return "неверный формат команды"
		}

		return take(params[1])
	case "одеть":
		if len(params) < 2 {
			return "неверный формат команды"
		}

		return player.wear(params[1])
	case "применить":
		if len(params) < 3 {
			return "неверный формат команды"
		}

		return apply(
			Item{
				name: params[1],
			}, Item{
				name: params[2],
			},
		)
	default:
		return "неизвестная команда"
	}
}

func initGame() {
	init_tasks()
	init_items()
	init_rooms()

	DOOR.status = "закрыт"

	INIT_ROOM = &KITCHEN
	player.room = INIT_ROOM

	player.items = []*Item{}
}

// < DONE
func main() {}

// > DONE
