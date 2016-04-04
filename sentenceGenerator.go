package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

var sentenceModel = []string{
	"11，22，33了44。\n1111，2222，不过是一场3333。\n你说4444，我说5555，最后不过6666。\n55，66，许我一场7777。\n一1一2一77，半3半4半88。",
	"你说11112222，后来33334444。\n5555，6666，终不敌7777。\n111，111，1122333"}
var fourWord = []string{"情深缘浅", "情深不寿", "莫失莫忘", "阴阳相隔", "如花美眷", "似水流年", "眉目如画", "曲终人散", "繁华落尽", "不诉离殇", "一世长安"}
var threeWord = []string{"冷时节", "空离别", "泪千行"}
var twoWord = []string{"朱砂", "天下", "杀伐", "人家", "韶华", "风华", "繁华", "血染", "墨染", "白衣",
	"素衣", "嫁衣", "倾城", "孤城", "空城", "旧城", "旧人", "伊人", "心疼", "春风", "古琴", "无情", "迷离", "奈何", "断弦", "焚尽", "散乱", "陌路", "乱世", "笑靥",
	"浅笑", "明眸", "轻叹", "烟火", "一生", "三生", "浮生", "桃花", "梨花", "落花", "烟花", "离殇", "情殇", "爱殇", "剑殇", "灼伤", "仓皇", "匆忙", "陌上", "清商", "焚香", "墨香", "微凉",
	"断肠", "痴狂", "凄凉", "黄梁", "未央", "成双", "无恙", "虚妄", "凝霜", "洛阳", "长安", "江南", "忘川", "千年", "纸伞", "烟雨", "回眸", "公子", "红尘", "红颜", "红衣", "红豆", "红线", "青丝", "青史",
	"青冢", "白发", "白首", "白骨", "黄土", "黄泉", "碧落", "紫陌"}

var oneWord = []string{"殇", "红", "离", "梦", "忖", "惨"}

func main() {
	var finalSentence string
	for i := 0; i < len(sentenceModel); i++ {
		newSentence := GenerateNewSentence(sentenceModel[i])
		finalSentence += "\n" + newSentence
	}
	show(finalSentence)
}

//GenerateNewSentence is a function used predetermined sentence model to generate the sentence;
func GenerateNewSentence(sentenceModel string) string {
	tempFourWordArray := make([]string, len(fourWord))
	copy(tempFourWordArray, fourWord)
	tempThreeWordArray := make([]string, len(threeWord))
	copy(tempThreeWordArray, threeWord)
	tempTwoWordArray := make([]string, len(twoWord))
	copy(tempTwoWordArray, twoWord)
	tempOneWordArray := make([]string, len(oneWord))
	copy(tempOneWordArray, oneWord)

	newSentence := sentenceModel
	substring := ""

	for i := 1; i < 9; i++ {
		substring = strconv.Itoa(i * 1111)
		for strings.Contains(newSentence, substring) {
			word := rand.Intn(len(tempFourWordArray))
			newSentence = strings.Replace(newSentence, substring, tempFourWordArray[word], -1)
			tempFourWordArray[word] = tempFourWordArray[len(tempFourWordArray)-1]
			tempFourWordArray = tempFourWordArray[:len(tempFourWordArray)-1]
		}

		substring = strconv.Itoa(i * 111)
		for strings.Contains(newSentence, substring) {
			word := rand.Intn(len(tempThreeWordArray))
			newSentence = strings.Replace(newSentence, substring, tempThreeWordArray[word], -1)
			tempThreeWordArray[word] = tempThreeWordArray[len(tempThreeWordArray)-1]
			tempThreeWordArray = tempThreeWordArray[:len(tempThreeWordArray)-1]
		}

		substring = strconv.Itoa(i * 11)
		for strings.Contains(newSentence, substring) {
			word := rand.Intn(len(tempTwoWordArray))
			newSentence = strings.Replace(newSentence, substring, tempTwoWordArray[word], -1)
			tempTwoWordArray[word] = tempTwoWordArray[len(tempTwoWordArray)-1]
			tempTwoWordArray = tempTwoWordArray[:len(tempTwoWordArray)-1]
		}
		substring = strconv.Itoa(i)
		for strings.Contains(newSentence, substring) {
			word := rand.Intn(len(tempOneWordArray))
			newSentence = strings.Replace(newSentence, substring, tempOneWordArray[word], -1)
			tempOneWordArray[word] = tempOneWordArray[len(tempOneWordArray)-1]
			tempOneWordArray = tempOneWordArray[:len(tempOneWordArray)-1]
		}

	}
	return newSentence
}

func show(sentence string) {
	fmt.Println(sentence)
}
