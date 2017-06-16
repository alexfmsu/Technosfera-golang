package main
import (
    "encoding/json"
    "gopkg.in/telegram-bot-api.v4"
    "io/ioutil"
    "log"
    "net/http"
    "os"
)
// для вендоринга используется GB
// сборка проекта осуществляется с помощью gb build
// установка зависимостей - gb vendor fetch gopkg.in/telegram-bot-api.v4
// установка зависимостей из манифеста - gb vendor restore
type Joke struct {
    ID   uint32 `json:"id"`
    Joke string `json:"joke"`
}
type JokeResponse struct {
    Type  string `json:"type"`
    Value Joke   `json:"value"`
}
var buttons = []tgbotapi.KeyboardButton{
    tgbotapi.KeyboardButton{Text: "Get Joke"},
}
// При старте приложения, оно скажет телеграму ходить с обновлениями по этому URL
const WebhookURL = "https://msu-go-2017.herokuapp.com/"
func getJoke() string {
    c := http.Client{}
    resp, err := c.Get("http://api.icndb.com/jokes/random?limitTo=[nerdy]")
    if err != nil {
        return "jokes API not responding"
    }
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    joke := JokeResponse{}
    err = json.Unmarshal(body, &joke)
    if err != nil {
        return "Joke error"
    }
    return joke.Value.Joke
}
func main() {
    // Heroku прокидывает порт для приложения в переменную окружения PORT
    port := os.Getenv("PORT")
    bot, err := tgbotapi.NewBotAPI("270497469:AAFBnCofRqpaZIcBD_Re2N5aVbTcvWe8XDw")
    if err != nil {
        log.Fatal(err)
    }
    bot.Debug = true
    log.Printf("Authorized on account %s", bot.Self.UserName)
    // Устанавливаем вебхук
    _, err = bot.SetWebhook(tgbotapi.NewWebhook(WebhookURL))
    if err != nil {
        log.Fatal(err)
    }
    updates := bot.ListenForWebhook("/")
    go http.ListenAndServe(":"+port, nil)
    // получаем все обновления из канала updates
    for update := range updates {
        var message tgbotapi.MessageConfig
        log.Println("received text: ", update.Message.Text)
        switch update.Message.Text {
        case "Get Joke":
            // Если пользователь нажал на кнопку, то придёт сообщение "Get Joke"
            message = tgbotapi.NewMessage(update.Message.Chat.ID, getJoke())
        default:
            message = tgbotapi.NewMessage(update.Message.Chat.ID, `Press "Get Joke" to receive joke`)
        }
        // В ответном сообщении просим показать клавиатуру
        message.ReplyMarkup = tgbotapi.NewReplyKeyboard(buttons)
        bot.Send(message)
    }
}
// package main

// import (
//  "log"
//  "net/http"
//  "os"
//  "fmt"
//  "github.com/gin-gonic/gin"
//  "game"
// )

// func main() {
//   port := os.Getenv("PORT")

//   if port == "" {
//       log.Fatal("$PORT must be set")
//   }

//   init_game()

//   router := gin.New()
//   router.Use(gin.Logger())
//   router.LoadHTMLGlob("templates/*.tmpl.html")
//   router.Static("/static", "static")

//   router.GET("/", func(c *gin.Context) {
//       c.HTML(http.StatusOK, "index.tmpl.html", nil)
//   })

//   fmt.Println("123\n\n\n\n")

//   router.Run(":" + port)
// }

// package main

// import (
// 	"strings"
// 	"sync"
// )

// type Command struct {
// 	player  *Player
// 	command string
// }

// var wg sync.WaitGroup

// // Rooms
// var KITCHEN Room
// var FLAT Room
// var HALL Room
// var STREET Room
// var HOME Room

// var ROOMS []*Room

// var INIT_ROOM *Room

// // Rooms

// // Tasks
// var TASKS []*task

// // Tasks

// // Items
// var BACKPACK = Item{
// 	name: "рюкзак",
// }

// var ITEMS map[string]*Item

