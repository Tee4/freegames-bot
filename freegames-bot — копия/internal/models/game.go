package models

import "time"

type Game struct {
	Title   string
	EndDate *time.Time
	Links   []GameLink
}

type GameLink struct {
	Store string
	URL   string
	Note  string
}
