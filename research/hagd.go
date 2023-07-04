package main

import (
	"container/list"
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	PokerColorHearts   = "H"
	PokerColorSpades   = "S"
	PokerColorDiamonds = "D"
	PokerColorClubs    = "C"
)

const (
	PokerHandSingle        = "Single"
	PokerHandPair          = "Pair"
	PokerHandTrips         = "Trips"
	PokerHandTwoTrips      = "TwoTrips"
	PokerHandThreePair     = "TreePair"
	PokerHandThreeWithTwo  = "ThreeWithTwo"
	PokerHandStraight      = "Straight"
	PokerHandBoom          = "Boom"
	PokerHandJokerBoom     = "JokerBoom"
	PokerHandStraightFlush = "StraightFlush"
	PokerHandPASS          = "PASS"
	PokerHandTribute       = "tribute"
	PokerHandBack          = "back"
	PokerHandError         = "Error"
	PokerHandMan           = "Manual"
)

const (
	PokerViewNum2 = "2"
	PokerViewNum3 = "3"
	PokerViewNum4 = "4"
	PokerViewNum5 = "5"
	PokerViewNum6 = "6"
	PokerViewNum7 = "7"
	PokerViewNum8 = "8"
	PokerViewNum9 = "9"
	PokerViewNumT = "T"
	PokerViewNumJ = "J"
	PokerViewNumQ = "Q"
	PokerViewNumK = "K"
	PokerViewNumA = "A"
	PokerViewNumB = "B"
	PokerViewNumR = "R"
)

const (
	PokerLevel2 = iota + 0
	PokerLevel3
	PokerLevel4
	PokerLevel5
	PokerLevel6
	PokerLevel7
	PokerLevel8
	PokerLevel9
	PokerLevel10
	PokerLevelJ
	PokerLevelQ
	PokerLevelK
	PokerLevelA
	PokerLevelB
	PokerLevelR
)

var mapViewToLevel map[string]int
var mapLevelToView map[int]string
var pokerColors []string
var pokerNumbers []string
var pokerJokers []string

func init() {
	mapViewToLevel = make(map[string]int, 0)
	mapViewToLevel[PokerViewNum2] = PokerLevel2
	mapViewToLevel[PokerViewNum3] = PokerLevel3
	mapViewToLevel[PokerViewNum4] = PokerLevel4
	mapViewToLevel[PokerViewNum5] = PokerLevel5
	mapViewToLevel[PokerViewNum6] = PokerLevel6
	mapViewToLevel[PokerViewNum7] = PokerLevel7
	mapViewToLevel[PokerViewNum8] = PokerLevel8
	mapViewToLevel[PokerViewNum9] = PokerLevel9
	mapViewToLevel[PokerViewNumT] = PokerLevel10
	mapViewToLevel[PokerViewNumJ] = PokerLevelJ
	mapViewToLevel[PokerViewNumQ] = PokerLevelQ
	mapViewToLevel[PokerViewNumK] = PokerLevelK
	mapViewToLevel[PokerViewNumA] = PokerLevelA
	mapViewToLevel[PokerViewNumB] = PokerLevelB
	mapViewToLevel[PokerViewNumR] = PokerLevelR

	mapLevelToView = make(map[int]string, 0)
	mapLevelToView[PokerLevel2] = PokerViewNum2
	mapLevelToView[PokerLevel3] = PokerViewNum3
	mapLevelToView[PokerLevel4] = PokerViewNum4
	mapLevelToView[PokerLevel5] = PokerViewNum5
	mapLevelToView[PokerLevel6] = PokerViewNum6
	mapLevelToView[PokerLevel7] = PokerViewNum7
	mapLevelToView[PokerLevel8] = PokerViewNum8
	mapLevelToView[PokerLevel9] = PokerViewNum9
	mapLevelToView[PokerLevel10] = PokerViewNumT
	mapLevelToView[PokerLevelJ] = PokerViewNumJ
	mapLevelToView[PokerLevelQ] = PokerViewNumQ
	mapLevelToView[PokerLevelK] = PokerViewNumK
	mapLevelToView[PokerLevelA] = PokerViewNumA
	mapLevelToView[PokerLevelB] = PokerViewNumB
	mapLevelToView[PokerLevelR] = PokerViewNumR

	pokerColors = getPokerColors()
	pokerNumbers = getPokerNumbers()
	pokerJokers = getPokerJokers()
}

func getPokerColors() []string {
	return []string{PokerColorSpades, PokerColorHearts, PokerColorClubs, PokerColorDiamonds}
}

func getPokerNumbers() []string {
	return []string{
		PokerViewNum2, PokerViewNum3, PokerViewNum4, PokerViewNum5, PokerViewNum6, PokerViewNum7, PokerViewNum8, PokerViewNum9, PokerViewNumT, PokerViewNumJ, PokerViewNumQ, PokerViewNumK, PokerViewNumA,
	}

}

func getPokerJokers() []string {
	return []string{
		PokerViewNumB, PokerViewNumR,
	}
}

type Card struct {
	Color      string
	ViewNumber string
	Level      int
	Wild       int
	Id         int // debug
}

func (c Card) Compare(a *Card) int {
	return c.Level - a.Level
}

func (c Card) Equal(a *Card) bool {
	return c.Color == a.Color && c.ViewNumber == a.ViewNumber && c.Level == a.Level
}

func (c Card) Next() Card {
	// Do Not Care Color
	level := (c.Level + 1) % 13
	return newCard(PokerColorDiamonds, mapLevelToView[level])
}

func (c Card) Prev() Card {
	// Do Not Care Color
	level := (c.Level - 1 + 13) % 13
	return newCard(PokerColorDiamonds, mapLevelToView[level])
}

func (c Card) Display() string {
	//return c.Color + c.ViewNumber + ":" + strconv.Itoa(c.Id)
	return c.Color + c.ViewNumber
}

func newCard(c, n string) Card {
	return Card{Color: c, ViewNumber: n, Level: mapViewToLevel[n]}
}

func newCardWithId(c, n string, id int) Card {
	return Card{Color: c, ViewNumber: n, Level: mapViewToLevel[n], Id: id}
}

type PokerHandCards []Card

// 两幅桌牌的管理
type PackCards struct {
	// 108
	Cards PokerHandCards
}

func (pc PackCards) Display() {

	for i, c := range pc.Cards {
		fmt.Println(i, c.Display())
	}
}

