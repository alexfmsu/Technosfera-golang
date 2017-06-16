// < DONE
package main

import (
	"gopkg.in/telegram-bot-api.v4"
	"net/http"
	"strings"
	"runtime"
	"log"
	"os"
	"time"
)

// > DONE

// < NOT DONE ? global
// Rooms
var Kitchen Room
var Hall Room
var Flat Room
var Street Room
var HOME Room // ? why
// Rooms
// > NOT DONE ? global

// < NOT DONE ? global
// Items
var Backpack = Item{
	name: "рюкзак",

	can_take: false,
	can_wear: true,
}

var Door = Item{
	name:      "дверь",
	positions: []string{"открыт", "закрыт"},
	status:    "закрыт",

	init_room: &Hall,

	can_take: false,
	can_wear: false,

	postfix: "а",
}

var Wardrobe = Item{
	name:      "шкаф",
	positions: []string{"открыт", "закрыт"},
	status:    "закрыт",

	init_room: &Flat,

	can_take: false,
	can_wear: true,

	postfix: "",
}

var Key = Item{
	apply_to: []*Item{&Door},

	init_room: &Flat,

	can_take: true,
	can_wear: false,
}

// > NOT DONE ? global
// Items

// < NOT DONE
func get_name(msg *tgbotapi.Message) string {
	var name string

	if username := msg.Chat.UserName; username != "" {
		name = username
	} else {
		name = msg.Chat.FirstName
	}

	return name
}

// > NOT DONE

const WebhookURL = "https://morning-dawn-22248.herokuapp.com/" // needs config

func add_button(buttons []tgbotapi.KeyboardButton, button_name string) []tgbotapi.KeyboardButton {
	buttons = append(buttons, tgbotapi.KeyboardButton{Text: button_name})

	return buttons
}

func (p *Player) get_options(buttons []tgbotapi.KeyboardButton) []tgbotapi.KeyboardButton {
	room := p.room

	for i := range room.neighbour {
		if room.neighbour[i].name != room.name {
			buttons = add_button(buttons, "идти "+room.neighbour[i].name)
		}
	}

	items := room.items
	for i := range items {
		for j := range items[i].items {
			if items[i].items[j].can_wear {
				buttons = add_button(buttons, "одеть "+items[i].items[j].name)
			}

			if items[i].items[j].can_take {
				buttons = add_button(buttons, "взять "+items[i].items[j].name)
			}
		}
	}

	if room == &Hall {
		buttons = add_button(buttons, "применить ключи дверь")
	}

	return buttons
}

var loginFormTmpl = `
<html>
	<body>
	<form action="/get_cookie" method="post">
		Login: <input type="text" name="login">
		Password: <input type="password" name="password">
		<input type="submit" value="Login">
	</form>
	</body>
</html>
`


func main() {
	var game Game
	game.initGame()

	port := os.Getenv("PORT")

	bot, err := tgbotapi.NewBotAPI("300176172:AAGKX-LiiGxyZklOWBF9Q8HfT2jdqYdJ5PM")

	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	_, err = bot.SetWebhook(tgbotapi.NewWebhook(WebhookURL))

	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		sessionID, err := r.Cookie("session_id")
		if err == http.ErrNoCookie {
			w.Write([]byte(loginFormTmpl))
			return
		} else if err != nil {
			PanicOnErr(err)
		}
		fmt.Fprint(w, "Welcome, "+sessionID.Value)
	})

	// http.HandleFunc("/get_cookie", func(w http.ResponseWriter, r *http.Request) {
	// 	r.ParseForm()
	// 	inputLogin := r.Form["login"][0]
	// 	expiration := time.Now().Add(365 * 24 * time.Hour)
	// 	cookie := http.Cookie{
	// 		Name:    "session_id",
	// 		Value:   inputLogin,
	// 		Expires: expiration,
	// 	}
	// 	http.SetCookie(w, &cookie)
	// 	http.Redirect(w, r, "/", http.StatusFound)
	// })

	updates := bot.ListenForWebhook("/")

	go http.ListenAndServe(":"+port, nil)

	for update := range updates {
		var name = get_name(update.Message)

		// var player = game.Players[name]

		ID := update.Message.Chat.ID
		MYNAME := name

		if strings.ToLower(update.Message.Text) == "новая игра" || update.Message.Text == "/start" {
			_, ok := game.Players[MYNAME]

			if !ok{
				game.addPlayer(NewPlayer(MYNAME))
				
			}else{
				game.Players[MYNAME].timer.Stop()
			}

			var buttons = []tgbotapi.KeyboardButton{}

			var message = tgbotapi.NewMessage(update.Message.Chat.ID, "ты в игре")
            buttons = add_button(buttons, "Новая игра")
            buttons = add_button(buttons, "осмотреться")
           	buttons = game.Players[MYNAME].get_options(buttons)

			message.ReplyMarkup = tgbotapi.NewReplyKeyboard(buttons)
			bot.Send(message)
			
			go game.Players[MYNAME].wait_player()
            

            go func() {
                var player = game.Players[MYNAME]

				for msg := range player.GetOutput() {
					var buttons = []tgbotapi.KeyboardButton{}

                    switch player.room {
                    case &Street:
                   	 	buttons = []tgbotapi.KeyboardButton{tgbotapi.KeyboardButton{Text: "Новая игра"}}
                        
                        message = tgbotapi.NewMessage(ID, "Поздравляем, вы прошли игру! Можете начать новую")
                    case nil:
                        message = tgbotapi.NewMessage(ID, msg)
                        
                        buttons = []tgbotapi.KeyboardButton{tgbotapi.KeyboardButton{Text: "Новая игра"}}
                    	
                    	player.remove()
                    default:
                        buttons = add_button(buttons, "Новая игра")
                        buttons = add_button(buttons, "осмотреться")
                        buttons = player.get_options(buttons)

                        message = tgbotapi.NewMessage(ID, msg)
                    } 
					
					message.ReplyMarkup = tgbotapi.NewReplyKeyboard(buttons)
					bot.Send(message)

					if player.room == &Street{
						game.initGame()

						return
					}
				}
			}()

		} else {
			game.Players[name].HandleInput(update.Message.Text)
			time.Sleep(time.Microsecond)
			runtime.Gosched()

			game.Players[name].timer.Stop()
			go game.Players[name].wait_player()
		}
	}
}
