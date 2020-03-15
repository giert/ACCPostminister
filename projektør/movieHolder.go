package projektør

type movieHolder struct {
	Movies []movie
}

func (mh movieHolder) Print() (result string) {
	for _, v := range mh.Movies {
		result += v.Print() + "\n"
	}
	return
}
