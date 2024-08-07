package zdpgo_gin_wallet

import (
	gin "github.com/zhangdapeng520/zdpgo_gin"
	gorm "github.com/zhangdapeng520/zdpgo_gorm"
	"net/http"
	"time"
)

func GetAccountUpdateHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 提取信息
		var req RequestAccount
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 查询
		var account GinWalletAccount
		db.Find(&GinWalletAccount{}, "username = ?", req.Username).First(&account)
		if account.Id == 0 {
			c.JSON(404, gin.H{"error": "account not found"})
			return
		}

		// 开启事务
		tx := db.Begin()

		// 修改余额
		oldMoney := account.Money
		addMoney := req.Money
		if req.Category == "支出" {
			addMoney = -addMoney
		}
		totalMoney := oldMoney + addMoney
		account.Money = totalMoney
		err := tx.Save(&account).Error
		if err != nil {
			tx.Rollback()
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		// 添加记录
		err = tx.Create(&GinWalletAccountRecord{
			Username:    req.Username,
			AddMoney:    addMoney,
			OldMoney:    oldMoney,
			TotalMoney:  account.Money,
			AddTime:     int(time.Now().Unix()),
			Category:    req.Category,
			Description: req.Description,
		}).Error
		if err != nil {
			tx.Rollback()
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		// 提交事务
		tx.Commit()

		// 返回
		c.JSON(http.StatusOK, nil)
	}
}