func NewPackCards() PackCards {
	pc := PackCards{
		Cards: []Card{},
	}

	id := 0

	for i := 0; i < 2; i++ {
		for _, n := range pokerNumbers {
			for _, c := range pokerColors {
				pc.Cards = append(pc.Cards, newCardWithId(c, n, id))
				id += 1
			}
		}
		pc.Cards = append(pc.Cards, newCardWithId(PokerColorSpades, PokerViewNumB, id))
		id += 1
		pc.Cards = append(pc.Cards, newCardWithId(PokerColorHearts, PokerViewNumR, id))
		id += 1
	}

	return pc
}

func (pc *PackCards) Shuffle() {

	rand.Seed(time.Now().UnixNano())

	for n := len(pc.Cards); n > 0; n-- {
		rand.Shuffle(len(pc.Cards), func(i, j int) {
			pc.Cards[i].Id, pc.Cards[i].Level, pc.Cards[i].ViewNumber, pc.Cards[i].Color, pc.Cards[j].Id, pc.Cards[j].Level, pc.Cards[j].ViewNumber, pc.Cards[j].Color =
				pc.Cards[j].Id, pc.Cards[j].Level, pc.Cards[j].ViewNumber, pc.Cards[j].Color, pc.Cards[i].Id, pc.Cards[i].Level, pc.Cards[i].ViewNumber, pc.Cards[i].Color

		})
	}
}

// 给选手发牌
func (pc *PackCards) DealCards(p1, p2, p3, p4 *Player) {

	for i := 0; i < 108; i += 4 {
		p1.AddCard(pc.Cards[i])
		p2.AddCard(pc.Cards[i+1])
		p3.AddCard(pc.Cards[i+2])
		p4.AddCard(pc.Cards[i+3])
	}

}

func (phc PokerHandCards) Len() int {
	return len(phc)
}
func (phc PokerHandCards) Less(i, j int) bool {
	return phc[i].Level < phc[j].Level
}

func (phc PokerHandCards) Swap(i, j int) {
	// @TODO
	// 这个与结构体赋值的区别是什么
	phc[i], phc[j] = phc[j], phc[i]
}

func (phc PokerHandCards) Display() string {
	sb := strings.Builder{}
	sb.WriteString("cards :\n")
	for _, c := range phc {
		sb.WriteString(c.Display())
		sb.WriteString(",")
	}
	sb.WriteString("\n")

	return sb.String()

}

type PokerHand struct {
	Name  string // 调试标识用
	Type  string // 牌型,自动计算出来或者应用赋值
	Level string // 当前牌的级别
	Num   int

	Cards PokerHandCards // 当前手牌集合

	minCard Card // 排序后最小的牌,  不计算配子
	maxCard Card // 排序后最大的牌, 不计算配子

	Trip Card // 统计3点数用
	Pair Card //  统计对子点数用过

	wildCard      Card //  配子
	wildCardNum   int  // 0-2
	wildCardIndex []int

	parseDone bool

	colorParsed   map[string]int
	viewNumParsed map[string]int

	parseResults []PokerHandCards // 分析完成后结果保存在这里, 输出所有可能的牌型
}

// -1 ph < phc
// 0 ph == phc
// 1 phc > phc
func (ph *PokerHand) Compare(phc *PokerHand) int {

	if ph.Type != phc.Type {
		if phc.Type == PokerHandBoom || phc.Type == PokerHandStraightFlush {
			return -1
		}
		if phc.Type == PokerHandPASS {
			return 1
		}
	} else {
		// sort
		ph.sort()
		phc.sort()

		if ph.Type == PokerHandThreeWithTwo {
			return ph.threeWithTwoCompare(phc)
		} else if ph.Type == PokerHandJokerBoom {
			return 1
		} else {
			// pass
			return ph.minCompare(phc)
		}
	}

	return 0
}

func (ph *PokerHand) minCompare(phc *PokerHand) int {
	return ph.Cards[0].Level - phc.Cards[0].Level
}

func (ph *PokerHand) threeWithTwoCompare(phc *PokerHand) int {
	ph.setPairAndTrip()
	phc.setPairAndTrip()
	fmt.Println("trip:")
	fmt.Println(ph.Trip.Display())
	fmt.Println(phc.Trip.Display())
	fmt.Println("end")
	return ph.Trip.Level - phc.Trip.Level
}

func (ph *PokerHand) setPairAndTrip() {

	fmt.Println(ph.Cards)
	if ph.Cards[1].Level != ph.Cards[2].Level {
		ph.Pair.ViewNumber = ph.Cards[0].ViewNumber
		ph.Pair.Level = ph.Cards[0].Level
		ph.Pair.Color = ph.Cards[0].Color
		ph.Trip.Color = ph.Cards[2].Color
		ph.Trip.ViewNumber = ph.Cards[2].ViewNumber
		ph.Trip.Level = ph.Cards[2].Level
	} else {
		ph.Trip.ViewNumber = ph.Cards[0].ViewNumber
		ph.Trip.Level = ph.Cards[0].Level
		ph.Trip.Color = ph.Cards[0].Color
		ph.Pair.ViewNumber = ph.Cards[3].ViewNumber
		ph.Pair.Level = ph.Cards[3].Level
		ph.Pair.Color = ph.Cards[3].Color
	}
}

func (ph *PokerHand) sort() {
	sort.Sort(ph.Cards)
}

func (ph *PokerHand) Display() string {
	//ph.sort()
	sb := strings.Builder{}
	sb.WriteString("name: ")
	sb.WriteString(ph.Name)
	sb.WriteString("\n")
	sb.WriteString("type: ")
	sb.WriteString(ph.Type)
	sb.WriteString("\n")
	sb.WriteString("card list:\n")
	for _, c := range ph.Cards {
		sb.WriteString(c.Display())
		sb.WriteString(",")
	}
	sb.WriteString("\n")

	return sb.String()
}

func (ph *PokerHand) ParseResult() {

	fmt.Println(">>>>>>>> ==== <<<<<<<<<")
	for k, v := range ph.parseResults {
		fmt.Println(k, v.Display())
	}
	fmt.Println(">>>>>>>> ==== <<<<<<<<<")

}

func (ph *PokerHand) Parse() {
	ph.sort()
	ph.parse()
}

