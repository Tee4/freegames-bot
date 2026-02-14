package gamerpower

// Giveaway — один giveaway из API GamerPower
// https://www.gamerpower.com/api-read
type Giveaway struct {
	ID              int    `json:"id"`
	Title           string `json:"title"`
	Worth           string `json:"worth"`
	Thumbnail       string `json:"thumbnail"`
	Image           string `json:"image"`
	Description     string `json:"description"`
	Instructions    string `json:"instructions"`
	OpenGiveawayURL string `json:"open_giveaway_url"`
	Platforms       string `json:"platforms"`
	EndDate         string `json:"end_date"`
	Type            string `json:"type"`
	Status          string `json:"status"`
}
