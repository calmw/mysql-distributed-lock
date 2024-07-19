package main

import (
	"fmt"
	"log"
	"mysql-distributed-lock/config"
	"mysql-distributed-lock/db"
	"mysql-distributed-lock/model"
	"mysql-distributed-lock/utils"
	"time"
)

func main() {
	config.InitConfig()
	db.InitTDB()
	//_ = db.Mysql.AutoMigrate(&model.BridgeOrder{})
	//for i := 0; i < 6; i++ {
	//	InsertOrder()
	//}
	OrderWriteLock()
	log.Println("锁表成功")
	log.Println(FailedTask())
	OrderWriteUnLock()
	log.Println("锁释放成功")
}

func FailedTask() error {
	log.Println("耗时任务开始")
	orders, err := FindFailedOrder()
	if err != nil {
		return err
	}
	for _, order := range orders {
		o := order
		fmt.Println(o.Hash)
		time.Sleep(time.Second * 5)
	}
	log.Println("耗时任务结束")
	return nil
}

func InsertOrder() {
	db.Mysql.Model(model.BridgeOrder{}).Create(&model.BridgeOrder{
		Data:       nil,
		Hash:       utils.UniqueId(),
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

func OrderWriteLock() {
	log.Println("锁表...")
	err := db.Mysql.Exec("lock tables bridge_orders write").Debug().Error
	if err != nil {
		log.Println(err)
	}
}

func OrderWriteUnLock() {
	log.Println("释放表锁...")
	err := db.Mysql.Exec("unlock tables").Debug().Error
	if err != nil {
		log.Println(err)
	}
}
