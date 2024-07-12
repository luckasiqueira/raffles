package controllers

import (
	"github.com/kataras/iris/v12"
	"raffles/internal/frontend"
	"raffles/internal/server/database"
	"raffles/utils/info"
	"raffles/utils/logger"
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
		logger.NewLog(err.Error())
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
