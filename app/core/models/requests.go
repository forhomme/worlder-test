package models

type RequestSensorParam struct {
	ID1      string `query:"id1"`
	ID2      string `query:"id2"`
	Duration string `query:"duration"`
	TimeFrom string `query:"time_from"`
	TimeTo   string `query:"time_to"`
	FilterPagination
}
