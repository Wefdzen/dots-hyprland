package weather

type Weather struct {
	City           string `json:"city"`
	Country        string `json:"country"`
	Celsius        int16  `json:"celsius"`
	EmojiOfWeather string `json:"emojiofweather"`
}
