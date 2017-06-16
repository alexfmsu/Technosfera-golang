package main

import (
	"strings"
	"sync"
	"time"
)

type Player struct {
	name string

	game  *Game
	room  *Room
	items []*Item
	tasks []*task

	in  *Chan
	out *Chan

	timer *time.Timer
}

type Chan struct {
	msg chan string
	sync.Mutex
}

type Command struct {
	player  *Player
	command string
	wg      *sync.WaitGroup
}

// > DONE

func NewPlayer(player_name string) *Player {
	var player Player

	player = Player{
		name: player_name,
		in: &Chan{
			msg: make(chan string),
		},
		out: &Chan{
			msg: make(chan string),
		},
		tasks: []*task{
			&task{
				name: "собрать рюкзак",
				is_done: func() bool {
					_, ok := player.get_item(&Backpack)
					return ok
				},
			},
			&task{
				name: "идти в универ",
				is_done: func() bool {
					return player.room == &Street
				},
			},
		},
	}

	return &player
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

func (game *Game) addPlayer(player *Player) {
	player_name := player.name

	game.Players[player_name] = player

	player.game = game
	player.room = game.init_room
}

func (player *Player) get_neightbours() []*Player {
	var in_players []*Player

	game := player.game
	Players := game.Players

	for _, v := range Players {
		if v.room == player.room && v.name != player.name {
			in_players = append(in_players, v)
		}
	}

	return in_players
}

func (player *Player) HandleInput(cmd string) {
	var wg sync.WaitGroup

	wg.Add(1)

	player.game.cmd_chan <- &Command{
		player:  player,
		command: cmd,
		wg:      &wg,
	}
}

func (player *Player) GetOutput() chan string {
	msg := player.out.msg

	return msg
}

// < NOT DONE tasks
func (player *Player) wear(item_name string) {
	var ret string

	room := player.room

	if item, ok := room.get_item(item_name); ok {
		player.add_item(item)
		room.remove_item(item)

		for i := range player.tasks {
			if player.tasks[i].is_done() {
				player.tasks = remove_task(player.tasks, i)

				break
			}
		}

		ret = "вы одели: " + item_name
	} else {
		ret = "нечего одевать"
	}

	player.out.msg <- ret
}

// > NOT DONE tasks

// < DONE
func (player *Player) add_item(item *Item) {
	player.items = append(player.items, item)
}

// > DONE

func (player *Player) handleCommand(s string, wg *sync.WaitGroup) {
	game := player.game
	wrong_cmd := game.messages["wrong_commnand"]

	params := strings.Split(s, " ")

	params[0] = strings.ToLower(params[0])

	switch params[0] {
	case "осмотреться":
		player.look_around()
	case "идти":
		if len(params) < 2 {
			player.out.msg <- wrong_cmd
		} else {
			params[1] = strings.ToLower(params[1])
			player.move(params[1])
		}
	case "взять":
		if len(params) < 2 {
			player.out.msg <- wrong_cmd
		} else {
			params[1] = strings.ToLower(params[1])
			player.take(params[1])
		}
	case "одеть":
		if len(params) < 2 {
			player.out.msg <- wrong_cmd
		} else {
			params[1] = strings.ToLower(params[1])
			player.wear(params[1])
		}
	case "применить":
		if len(params) < 3 {
			player.out.msg <- wrong_cmd
		} else {
			player.apply(
				Item{
					name: params[1],
				}, Item{
					name: params[2],
				},
			)
		}
	case "сказать":
		if len(params) < 2 {
			player.out.msg <- wrong_cmd
		} else {
			player.say(params[1:])
		}
	case "сказать_игроку":
		if len(params) < 2 {
			player.out.msg <- wrong_cmd
		} else {
			player.say_player(params[1], params[2:])
		}
	default:
		player.out.msg <- wrong_cmd
	}

	wg.Done()
}

func (player *Player) wait_player() {
	game := player.game
	timeout := game.wait_timeout

	player.timer = time.NewTimer(time.Second * timeout)

	<-player.timer.C

	player.out.msg <- "игрок " + player.name + " был выкинут из игры за неактивность"

	player.room = nil
}

func (p *Player) remove() {
	items := p.items

	for i := range items {
		if items[i].can_take || items[i].can_wear {
			room := items[i].init_room
			
			room.items = append(room.items, items[i])
		}
	}

	p.items = []*Item{}

	game := p.game

	delete(game.Players, p.name)
}
