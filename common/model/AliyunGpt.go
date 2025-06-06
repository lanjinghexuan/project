package model

import "time"

type Aliyungpt struct {
	Id        int32     `gorm:"column:id;type:int;primaryKey;not null;" json:"id"`
	TaskNo    string    `gorm:"column:task_no;type:varchar(255);comment:任务编号;default:NULL;" json:"task_no"` // 任务编号
	Content   string    `gorm:"column:content;type:varchar(255);default:NULL;" json:"content"`
	Status    int32     `gorm:"column:status;type:int;comment:0未生成 1生成中 2 已完成 3失败;default:0;" json:"status"` // 0未生成 1生成中 2 已完成 3失败
	Request   string    `gorm:"column:request;type:varchar(255);default:NULL;" json:"request"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;default:NULL;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;default:NULL;" json:"updated_at"`
	Errors    string    `gorm:"column:errors;type:varchar(255);default:NULL;" json:"errors"`
}
