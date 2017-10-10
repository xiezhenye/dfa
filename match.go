package dfa

import "io"

func (m *M) Match(src []byte) (size, Label int, matched bool) {
	return m.MatchAt(src, 0)
}

// Match greedily matches the DFA against src.
func (m *M) MatchAt(src []byte, p int) (size, Label int, matched bool) {
	var (
		s, matchedState *S
		sid             = 0
	)
	pos := p
	matchedPos := pos
	for sid >= 0 {
		s = &m.States[sid]
		if s.Label >= defaultFinal {
			matchedState = s
			matchedPos = pos
		}
		if pos < len(src) {
			sid = s.next(src[pos])
			if sid >= 0 {
				pos++
			}
		} else {
			break
		}
	}
	if matchedState != nil {
		return matchedPos, matchedState.Label.toExternal(), true
	}
	return matchedPos - p, -1, false
}

func (m *M) MatchReader(src io.Reader, dest *[]byte) (size, Label int, matched bool, err error) {
	var (
		s, matchedState *S
		sid             = 0
	)
	pos := 0
	matchedPos := pos
	var t [1]byte
	for sid >= 0 {
		s = &m.States[sid]
		if s.Label >= defaultFinal {
			matchedState = s
			matchedPos = pos
		}
		_, err := src.Read(t[:])
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return 0, -1, false, err
			}
		}
		sid = s.next(t[0])
		if sid >= 0 {
			pos++
		}
		*dest = append(*dest, t[0])
	}
	if matchedState != nil {
		return matchedPos, matchedState.Label.toExternal(), true, nil
	}
	return 0, -1, false, nil
}

func (m *FastM) Match(src []byte) (size, Label int, matched bool) {
	return m.MatchAt(src, 0)
}

// Match greedily matches the DFA against src.
func (m *FastM) MatchAt(src []byte, p int) (size, Label int, matched bool) {
	cur := &m.States[0]
	pos := p
	matchedPos := pos
	for {
		if cur.Label >= 0 {
			matchedPos = pos
			Label = cur.Label
			matched = true
		}
		if pos == len(src) {
			break
		}
		if cur = cur.Trans[src[pos]]; cur == nil {
			break
		}
		pos++
	}
	size = matchedPos - p
	return
}

func (m *FastM) MatchReader(src io.Reader, dest *[]byte) (size, Label int, matched bool, err error) {
	cur := &m.States[0]
	pos := 0
	matchedPos := pos
	var t [1]byte
	for {
		if cur.Label >= 0 {
			matchedPos = pos
			Label = cur.Label
			matched = true
		}
		_, err := src.Read(t[:])
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return 0, -1, false, err
			}
		}
		if cur = cur.Trans[t[0]]; cur == nil {
			break
		}
		*dest = append(*dest, t[0])
		pos++
	}
	size = matchedPos
	return
}
