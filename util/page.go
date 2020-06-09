package util

type Pageable interface {
	GetPageNo() int32
	GetPageSize() int32
}

type DefaultPage struct {
	pageNo,
	pageSize int32
}

func (p *DefaultPage) GetPageNo() int32 {
	return p.pageNo
}

func (p *DefaultPage) GetPageSize() int32 {
	return p.pageSize
}
