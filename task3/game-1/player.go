// < DONE
package main

import (
	"strings"
	"sync"
)

type Player struct {
	name  string
	room  *Room
	items []*Item
	Tasks []*task
	in    *Chan
	out   *Chan
}

type Chan struct {
	msg chan string
	sync.Mutex
}

// > DONE

var PLAYERS []*Player

func NewPlayer(player_name string) *Player {
	var p Player

	p = Player{
		name: player_name,
		room: &KITCHEN,
		in: &Chan{
			msg: make(chan string),
		},
		out: &Chan{
			msg: make(chan string),
		},
		Tasks: []*task{
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
		},
	}

	return &p
}

// < DONE
func (player *Player) get_item(item *Item) (*Item, bool) {
	for i := range player.items {
		if player.items[i].name == item.name {
			return player.items[i], true
		}
	}

	return nil, false
}

func addPlayer(p *Player) {
	PLAYERS = append(PLAYERS, p)
}

func (p *Player) say(s []string) {
	var ret string

	ret = p.name + " говорит: " + strings.Join(s, " ")

	for i := range PLAYERS {
		PLAYERS[i].out.msg <- ret
	}
}

func (p *Player) get_neightbours() []*Player {
	var pl []*Player

	for i := range PLAYERS {
		if PLAYERS[i].room == p.room {
			pl = append(pl, PLAYERS[i])
		}
	}

	return pl
}

func (p *Player) say_player(player_name string, s []string) {
	var ret string

	pl := p.get_neightbours()
	for i := range pl {
		if pl[i].name == player_name {
			if len(s) > 0 {
				ret = p.name + " говорит вам: " + strings.Join(s, " ")
			} else {
				ret = p.name + " выразительно молчит, смотря на вас"
			}

			pl[i].out.msg <- ret

			return
		}
	}

	ret = "тут нет такого игрока"

	p.out.msg <- ret
}

func (p *Player) HandleInput(cmd string) {
	wg.Add(2)

	IN <- &Command{
		player:  p,
		command: cmd,
	}

	wg.Wait()
}

func (p *Player) GetOutput() chan string {
	msg := p.out.msg

	return msg
}

// > DONE

// < NOT DONE tasks
func (p *Player) wear(item_name string) {
	var ret string

	if item, ok := p.room.has_item(item_name); ok {
		p.add_item(item)
		p.room.remove_item(item)

		for i := range p.Tasks {
			if p.Tasks[i].is_done() {
				p.Tasks = remove_task(p.Tasks, i)

				break
			}
		}

		ret = "вы одели: " + item_name
	} else {
		ret = "нечего одевать"
	}

	p.out.msg <- ret
}

// > NOT DONE tasks

// < DONE
func (player *Player) add_item(item *Item) {
	player.items = append(player.items, item)
}

// > DONE

func (p *Player) handleCommand(s string) {
	params := strings.Split(s, " ")

	switch params[0] {
	case "осмотреться":
		p.look_around()
	case "идти":
		if len(params) < 2 {
			p.out.msg <- "неверный формат команды"
		} else {
			p.move(params[1])
		}
	case "взять":
		if len(params) < 2 {
			p.out.msg <- "неверный формат команды"
		} else {
			p.take(params[1])
		}
	case "одеть":
		if len(params) < 2 {
			p.out.msg <- "неверный формат команды"
		} else {
			p.wear(params[1])
		}
	case "применить":
		if len(params) < 3 {
			p.out.msg <- "неверный формат команды"
		} else {
			p.apply(
				Item{
					name: params[1],
				}, Item{
					name: params[2],
				},
			)
		}
	case "сказать":
		if len(params) < 2 {
			p.out.msg <- "неверный формат команды"
		} else {
			p.say(params[1:])
		}
	case "сказать_игроку":
		if len(params) < 2 {
			p.out.msg <- "неверный формат команды"
		} else {
			p.say_player(params[1], params[2:])
		}
	default:
		p.out.msg <- "неизвестная команда"
	}

	wg.Done()

}
