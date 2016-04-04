package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var sentenceModel = "xx，xx，xx了xx。\nxxxx，xxxx，不过是一场xxxx。\n你说xxxx，我说xxxx，最后不过xxxx。\nxx，xx，许我一场xxxx。\n一x一x一xx，半x半x半xx。\n你说xxxxxxxx，后来xxxxxxxx。\nxxxx，xxxx，终不敌xxxx。\n"
var twoWord = []string{"朱砂", "天下", "杀伐", "人家", "韶华", "风华", "繁华", "血染", "墨染", "白衣",
	"素衣", "嫁衣", "倾城", "孤城", "空城", "旧城", "旧人", "伊人", "心疼", "春风", "古琴", "无情", "迷离", "奈何", "断弦", "焚尽", "散乱", "陌路", "乱世", "笑靥",
	"浅笑", "明眸", "轻叹", "烟火", "一生", "三生", "浮生", "桃花", "梨花", "落花", "烟花", "离殇", "情殇", "爱殇", "剑殇", "灼伤", "仓皇", "匆忙", "陌上", "清商", "焚香", "墨香", "微凉",
	"断肠", "痴狂", "凄凉", "黄梁", "未央", "成双", "无恙", "虚妄", "凝霜", "洛阳", "长安", "江南", "忘川", "千年", "纸伞", "烟雨", "回眸", "公子", "红尘", "红颜", "红衣", "红豆", "红线", "青丝", "青史",
	"青冢", "白发", "白首", "白骨", "黄土", "黄泉", "碧落", "紫陌"}

var fourWord = []string{"情深缘浅", "情深不寿", "莫失莫忘", "阴阳相隔", "如花美眷", "似水流年", "眉目如画", "曲终人散", "繁华落尽", "不诉离殇", "一世长安"}
var oneWord = []string{"殇", "红", "离", "梦"}

func main() {
	newSentence := sentenceModel

	for strings.Contains(newSentence, "xxxx") {
		newSentence = strings.Replace(newSentence, "xxxx", fourWord[rand.Intn(len(fourWord))], 1)
	}
	for strings.Contains(newSentence, "xx") {
		newSentence = strings.Replace(newSentence, "xx", twoWord[rand.Intn(len(twoWord))], 1)
	}
	for strings.Contains(newSentence, "x") {
		newSentence = strings.Replace(newSentence, "x", oneWord[rand.Intn(len(oneWord))], 1)
	}

	show(newSentence)
}

func show(sentence string) {
	fmt.Println(sentence)
}
