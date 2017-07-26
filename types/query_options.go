package types

type QueryOptions interface {
	GetQueuable() bool
	SetQueuable(bool)
	GetFrom() int
	SetFrom(int)
	GetSize() int
	SetSize(int)
	GetScroll() string
	SetScroll(string)
	GetScrollId() string
	SetScrollId(string)
	GetVolatile() VolatileData
	SetVolatile(VolatileData)
	GetRefresh() string
	SetRefresh(string)
	GetIfExist() string
	SetIfExist(string)
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
}

func (qo queryOptions) GetQueuable() bool {
	return qo.queuable
}

func (qo *queryOptions) SetQueuable(queuable bool) {
	qo.queuable = queuable
}

func (qo queryOptions) GetFrom() int {
	return qo.from
}

func (qo *queryOptions) SetFrom(from int) {
	qo.from = from
}

func (qo queryOptions) GetSize() int {
	return qo.size
}

func (qo *queryOptions) SetSize(size int) {
	qo.size = size
}

func (qo queryOptions) GetScroll() string {
	return qo.scroll
}

func (qo queryOptions) SetScroll(scroll string) {
	qo.scroll = scroll
}

func (qo queryOptions) GetScrollId() string {
	return qo.scrollId
}

func (qo *queryOptions) SetScrollId(scrollId string) {
	qo.scrollId = scrollId
}

func (qo queryOptions) GetVolatile() VolatileData {
	return qo.volatile
}

func (qo *queryOptions) SetVolatile(volatile VolatileData) {
	qo.volatile = volatile
}

func (qo queryOptions) GetRefresh() string {
	return qo.refresh
}

func (qo *queryOptions) SetRefresh(refresh string) {
	qo.refresh = refresh
}

func (o queryOptions) GetIfExist() string {
	return o.ifExist
}

func (o *queryOptions) SetIfExist(ifExist string) {
	o.ifExist = ifExist
}

func NewQueryOptions() *queryOptions {
	return &queryOptions{
		size: 10,
	}
}
