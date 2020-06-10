package util

type Pageable interface {
	GetPageNo() int64
	GetPageSize() int64
	GetOffset() int64
}

type DefaultPage struct {
	pageNo,
	pageSize int64
}

func NewDefaultPage(pageNo int64, pageSize int64) *DefaultPage {
	return &DefaultPage{pageNo: pageNo, pageSize: pageSize}
}

func (p *DefaultPage) GetOffset() int64 {
	return (p.pageNo - 1) * p.pageSize
}

func (p *DefaultPage) GetPageNo() int64 {
	return p.pageNo
}

func (p *DefaultPage) GetPageSize() int64 {
	return p.pageSize
}
