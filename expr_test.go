package dfa

import "testing"

func TestStr(t *testing.T) {
	s := Str("hello").As(100)
	n, l, ok := s.Match([]byte("hello"))
	if !ok || n != 5 || l != 100 {
		t.Error("match fail")
	}
	n, l, ok = s.Match([]byte("hello world"))
	if !ok || n != 5 || l != 100 {
		t.Error("match fail")
	}
	n, l, ok = s.Match([]byte("!hello"))
	if ok {
		t.Error("match should fail")
	}
}

func TestChar(t *testing.T) {
	c := Char("+-*/").As(101)
	n, l, ok := c.Match([]byte("+"))
	if !ok || n != 1 || l != 101 {
		t.Error("match fail")
	}
	n, l, ok = c.Match([]byte("*1"))
	if !ok || n != 1 || l != 101 {
		t.Error("match fail")
	}
	n, l, ok = c.Match([]byte("1*"))
	if ok {
		t.Error("match should fail")
	}
}

func TestBetween(t *testing.T) {
	a := Between(rune('a'), rune('z')).As(101)
	n, l, ok := a.Match([]byte("a+b"))
	if !ok || n != 1 || l != 101 {
		t.Error("match fail")
	}
	n, l, ok = a.Match([]byte("111"))
	if ok {
		t.Error("match should fail")
	}
	a = Between(rune(0x00), rune(0x20)).As(101)
	n, l, ok = a.Match([]byte(" "))
	if !ok || n != 1 || l != 101 {
		t.Error("match fail")
	}
	a = Between(rune('一'), rune('龥')).As(102)
	n, l, ok = a.Match([]byte("中"))
	if !ok || n != 3 || l != 102 {
		t.Error("match fail")
	}
}

func TestBetweenByte(t *testing.T) {
	a := BetweenByte('a', 'z').As(101)
	n, l, ok := a.Match([]byte("a+b"))
	if !ok || n != 1 || l != 101 {
		t.Error("match fail")
	}
	n, l, ok = a.Match([]byte("111"))
	if ok {
		t.Error("match should fail")
	}
	a = BetweenByte('\x00', '\x20').As(101)
	n, l, ok = a.Match([]byte(" "))
	if !ok || n != 1 || l != 101 {
		t.Error("match fail")
	}
}

func TestCharClass(t *testing.T) {
	a := CharClass("Han").As(102)
	n, l, ok := a.Match([]byte("中"))
	if !ok || n != 3 || l != 102 {
		t.Error("match fail")
	}
	n, l, ok = a.Match([]byte("english"))
	if ok {
		t.Error("match should fail")
	}
}

func TestAnd(t *testing.T) {
	a := BetweenByte('a', 'e')
	b := BetweenByte('c', 'g')
	c := And(a, b)
	n, _, ok := c.Match([]byte("cc"))
	if !ok || n != 1 {
		t.Error("match fail")
	}
	n, _, ok = c.Match([]byte("aa"))
	if ok {
		t.Error("match should fail")
	}
	n, _, ok = c.Match([]byte("ff"))
	if ok {
		t.Error("match should fail")
	}
}

func TestCon(t *testing.T) {
	a := BetweenByte('a', 'z')
	b := BetweenByte('0', '9')
	c := Con(a, b)
	n, _, ok := c.Match([]byte("a1"))
	if !ok || n != 2 {
		t.Error("match fail")
	}
	n, _, ok = c.Match([]byte("aa"))
	if ok {
		t.Error("match should fail")
	}
	n, _, ok = c.Match([]byte("11"))
	if ok {
		t.Error("match should fail")
	}
}

func TestOr(t *testing.T) {
	ae := BetweenByte('a', 'e')
	cg := BetweenByte('c', 'g')
	a := Or(ae, cg)
	n, _, ok := a.Match([]byte("cat"))
	if !ok || n != 1 {
		t.Error("match fail")
	}
	n, _, ok = a.Match([]byte("egg"))
	if !ok || n != 1 {
		t.Error("match fail")
	}
	n, _, ok = a.Match([]byte("bad"))
	if !ok || n != 1 {
		t.Error("match fail")
	}
	n, _, ok = a.Match([]byte("fat"))
	if !ok || n != 1 {
		t.Error("match fail")
	}
	n, _, ok = a.Match([]byte("add"))
	if !ok || n != 1 {
		t.Error("match fail")
	}
	n, _, ok = a.Match([]byte("get"))
	if !ok || n != 1 {
		t.Error("match fail")
	}
	n, _, ok = a.Match([]byte("123"))
	if ok {
		t.Error("match should fail")
	}
	n, _, ok = a.Match([]byte("hat"))
	if ok {
		t.Error("match should fail")
	}
}

func TestOr2(t *testing.T) {
	nm := BetweenByte('a', 'e').AtLeast(1).As(100)
	kw := Str("hello").As(200)
	a := Or(nm, kw)
	n, l, ok := a.Match([]byte("abc"))
	if !ok || n != 3 || l != 100 {
		t.Error("match fail")
	}
	n, l, ok = a.Match([]byte("hello"))
	if !ok || n != 5 || l != 200 {
		t.Error("match fail")
	}
}

