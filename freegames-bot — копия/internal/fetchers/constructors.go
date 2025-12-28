package fetchers

func NewEpic() *EpicFetcher {
	return &EpicFetcher{}
}

func NewSteam() *SteamFetcher {
	return &SteamFetcher{}
}

func NewGOG() *GOGFetcher {
	return &GOGFetcher{}
}

func NewGamerPower() *GamerPowerFetcher {
	return &GamerPowerFetcher{}
}

func NewIndieGala() *IndieGalaFetcher {
	return &IndieGalaFetcher{}
}

func NewFreeGameLoot() *FreeGameLootFetcher {
	return &FreeGameLootFetcher{}
}

func NewItch() *ItchFetcher {
	return &ItchFetcher{}
}
