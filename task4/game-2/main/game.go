// < DONE
package main

import "time"
// > DONE

// < NOT DONE except init_room, items, cmd_chan
type Game struct {
	Players map[string]*Player

	rooms     []*Room
	init_room *Room

	items map[string]*Item

	cmd_chan chan *Command

	messages map[string]string

	wait_timeout time.Duration
}

// > NOT DONE except init_room, items, cmd_chan

func (g *Game) initGame() {
	g.Players = map[string]*Player{}
	g.init_rooms()
	g.init_items()

	Door.status = "закрыт"

	g.cmd_chan = make(chan *Command)

	g.init_room = &Kitchen

	g.messages = map[string]string{
		"wrong_command": "неверный формат команды",
	}

	g.wait_timeout = 60

	go g.run()
}

func (g *Game) run() {
	for cmd := range g.cmd_chan {
		cmd.player.handleCommand(cmd.command, cmd.wg)

		cmd.wg.Wait()
	}
}
