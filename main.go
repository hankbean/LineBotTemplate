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
	"time"
	
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
					if message.Text=="早餐" || message.Text=="午餐"  || message.Text=="晚餐"{
						//bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("早中晚？(有bug晚點修 可以催稿)")).Do()
						if message.Text=="早餐"{
						rand.Seed(time.Now().UnixNano()) // Try changing this number!
						answers := []string{
							"全家",
							"愛瘋牛排",
							"華美自助餐",
							"華美丼飯",
							"鐵板便當",
							"滷味",
							"7-11",
							"全家",
							"台南意麵",
							"淡江炒飯",
							"Xbuger",
							"赤鳥家",
							"要減肥了",
							"台北城",
							"台北煮",
							"十全",
							"麥當勞",
							"學餐",
							"阿羅哈",
							"大Q",
							"大紅袍",
						}
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(answers[rand.Intn(len(answers))])).Do()
						} else if message.Text=="午餐"{
						rand.Seed(time.Now().UnixNano()) // Try changing this number!
						answers := []string{
							"感恩",
							"愛瘋",
							"華美",
							"鐵板",
							"滷味",
							"7-11",
							"全家",
							"台南意麵",
							"淡江",
							"輔大",
							"赤鳥家",
							"要減肥了",
							"台北城",
							"台北煮",
							"十全",
							"麥當勞",
							"學餐",
							"阿羅哈",
							"大Q",
							"大紅袍",
						}
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(answers[rand.Intn(len(answers))])).Do()
						} else if message.Text=="晚餐"{
							rand.Seed(time.Now().UnixNano()) // Try changing this number!
						answers := []string{
							"感恩",
							"愛瘋",
							"華美",
							"鐵板",
							"滷味",
							"7-11",
							"全家",
							"台南意麵",
							"淡江炒飯",
							"輔大豬排",
							"赤鳥家",
							"要減肥了",
							"台北城",
							"台北煮",
							"十全",
							"麥當勞",
							"學餐",
							"宵夜快餐",
							"大Q",
							"大紅袍",
						}
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("早中晚？(有bug晚點修 可以催稿)")).Do()
						} else {
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("指令錯誤")).Do()
						}


					} else if message.Text=="456"{
						rand.Seed(time.Now().UnixNano()) // Try changing this number!
						answers := []string{
							"感恩",
							"愛瘋",
							"華美",
							"鐵板",
							"滷味",
							"7-11",
							"全家",
							"台南意麵",
							"淡江",
							"輔大",
							"赤鳥家",
							"要減肥了",
							"台北城",
							"台北煮",
							"十全",
							"麥當勞",
							"學餐",
							"阿羅哈",
							"大Q",
							"大紅袍",
						}
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(answers[rand.Intn(len(answers))])).Do()
					} else if message.Text=="789"{
							rand.Seed(time.Now().UnixNano()) // Try changing this number!
						answers := []string{
							"感恩",
							"愛瘋",
							"華美",
							"鐵板",
							"滷味",
							"7-11",
							"全家",
							"台南意麵",
							"淡江炒飯",
							"輔大豬排",
							"赤鳥家",
							"要減肥了",
							"台北城",
							"台北煮",
							"十全",
							"麥當勞",
							"學餐",
							"宵夜快餐",
							"大Q",
							"大紅袍",
						}
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(answers[rand.Intn(len(answers))])).Do()
					} else if message.Text=="今日運勢"{
						rand.Seed(time.Now().UnixNano()) // Try changing this number!
						answers := []string{
							"愚人",
							"魔術師",
							"女教皇",
							"皇后",
							"國王",
							"教皇",
							"戀人",
							"戰車",
							"力量",
							"隱者",
							"命運之輪",
							"正義",
							"吊人",
							"死神",
							"節制",
							"惡魔",
							"塔",
							"星星",
							"月亮",
							"太陽",
							"審判",
							"世界",
						}
						turn := []string{
							"正位",
							"逆位",
						}
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(turn[rand.Intn(len(turn))]+answers[rand.Intn(len(answers))])).Do()
					
					} else if message.Text=="艾路"{
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("卡娜赫拉")).Do()

					} else if message.Text=="厨ニ" || message.Text=="廚二"{
						rand.Seed(time.Now().UnixNano()) // Try changing this number!
						answers := []string{
							"爆裂吧，現實。迸裂吧，精神。放逐這個世界！",
							"闇のほのにだかえで　消えろ!!",
							"闇の炎に抱かれて死ね！",
							"El Psy Congroo",
							"隐藏着黑暗力量的钥匙啊!",
							"我要代表月亮，消灭你！~",
							"既然你诚心诚意的请教了！我们就大发慈悲的告诉你！",
							"生きているものなら、神様も杀して见せる。",
							"只要是活着的东西，就算是神我也杀给你看 。",
							"僕は新世界の神となる!",
							"真実はいつも一つ！！",
							"真相只有一个！！",
							"人被杀，就会死",
							"东中出身 凉宫ハルヒ ただの人间には兴味ありません この中に宇宙人 ·未来人·超能力者がいだら あたしのところに来なさいっ 以上 ",
							"你已經死了！",
							"我要成為新世界的神！",
							"愉悅！",
							"你那無聊的幻想由我來打破！",
							"一切都是命運石之門的選擇！",
							"我不做人啦！JOJO",
						}
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(answers[rand.Intn(len(answers))])).Do()

					} else {
						//bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("指令清單\r\n"+"早餐/午餐/晚餐\r\n今日運勢")).Do()
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("魔法指令有誤")).Do()
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
				//bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text+"\r\n"+"嗨啊早餐吃了沒")).Do()
			}
		}
	}
}
