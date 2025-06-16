package model

import "github.com/lanjinghexuan/project/common/gload"

type VideoGoodsClassifiy struct {
	Id            int32  `gorm:"column:id;type:int;primaryKey;not null;" json:"id"`
	ClassifiyName string `gorm:"column:classifiy_name;type:varchar(50);comment:分类名称;default:NULL;" json:"classifiy_name"` // 分类名称
	Pid           int32  `gorm:"column:pid;type:int;comment:父级ID;default:0;" json:"pid"`                                  // 父级ID
	Sort          int32  `gorm:"column:sort;type:int;comment:排序;default:NULL;" json:"sort"`                               // 排序
}

func (c VideoGoodsClassifiy) GetClass(pid int32) (list []*VideoGoodsClassifiy, err error) {
	err = gload.DB.Table("video_goods_classifiy").Where("pid=?", pid).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}
