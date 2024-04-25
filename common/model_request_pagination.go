package common

import (
	"fmt"
	"strconv"
)

type Pagination struct {
	Current_Page int `json:"current-page"`
	Limit        int `json:"limit"`
	Total_Item   int `json:"total-item"`
	Total_Page   int `json:"total-page"`
}

func (p *Pagination) Process(page, total string) error {
	if page == "" {
		p.Current_Page = 1
	} else {
		number, err := strconv.Atoi(page)
		if err != nil {
			return fmt.Errorf("page is not a number: %v", err)
		}
		p.Current_Page = number
	}
	if total == "" {
		p.Limit = 10
	} else {
		number, err := strconv.Atoi(total)
		if err != nil {
			return fmt.Errorf("total is not a number: %v", err)
		}
		p.Limit = number
	}
	return nil
}
