package projektør

type movie struct {
	ID     int
	Title  string
	IMDB   string
	Link   string
	Sub    string
	Status int
}

const (
	unwatched = 0
	watching  = 1
	watched   = 2
)

func (m movie) Print() string {
	return m.Title
}
