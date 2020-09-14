package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)


func main() {
  fmt.Println("Select the window you want to keep active to front\n10秒内目标窗口为前台");
  time.Sleep(10 * time.Second);
  fmt.Println("Start drawing square 鼠标画块儿开始");
  robotgo.SetMouseDelay(500);
  var round = 0;
  for {
    round++;
    fmt.Printf("Round%d\n", round);
    robotgo.MoveMouseSmooth(500, 500);
    robotgo.MoveMouseSmooth(500, 1000);
    robotgo.MoveMouseSmooth(1000, 1000);
    robotgo.MoveMouseSmooth(1000, 500);
    robotgo.MoveMouseSmooth(500, 500);
    time.Sleep(10 * time.Second);
  }
}
