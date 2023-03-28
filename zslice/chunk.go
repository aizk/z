package zslice

func Chunk[T any](r []T, n int) (s [][]T) {
	l := len(r)
	if l == 0 || n == 0 {
		return nil
	}

	if l < n {
		s = [][]T{r}
		return
	}

	// c round
	c := l / n
	if l %n != 0 {
		c += 1
	}

	s = make([][]T, c)

	i := 0
	j := 0
	k := n
	// build chunk
	for i < c {
		s[i] = r[j:k]
		i += 1
		j += n
		k += n
		if k >= l {
			k = l
		}
	}

	return
}
