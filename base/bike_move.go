package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

var road = [9][9]int{
	{0, 0, 0, 0, 0, 0, 0, 0},
	{1, 1, 1, 1, 0, 0, 0, 0},
	{1, 0, 0, 1, 0, 0, 0, 0},
	{1, 0, 0, 1, 1, 1, 1, 1},
	{1, 0, 0, 0, 0, 0, 0, 1},
	{1, 1, 1, 1, 1, 1, 0, 1},
	{0, 0, 0, 0, 0, 1, 0, 1},
	{0, 0, 0, 0, 0, 1, 1, 1},
	{0, 0, 0, 0, 0, 0, 0, 1},
}

type point struct {
	x, y int
}

type station struct {
	point
	bikeNum int
	name    string
}

type player struct {
	point             //出发站点
	destination point //目的地，目的地从可通行区域来
	status      int   //0:扫码，1：骑行，2：还车
	name        string
}

func (s *station) getStatus() string {
	return strings.Join([]string{s.name, strconv.Itoa(s.bikeNum)}, "车")
}

//货车
type van struct {
	point
	bikeNum int
	name    string
}

//A站
var aStation = &station{
	point:   point{1, 1},
	bikeNum: 30,
	name:    "A站",
}

//B站
var bStation = &station{
	point:   point{8, 8},
	bikeNum: 40,
	name:    "B站",
}

//C站
var cStation = &station{
	point:   point{4, 3},
	bikeNum: 30,
	name:    "C站",
}

var allStation = []*station{aStation, bStation, cStation}
var playerBike []*player

func getRoadBikeStatus() string {

	return ""
}

func start() {

}

func playerRide(time int) {
	if time == 0 {
		start()
		return
	}

	passenger := rand.Intn(len(allStation))
	station := allStation[passenger]
	station.bikeNum -= 1
	player := &player{
		point:  point{station.x, station.y},
		name:   "",
		status: 0,
	}

	playerBike = append(playerBike, player)
}

func getTime(i int) string {
	return strings.Join([]string{"第", strconv.Itoa(i), "秒"}, "")
}

func main() {
	time := 201

	for i := 0; i < time; i++ {
		playerRide(i)
		vanMove()
		roadBikeNum := getRoadBikeStatus()
		timeFmt := getTime(i)
		fmt.Printf("%s,%s,%s,%s,%s \n", timeFmt, allStation[0].getStatus(), allStation[1].getStatus(), allStation[2].getStatus(), roadBikeNum)
	}
}

func vanMove() {

}