func (ph *PokerHand) parse() {

	if ph.parseDone {
		return
	}

	ph.Num = len(ph.Cards)
	ph.wildCard = newCard(PokerColorHearts, ph.Level)
	ph.minCard = Card{Level: PokerLevelR}
	ph.maxCard = Card{Level: PokerLevel2}
	ph.wildCardIndex = nil
	ph.wildCardNum = 0
	ph.wildCardIndex = make([]int, 2)
	ph.viewNumParsed = make(map[string]int, 10)
	ph.colorParsed = make(map[string]int, 10)

	for i := 0; i < ph.Num; i += 1 {

		// 统计通配牌的情况
		if ph.Cards[i].Equal(&ph.wildCard) {
			ph.wildCardNum += 1
			ph.wildCardIndex = append(ph.wildCardIndex, i)
		} else if ph.Cards[i].Compare(&ph.minCard) <= 0 {
			ph.minCard = ph.Cards[i]
		} else if ph.Cards[i].Compare(&ph.maxCard) >= 0 {
			ph.maxCard = ph.Cards[i]
		}

		// 此处不统计 wildCard
		if !ph.Cards[i].Equal(&ph.wildCard) {
			// 花色统计
			if _, ok := ph.colorParsed[ph.Cards[i].Color]; ok {
				ph.colorParsed[ph.Cards[i].Color] += 1
			} else {
				ph.colorParsed[ph.Cards[i].Color] = 1
			}

			// 统计数字情况
			if _, ok := ph.viewNumParsed[ph.Cards[i].ViewNumber]; ok {
				ph.viewNumParsed[ph.Cards[i].ViewNumber] += 1
			} else {
				ph.viewNumParsed[ph.Cards[i].ViewNumber] = 1
			}
		}

	}
	ph.parseDone = true
}

func (ph *PokerHand) needManual() bool {

	if ph.Num != 6 {
		return false
	}

	c := 0
	cards := make([]Card, 0)
	for k, v := range ph.viewNumParsed {
		fmt.Println(k, v)
		if v > c {
			c = v
		}
		cards = append(cards, newCard(PokerColorDiamonds, k))
	}

	diff := cards[0].Level - cards[1].Level
	if len(ph.viewNumParsed) == 2 && c == 2 && (diff == 1 || diff == -1) && ph.wildCardNum == 2 {
		return true
	}
	return false
}

func (ph *PokerHand) TypeAutoCheck() string {

	// prepare
	ph.sort()
	ph.parse()

	// exception checking
	if ph.needManual() {
		return PokerHandMan
	}

	// typeChecking
	if ph.Num == 0 {
		return PokerHandPASS
	} else if ph.Num == 1 {
		return PokerHandSingle
	} else if ph.Num == 2 && ph.isPair() {
		return PokerHandPair
	} else if ph.Num == 3 && ph.isTrip() {
		return PokerHandTrips
	} else if ph.Num == 4 && ph.isBomb() {
		return PokerHandBoom
	} else if ph.Num == 4 && ph.isJokerBomb() {
		return PokerHandJokerBoom
	} else if ph.Num == 5 && ph.isStraightFlush() {
		return PokerHandStraightFlush
	} else if ph.Num == 5 && ph.isStraight() {
		return PokerHandStraight
	} else if ph.Num == 5 && ph.isThreeTwo() {
		return PokerHandThreeWithTwo
	} else if ph.Num == 5 && ph.isBomb() {
		return PokerHandBoom
	} else if ph.Num == 6 && ph.isTwoTrip() {
		return PokerHandTwoTrips
	} else if ph.Num == 6 && ph.isThreePair() {
		return PokerHandThreePair
	} else if ph.Num == 6 && ph.isBomb() {
		return PokerHandBoom
	} else if ph.Num >= 7 && ph.Num <= 10 && ph.isBomb() {
		return PokerHandBoom
	}
	return PokerHandError
}

func (ph *PokerHand) isPair() bool {

	if ph.Num != 2 {
		return false
	}
	if len(ph.viewNumParsed) == 1 {
		return true
	}
	return false
}

func (ph *PokerHand) isThreePairContinuous() bool {
	ph.parseResults = make([]PokerHandCards, 0)
	tempCollect := make([]Card, 0)
	cards := make([]Card, 0)
	for k, v := range ph.viewNumParsed {
		fmt.Println("===", k, v)
		cards = append(cards, newCard(PokerColorDiamonds, k))
	}
	// temp function
	getValidWildNum := func() int {
		if ph.wildCardNum > 0 {
			return 1
		}
		return 0
	}
	findTargetCard := func(c *Card) bool {
		for i := 0; i < len(cards); i += 1 {
			if cards[i].Compare(c) == 0 {
				tempCollect = append(tempCollect, cards[i])
				return true
			}
		}
		return false
	}

	targetCount := 3
	currentCard := ph.minCard
	cardA := newCard(PokerColorDiamonds, PokerViewNumA)
	ca := ph.findCard(&cardA)
	// 翻转用
	if ca && ph.minCard.Prev().ViewNumber == PokerViewNumA {
		currentCard = currentCard.Prev()
	}

	isCont := false
	for i := getValidWildNum(); i >= 0; i -= 1 {
		count := 0
		wildCardNum := getValidWildNum()
		fmt.Println("CURRENT  == ", currentCard.Display())
		fmt.Println("cards ", ph.Cards.Display())
		n := currentCard
		tempCards := make([]Card, 0) // @TODO 打印所有可能的结果状态，后面需要展示在前端给用户,并要求确认
		for fc := true; fc || wildCardNum > 0; n = n.Next() {
			fmt.Println("next ", n.Display())
			fmt.Println("count ", count)
			// 345
			if count == targetCount {
				break
			}
			fc = findTargetCard(&n)
			if fc {
				count += 1
				tempCards = append(tempCards, n)
				continue
			}
			if !fc && wildCardNum > 0 {
				wildCardNum -= 1
				count += 1
				tempCards = append(tempCards, n)
				fc = true // replace
				continue
			}
		}
		// multicase
		if count == targetCount && len(tempCollect) == len(cards) {
			isCont = true
			ph.parseResults = append(ph.parseResults, tempCards)
		} else {
			fmt.Println("missing ", n.Display())
		}

		currentCard = currentCard.Prev()
	}
	return isCont
}
func (ph *PokerHand) isThreePair() bool {

	if len(ph.Cards) != 6 {
		return false
	}

	// 334455
	// 33445 *2
	// 3445 *2*2
	// 22344 *2
	// 3344 *2*2 =>? 223344 334455 333444
	// 2244 *2*2
	// 4444 *2*2 ==> 炸弹

	c := 0
	for k, v := range ph.viewNumParsed {
		fmt.Println("===", k, v)
		if v > c {
			c = v
		}
	}
	if (len(ph.viewNumParsed) == 3 && c == 2) || (len(ph.viewNumParsed) == 2 && c == 2 && ph.wildCardNum == 2) {
		if ph.isThreePairContinuous() {
			return true
		}
		return false
	}
	return false
}

func (ph *PokerHand) isTrip() bool {

	if len(ph.Cards) != 3 {
		return false
	}
	// 333
	// 33 *2
	// 3 *2*2
	if len(ph.viewNumParsed) == 1 {
		return true
	}
	return false
}

