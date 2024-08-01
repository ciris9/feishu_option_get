package model

import (
	"fmt"
	"option_get/constants"
)

type BiTableRecordsResponse struct {
	Msg  string              `json:"msg"`
	Code int64               `json:"code"`
	Data *BiTableRecordsData `json:"data"`
}

type BiTableRecordsData struct {
	Total   int64                 `json:"total"`
	HasMore bool                  `json:"has_more"`
	Items   []*BiTableRecordsItem `json:"items"`
}

type BiTableRecordsItem struct {
	Fields   map[string]any `json:"fields"`
	RecordID string         `json:"record_id"`
}

func NewBiTableRecordsResp() *BiTableRecordsResponse {
	return &BiTableRecordsResponse{
		Msg:  "",
		Code: 0,
		Data: &BiTableRecordsData{
			Total:   0,
			HasMore: false,
			Items:   []*BiTableRecordsItem{},
		},
	}
}

type BiTableRecordsReq struct {
	AppToken string
	TableID  string
}

func NewBiTableRecordsReq() *BiTableRecordsReq {
	return &BiTableRecordsReq{
		AppToken: "",
		TableID:  "",
	}
}

func (b *BiTableRecordsResponse) String() {
	for _, item := range b.Data.Items {
		for _, field := range item.Fields {
			fmt.Printf("%#v \n", field)
		}
	}
}

func (b *BiTableRecordsResponse) GetRecordIDByName(tag string) string {
	for _, item := range b.Data.Items {
		optionSlice := item.Fields[constants.OptionsKey]
		if optionSlice == nil {
			return ""
		}
		for _, option := range optionSlice.([]interface{}) {
			name := option.(map[string]interface{})[constants.TextKey].(string)
			if name == tag {
				return item.RecordID
			}
		}
	}
	return ""
}
