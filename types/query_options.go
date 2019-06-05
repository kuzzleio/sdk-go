// Copyright 2015-2018 Kuzzle
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

// QueryOptions provides a public access to queryOptions private struct
type QueryOptions interface {
	Queuable() bool
	SetQueuable(bool) *queryOptions
	From() int
	SetFrom(int) *queryOptions
	Size() int
	SetSize(int) *queryOptions
	Scroll() string
	SetScroll(string) *queryOptions
	ScrollId() string
	SetScrollId(string) *queryOptions
	Volatile() VolatileData
	SetVolatile(VolatileData) *queryOptions
	Refresh() string
	SetRefresh(string) *queryOptions
	IfExist() string
	SetIfExist(string) *queryOptions
	IncludeTrash() bool
	SetIncludeTrash(bool) *queryOptions
	RetryOnConflict() int
	SetRetryOnConflict(int) *queryOptions
	Start() int
	SetStart(int) *queryOptions
	End() int
	SetEnd(int) *queryOptions
	Count() int
	SetCount(int) *queryOptions
	Sort() string
	SetSort(string) *queryOptions
	Match() string
	SetMatch(string) *queryOptions
	Ch() bool
	SetCh(bool) *queryOptions
	Incr() bool
	SetIncr(bool) *queryOptions
	Nx() bool
	SetNx(bool) *queryOptions
	Xx() bool
	SetXx(bool) *queryOptions
	Ex() int
	SetEx(int) *queryOptions
	Px() int
	SetPx(int) *queryOptions
	Limit() []int
	SetLimit([]int) *queryOptions
	Aggregate() string
	SetAggregate(string) *queryOptions
	Weights() []int
	SetWeights([]int) *queryOptions
	Type() string
	SetType(string) *queryOptions
	By() string
	SetBy(string) *queryOptions
	Direction() string
	SetDirection(string) *queryOptions
	Get() []string
	SetGet([]string) *queryOptions
	Alpha() bool
	SetAlpha(bool) *queryOptions
	Unit() string
	SetUnit(string) *queryOptions
	Withdist() bool
	SetWithdist(bool) *queryOptions
	Withcoord() bool
	SetWithcoord(bool) *queryOptions
	Reset() bool
	ID() string
}

type queryOptions struct {
	queuable        bool
	from            int
	size            int
	scroll          string
	scrollId        string
	volatile        VolatileData
	refresh         string
	ifExist         string
	includeTrash    bool
	retryOnConflict int
	start           int
	end             int
	count           int
	sort            string
	match           string
	ch              bool
	incr            bool
	nx              bool
	xx              bool
	ex              int
	px              int
	limit           []int
	aggregate       string
	weights         []int
	colType         string // type would conflict with the Golang keyword
	by              string
	direction       string
	get             []string
	alpha           bool
	unit            string
	withcoord       bool
	withdist        bool
	reset           bool
	_id             string
}

func (qo queryOptions) Queuable() bool {
	return qo.queuable
}

func (qo *queryOptions) SetQueuable(queuable bool) *queryOptions {
	qo.queuable = queuable
	return qo
}

func (qo queryOptions) From() int {
	return qo.from
}

func (qo *queryOptions) SetFrom(from int) *queryOptions {
	qo.from = from
	return qo
}

func (qo queryOptions) Size() int {
	return qo.size
}

func (qo *queryOptions) SetSize(size int) *queryOptions {
	qo.size = size
	return qo
}

func (qo queryOptions) Scroll() string {
	return qo.scroll
}

func (qo *queryOptions) SetScroll(scroll string) *queryOptions {
	qo.scroll = scroll
	return qo
}

func (qo queryOptions) ScrollId() string {
	return qo.scrollId
}

func (qo *queryOptions) SetScrollId(scrollId string) *queryOptions {
	qo.scrollId = scrollId
	return qo
}

func (qo queryOptions) Volatile() VolatileData {
	return qo.volatile
}

func (qo *queryOptions) SetVolatile(volatile VolatileData) *queryOptions {
	qo.volatile = volatile
	return qo
}

func (qo queryOptions) Refresh() string {
	return qo.refresh
}

func (qo *queryOptions) SetRefresh(refresh string) *queryOptions {
	qo.refresh = refresh
	return qo
}

func (qo queryOptions) IfExist() string {
	return qo.ifExist
}

func (qo *queryOptions) SetIfExist(ifExist string) *queryOptions {
	qo.ifExist = ifExist
	return qo
}

func (qo *queryOptions) IncludeTrash() bool {
	return qo.includeTrash
}

func (qo *queryOptions) SetIncludeTrash(includeTrash bool) *queryOptions {
	qo.includeTrash = includeTrash
	return qo
}

