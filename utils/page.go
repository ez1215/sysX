package utils

type Page struct {
	//当前页码
	PageNo  int64
	//页面显示数量
	PageSize int64
	//总页数
	TotalPage int64
	//总记录数
	TotalCount int64
	//是否是第一页
	FirstPage  bool
	//是否是最后一页
	LastPage   bool
	//封装数据
	List interface{}
}

func PageUtil(count int64, pageNo int64, pageSize int64, list interface{}) Page  {
	tp := count / pageSize
	if count % pageSize > 0 {
		tp = count / pageSize + 1
	}
	return Page{PageNo: pageNo, PageSize: pageSize, TotalPage: tp, TotalCount: count,FirstPage: pageNo == 1,LastPage: pageNo == tp, List: list}
}
