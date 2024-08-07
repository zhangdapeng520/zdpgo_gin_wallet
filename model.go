package zdpgo_gin_wallet

// GinWalletAccount 钱包账户表
type GinWalletAccount struct {
	Id       int     `json:"id"`
	Username string  `json:"username" gorm:"column:username;unique;"`
	Money    float64 `json:"money"`
	AddTime  int     `json:"add_time" gorm:"type:int"`
}

// GinWalletAccountRecord 钱包账户变动记录表
type GinWalletAccountRecord struct {
	Id          int     `json:"id"`
	Username    string  `json:"username"`
	AddMoney    float64 `json:"add_money" gorm:"type:decimal"`
	OldMoney    float64 `json:"old_money" gorm:"type:decimal"`
	TotalMoney  float64 `json:"total_money" gorm:"type:decimal"`
	AddTime     int     `json:"add_time" gorm:"type:int"`
	Category    string  `json:"category" gorm:"type:varchar(2);default:'支出'"` // 类型：支出/收入
	Description string  `json:"description"`
}
