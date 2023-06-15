package model

// import (
// 	_ "xorm.io/xorm"
// )

type WkData struct {
	Id    int    `xorm:"'ID'"`
	Data1 string `xorm:"'DATA_1'"`
	Data2 string `xorm:"'DATA_2'"`
	Data3 string `xorm:"'DATA_3'"`
	Data4 string `xorm:"'DATA_4'"`
	Data5 string `xorm:"'DATA_5'"`
	Data6 string `xorm:"'DATA_6'"`
	Data7 string `xorm:"'DATA_7'"`
}

func (m *WkData) TableName() string {
	return `"WK_DATA"`
}
