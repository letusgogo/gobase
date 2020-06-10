package util

type Pageable interface {
	GetPageNo() int64
	GetPageSize() int64
	GetOffset() int64
}

type GormPage struct {
	pageNo,
	pageSize int64
}

func NewGormPage(pageNo int64, pageSize int64) *GormPage {
	return &GormPage{pageNo: pageNo, pageSize: pageSize}
}

func (p *GormPage) GetOffset() int64 {
	// gorm offset == -1 的时候是取消 offset 限制
	if p.pageSize <= 0 {
		return -1
	}
	return (p.pageNo - 1) * p.pageSize
}

func (p *GormPage) GetPageNo() int64 {
	return p.pageNo
}

func (p *GormPage) GetPageSize() int64 {
	// gorm pageSize == -1 的时候是取消 limit 限制
	if p.pageSize <= 0 {
		return -1
	}
	return p.pageSize
}
