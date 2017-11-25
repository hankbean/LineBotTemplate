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
							"El Psy Congroo！",
							"隐藏着黑暗力量的钥匙啊!",
							"我要代表月亮，消灭你！~",
							"既然你誠心誠意的發問了,我們就大發慈悲的告訴你,為了防止世界被破壞,為了守護世界的和平,貫徹愛與真實的邪惡,可愛又迷人的反派角色,武藏！小次郎！我們是穿梭在銀河中的火箭隊,白洞、白色的明天正等著我們,就是這樣喵！",
							"生きているものなら、神様も杀して见せる。",
							"只要是活著的東西，就算是神我也殺給你看",
							"我對普通的人類沒有興趣，你們當中要是有外星人、未來人、異世界人以及超能力者的話，就儘管來找我吧！以上。",
							"由統括這個銀河系的資訊統合思念體，製造出來與有機生命體接觸用的聯繫裝置外星人，就是…我。",
							"這是禁止事項",
							"僕は新世界の神となる!",
							"真実はいつも一つ！！",
							"真相只有一个！！",
							"人被殺，就會死",
							"东中出身 凉宫ハルヒ ただの人间には兴味ありません この中に宇宙人 ·未来人·超能力者がいだら あたしのところに来なさいっ 以上 ",
							"你已經死了！",
							"我要成為新世界的神！",
							"愉悅！",
							"你那無聊的幻想由我來打破！",
							"一切都是命運石之門的選擇！",
							"我不做人啦！JOJO！",
						}
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(answers[rand.Intn(len(answers))])).Do()

					} else if message.Text=="豆豆" || message.Text=="吳浩宇"{
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("我不是變態...我是豆神!!")).Do()

					} else if message.Text=="祥瑀" || message.Text=="黃祥瑀"{
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("約炮小王子")).Do()

					} else if message.Text=="于姿婷" || message.Text=="黃玥萱"{
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("烤土司小公主")).Do()

					} else if message.Text=="博榮" || message.Text=="榮榮" || message.Text=="王博榮"{
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("封杯小王子")).Do()

					} else if message.Text=="躍萱" || message.Text=="黃躍萱"{
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("黃色洨話冠軍錦標賽第一屆傳承人")).Do()

					} else if message.Text=="萱萱"{
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你找躍萱還是玥萱")).Do()

					} else if message.Text=="大腸包小腸"{
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("チンチン大きいです")).Do()

					} else {
						rand.Seed(time.Now().UnixNano()) // Try changing this number!
						answers := []string{
							"在你面前閉氣的話，就會不能呼吸喔",
							"跟你在一起時，回憶一天前的事，就像回想昨天的事情",
							"二個人比一個人還多唷",
							"El Psy Congroo！",
							"隐藏着黑暗力量的钥匙啊!",
							"我要代表月亮，消灭你！~",
							"既然你誠心誠意的發問了,我們就大發慈悲的告訴你,為了防止世界被破壞,為了守護世界的和平,貫徹愛與真實的邪惡,可愛又迷人的反派角色,武藏！小次郎！我們是穿梭在銀河中的火箭隊,白洞、白色的明天正等著我們,就是這樣喵！",
							"你不在的這十二個月，對我來說就如同一年般長",
							"跟你通話的那個晚上，確實聽到了你的聲音",
							"不知道為什麼，把眼睛蒙上後什麼都看不到",
							"大部分的雞蛋料理，都要用上雞蛋喔",
							"在非洲 每一分鐘就有60秒過去",
							"出生時，大家都是裸體的喔",
							"0.0",
							"0.0",
							"人被殺，就會死",
							"台灣競爭力低落，在美國就連小學生都會說流利的英語",
							"0.0",
							"0.0",
							"愉悅！",
							"w",
							"那天看中醫 醫師問我 是不是冬天都會感到特別冷 喝太多水就想尿尿",
							"只要每天省下買一杯奶茶的錢，十天後就能買十杯奶茶",
						}
						if (rand.Intn(10)==1){
							bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(answers[rand.Intn(len(answers))])).Do()
						}

						//bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("指令清單\r\n"+"早餐/午餐/晚餐\r\n今日運勢")).Do()
						

						//bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("魔法指令有誤")).Do()
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
