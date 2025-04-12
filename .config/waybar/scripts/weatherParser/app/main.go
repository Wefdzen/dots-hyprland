package main

import (
	"fmt"

	"github.com/Wefdzen/WayWeather/utils"
)

func main() {
	res, err := utils.SendReq()
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	// You can custom Output of programm.
	fmt.Printf("%v  %v°C\n", res.EmojiOfWeather, res.Celsius)
}
