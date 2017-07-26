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

func (qo queryOptions) SetScroll(scroll string) *queryOptions {
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

func (o queryOptions) GetIfExist() string {
	return o.ifExist
}

func (o *queryOptions) SetIfExist(ifExist string) *queryOptions {
	o.ifExist = ifExist
	return qo
}

func NewQueryOptions() *queryOptions {
	return &queryOptions{
		size: 10,
	}
}