func (qo queryOptions) RetryOnConflict() int {
	return qo.retryOnConflict
}

func (qo *queryOptions) SetRetryOnConflict(retryOnConflict int) *queryOptions {
	qo.retryOnConflict = retryOnConflict
	return qo
}

func (qo queryOptions) Start() int {
	return qo.start
}

func (qo *queryOptions) SetStart(start int) *queryOptions {
	qo.start = start
	return qo
}

func (qo queryOptions) End() int {
	return qo.end
}

func (qo *queryOptions) SetEnd(end int) *queryOptions {
	qo.end = end
	return qo
}

func (qo queryOptions) Count() int {
	return qo.count
}

func (qo *queryOptions) SetCount(count int) *queryOptions {
	qo.count = count
	return qo
}

func (qo queryOptions) Sort() string {
	return qo.sort
}

func (qo *queryOptions) SetSort(sort string) *queryOptions {
	qo.sort = sort
	return qo
}

func (qo queryOptions) Match() string {
	return qo.match
}

func (qo *queryOptions) SetMatch(match string) *queryOptions {
	qo.match = match
	return qo
}

func (qo queryOptions) Ch() bool {
	return qo.ch
}

func (qo *queryOptions) SetCh(ch bool) *queryOptions {
	qo.ch = ch
	return qo
}

func (qo queryOptions) Incr() bool {
	return qo.incr
}

func (qo *queryOptions) SetIncr(incr bool) *queryOptions {
	qo.incr = incr
	return qo
}

func (qo queryOptions) Nx() bool {
	return qo.nx
}

func (qo *queryOptions) SetNx(nx bool) *queryOptions {
	qo.nx = nx
	return qo
}

func (qo queryOptions) Ex() int {
	return qo.ex
}

func (qo *queryOptions) SetEx(ex int) *queryOptions {
	qo.ex = ex
	return qo
}

func (qo queryOptions) Px() int {
	return qo.px
}

func (qo *queryOptions) SetPx(px int) *queryOptions {
	qo.px = px
	return qo
}

func (qo queryOptions) Xx() bool {
	return qo.xx
}

func (qo *queryOptions) SetXx(xx bool) *queryOptions {
	qo.xx = xx
	return qo
}

func (qo queryOptions) Limit() []int {
	return qo.limit
}

func (qo *queryOptions) SetLimit(limit []int) *queryOptions {
	qo.limit = limit
	return qo
}

func (qo queryOptions) Aggregate() string {
	return qo.aggregate
}

func (qo *queryOptions) SetAggregate(aggregate string) *queryOptions {
	qo.aggregate = aggregate
	return qo
}

func (qo queryOptions) Weights() []int {
	return qo.weights
}

func (qo *queryOptions) SetWeights(weights []int) *queryOptions {
	qo.weights = weights
	return qo
}

func (qo queryOptions) Type() string {
	return qo.colType
}

func (qo *queryOptions) SetType(colType string) *queryOptions {
	qo.colType = colType
	return qo
}

func (qo *queryOptions) By() string {
	return qo.by
}

func (qo *queryOptions) SetBy(by string) *queryOptions {
	qo.by = by
	return qo
}

func (qo *queryOptions) Direction() string {
	return qo.direction
}

func (qo *queryOptions) SetDirection(direction string) *queryOptions {
	qo.direction = direction
	return qo
}

func (qo *queryOptions) Get() []string {
	return qo.get
}

func (qo *queryOptions) SetGet(get []string) *queryOptions {
	qo.get = get
	return qo
}

func (qo *queryOptions) Alpha() bool {
	return qo.alpha
}

func (qo *queryOptions) SetAlpha(alpha bool) *queryOptions {
	qo.alpha = alpha
	return qo
}

func (qo *queryOptions) Unit() string {
	return qo.unit
}

func (qo *queryOptions) SetUnit(unit string) *queryOptions {
	qo.unit = unit
	return qo
}

func (qo *queryOptions) Withcoord() bool {
	return qo.withcoord
}

func (qo *queryOptions) SetWithcoord(withcoord bool) *queryOptions {
	qo.withcoord = withcoord
	return qo
}

func (qo *queryOptions) Withdist() bool {
	return qo.withdist
}

func (qo *queryOptions) SetWithdist(withdist bool) *queryOptions {
	qo.withdist = withdist
	return qo
}

func (qo *queryOptions) Reset() bool {
	return qo.reset
}

func (qo *queryOptions) ID() string {
	return qo._id
}

// NewQueryOptions instanciates a new QueryOptions with default values
func NewQueryOptions() *queryOptions {
	return &queryOptions{
		size:    10,
		ifExist: "error",
	}
}
