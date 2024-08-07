package zdpgo_gin_wallet

import (
	gin "github.com/zhangdapeng520/zdpgo_gin"
	gorm "github.com/zhangdapeng520/zdpgo_gorm"
	"net/http"
	"time"
)

func GetAccountAddHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req RequestAccount
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db.Create(&GinWalletAccount{
			Username: req.Username,
			Money:    req.Money,
			AddTime:  int(time.Now().Unix()),
		})

		c.JSON(http.StatusOK, nil)
	}
}