func TestAtLeast(t *testing.T) {
	a := BetweenByte('a', 'z').AtLeast(3)
	n, _, ok := a.Match([]byte("add"))
	if !ok || n != 3 {
		t.Error("match fail")
	}
	n, _, ok = a.Match([]byte("flag"))
	if !ok || n != 4 {
		t.Error("match fail")
	}
	n, _, ok = a.Match([]byte("xx"))
	if ok {
		t.Error("match should fail")
	}
	n, _, ok = a.Match([]byte("123"))
	if ok {
		t.Error("match should fail")
	}
	a = BetweenByte('a', 'z').AtLeast(0)
	n, _, ok = a.Match([]byte("add"))
	if !ok || n != 3 {
		t.Error("match fail")
	}
	n, _, ok = a.Match([]byte("123"))
	if !ok || n != 0 {
		t.Error("match fail")
	}
}

func TestAtMost(t *testing.T) {
	a := BetweenByte('a', 'z').AtMost(3)
	n, _, ok := a.Match([]byte("add"))
	if !ok || n != 3 {
		t.Error("match fail")
	}
	n, _, ok = a.Match([]byte("123"))
	if !ok || n != 0 {
		t.Error("match fail")
	}
	n, _, ok = a.Match([]byte("flag"))
	if !ok || n != 3 {
		t.Error("match fail")
	}
}

func TestLoop(t *testing.T) {
	a := BetweenByte('0', '9').Loop(IfNot('0'))
	n, _, ok := a.Match([]byte("11101"))
	if !ok && n != 3 {
		t.Error("match fail")
	}
	n, _, ok = a.Match([]byte("011"))
	if !ok && n != 1 {
		t.Error("match fail")
	}
	n, _, ok = a.Match([]byte("111"))
	if !ok && n != 3 {
		t.Error("match fail")
	}
}

func TestOptional(t *testing.T) {
	a := BetweenByte('a', 'z')
	b := BetweenByte('0', '9')
	c := Con(a, Optional(b))
	n, _, ok := c.Match([]byte("a"))
	if !ok && n != 1 {
		t.Error("match fail")
	}
	n, _, ok = c.Match([]byte("a9"))
	if !ok && n != 2 {
		t.Error("match fail")
	}
	n, _, ok = c.Match([]byte("9a"))
	if ok {
		t.Error("match should fail")
	}
}

func TestComplement(t *testing.T) {
	//a := Con(BetweenByte('a', 'z'), BetweenByte('a', 'z').Complement())
	//a.SaveSVG("/tmp/a.svg")
	//n, _, ok := a.Match([]byte("a"))
	//if !ok && n != 1 {
	//	t.Error("match fail")
	//}
	//n, _, ok = a.Match([]byte("a9"))
	//if !ok && n != 2 {
	//	t.Error("match fail")
	//}
	//n, _, ok = a.Match([]byte("aaa"))
	//println(n)
	//if ok {
	//	t.Error("match should fail")
	//}
}

func TestExclude(t *testing.T) {
	a := BetweenByte('a', 'z').Exclude(Char("aeiou"))
	n, _, ok := a.Match([]byte("b"))
	if !ok && n != 1 {
		t.Error("match fail")
	}
	n, _, ok = a.Match([]byte("a"))
	if ok {
		t.Error("match should fail")
	}
	n, _, ok = a.Match([]byte("9a"))
	if ok {
		t.Error("match should fail")
	}
}

func TestRepeat(t *testing.T) {
	a := BetweenByte('a', 'z')
	b := a.Repeat()
	n, _, ok := b.Match([]byte("b"))
	if !ok && n != 1 {
		t.Error("match fail")
	}
	n, _, ok = b.Match([]byte("bbb"))
	if !ok && n != 3 {
		t.Error("match fail")
	}
	n, _, ok = b.Match([]byte(""))
	if !ok && n != 0 {
		t.Error("match fail")
	}
	n, _, ok = b.Match([]byte("1"))
	if !ok && n != 0 {
		t.Error("match fail")
	}

	b = a.Repeat(3)
	n, _, ok = b.Match([]byte("bbb"))
	if !ok && n != 3 {
		t.Error("match fail")
	}
	n, _, ok = b.Match([]byte("bb"))
	if ok {
		t.Error("match should fail")
	}
	n, _, ok = b.Match([]byte("bbbb"))
	if !ok && n != 3 {
		t.Error("match fail")
	}

	b = a.Repeat(2, 3)
	n, _, ok = b.Match([]byte("bbb"))
	if !ok && n != 3 {
		t.Error("match fail")
	}
	n, _, ok = b.Match([]byte("bb3"))
	if !ok && n != 2 {
		t.Error("match fail")
	}
	n, _, ok = b.Match([]byte("bbbb"))
	if !ok && n != 3 {
		t.Error("match fail")
	}
}

func TestInvalidPrefix(t *testing.T) {
	a := BetweenByte('a', 'z')
	b := Or(a.Repeat(2).As(1), a.Repeat(3).As(2))

	c := b.InvalidPrefix().Minimize()
	//println(c.String())
	n, _, ok := c.Match([]byte("x0"))
	if !ok || n != 2 {
		t.Error("match fail")
	}
	n, _, ok = c.Match([]byte("xx"))
	if ok {
		t.Error("match should fail")
	}
	n, _, ok = c.Match([]byte("abc"))
	if ok {
		t.Error("match should fail")
	}
}
