package entity

type URL struct {
	ID          string
	OriginalURL string
}

func NewURL(ID string, OriginalURL string) *URL {
	return &URL{ID, OriginalURL}
}
