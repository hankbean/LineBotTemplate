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
	"strconv"
	
	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

func main() {
	f, err := os.OpenFile("logfile.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
    if err != nil {
    	log.Fatalf("file open error : %v", err)
    }
    defer f.Close()
    log.SetOutput(f)
    log.Println("This is a test log entry")

	//var err error
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

					fmt.Println(message.Text)
					if message.Text=="早餐" || message.Text=="午餐"  || message.Text=="晚餐"|| message.Text=="吃什麼"{
						//bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("早中晚？(有bug晚點修 可以催稿)")).Do()
						if message.Text=="早餐" || message.Text=="午餐"  || message.Text=="晚餐"|| message.Text=="吃什麼"{
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
						} else if message.Text=="a午餐"{
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
						} else if message.Text=="a晚餐"{
							bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("早中晚？(有bug晚點修 可以催稿)")).Do()
						} else {
							// bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("指令錯誤")).Do()
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
					} else if message.Text=="抽牌" || message.Text=="抽大牌"{
						rand.Seed(time.Now().UnixNano()) // Try changing this number!
						
						mesText := "";

						turn := []string{
							"正位",
							"逆位",
						}

						majorArcana := []string{
							"愚人",
							"魔術師",
							"女教皇",
							"皇后",
							"皇帝",
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
						
						minorArcanaName := []string{
							"劍",
							"杖",
							"杯",
							"幣",
						}

						minorArcanaNum := []string{
							"1",
							"2",
							"3",
							"4",
							"5",
							"6",
							"7",
							"8",
							"9",
							"10",
							"侍從",
							"騎士",
							"皇后",
							"國王",
						}

						var ifNum = rand.Intn(78-1)
						fmt.Println(ifNum)
						if (message.Text=="抽大牌" || ifNum >= (1-1) && ifNum < (22-1)){
							mesText = turn[rand.Intn(len(turn))] + majorArcana[rand.Intn(len(majorArcana))]
						} else {
							mesText = turn[rand.Intn(len(turn))] + minorArcanaName[rand.Intn(len(minorArcanaName))] + minorArcanaNum[rand.Intn(len(minorArcanaNum))]
						}

						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(mesText)).Do()

					} else if message.Text=="骰子卡" {
						rand.Seed(time.Now().UnixNano())
						mesText := "" ;
						starNum := strconv.Itoa(rand.Intn(11)+1);
						signNum := strconv.Itoa(rand.Intn(11)+1);
						palaceNum := strconv.Itoa(rand.Intn(11)+1);
						star := []string {
							"月亮",
							"水星",
							"金星",
							"地球",
							"火星",
							"木星",
							"土星",
							"天王星",
							"海王星",
							"冥王星",
							"凱隆星",
							"北交點",
						} 
						mesText = star[starNum] + "***" + signNum + "***" + palaceNum;
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(mesText)).Do()

					} else if message.Text=="進階骰子卡" {
						rand.Seed(time.Now().UnixNano())
						mesText := "";
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(mesText)).Do()
					
					} else if message.Text=="艾路"{
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("卡娜赫拉")).Do()

					} else if message.Text=="中二" || message.Text=="廚二"{
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

					} else if message.Text=="文大吃吃" || message.Text=="吃吃精靈" || message.Text=="文大吃吃精靈" || message.Text=="吃吃"{
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("額咪啊")).Do()

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
							"直線七秒很簡單啊 我三秒就過了",
							"每對父母平均一人有一顆睪丸",
							"人被殺，就會死",
							"台灣競爭力低落，在美國就連小學生都會說流利的英語",
							"南半球的一隻蝴蝶拍了拍翅膀 他就上升了一點點",
							"不要怨恨自己沒幹到正妹 連正妹自己都沒幹過自己",
							"愉悅！",
							"w",
							"那天看中醫 醫師問我 是不是冬天都會感到特別冷 喝太多水就想尿尿",
							"每當你往上踩一層階梯   你與天空的距離就會縮短",
							"誰能知道這18歲的少女，4年前居然才14歲",
							"為什麼自古以來紅顏多薄命呢 因為根本沒人在乎醜的人活多久",
							"為什麼襪子總會不見一隻 因為不見兩隻你根本不會發現",
							"根據研究，在北京呼吸霾霧很可怕，每過1小時，會折損3600秒的壽命",
							// "為什麼吳浩宇那麼醜，不知道，也沒人想知道",
							"戒菸很簡單，我已經戒過好幾次了",
							"當黑人演員很慘，他們永遠只能演黑人",
							"人要是死了 那就真的死了",
							"明天過了，後天就會來到",
							"有些東西就是要在打折的時候買比較便宜",
							"人就跟樹一樣 你拿斧頭砍他 他就會死掉",
							"據統計，未婚生子的人數中有高機率為女性",
							"當你的左臉被人打，那你的左臉就會痛",
							"聽見你的名字時，會先想到你的樣子。",
							"只要跟你在一起，每天就是Everyday！",
							"放棄的話，就代表Give UP了喔！",
							"對今天解決不了的事情 不要太擔心 因為明天也解決不了",
							"心理學發現，聰明的人學東西比較快",
							"世上無難事 只要肯放棄",
							"努力不一定成功 但不努力一定很爽",
							"如果你氣憋得夠久 就可以睡到永遠",
							"統計顯示，所有的父母的年齡都比他們的親生孩子更大，而至今沒有任何科學研究去解釋這件事的原因。",
							"在美國，如果你不吃中餐，就會餓",
							"聽說20歲就可以交到女朋友了，現在我完成一半了，我20歲",
							"人如果半夜睡不著覺，就會變得格外清醒",
							"當你在洗澡時，有高機率是裸體",
							"若要人不知，除非你不要跟他講",
						}
						var ifNum = rand.Intn(9)
						if (ifNum==0){
							bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(answers[rand.Intn(len(answers))])).Do()
						// } else {
							//bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("debug: "+ifNum)).Do()
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
