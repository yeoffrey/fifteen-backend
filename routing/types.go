package routing

type Mode = string

const (
	Biking Mode = "bike"
	Walk   Mode = "walk"
)

type Route struct {
	Distance int     `json:"distance"`
	Duration float64 `json:"duration"`
	Mode     Mode    `json:"mode"`
}
