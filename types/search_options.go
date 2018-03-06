package types

type SearchOptions struct {
	Type   string
	From   int
	Size   int
	Scroll string
}

func NewSearchOptions(t string, from int, size int, scroll string) *SearchOptions {
	return &SearchOptions{
		Type:   t,
		From:   from,
		Size:   size,
		Scroll: scroll,
	}
}

func (so *SearchOptions) SetType(nt string) *SearchOptions {
	so.Type = nt
	return so
}

func (so *SearchOptions) SetFrom(nf int) *SearchOptions {
	so.From = nf
	return so
}

func (so *SearchOptions) SetSize(ns int) *SearchOptions {
	so.Size = ns
	return so
}

func (so *SearchOptions) SetScroll(ns string) *SearchOptions {
	so.Scroll = ns
	return so
}