func (ph *PokerHand) isTwoTrip() bool {
	if ph.Num != 6 {
		return false
	}
	// 333444
	// 33344 *2
	// 3344 *2*2 => ? 333444 334455 223344
	// 3334 *2*2
	c := 0
	cards := make([]Card, 0)
	for k, v := range ph.viewNumParsed {
		fmt.Println(k, v)
		if v > c {
			c = v
		}
		cards = append(cards, newCard(PokerColorDiamonds, k))
	}

	diff := cards[0].Level - cards[1].Level
	if len(ph.viewNumParsed) == 2 && c <= 3 && (diff == 1 || diff == -1) {
		return true
	}

	return false
}

func (ph *PokerHand) isBomb() bool {

	if ph.Num < 4 {
		return false
	}
	// 33 *2*2
	// 333 *2
	// 333 *2*2
	// 3333
	// 33333
	// 33333 *2
	if len(ph.viewNumParsed) == 1 {
		return true
	}
	return false
}

func (ph *PokerHand) isJokerBomb() bool {
	if ph.Num != 4 {
		return false
	}

	hasB := false
	if v, ok := ph.viewNumParsed["B"]; ok {
		if v == 2 {
			hasB = true
		}
	}
	hasR := false
	if v, ok := ph.viewNumParsed["R"]; ok {
		if v == 2 {
			hasR = true
		}
	}
	if hasR && hasB {
		return true
	}
	return false
}

func (ph *PokerHand) findCard(c *Card) bool {

	for i := 0; i < ph.Num; i += 1 {
		// skip wildCard
		if !ph.Cards[i].Equal(&ph.wildCard) && ph.Cards[i].Compare(c) == 0 {
			return true
		}
	}
	return false
}
func (ph *PokerHand) isStraightContinuous() bool {

	// @TODO
	// ACE 的翻转的问题

	ph.parseResults = make([]PokerHandCards, 0) // 初始化结果数组
	isCont := false
	// 345 *2*3
	// 34567 *2
	currentCard := ph.minCard
	fmt.Println(ph.Display())
	for i := ph.wildCardNum; i >= 0; i -= 1 {
		count := 0
		wildCardNum := ph.wildCardNum
		fmt.Println("+++++++++++++ ", currentCard.Display())
		fmt.Println("========", ph.Cards.Display())
		n := currentCard
		cards := make([]Card, 0) // @TODO 打印所有可能的结果状态，后面需要展示在前端给用户,并要求确认
		for fc := true; fc || wildCardNum > 0; n = n.Next() {
			fmt.Println("next ", n.Display())
			fmt.Println("count ", count)
			fmt.Println("wildCardNum", wildCardNum)
			fc = ph.findCard(&n)
			if fc {
				count += 1
				cards = append(cards, n)
				continue
			}
			if !fc && wildCardNum > 0 {
				wildCardNum -= 1
				count += 1
				cards = append(cards, n)
				fc = true // replace
				continue
			}

		}
		// multicase
		if count == 5 {
			isCont = true
			// save cases
			ph.parseResults = append(ph.parseResults, cards)
		} else {
			fmt.Println("missing ", n.Display())
		}

		fmt.Println("BEGIN-RC:///")
		for _, c := range cards {
			fmt.Println("RC >>> ", c.Display())
		}
		fmt.Println("END-RC:///")

		currentCard = currentCard.Prev()
	}

	return isCont
}

func (ph *PokerHand) isStraight() bool {
	if ph.Num != 5 {
		return false
	}

	// Y 34567
	// Y 4567 *2

	// Y 345 *2*2
	// F 234 *2*2

	if (len(ph.viewNumParsed) == 5 || (len(ph.viewNumParsed) == 4 && ph.wildCardNum == 1) || (len(ph.viewNumParsed) == 3 && ph.wildCardNum == 2)) && ph.isStraightContinuous() {
		return true
	}
	return false
}

func (ph *PokerHand) isStraightFlush() bool {
	if ph.Num != 5 {
		return false
	}
	if ph.isStraight() && len(ph.colorParsed) == 1 {
		return true
	}
	return false
}

func (ph *PokerHand) isThreeTwo() bool {

	if ph.Num != 5 {
		return false
	}
	// 33344
	// 3344 *2
	// 3334 *2
	// 332 *2*2
	// 333 *2*2 ==> 炸弹
	// 222 *2*2  ==> 炸弹
	c := 0
	key := ""
	for k, v := range ph.viewNumParsed {
		fmt.Println(k, v)
		if v > c {
			c = v
			key = k
		}
	}
	if (len(ph.viewNumParsed) == 2 && c <= 3) || (len(ph.viewNumParsed) == 1 && c == 3 && key != ph.Level) {
		return true
	}
	return false
}

type Player struct {
	Location int
	Name     string
	Role     int   // 0 自己局 ， 1 对方局
	Score    []int // 0 1
	Level    string
	Cards    PokerHandCards

	// 选牌组合
	dfsCards      PokerHandCards
	dfsMasks      []int
	dfsResults    []PokerHandCards
	dfsTargetNum  int
	dfsCardDomain int
	dfsTargetCard PokerHandCards
	//

	pokerHand PokerHand
}

func (p *Player) AddCard(c Card) {
	p.Cards = append(p.Cards, c)
}

func (p *Player) Sort() {
	sort.Sort(p.Cards)

}

func (p *Player) Display() string {
	sb := strings.Builder{}

	sb.WriteString("location: ")
	sb.WriteString(strconv.Itoa(p.Location))
	sb.WriteString("\n")
	sb.WriteString("name: ")
	sb.WriteString(p.Name)
	sb.WriteString("\ncards:\n")
	p.Sort()
	for _, c := range p.Cards {
		sb.WriteString(c.Display())
		sb.WriteString(",")
	}
	sb.WriteString("\n")
	return sb.String()
}

func (p *Player) dfs(u, m int) {

	if u == p.dfsTargetNum {
		cards := make([]Card, p.dfsTargetNum)
		for i := 0; i < p.dfsTargetNum; i++ {
			cards[i] = p.dfsTargetCard[i]
		}
		p.dfsResults = append(p.dfsResults, cards)
		return
	}
	for i := u; i < p.dfsCardDomain; i++ {

		if p.dfsMasks[i] == 0 && i >= m {
			p.dfsTargetCard[u] = p.dfsCards[i]
			p.dfsMasks[i] = 1
			p.dfs(u+1, i)
			p.dfsMasks[i] = 0
		}
	}
}

func (p *Player) Parse() {
	p.pokerHand = PokerHand{
		Level: p.Level,
		Cards: p.Cards,
		Name:  "player",
	}
	p.pokerHand.Parse()
}

