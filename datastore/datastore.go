package datastore

type Storage struct {
	Data map[string]Entry
}

type Entry interface {
	Print() string
}

func (s Storage) ListMovies() (result string) {
	for _, v := range s.Data {
		result += v.Print() + "\n"
	}
	return
}
