package snowflake

import (
	"errors"
	"github.com/sony/sonyflake"
	"go.uber.org/zap"
	"time"
)

var (
	//sonyFlake是一个指向sonyflake.Sonyflake类型的指针变量。sonyflake.Sonyflake是一个用于生成唯一ID的库，
	//通常用于分布式系统中。通过使用这个库，你可以生成在整个分布式系统中具有全局唯一性的ID。
	snowyFlake *sonyflake.Sonyflake
	//sonyMachineID是一个无符号整数（uint16）变量，它用于标识当前机器或节点的唯一ID。
	snowMachineID uint16
)

// 返回机械码
func getMachineID() (uint16, error) {
	return snowMachineID, nil
}

func Init(machineId uint16) (err error) {
	snowMachineID = machineId
	t, _ := time.Parse("2006-01-02", "2023-07-05")
	settings := sonyflake.Settings{
		StartTime: t,
		MachineID: getMachineID,
	}
	snowyFlake = sonyflake.NewSonyflake(settings)
	if snowyFlake == nil {
		return errors.New("init snowflake failed")
	}
	return nil
}

//生成用户id
func GetID() (id uint64, err error) {
	id, err = snowyFlake.NextID()
	if err != nil {
		zap.L().Error("生成用户id十遍")
		return 0000, err
	}
	return
}