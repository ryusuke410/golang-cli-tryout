package fp

func MapE[S any, T any](in []S, f func(S, int) (T, error)) ([]T, error) {
	result := make([]T, 0, len(in))
	for i, v := range in {
		t, err := f(v, i)
		if err != nil {
			return nil, err
		}
		result = append(result, t)
	}
	return result, nil
}
