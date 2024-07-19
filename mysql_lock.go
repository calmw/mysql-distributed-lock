package main

import (
	"fmt"
	"log"
	"mysql-distributed-lock/config"
	"mysql-distributed-lock/db"
	"mysql-distributed-lock/model"
	"sync"
	"time"
)

func main() {
	config.InitConfig()
	db.InitTDB()
	_ = db.Mysql.AutoMigrate(&model.BridgeOrder{})

	log.Println(FailedTask())
}

func FailedTask() error {
	orders, err := FindFailedOrder()
	if err != nil {
		return err
	}
	wg := sync.WaitGroup{}
	for _, order := range orders {

		wg.Add(1)
		o := order
		go func() {
			fmt.Println(o.Status)
			time.Sleep(time.Second * 3)
			wg.Done()
		}()
	}
	wg.Wait()
	return nil
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
	return db.Mysql.Exec("unlock tables bridge_orders").Debug().Error
}
