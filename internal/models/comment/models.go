package models

type Comment struct {
	ID      int     `json:"id"`
	MusicID int     `json:"music_id"`
	UserID  int     `json:"user_id"`
	Content string  `json:"content"`
	Rating  float64 `json:"rating"`
}
