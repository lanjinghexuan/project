package model

type VideoUser struct {
	Id            int32   `gorm:"column:id;type:int;primaryKey;not null;" json:"id"`
	Name          string  `gorm:"column:name;type:varchar(20);comment:名称;not null;" json:"name"`                          // 名称
	NickName      string  `gorm:"column:nick_name;type:varchar(20);comment:昵称;not null;" json:"nick_name"`                // 昵称
	UserCode      string  `gorm:"column:user_code;type:varchar(20);comment:编号;not null;" json:"user_code"`                // 编号
	Signature     string  `gorm:"column:signature;type:varchar(50);comment:签名;default:NULL;" json:"signature"`            // 签名
	Sex           string  `gorm:"column:sex;type:varchar(10);comment:性别;not null;" json:"sex"`                            // 性别
	IpAddress     string  `gorm:"column:ip_address;type:varchar(50);comment:IP地址;default:NULL;" json:"ip_address"`        // IP地址
	Constellation string  `gorm:"column:constellation;type:varchar(20);comment:星座;default:NULL;" json:"constellation"`    // 星座
	AttendCount   float32 `gorm:"column:attend_count;type:float(11, 2);comment:关注数;default:0.00;" json:"attend_count"`    // 关注数
	FansCount     float32 `gorm:"column:fans_count;type:float(11, 2);comment:粉丝数;default:1.00;" json:"fans_count"`        // 粉丝数
	ZanCount      float32 `gorm:"column:zan_count;type:float(11, 2);comment:点赞数;default:0.00;" json:"zan_count"`          // 点赞数
	Status        string  `gorm:"column:status;type:varchar(20);comment:用户状态;default:1;" json:"status"`                   // 用户状态
	AvatorFileId  int32   `gorm:"column:avator_file_id;type:int;comment:头像关联id;default:NULL;" json:"avator_file_id"`      // 头像关联id
	AuthriryInfo  string  `gorm:"column:authriry_info;type:varchar(50);comment:认证信息;default:NULL;" json:"authriry_info"`  // 认证信息
	Mobile        string  `gorm:"column:mobile;type:varchar(11);comment:手机号;default:NULL;" json:"mobile"`                 // 手机号
	RealNameAuth  string  `gorm:"column:real_name_auth;type:varchar(20);comment:实名认证状态;default:2;" json:"real_name_auth"` // 实名认证状态
	Age           int32   `gorm:"column:age;type:int;comment:年龄;default:NULL;" json:"age"`                                // 年龄
	OnlineStatus  string  `gorm:"column:online_status;type:varchar(20);comment:在线状态;default:NULL;" json:"online_status"`  // 在线状态
	AuthrityType  string  `gorm:"column:authrity_type;type:varchar(20);comment:认证类型;default:NULL;" json:"authrity_type"`  // 认证类型
	Level         int32   `gorm:"column:level;type:int;comment:用户等级;default:NULL;" json:"level"`                          // 用户等级
	Balance       int32   `gorm:"column:balance;type:int;comment:余额;default:NULL;" json:"balance"`                        // 余额
}