func (p *Player) DFS(cards []Card, num int) []PokerHandCards {

	p.dfsCards = cards

	p.dfsCardDomain = len(p.dfsCards)
	p.dfsTargetNum = num
	p.dfsMasks = make([]int, p.dfsCardDomain)
	p.dfsTargetCard = make([]Card, p.dfsTargetNum)
	p.dfsResults = make([]PokerHandCards, 0)
	p.dfs(0, 0)
	return p.dfsResults
}

func (p *Player) querySingle() []PokerHandCards {
	cards := make([]PokerHandCards, 0)
	for _, c := range p.Cards {
		cds := make([]Card, 0)
		cds = append(cds, c)
		cards = append(cards, cds)
	}
	return cards
}

func (p *Player) getCards(viewNum string) PokerHandCards {

	fmt.Println(p.pokerHand.wildCard.Display())
	cards := make([]Card, 0)
	targetCard := newCard(PokerColorDiamonds, viewNum)
	for _, c := range p.Cards {
		if c.Compare(&targetCard) == 0 || c.Equal(&p.pokerHand.wildCard) {
			cards = append(cards, c)
		}
	}
	return cards
}

func (p *Player) querySamePokerHandCards(num int) []PokerHandCards {
	p.Parse()
	pairHands := make([]PokerHandCards, 0)
	numbers := getPokerNumbers()
	for _, viewNum := range numbers {

		card := p.getCards(viewNum)
		fmt.Println("card ", card.Display())
		if len(card) < num {
			continue
		}
		phc := p.DFS(card, num)
		for _, pc := range phc {
			fmt.Println("**** ", pc.Display())
			pairHands = append(pairHands, pc)
		}
	}
	return pairHands
}
func (p *Player) queryPair() []PokerHandCards {
	return p.querySamePokerHandCards(2)
}

func (p *Player) queryTrip() []PokerHandCards {
	return p.querySamePokerHandCards(3)
}

func (p *Player) queryStraight() []PokerHandCards {
	cards := []PokerHandCards{}
	return cards

}

func testList() {
	l := list.New()
	l.PushBack(newCard(PokerColorDiamonds, PokerViewNum2))
	l.PushBack(newCard(PokerColorDiamonds, PokerViewNum4))
	l.PushBack(newCard(PokerColorDiamonds, PokerViewNum3))

	for v := l.Front(); v != nil; v = v.Next() {
		fmt.Printf("color %s value %d dispaly %s \n", v.Value.(Card).Color, v.Value.(Card).Level, v.Value.(*Card).Display())
	}
}

func testSort() {

	cards := []Card{
		newCard(PokerColorDiamonds, PokerViewNum2),
		newCard(PokerColorDiamonds, PokerViewNum4),
		newCard(PokerColorDiamonds, PokerViewNum3),
	}
	sort.Sort(PokerHandCards(cards))
	fmt.Println(cards)
}

func testPokerHandPair() {

	// 88
	cardA := []Card{
		newCard(PokerColorHearts, PokerViewNum8),
		newCard(PokerColorDiamonds, PokerViewNum8),
	}
	ph := PokerHand{
		Type:  PokerHandPair,
		Cards: cardA,
	}
	fmt.Println(ph.Display())

}

func testPokerHand32() {
	// 33322
	cardA := []Card{
		newCard(PokerColorDiamonds, PokerViewNum2),
		newCard(PokerColorHearts, PokerViewNum3),
		newCard(PokerColorDiamonds, PokerViewNum3),
		newCard(PokerColorDiamonds, PokerViewNum3),
		newCard(PokerColorDiamonds, PokerViewNum2),
	}
	ph := PokerHand{
		Name:  "cardA",
		Type:  PokerHandThreeWithTwo,
		Cards: cardA,
	}
	fmt.Println(ph.Display())

}

func testCompareHand32() {
	// 33322
	cardA := []Card{
		newCard(PokerColorDiamonds, PokerViewNum2),
		newCard(PokerColorHearts, PokerViewNum3),
		newCard(PokerColorDiamonds, PokerViewNum3),
		newCard(PokerColorDiamonds, PokerViewNum3),
		newCard(PokerColorDiamonds, PokerViewNum2),
	}

	// 22233
	cardB := []Card{
		newCard(PokerColorHearts, PokerViewNum2),
		newCard(PokerColorHearts, PokerViewNum2),
		newCard(PokerColorHearts, PokerViewNum3),
		newCard(PokerColorClubs, PokerViewNum3),
		newCard(PokerColorDiamonds, PokerViewNum2),
	}
	ph := PokerHand{
		Name:  "cardA",
		Type:  PokerHandThreeWithTwo,
		Cards: cardA,
	}

	phc := PokerHand{
		Type:  PokerHandThreeWithTwo,
		Name:  "cardB",
		Cards: cardB,
	}

	fmt.Println(ph.Display())
	fmt.Println(phc.Display())

	c := ph.Compare(&phc)

	fmt.Println(ph.Display())
	fmt.Println(phc.Display())
	fmt.Println("code ", c)

	if c < 0 {
		fmt.Println("card  A  less than card b")
	} else if c > 0 {
		fmt.Println("card  A  greater than card b")
	} else {
		fmt.Println("card  A  equal card b")
	}

}

func testPokerHandStraight() {
	// 10,JQKA
	cardA := []Card{
		newCard(PokerColorDiamonds, PokerViewNumT),
		newCard(PokerColorHearts, PokerViewNumQ),
		newCard(PokerColorDiamonds, PokerViewNumK),
		newCard(PokerColorDiamonds, PokerViewNumJ),
		newCard(PokerColorDiamonds, PokerViewNumA),
	}

	// 34567
	cardB := []Card{
		newCard(PokerColorHearts, PokerViewNum7),
		newCard(PokerColorHearts, PokerViewNum6),
		newCard(PokerColorHearts, PokerViewNum5),
		newCard(PokerColorClubs, PokerViewNum3),
		newCard(PokerColorDiamonds, PokerViewNum4),
	}
	ph := PokerHand{
		Name:  "cardA",
		Type:  PokerHandStraight,
		Cards: cardA,
	}

	phc := PokerHand{
		Type:  PokerHandStraight,
		Name:  "cardB",
		Cards: cardB,
	}

	fmt.Println(ph.Display())
	fmt.Println(phc.Display())

	c := ph.Compare(&phc)

	fmt.Println(ph.Display())
	fmt.Println(phc.Display())
	fmt.Println("code ", c)

	if c < 0 {
		fmt.Println("card  A  less than card b")
	} else if c > 0 {
		fmt.Println("card  A  greater than card b")
	} else {
		fmt.Println("card  A  equal card b")
	}

}

