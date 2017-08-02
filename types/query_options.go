package types

type QueryOptions interface {
	GetQueuable() bool
	SetQueuable(bool) *queryOptions
	GetFrom() int
	SetFrom(int) *queryOptions
	GetSize() int
	SetSize(int) *queryOptions
	GetScroll() string
	SetScroll(string) *queryOptions
	GetScrollId() string
	SetScrollId(string) *queryOptions
	GetVolatile() VolatileData
	SetVolatile(VolatileData) *queryOptions
	GetRefresh() string
	SetRefresh(string) *queryOptions
	GetIfExist() string
	SetIfExist(string) *queryOptions
	GetStart() int
	SetStart(int) *queryOptions
	GetEnd() int
	SetEnd(int) *queryOptions
	GetCount() int
	SetCount(int) *queryOptions
	GetSort() string
	SetSort(string) *queryOptions
	GetMatch() string
	SetMatch(string) *queryOptions
	GetCh() bool
	SetCh(bool) *queryOptions
	GetIncr() bool
	SetIncr(bool) *queryOptions
	GetNx() bool
	SetNx(bool) *queryOptions
	GetXx() bool
	SetXx(bool) *queryOptions
	GetLimit() []int
	SetLimit([]int) *queryOptions
}

type queryOptions struct {
	queuable bool
	from     int
	size     int
	scroll   string
	scrollId string
	volatile VolatileData
	refresh  string
	ifExist  string
	start    int
	end      int
	count    int
	sort     string
	match    string
	ch       bool
	incr     bool
	nx       bool
	xx       bool
	limit    []int
}

func (qo queryOptions) GetQueuable() bool {
	return qo.queuable
}

func (qo *queryOptions) SetQueuable(queuable bool) *queryOptions {
	qo.queuable = queuable
	return qo
}

func (qo queryOptions) GetFrom() int {
	return qo.from
}

func (qo *queryOptions) SetFrom(from int) *queryOptions {
	qo.from = from
	return qo
}

func (qo queryOptions) GetSize() int {
	return qo.size
}

func (qo *queryOptions) SetSize(size int) *queryOptions {
	qo.size = size
	return qo
}

func (qo queryOptions) GetScroll() string {
	return qo.scroll
}

func (qo *queryOptions) SetScroll(scroll string) *queryOptions {
	qo.scroll = scroll
	return qo
}

func (qo queryOptions) GetScrollId() string {
	return qo.scrollId
}

func (qo *queryOptions) SetScrollId(scrollId string) *queryOptions {
	qo.scrollId = scrollId
	return qo
}

func (qo queryOptions) GetVolatile() VolatileData {
	return qo.volatile
}

func (qo *queryOptions) SetVolatile(volatile VolatileData) *queryOptions {
	qo.volatile = volatile
	return qo
}

func (qo queryOptions) GetRefresh() string {
	return qo.refresh
}

func (qo *queryOptions) SetRefresh(refresh string) *queryOptions {
	qo.refresh = refresh
	return qo
}

func (qo queryOptions) GetIfExist() string {
	return qo.ifExist
}

func (qo *queryOptions) SetIfExist(ifExist string) *queryOptions {
	qo.ifExist = ifExist
	return qo
}

func (qo queryOptions) GetStart() int {
	return qo.start
}

func (qo *queryOptions) SetStart(start int) *queryOptions {
	qo.start = start
	return qo
}

func (qo queryOptions) GetEnd() int {
	return qo.end
}

func (qo *queryOptions) SetEnd(end int) *queryOptions {
	qo.end = end
	return qo
}

func (qo queryOptions) GetCount() int {
	return qo.count
}

func (qo *queryOptions) SetCount(count int) *queryOptions {
	qo.count = count
	return qo
}

func (qo queryOptions) GetSort() string {
	return qo.sort
}

func (qo *queryOptions) SetSort(sort string) *queryOptions {
	qo.sort = sort
	return qo
}

func (qo queryOptions) GetMatch() string {
	return qo.match
}

func (qo *queryOptions) SetMatch(match string) *queryOptions {
	qo.match = match
	return qo
}

func (qo queryOptions) GetCh() bool {
	return qo.ch
}

func (qo *queryOptions) SetCh(ch bool) *queryOptions {
	qo.ch = ch
	return qo
}

func (qo queryOptions) GetIncr() bool {
	return qo.incr
}

func (qo *queryOptions) SetIncr(incr bool) *queryOptions {
	qo.incr = incr
	return qo
}

func (qo queryOptions) GetNx() bool {
	return qo.nx
}

func (qo *queryOptions) SetNx(nx bool) *queryOptions {
	qo.nx = nx
	return qo
}

func (qo queryOptions) GetXx() bool {
	return qo.xx
}

func (qo *queryOptions) SetXx(xx bool) *queryOptions {
	qo.xx = xx
	return qo
}

func (qo queryOptions) GetLimit() []int {
	return qo.limit
}

func (qo *queryOptions) SetLimit(limit []int) *queryOptions {
	qo.limit = limit
	return qo
}

func NewQueryOptions() *queryOptions {
	return &queryOptions{
		size: 10,
	}
}
