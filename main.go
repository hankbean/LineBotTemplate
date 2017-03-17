// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"math/rand"
	
	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

func main() {
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
				case *linebot.TextMessage:
					//if _, err = 
					if message.Text=="早餐"{
						rand.Seed(42) // Try changing this number!
						answers := []string{
							"It is certain",
							"It is decidedly so",
							"Without a doubt",
							"Yes definitely",
							"You may rely on it",
							"As I see it yes",
							"Most likely",
							"Outlook good",
							"Yes",
							"Signs point to yes",
							"Reply hazy try again",
							"Ask again later",
							"Better not tell you now",
							"Cannot predict now",
							"Concentrate and ask again",
							"Don't count on it",
							"My reply is no",
							"My sources say no",
							"Outlook not so good",
							"Very doubtful",
	}
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(answers[rand.Intn(len(answers)))).Do()
					} else if message.Text=="午餐"{
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("不要吃啦")).Do()
					} else {
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text+"\r\n"+"嗨啊早餐吃了沒")).Do()
					}
					//; err != nil {
					//	log.Print(err)
					//}
					//fmt.Println("hello world")
					//fmt.Println(message.ID)
				//case "今日運勢":
				//	if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("節制")).Do(); err != nil {
				//		log.Print(err)
				//	}
			}
		}
	}
}