func testPackCards() {

	fmt.Println("===== ")
	pc := NewPackCards()
	pc.Display()
	pc.Shuffle()
	pc.Display()
}

func testDealCard() {
	fmt.Println("+++++")
	pc := NewPackCards()
	pc.Shuffle()
	pc.Display()

	p1 := &Player{Name: "player1", Location: 1, Cards: []Card{}}
	p2 := &Player{Name: "player2", Location: 2, Cards: []Card{}}
	p3 := &Player{Name: "player3", Location: 3, Cards: []Card{}}
	p4 := &Player{Name: "player4", Location: 4, Cards: []Card{}}
	pc.DealCards(p1, p2, p3, p4)
	fmt.Println(p1.Display())
	fmt.Println(p2.Display())
	fmt.Println(p3.Display())
	fmt.Println(p4.Display())
}

// generate Cases

func testCardsChoice() {

	cards := []Card{
		newCardWithId(PokerColorDiamonds, PokerViewNum3, 0),
		newCardWithId(PokerColorSpades, PokerViewNum4, 1),
		newCardWithId(PokerColorHearts, PokerViewNum4, 2),
		newCardWithId(PokerColorDiamonds, PokerViewNum4, 3),
		newCardWithId(PokerColorClubs, PokerViewNum4, 4),
		newCardWithId(PokerColorClubs, PokerViewNum5, 5),
		newCardWithId(PokerColorDiamonds, PokerViewNum5, 6),
		newCardWithId(PokerColorHearts, PokerViewNum6, 7),
		newCardWithId(PokerColorClubs, PokerViewNum6, 8),
		newCardWithId(PokerColorSpades, PokerViewNum7, 9),
		newCardWithId(PokerColorDiamonds, PokerViewNum7, 10),
		newCardWithId(PokerColorHearts, PokerViewNum8, 11),
		newCardWithId(PokerColorClubs, PokerViewNum8, 12),
		newCardWithId(PokerColorDiamonds, PokerViewNum8, 13),
		newCardWithId(PokerColorHearts, PokerViewNum9, 14),
		newCardWithId(PokerColorSpades, PokerViewNum9, 15),
		newCardWithId(PokerColorClubs, PokerViewNumT, 16),
		newCardWithId(PokerColorHearts, PokerViewNumT, 17),
		newCardWithId(PokerColorClubs, PokerViewNumJ, 18),
		newCardWithId(PokerColorHearts, PokerViewNumQ, 19),
		newCardWithId(PokerColorSpades, PokerViewNumK, 20),
		newCardWithId(PokerColorClubs, PokerViewNumK, 21),
		newCardWithId(PokerColorHearts, PokerViewNumA, 22),
		newCardWithId(PokerColorDiamonds, PokerViewNumA, 23),
		newCardWithId(PokerColorHearts, PokerViewNum2, 24),
		newCardWithId(PokerColorDiamonds, PokerViewNum2, 25),
		newCardWithId(PokerColorHearts, PokerViewNumR, 26),
	}
	p1 := &Player{Name: "player1", Location: 1, Cards: cards, Level: PokerViewNum2}
	//s := p1.querySingle()
	s := p1.queryTrip()

	for i, ss := range s {
		fmt.Println("index ", i, ss.Display())
	}

	//p1.Display()

	fmt.Println("+++++++++++++")

}

func testThreeWithTwo() {
	// 33344
	// Level 2
	t1 := []Card{
		newCardWithId(PokerColorDiamonds, PokerViewNum3, 0),


		newCardWithId(PokerColorDiamonds, PokerViewNum3, 1),
		newCardWithId(PokerColorSpades, PokerViewNum3, 2),
		newCardWithId(PokerColorSpades, PokerViewNum4, 3),
		newCardWithId(PokerColorSpades, PokerViewNum4, 4),
	}
	ph1 := PokerHand{Name: "t1:33344", Level: PokerViewNum2, Cards: t1}
	fmt.Println(ph1.TypeAutoCheck())
	fmt.Println(ph1.Display())

	// 3344 *2
	// Level 2
	t2 := []Card{
		newCardWithId(PokerColorDiamonds, PokerViewNum3, 0),
		newCardWithId(PokerColorDiamonds, PokerViewNum3, 1),
		newCardWithId(PokerColorSpades, PokerViewNum4, 3),
		newCardWithId(PokerColorSpades, PokerViewNum4, 4),
		newCardWithId(PokerColorHearts, PokerViewNum2, 2),
	}
	ph2 := PokerHand{Name: "t2:3344*2", Level: PokerViewNum2, Cards: t2}
	fmt.Println(ph2.TypeAutoCheck())
	fmt.Println(ph2.Display())

	// Level 2
	t3 := []Card{
		newCardWithId(PokerColorDiamonds, PokerViewNum2, 1),
		newCardWithId(PokerColorSpades, PokerViewNum2, 3),
		newCardWithId(PokerColorSpades, PokerViewNum2, 4),
		newCardWithId(PokerColorHearts, PokerViewNum2, 2),
		newCardWithId(PokerColorHearts, PokerViewNum2, 4),
	}
	ph3 := PokerHand{Name: "t3:222 *2*2", Level: PokerViewNum2, Cards: t3}
	fmt.Println(ph3.TypeAutoCheck())
	fmt.Println(ph3.Display())

	// 3334 *2
	// Level 2
	t4 := []Card{
		newCardWithId(PokerColorDiamonds, PokerViewNum3, 0),
		newCardWithId(PokerColorDiamonds, PokerViewNum3, 1),
		newCardWithId(PokerColorSpades, PokerViewNum3, 3),
		newCardWithId(PokerColorSpades, PokerViewNum4, 4),
		newCardWithId(PokerColorHearts, PokerViewNum2, 2),
	}
	ph4 := PokerHand{Name: "t4:3334*2", Level: PokerViewNum2, Cards: t4}
	fmt.Println(ph4.TypeAutoCheck())
	fmt.Println(ph4.Display())

}
func testTypeChecking() {

	// 34567
	// Level 2
	t3 := []Card{
		newCardWithId(PokerColorDiamonds, PokerViewNum3, 0),
		newCardWithId(PokerColorDiamonds, PokerViewNum4, 1),
		newCardWithId(PokerColorSpades, PokerViewNum6, 3),
		newCardWithId(PokerColorSpades, PokerViewNum7, 4),
		newCardWithId(PokerColorHearts, PokerViewNum5, 2),
	}
	ph3 := PokerHand{Name: "t3:34567", Level: PokerViewNum2, Cards: t3}
	fmt.Println(ph3.TypeAutoCheck())
	fmt.Println(ph3.Display())

	// 同花
	// 34567
	// Level 2
	t4 := []Card{
		newCardWithId(PokerColorDiamonds, PokerViewNum3, 0),
		newCardWithId(PokerColorDiamonds, PokerViewNum4, 1),
		newCardWithId(PokerColorDiamonds, PokerViewNum6, 3),
		newCardWithId(PokerColorDiamonds, PokerViewNum7, 4),
		newCardWithId(PokerColorDiamonds, PokerViewNum5, 2),
	}
	ph4 := PokerHand{Name: "t4:34567", Level: PokerViewNum2, Cards: t4}
	fmt.Println(ph4.TypeAutoCheck())
	fmt.Println(ph4.Display())

	// 同花
	// 34567
	// Level 2
	t5 := []Card{
		newCardWithId(PokerColorDiamonds, PokerViewNum3, 0),
		newCardWithId(PokerColorDiamonds, PokerViewNum4, 1),
		newCardWithId(PokerColorHearts, PokerViewNum2, 3),
		newCardWithId(PokerColorDiamonds, PokerViewNum7, 4),
		newCardWithId(PokerColorDiamonds, PokerViewNum5, 2),
	}
	ph5 := PokerHand{Name: "t5:34567", Level: PokerViewNum2, Cards: t5}
	fmt.Println(ph5.TypeAutoCheck())
	fmt.Println(ph5.Display())

	// Error
	// 4567
	// Level 2
	t6 := []Card{
		newCardWithId(PokerColorDiamonds, PokerViewNum3, 0),
		newCardWithId(PokerColorDiamonds, PokerViewNum4, 1),
		newCardWithId(PokerColorHearts, PokerViewNum2, 3),
		newCardWithId(PokerColorDiamonds, PokerViewNum7, 4),
	}
	ph6 := PokerHand{Name: "t6:34567", Level: PokerViewNum2, Cards: t6}
	fmt.Println(ph6.TypeAutoCheck())
	fmt.Println(ph6.Display())

	// Bomb
	// 555 *2*2
	// Level 2
	t7 := []Card{
		newCardWithId(PokerColorDiamonds, PokerViewNum5, 0),
		newCardWithId(PokerColorDiamonds, PokerViewNum5, 1),
		newCardWithId(PokerColorHearts, PokerViewNum2, 3),
		newCardWithId(PokerColorHearts, PokerViewNum2, 2),
		newCardWithId(PokerColorSpades, PokerViewNum5, 4),
	}
	ph7 := PokerHand{Name: "t7:555 *2*2", Level: PokerViewNum2, Cards: t7}
	fmt.Println(ph7.TypeAutoCheck())
	fmt.Println(ph7.Display())

}

