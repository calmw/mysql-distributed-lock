package main

import (
	"fmt"
	"log"
	"mysql-distributed-lock/config"
	"mysql-distributed-lock/db"
	"mysql-distributed-lock/model"
	"time"
)

func main() {
	config.InitConfig()
	db.InitTDB()
	//_ = db.Mysql.AutoMigrate(&model.BridgeOrder{})
	//InsertOrder()
	log.Println(OrderWriteLock())
	log.Println("锁表成功")
	log.Println(FailedTask())
	log.Println(OrderWriteUnLock())
	log.Println("锁释放成功")
}

func FailedTask() error {
	orders, err := FindFailedOrder()
	if err != nil {
		return err
	}
	for _, order := range orders {
		o := order
		fmt.Println(o.Hash)
		time.Sleep(time.Second * 5)
	}
	return nil
}

func InsertOrder() {
	db.Mysql.Model(model.BridgeOrder{}).Create(&model.BridgeOrder{
		Data:       nil,
		Hash:       "0xaaaaaaa",
		VoteStatus: false,
		Status:     false,
	})

}

func FindFailedOrder() ([]model.BridgeOrder, error) {
	var orders []model.BridgeOrder
	err := db.Mysql.Model(model.BridgeOrder{}).Where("`status`=?", false).Find(&orders).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return orders, nil
}

func OrderWriteLock() error {
	log.Println("锁表")
	return db.Mysql.Exec("lock tables bridge_orders write").Debug().Error
}

func OrderWriteUnLock() error {
	log.Println("释放表锁")
	return db.Mysql.Exec("unlock tables").Debug().Error
}
