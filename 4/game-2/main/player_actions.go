// < DONE
package main

import (
	"strings"
)

// > DONE

func (player *Player) look_around() {
	var ret string

	game := player.game
	room := player.room

	if room == game.init_room {
		ret += room.environment + ", " + player.get_items() + ", " + "надо " + player.get_tasks()
	} else {
		ret += player.get_items()
	}

	ret += ". "

	ret += player.can_move()

	pl := player.get_neightbours()

	if len(pl) > 0 {
		ret += ". Кроме вас тут ещё "

		var pl_s []string

		for i := range pl {
			if pl[i].name != player.name {
				pl_s = append(pl_s, pl[i].name)
			}
		}

		ret += join(", ", pl_s)
	} else {
		ret += ". Кроме вас тут никого нет"
	}

	player.out.msg <- ret
}

func (player *Player) can_move() string {
	var ret string

	room := player.room

	if room.has_neighbour() {
		var str []string

		for i := range room.neighbour {
			str = append(str, (*room.neighbour[i]).name)
		}

		ret = "можно пройти - " + strings.Join(str, ", ")
	} else {
		ret = "нельзя никуда пройти"
	}

	return ret
}

func (player *Player) move(room_name string) {
	var ret string

	if room, ok := player.get_neighbour(room_name); ok {

		if ok, fault_msg := room.in_condition(); !ok {
			player.out.msg <- fault_msg

			return
		}

		player.room = room

		ret = room.move_print + ". " + player.can_move()
	} else {
		ret = "нет пути в " + room_name
	}

	player.out.msg <- ret
}

func (player *Player) take(item_name string) {
	var ret string

	game := player.game
	room := player.room

	if _, ok := player.get_item(&Backpack); ok {
		if item, ok := room.get_item(item_name); ok {
			player.add_item(item)
			room.remove_item(item)

			ret = "предмет добавлен в инвентарь: " + item.name

			for _, v := range game.Players {
				if v.name != player.name{
					v.out.msg <- ""
				}
			}
		} else {
			ret = "нет такого"
		}
	} else {
		ret = "некуда класть"
	}

	player.out.msg <- ret
}

func (player *Player) apply(subj Item, obj Item) {
	var ret string

	game := player.game

	if _, ok := player.get_item(&subj); !ok {
		ret = "нет предмета в инвентаре - " + subj.name

		player.out.msg <- ret

		return
	}

	if game.item_exists(subj.name) {
		items := game.items

		objX := items[subj.name]

		for i := range objX.apply_to {
			if objX.apply_to[i].name == obj.name {
				objY := items[obj.name]

				if objY.status == "закрыт" && objY.has_position("открыт") {
					items[obj.name].status = "открыт"

					ret = items[obj.name].name + " " + items[obj.name].status + items[obj.name].postfix

					player.out.msg <- ret

					return
				}
			}
		}
	}

	ret = "не к чему применить"

	player.out.msg <- ret
}

func (player *Player) say_player(player_name string, s []string) {
	var ret string

	in_players := player.get_neightbours()

	for i := range in_players {
		if in_players[i].name == player_name {
			if len(s) > 0 {
				ret = player.name + " говорит вам: " + strings.Join(s, " ")
			} else {
				ret = player.name + " выразительно молчит, смотря на вас"
			}

			in_players[i].out.msg <- ret

			return
		}
	}

	ret = "тут нет такого игрока"

	player.out.msg <- ret
}

func (player *Player) say(messages []string) {
	var ret string

	game := player.game

	ret = player.name + " говорит: " + strings.Join(messages, " ")

	for _, v := range game.Players {
		v.out.msg <- ret
	}
}
