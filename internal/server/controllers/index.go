package controllers

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"raffles/internal/frontend"
	"raffles/internal/server/database"
	"raffles/utils/info"
	"strconv"
)

func Index(c iris.Context) {
	drawinfo := drawStatus()
	buyURL := info.Env["BUY_URL"]
	c.RenderComponent(frontend.Index(drawinfo, buyURL))
}

func drawStatus() []database.DrawInfo {
	taken := database.DrawTaken()
	var drawinfo []database.DrawInfo
	total, err := strconv.Atoi(info.Env["TOTAL_NUMS"])
	if err != nil {
		fmt.Println(err.Error())
		return drawinfo
	}
	numbers := make([]int, total)
	for i := range numbers {
		var draw database.DrawInfo
		draw.Number = i + 1
		draw.Available = !contains(taken, draw.Number)
		drawinfo = append(drawinfo, draw)
	}
	return drawinfo
}

func contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
