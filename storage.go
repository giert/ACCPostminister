package main

type storage struct {
	Data map[string]entry
}

func (s storage) listMovies() (result string) {
	for _, v := range s.Data {
		result += v.println() + "\n"
	}
	return
}
