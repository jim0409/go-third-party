package models

import (
	"strings"

	"github.com/fatih/structs"
	"gorm.io/gorm"
)

type MessageTable struct {
	gorm.Model
	Name       string `gorm:"column:name"`
	Age        string `gorm:"column:age"`
	Type       string `gorm:"column:type;comment:顯示使用的訊息種類;default: text"`
	UserID     string `gorm:"user_id"`
	GroupID    string `gorm:"column:group_id"`
	TargetID   string `gorm:"column:target_id"`
	Text       string `gorm:"column:text"`
	Template   string `gorm:"template"`
	SourceText string `gorm:"source_txt;comment:如果有取代字，可以用來替換"`
	OriginalID string `gorm:"original_id;comment:用來表示該則留言是回覆之前哪一則留言用"`
	FileIDs    string `gorm:"file_ids"`

	TableName string `gorm:"-"`
}

/*
1. 創建訊息群組
2. 增加群組人員訊息
3. 管制DB查詢位置
*/

type GroupMessage interface {
	createGroupMsgTabel(string) error
	InsertRecrods(string, string, string) error
	QueryTable(string, []string, int, int) ([]map[string]interface{}, error)
}

func (o *Operation) createGroupMsgTabel(tbname string) error {
	return o.DB.Table(tbname).AutoMigrate(&MessageTable{TableName: tbname})
}

func (o *Operation) InsertRecrods(tbname, name string, age string) error {
	return o.DB.Table(tbname).Create(&MessageTable{Name: name, Age: age}).Error
}

// QueryRecord( group string, filter []string})
func (o *Operation) QueryTable(tbname string, filter []string, limit int, offset int) ([]map[string]interface{}, error) {
	var records []MessageTable
	var filters = "*"

	if filter != nil {
		filters = strings.Join(filter, ",")
	}

	if err := o.DB.Table(tbname).Select(filters).Offset(offset).Limit(limit).Scan(&records).Error; err != nil {
		return nil, err
	}

	res := make([]map[string]interface{}, len(records))
	for i, v := range records {
		res[i] = structs.Map(v)
	}

	return res, nil
}