func testThreePair() {
	// ThreePair
	// 334455
	// Level 2
	t8 := []Card{
		newCardWithId(PokerColorDiamonds, PokerViewNum3, 0),
		newCardWithId(PokerColorDiamonds, PokerViewNum3, 1),
		newCardWithId(PokerColorHearts, PokerViewNum5, 3),
		newCardWithId(PokerColorHearts, PokerViewNum5, 2),
		newCardWithId(PokerColorSpades, PokerViewNum4, 4),
		newCardWithId(PokerColorSpades, PokerViewNum4, 5),
	}
	ph8 := PokerHand{Name: "t8:334455 *2*2", Level: PokerViewNum2, Cards: t8}
	fmt.Println(ph8.TypeAutoCheck())
	fmt.Println(ph8.Display())
	ph8.ParseResult()

	// ThreePair
	// 33455 *2
	// Level 2
	fmt.Println("HHHHHHHHHHHHHHHHHHHHHHHH")
	t9 := []Card{
		newCardWithId(PokerColorDiamonds, PokerViewNum3, 0),
		newCardWithId(PokerColorDiamonds, PokerViewNum3, 1),
		newCardWithId(PokerColorHearts, PokerViewNum5, 3),
		newCardWithId(PokerColorHearts, PokerViewNum5, 2),
		newCardWithId(PokerColorSpades, PokerViewNum4, 4),
		newCardWithId(PokerColorHearts, PokerViewNum2, 5),
	}
	ph9 := PokerHand{Name: "t9:33455 *2", Level: PokerViewNum2, Cards: t9}
	fmt.Println(ph9.TypeAutoCheck())
	fmt.Println(ph9.Display())
	fmt.Println("%%%%%%%%%%%%%%%%%%%%%")
	ph9.ParseResult()

	// ThreePair
	// 4455 *2*2
	// Level 2
	t1 := []Card{
		newCardWithId(PokerColorDiamonds, PokerViewNum4, 0),
		newCardWithId(PokerColorDiamonds, PokerViewNum4, 1),
		newCardWithId(PokerColorHearts, PokerViewNum5, 3),
		newCardWithId(PokerColorHearts, PokerViewNum5, 2),
		newCardWithId(PokerColorHearts, PokerViewNum2, 4),
		newCardWithId(PokerColorHearts, PokerViewNum2, 5),
	}
	ph1 := PokerHand{Name: "t1:4455 *2*2", Level: PokerViewNum2, Cards: t1}
	fmt.Println(ph1.TypeAutoCheck())
	fmt.Println(ph1.Display())
	ph1.ParseResult()

	// ThreePair
	// 3355 *2*2
	// Level 2
	t2 := []Card{
		newCardWithId(PokerColorDiamonds, PokerViewNum3, 0),
		newCardWithId(PokerColorDiamonds, PokerViewNum3, 1),
		newCardWithId(PokerColorHearts, PokerViewNum5, 3),
		newCardWithId(PokerColorHearts, PokerViewNum5, 2),
		newCardWithId(PokerColorHearts, PokerViewNum2, 4),
		newCardWithId(PokerColorHearts, PokerViewNum2, 5),
	}
	ph2 := PokerHand{Name: "t2:3355 *2*2", Level: PokerViewNum2, Cards: t2}
	fmt.Println(ph2.TypeAutoCheck())
	fmt.Println(ph2.Display())
	ph2.ParseResult()
	// ThreePair
	// 3335 *2*2
	// Level 2
	t3 := []Card{
		newCardWithId(PokerColorDiamonds, PokerViewNum3, 0),
		newCardWithId(PokerColorDiamonds, PokerViewNum3, 1),
		newCardWithId(PokerColorHearts, PokerViewNum3, 3),
		newCardWithId(PokerColorHearts, PokerViewNum5, 2),
		newCardWithId(PokerColorHearts, PokerViewNum2, 4),
		newCardWithId(PokerColorHearts, PokerViewNum2, 5),
	}
	ph3 := PokerHand{Name: "t3:3335 *2*2", Level: PokerViewNum2, Cards: t3}
	fmt.Println(ph3.TypeAutoCheck())
	fmt.Println(ph3.Display())
	ph3.ParseResult()

	// @TODO
	// ThreePair
	// AA2233
	// Level 2
	t4 := []Card{
		newCardWithId(PokerColorDiamonds, PokerViewNum3, 0),
		newCardWithId(PokerColorDiamonds, PokerViewNum3, 1),
		newCardWithId(PokerColorSpades, PokerViewNum2, 3),
		newCardWithId(PokerColorSpades, PokerViewNum2, 2),
		newCardWithId(PokerColorHearts, PokerViewNumA, 4),
		newCardWithId(PokerColorHearts, PokerViewNumA, 5),
	}
	ph4 := PokerHand{Name: "t4:AA2233", Level: PokerViewNum2, Cards: t4}
	fmt.Println(ph4.TypeAutoCheck())
	fmt.Println(ph4.Display())
	ph4.ParseResult()
	// QQKKAA
	// ThreePair
	// QQKKAA
	// Level 2
	t5 := []Card{
		newCardWithId(PokerColorDiamonds, PokerViewNumK, 0),
		newCardWithId(PokerColorDiamonds, PokerViewNumK, 1),
		newCardWithId(PokerColorSpades, PokerViewNumQ, 3),
		newCardWithId(PokerColorSpades, PokerViewNumQ, 2),
		newCardWithId(PokerColorHearts, PokerViewNumA, 4),
		newCardWithId(PokerColorHearts, PokerViewNumA, 5),
	}
	ph5 := PokerHand{Name: "t5:QQKKAA", Level: PokerViewNum2, Cards: t5}
	fmt.Println(ph5.TypeAutoCheck())
	fmt.Println(ph5.Display())
	ph5.ParseResult()
	// QQKKAA
	// ThreePair
	// 223344
	// Level 2
	t6 := []Card{
		newCardWithId(PokerColorDiamonds, PokerViewNum3, 0),
		newCardWithId(PokerColorDiamonds, PokerViewNum3, 1),
		newCardWithId(PokerColorSpades, PokerViewNum2, 3),
		newCardWithId(PokerColorSpades, PokerViewNum2, 2),
		newCardWithId(PokerColorHearts, PokerViewNum4, 4),
		newCardWithId(PokerColorHearts, PokerViewNum4, 5),
	}
	ph6 := PokerHand{Name: "t6:223344 ", Level: PokerViewNum2, Cards: t6}
	fmt.Println(ph6.TypeAutoCheck())
	fmt.Println(ph6.Display())
	ph6.ParseResult()
}

