package mocks

type Spy struct {
	CallTimes int
	Called    bool
	Returns   []any
	Params    []any
}

func (s *Spy) ExpectedParams(params []any) bool {
	return equalSlices(s.Params, params)
}

func (s *Spy) ExpectedReturns(returns []any) bool {
	return equalSlices(s.Returns, returns)
}

func equalSlices(wantSlice, gotSlice []any) bool {
	for i := 0; i < len(wantSlice); i++ {
		want := wantSlice[i]
		got := gotSlice[i]

		if want != got {
			return false
		}
	}

	return true
}
