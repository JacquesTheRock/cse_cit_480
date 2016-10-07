package entity

type Mail struct {
	ID      int64  `json:"id"`
	Src     string `json:"src"`
	Dest    string `json:"dest"`
	Prev    int64  `json:"prev_id"`
	Date    string `json:"date"`
	Message string `json:"message"`
}