func testContinue() {

	// 同花
	// 34567
	// Level 2
	t5 := []Card{
		newCardWithId(PokerColorDiamonds, PokerViewNum3, 0),
		newCardWithId(PokerColorDiamonds, PokerViewNum4, 1),
		newCardWithId(PokerColorHearts, PokerViewNum2, 3),
		newCardWithId(PokerColorDiamonds, PokerViewNum7, 4),
		newCardWithId(PokerColorDiamonds, PokerViewNum5, 2),
	}
	ph5 := PokerHand{Name: "t5:34267", Level: PokerViewNum2, Cards: t5}
	fmt.Println(ph5.TypeAutoCheck())
	fmt.Println(ph5.Display())
	ph5.ParseResult()

	// 不是顺子
	// 34568
	// Level 2
	t6 := []Card{
		newCardWithId(PokerColorDiamonds, PokerViewNum3, 0),
		newCardWithId(PokerColorDiamonds, PokerViewNum4, 1),
		newCardWithId(PokerColorHearts, PokerViewNum2, 3),
		newCardWithId(PokerColorDiamonds, PokerViewNum8, 4),
		newCardWithId(PokerColorDiamonds, PokerViewNum5, 2),
	}
	ph6 := PokerHand{Name: "t6:34268", Level: PokerViewNum2, Cards: t6}
	fmt.Println(ph6.TypeAutoCheck())
	fmt.Println(ph6.Display())
	ph6.ParseResult()

	// 顺子不是同花
	// 458
	// Level 2
	t7 := []Card{
		newCardWithId(PokerColorHearts, PokerViewNum2, 0),
		newCardWithId(PokerColorDiamonds, PokerViewNum4, 1),
		newCardWithId(PokerColorHearts, PokerViewNum2, 3),
		newCardWithId(PokerColorHearts, PokerViewNum8, 4),
		newCardWithId(PokerColorDiamonds, PokerViewNum5, 2),
	}
	ph7 := PokerHand{Name: "t7:458 *2*2", Level: PokerViewNum2, Cards: t7}
	fmt.Println(ph7.TypeAutoCheck())
	fmt.Println(ph7.Display())
	ph7.ParseResult()

	// 顺子不是同花
	// 456 *2*2
	// Level 2

	fmt.Println("==========================")
	t8 := []Card{
		newCardWithId(PokerColorHearts, PokerViewNum2, 0),
		newCardWithId(PokerColorDiamonds, PokerViewNum4, 1),
		newCardWithId(PokerColorHearts, PokerViewNum2, 3),
		newCardWithId(PokerColorHearts, PokerViewNum6, 4),
		newCardWithId(PokerColorDiamonds, PokerViewNum5, 2),
	}
	ph8 := PokerHand{Name: "t8:456 *2*2", Level: PokerViewNum2, Cards: t8}
	fmt.Println(ph8.TypeAutoCheck())
	fmt.Println(ph8.Display())
	ph8.ParseResult()

	fmt.Println("--------------")
	t9 := []Card{
		newCardWithId(PokerColorHearts, PokerViewNum2, 0),
		newCardWithId(PokerColorHearts, PokerViewNum2, 3),
		newCardWithId(PokerColorDiamonds, PokerViewNum4, 1),
		newCardWithId(PokerColorHearts, PokerViewNum8, 4),
		newCardWithId(PokerColorDiamonds, PokerViewNum5, 2),
	}
	ph9 := PokerHand{Name: "t9:458 *2*2", Level: PokerViewNum2, Cards: t9}
	fmt.Println(ph9.TypeAutoCheck())
	fmt.Println(ph9.Display())
	ph9.ParseResult()
}

func main() {
	//testList()
	//testSort()
	//testPokerHandPair()
	//testPokerHand32()
	//testCompareHand32()
	//testPokerHandStraight()
	//testPackCards()
	//testDealCard()

	//testTypeChecking()

	// @TODO
	// 1 连续性检测  done
	// 2 生成遍历规则
	// 3 翻转
	// 4 保存所有可能的牌型 done=

	//testContinue()
	//testThreeWithTwo()

	//testThreePair()
	testCardsChoice()
}
