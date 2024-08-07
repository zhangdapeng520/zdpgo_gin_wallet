package main

import (
	gin "github.com/zhangdapeng520/zdpgo_gin"
	ginWallet "github.com/zhangdapeng520/zdpgo_gin_wallet"
	gorm "github.com/zhangdapeng520/zdpgo_gorm"
	_ "github.com/zhangdapeng520/zdpgo_mysql"
)

func main() {
	db, err := gorm.Open(
		"mysql",
		"root:root@tcp(127.0.0.1:3306)/blog?charset=utf8",
	)
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&ginWallet.GinWalletAccount{})
	db.AutoMigrate(&ginWallet.GinWalletAccountRecord{})
	defer db.Close()
	r := gin.Default()

	r.POST("/wallet/account/", ginWallet.GetAccountAddHandler(db))

	r.Run(":8888")
}
