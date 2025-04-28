package in

import (
	"errors"
	"kredi-plus.com/be/lib/helper"
	"strings"
)

type AbstractRequest struct {
	Page    int      `json:"page"`
	Limit   int      `json:"limit"`
	Search  string   `json:"search"`
	SortStr string   `json:"sort"`
	Sort    []string `json:"-"`
}

func (model *AbstractRequest) GetSort(defaultSort string) (err error) {
	model.Sort = []string{}
	arrAvailableSort := []string{"ASC", "DESC"}
	if model.SortStr == "" {
		model.Sort = []string{defaultSort}
	} else {
		arrSort := strings.Split(model.SortStr, ",")
		for _, value := range arrSort {
			arrSortValue := strings.Split(value, " ")
			if len(arrSortValue) == 1 {
				model.Sort = append(model.Sort, arrSortValue[0]+" DESC")
			} else {
				arrSortValue[1] = strings.ToUpper(arrSortValue[1])
				if !helper.CheckDataOnSlice(arrAvailableSort[1], arrAvailableSort) {
					return errors.New("sort type must be ASC or DESC")
				}
				model.Sort = append(model.Sort, strings.Join(arrSortValue, " "))
			}
		}
	}
	return
}

func (model AbstractRequest) GetOffset() int {
	return (model.Page - 1) * model.Limit
}

func (input AbstractRequest) GetLimit() int {
	if input.Limit == 0 {
		return -1
	}
	return input.Limit
}
