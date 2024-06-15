package models

type Room struct {
	ID      int     `db:"id"`
	Name    string  `db:"name"`
	Desc    string  `db:"description"`
	Price   float32 `db:"price"`
	ImgPath string  `db:"img_path"`
}
