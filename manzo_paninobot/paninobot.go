// telemanzobot project main.go
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

var bot *tb.Bot

// UserInfo serve per contenere informazioni sugli utenti
type UserInfo struct {
	firstname string
	lastname  string
	phone     string
	lat       float32
	lng       float32
}

var userinfo map[int]*UserInfo // ID->userinfo

func main() {

	b, err := tb.NewBot(tb.Settings{
		Token:  os.Getenv("TELEGRAM_API_TOKEN"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}
	bot = b
	userinfo = make(map[int]*UserInfo)

	bot.Handle(tb.OnQuery, menuHandler)
	bot.Handle("/help", helpHandler)
	bot.Handle("/consegna", consegnaHandler)
	bot.Handle("/conferma", confermaHandler)
	bot.Handle(tb.OnLocation, locationHandler)
	bot.Handle(tb.OnContact, contactHandler)

	bot.Start()
}

func checkUser(userid int) {
	_, ok := userinfo[userid]
	if !ok {
		userinfo[userid] = &UserInfo{}
	}
}

func menuHandler(q *tb.Query) {
	//TODO: get menu json from a webservice
	menu := tb.Results{
		&tb.ArticleResult{Title: "🍔 hamburgher", Text: "hamburger gustoso a soli 4€", HideURL: true},
		&tb.ArticleResult{Title: "🍕 pizza", Text: "trancio di margherita 3€", HideURL: true},
		&tb.ArticleResult{Title: "🍟 patatine", Text: "patatine, offerta solo 2€", HideURL: true},
		&tb.ArticleResult{Title: "🍰 torta", Text: "cheesecake al cioccolato! 3.50 €", HideURL: true},
	}

	for i := range menu {
		menu[i].SetResultID(strconv.Itoa(i)) // It's needed to set a unique string ID for each result
	}

	err := bot.Answer(q, &tb.QueryResponse{
		Results:   menu,
		CacheTime: 60, // a minute
	})

	if err != nil {
		log.Println(err)
	}
}

func helpHandler(m *tb.Message) {
	checkUser(m.Sender.ID)
	log.Printf("helpHandler: user=%s %s\n", m.Sender.FirstName, m.Sender.LastName)
	bot.Send(m.Sender, "usage: @nomebot /consegna /conferma")
}

func locationHandler(m *tb.Message) {

	userid := m.Sender.ID
	checkUser(userid)
	userinfo[userid].lat = m.Location.Lat
	userinfo[userid].lng = m.Location.Lng
	log.Printf("locationHandler: userid=%d pos=%f,%f\n", userid, m.Location.Lat, m.Location.Lng)

	bot.Send(m.Sender, fmt.Sprintf("ok %s ho la tua posizione", m.Sender.Username))
}

func contactHandler(m *tb.Message) {
	userid := m.Sender.ID
	checkUser(userid)
	userinfo[userid].phone = m.Contact.PhoneNumber
	userinfo[userid].firstname = m.Contact.FirstName
	userinfo[userid].lastname = m.Contact.LastName

	log.Printf("contactHandler: user=%s cellphone=%s\n", m.Sender.Username, m.Contact.PhoneNumber)
	bot.Send(m.Sender, fmt.Sprintf("ok, %s ho il tuo numero di telefono: %s", m.Sender.FirstName, m.Contact.PhoneNumber))
}

func consegnaHandler(m *tb.Message) {
	checkUser(m.Sender.ID)
	//get user information
	replyBtn1 := tb.ReplyButton{Text: "telefono", Contact: true}
	replyBtn2 := tb.ReplyButton{Text: "posizione", Location: true}
	replyBtn3 := tb.ReplyButton{Text: "cancel"}
	replyKeys := [][]tb.ReplyButton{
		[]tb.ReplyButton{replyBtn1, replyBtn2, replyBtn3},
		// ...
	}
	replyMarkup := tb.ReplyMarkup{ReplyKeyboard: replyKeys}
	bot.Send(m.Sender, "ciao, dicci chi sei!", &replyMarkup)
}

func confermaHandler(m *tb.Message) {

	userid := m.Sender.ID
	checkUser(userid)
	var msg string
	if info, ok := userinfo[userid]; ok {
		//TODO: gestisce pagamento
		time1 := time.Now().Local().Add(time.Minute * time.Duration(15))
		time2 := time1.Add(time.Minute * time.Duration(30))
		msg = fmt.Sprintf("ok, %s! Il tuo ordine arriverà tra le %s e le %s", info.firstname, time1.Format(time.Kitchen), time2.Format(time.Kitchen))
		//chiama webservice per processare ordine
	} else {
		msg = fmt.Sprintf("ciao %s ! usa /consegna per fornire le informazioni mancanti", m.Sender.FirstName)
	}
	bot.Send(m.Sender, msg)
}
