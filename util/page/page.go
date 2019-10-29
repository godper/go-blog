package page

//Page 页码
type Page struct {
	PageNum  int `json:"page_num" `
	PageSize int `json:"page_size"`
}

//NewPage Newpage实例
func NewPage(pageSize int) *Page {
	return &Page{
		PageNum:  1,
		PageSize: pageSize,
	}
}

//OffsetLimit offsetlimit
func (p *Page) OffsetLimit(pageNum int, pageSize int) (offset int, limit int) {
	if pageNum <= 0 {
		pageNum = p.PageNum
	}
	if pageSize <= 0 {
		pageSize = p.PageSize
	}

	offset = (pageNum - 1) * pageSize
	limit = pageSize
	return
}