// var DOOR = Item{
// 	name:      "дверь",
// 	positions: []string{"открыт", "закрыт"},
// 	status:    "закрыт",
// 	postfix:   "а",
// }

// var WARDROBE = Item{
// 	name:      "шкаф",
// 	positions: []string{"открыт", "закрыт"},
// 	status:    "закрыт",
// 	postfix:   "",
// }

// var KEY = Item{
// 	apply_to: []*Item{&DOOR},
// }

// // Items

// func (p *Player) look_around() {
// 	var ret string

// 	if p.room == INIT_ROOM {
// 		ret += p.room.environment + ", " + p.get_items() + ", " + "надо " + p.get_tasks() + ". "
// 	} else {
// 		ret += p.get_items() + ". "
// 	}

// 	ret += p.can_move()

// 	pl := p.get_neightbours()

// 	if len(pl) > 1 {
// 		ret += ". Кроме вас тут ещё "

// 		for i := range pl {
// 			if pl[i].name != p.name {
// 				ret += pl[i].name
// 			}
// 		}
// 	} else if len(PLAYERS) != 1 {
// 		ret += "Кроме вас тут никого нет"
// 	}

// 	p.out.msg <- ret
// }

// func (p *Player) can_move() string {
// 	var ret string

// 	if p.room.has_neighbour() {
// 		var str []string

// 		for i := range p.room.neighbour {
// 			str = append(str, (*p.room.neighbour[i]).name)
// 		}

// 		ret = "можно пройти - " + strings.Join(str, ", ")
// 	} else {
// 		ret = "нельзя никуда пройти"
// 	}

// 	return ret
// }

// func (p *Player) move(room_name string) {
// 	var ret string

// 	if room, ok := p.get_neighbour(room_name); ok {

// 		if ok, fault_msg := room.in_condition(); !ok {
// 			p.out.msg <- fault_msg

// 			return
// 		}

// 		p.room = room

// 		ret = p.room.move_print + ". " + p.can_move()
// 	} else {
// 		ret = "нет пути в " + room_name
// 	}

// 	p.out.msg <- ret
// }

// func (p *Player) take(item_name string) {
// 	var ret string

// 	if _, ok := p.get_item(&BACKPACK); ok {
// 		if item, ok := p.room.has_item(item_name); ok {
// 			p.add_item(item)
// 			p.room.remove_item(item)

// 			ret = "предмет добавлен в инвентарь: " + item.name
// 		} else {
// 			ret = "нет такого"
// 		}
// 	} else {
// 		ret = "некуда класть"
// 	}

// 	p.out.msg <- ret
// }

// func (p *Player) apply(subj Item, obj Item) {
// 	var ret string

// 	if _, ok := p.get_item(&subj); !ok {
// 		ret = "нет предмета в инвентаре - " + subj.name

// 		p.out.msg <- ret

// 		return
// 	}

// 	if item_exists(subj.name) {
// 		objX := ITEMS[subj.name]

// 		for i := range objX.apply_to {
// 			if objX.apply_to[i].name == obj.name {
// 				objY := ITEMS[obj.name]

// 				if objY.status == "закрыт" && objY.has_position("открыт") {
// 					ITEMS[obj.name].status = "открыт"

// 					ret = ITEMS[obj.name].name + " " + ITEMS[obj.name].status + ITEMS[obj.name].postfix

// 					p.out.msg <- ret

// 					return
// 				}
// 			}
// 		}
// 	}

// 	ret = "не к чему применить"

// 	p.out.msg <- ret
// }

// var IN chan *Command

// func run() {
// 	for cmd := range IN {
// 		cmd.player.handleCommand(cmd.command)

// 		wg.Done()
// 	}
// }

// func initGame() {
// 	init_rooms()
// 	init_items()

// 	DOOR.status = "закрыт"

// 	IN = make(chan *Command)

// 	INIT_ROOM = &KITCHEN

// 	PLAYERS = make([]*Player, 0)

// 	go run()
// }

// func main() {}
