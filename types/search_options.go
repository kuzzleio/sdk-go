package types

// SearchOptions options for search functions
type SearchOptions struct {
	Type   string
	From   *int
	Size   *int
	Scroll string
}

//SetType Type setter
func (so *SearchOptions) SetType(nt string) {
	so.Type = nt
}

//SetFrom From setter
func (so *SearchOptions) SetFrom(nf int) {
	so.From = &nf
}

//SetSize Size setter
func (so *SearchOptions) SetSize(ns int) {
	so.Size = &ns
}

//SetScroll Scroll setter
func (so *SearchOptions) SetScroll(ns string) {
	so.Scroll = ns
}
