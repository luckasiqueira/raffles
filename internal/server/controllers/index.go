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
	}
	for i := 1; i <= total; i++ {
		for _, j := range taken {
			var draw database.DrawInfo
			if i != j {
				draw.Available = true
			}
			draw.Number = i
			drawinfo = append(drawinfo, draw)
			break
		}
	}
	return drawinfo
}
