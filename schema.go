package zdpgo_gin_wallet

type RequestAccount struct {
	Username string  `json:"username" binding:"required,gte=3,lte=36"`
	Money    float64 `json:"money" binding:"gte=0,lte=10000000"`
}
