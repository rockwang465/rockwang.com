package main

import (
	"fmt"
)

type nodeConsts struct {
	TFDF int32 //engine-timespace-face-replica-factor
}

func main() {
	//token := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjIsImV4cCI6MTU5NTY2Nzg1NSwiaWF0IjoxNTk1NjY2MDU1LCJpc3MiOiJyb2NrIiwic3ViIjoidXNlciB0b2tlbiJ9.2mJNWgPmletpwK6BBIH0D-4F_vKXmhRAKgqoI7bMEq4"
	//fmt.Printf("[%v]", token[7:])

	Consts := map[string]interface{}{
		"access-control-process-service-nosense-tasks":            718,
		"access-control-process-service-tasks":                    718,
		"engine-alert-feature-db-dbs":                             101,
		"engine-alert-feature-db-proxy-node-tps":                  160,
		"engine-alert-feature-db-worker-node-capacity":            10,
		"engine-image-face-process-service-node-tps":              120,
		"engine-image-ingress-service-node-devices":               718,
		"engine-image-ingress-service-node-tps":                   160,
		"engine-image-pach-process-service-node-tps":              0,
		"engine-timespace-face-feature-db-proxy-node-tps":         5,
		"engine-timespace-face-feature-db-worker-node-capacity":   40,
		"engine-timespace-face-replica-factor":                    2,
		"engine-timespace-ped-feature-db-proxy-node-tps":          0,
		"engine-timespace-ped-feature-db-worker-node-capacity":    0,
		"engine-timespace-ped-replica-factor":                     0,
		"engine-video-action-process-service-node-streams":        0,
		"engine-video-bfsu-process-service-node-streams":          0,
		"engine-video-crowd-oversea-process-service-node-streams": 0,
		"engine-video-face-process-service-node-streams":          0,
		"engine-video-headshoulder-process-service-node-streams":  0,
		"engine-video-pach-process-service-node-streams":          0,
		"iva_feature_encryption_key":                              "796D097DF8616C25",
		"tailing-detection-comparison-service-tasks":              0,
	}

	// 获取license Consts engine-timespace-face-replica-factor 的值
	nodeConst := &nodeConsts{
		TFDF: 0,
	}

	var quotas int32 = 10

	for k, v := range Consts {
		//glog.Infof("nodeInfo.Consts: %#v", Consts)
		fmt.Println(k, v)
		fmt.Printf("%T,%T\n", k, v)
		switch k {
		case "engine-timespace-face-replica-factor":
			if value, ok := v.(int); ok {  // 实际我们的sophon中获取的是float64类型的，不是int类型
				nodeConst.TFDF = int32(value)
				fmt.Println("-------------------------------------------------------------->")
				fmt.Printf("nodeConst.TFDF: %T, %#v", nodeConst.TFDF, nodeConst.TFDF)
			}
		}
	}

	fmt.Println(quotas / nodeConst.TFDF)

}
