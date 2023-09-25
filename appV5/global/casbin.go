package global

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"log"
	"os"
)

var (
	Enforcer *casbin.Enforcer
)

// 创建casbin的enforcer
func SetupCasbinEnforcer() error {
	dir, _ := os.Getwd()
	modelPath := dir + "/appV3/config/rbac_model.conf"
	csvPath := dir + "/appV3/config/rbac2.csv"
	fmt.Println("modelPath:" + modelPath)
	fmt.Println("csvPath:" + csvPath)
	var errC error
	Enforcer, errC = casbin.NewEnforcer(modelPath, csvPath)
	//fmt.Printf("RBAC test start\n") // output for debug
	if errC != nil {
		//fmt.Println(errC)
		log.Fatalf("SetupCasbinEnforcer err: %v", errC)
		return errC
	} else {
		Enforcer.EnableLog(true)
		return nil
	}
}
