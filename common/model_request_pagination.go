package common

type Pagination struct {
	Current_Page int   `json:"current-page"`
	Limit        int   `json:"limit"`
	Total_Item   int64 `json:"total-item"`
	Total_Page   int64 `json:"total-page"`
}

func (p *Pagination) Process() {
	if p.Current_Page < 1 {
		p.Current_Page = 1
	}
	if p.Limit < 0 || p.Limit >= 100 {
		p.Limit = 10
	}
}
