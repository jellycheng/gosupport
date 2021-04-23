package utils

import (
	"fmt"
	"strconv"
)

/*
	页码处理
	调用示例：
	pageObj := utils.NewPage()
	pageObj.MustSetPage(1).MustSetPageSize(15)
	fmt.Println(pageObj.GetLimit())
 */
type Page struct {
	page int64   //页码，起始页码从1开始，计算limit或offset时，如果值小于1则强制修改为1
	pageSize int64  //每页记录数，默认15,计算limit或offset时，如果值小于1则强制修改为15

}

func NewPage() *Page {
	res := &Page{page: 1, pageSize: 15}
	return res
}
func (p *Page)SetPage(pageVal interface{}) (*Page, error)  {
	if v, err := strconv.ParseInt(fmt.Sprintf("%v", pageVal), 10, 64);err==nil {
		p.page = v
		return p,nil
	} else {
		return p,err
	}
}

func (p *Page)MustSetPage(pageVal interface{}) (*Page)  {
	p.SetPage(pageVal)
	return p
}

func (p *Page)GetPage() int64 {
	return p.page
}

func (p *Page)GetPage2int() int {
	return int(p.page)
}

func (p *Page)SetPageSize(pageSizeVal interface{}) (*Page, error)  {
	if v, err := strconv.ParseInt(fmt.Sprintf("%v", pageSizeVal), 10, 64);err==nil {
		p.pageSize = v
		return p,nil
	} else {
		return p,err
	}
}

func (p *Page)MustSetPageSize(pageSizeVal interface{}) (*Page)  {
	p.SetPageSize(pageSizeVal)
	return p
}

func (p *Page)GetPageSize() int64 {
	return p.pageSize
}

func (p *Page)GetPageSize2int() int {
	return int(p.pageSize)
}

//返回mysql limit部分的内容，格式："0,15"
func (p *Page)GetLimitStr() string {
	limit := p.GetLimit()
	return fmt.Sprintf("%d,%d", limit[0], limit[1])
}

func (p *Page)GetLimit() []int64 {
	var limit []int64
	page := p.page
	if page < 1 {
		page = 1
	}
	pageSize := p.pageSize
	if pageSize < 1 {
		pageSize = 15
	}
	limit = append(limit, (page - 1) * pageSize, pageSize)
	return limit
}

func (p *Page) GetStartLimit() int64  {
	limit := p.GetLimit()
	return limit[0]
}

func (p *Page) GetStartLimit2int() int {
	return int(p.GetStartLimit())
}

func (p *Page) GetEndLimit() int64  {
	limit := p.GetLimit()
	return limit[1]
}

func (p *Page) GetEndLimit2int() int {
	return int(p.GetEndLimit())
}

//获取mysql offset 部分的值： SELECT * FROM 表名 OFFSET 偏移量 LIMIT 每页记录数
func (p *Page) GetOffset() int64  {
	return p.GetStartLimit()
}

func (p *Page) GetOffset2int() int {
	return int(p.GetStartLimit())
}

func (p *Page) GormDbGetLimit() int64 {
	pageSize := p.pageSize
	if pageSize < 1 {
		pageSize = 15
	}
	return pageSize
}

func (p *Page) GormDbGetLimit2int() int {
	return int(p.GormDbGetLimit())
}

// 获取总页数，count=总记录数，pageSize=每页记录数
func TotalPage(count int, pageSize int) int {
	result := count / pageSize
	yu := count % pageSize
	if yu > 0 {
		result = result + 1
	}
	return result
}
