package server

import (
	"context"
	"fmt"
	"github.com/golang/glog"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"

	"gitlab.sz.sensetime.com/galaxias/infra-sophon-service/api"
	"gitlab.sz.sensetime.com/galaxias/infra-sophon-service/pkg/license"
)

type nodeQuotas struct {
	VPS  map[string]int32
	AFD  int32
	TFD  int32
	AC   int32
	TD   int32
	STFD int32
	SDB  int32
	IIS  int32
}

type nebulaConfig struct {
	VPS  map[string]string
	STFD map[string]string
	TFD  map[string]string
	IPS  map[string]string
	DelSvc []string
}

type nvidiaCard struct {
	CardTotal  int
	RestCardNo int
	UsedCardNo []int
}

func (s *SophonServer) ExecLinuxCommand(strCmd string)(string, error){
	cmd := exec.Command("/bin/bash", "-c", strCmd)

	stdout, _ := cmd.StdoutPipe()
	if err := cmd.Start(); err != nil{
		//glog.Fatal("Execute failed when Start:" + err.Error())
		return "", err
	}

	outBytes, _ := ioutil.ReadAll(stdout)
	stdout.Close()

	if err := cmd.Wait(); err != nil {
		//fmt.Println("Execute failed when Wait:" + err.Error())
		return "", err
	}
	return string(outBytes), nil
}

func (s *SophonServer) SaveDeployInfo(NC map[string]string, name string, configVlaue string) error {
	instance, err := s.helmClient.GetInstance(name, 0)
	if err != nil {
		return err
	}
	namespace := instance.Namespace
	chart := fmt.Sprintf("%s/charts/%s-%s.tgz", s.config.Chartmuseum, instance.AppName, instance.Version)
	NC = map[string]string{"name":name, "config":configVlaue, "namespace":namespace, "chart":chart}
	fmt.Println()
	return nil
}

func (s *SophonServer) ReDeployInstance(name string, configVlaue string) error {
	instance, err := s.helmClient.GetInstance(name, 0)
	namespace := instance.Namespace
	chart := fmt.Sprintf("%s/charts/%s-%s.tgz", s.config.Chartmuseum, instance.AppName, instance.Version)
	//name, namespace, chart, values, req.RecreatePods
	if err != nil {
		return err
	}
	//s.helmClient.DeployInstance(name, namespace, chart, configVlaue, true)
	if err := s.helmClient.DeployInstance(name, namespace, chart, configVlaue, true); err != nil {
		return err
	}
	return nil

}

func (s *SophonServer) GetInstanceConfig(name string) (string, string, error) {
	instance, err := s.helmClient.GetInstance(name, 0)
	if err != nil {
		fmt.Println("get instance config err !!!!!!!!!!!!!------------------")
		return "", "", err
	}
	return instance.Name, instance.Config, nil
}

func ConfigMatch(option string, config string, replicasNum int) string {
	cfg := map[string]interface{}{}
	err := yaml.Unmarshal([]byte(config), cfg)
	if err != nil {
		glog.Fatal(err)
		return ""
	}

	for k, v := range cfg {
		switch v2 := v.(type) {
		case map[interface {}]interface {}:
			//fmt.Println(k,v2["replicas"])
			if k == option{
				v2["replicas"] = replicasNum
				//tmpMap["crowd_oversea_worker"] = v2
			}
		}
	}

	res ,err := yaml.Marshal(cfg)
	if err != nil {
		glog.Fatal(err)
		return ""
	}
	return string(res)

	////var exp string = option + `\:\n\s+replicas\:\s+\d`  (?P<id>\d)abc(?P=id)  (?P<id>abc)
	//var exp string = `(?s)(?P<X>` + option + `:.*?replicas: )(?P<Y>.*?)\s`
	//optionRepli := regexp.MustCompile(exp)
	//optionStr := optionRepli.FindAllStringSubmatch(config, -1)
	////opeplis := regexp.MustCompile(`\d`)
	//if optionStr != nil {
	//	repl := regexp.MustCompile(`(replicas:)\s+(\d)`)
	//	replStr := repl.FindAllStringSubmatch(optionStr[0][0], -1)
	//	replCount := regexp.MustCompile(`\d`)
	//	if replStr != nil {
	//		replconfig := replCount.ReplaceAllLiteralString(replStr[0][0], strconv.Itoa(replicasNum))
	//		optionconfig := repl.ReplaceAllLiteralString(optionStr[0][0], replconfig)
	//		optionConfigNew := optionRepli.ReplaceAllLiteralString(config, optionconfig)
	//		fmt.Printf("glog Rock: ---------------------> return [%v] optionConfigNew-------> :%v",option, optionConfigNew)
	//		return optionConfigNew
	//	}
	//	//optionlastStr := opeplis.ReplaceAllLiteralString(optionStr[0][0], strconv.Itoa(replicasNum))
	//} else {
	//	fmt.Println("error error error error error error error error error error error error error error error error error error error error ==========1")
	//	return ""
	//}
	//fmt.Println("error error error error error error error error error error error error error error error error error error error error ==========2")
	//return ""
}

func (s *SophonServer) AnalysisQuotas(req license.ServerType) (*nodeQuotas, error) {
	ctl, err := license.NewServiceCtl()
	nodeInfo, err := ctl.GetCAStatus(req)
	//fmt.Println("nodeInfo:", nodeInfo)
	// &{false true 100 0 1595214598 alone a3aa3f58-3bd1-4b04-bb93-8513fac42151 IVA-VIPER 1335129808 99991231 sensetime_SC [22000] map[access-control-process-service-nodes:{1 2} engine-alert-feature-db-proxy-nodes:{1 2} engine-alert-feature-db-worker-nodes:{1 2} engine-image-face-process-service-nodes:{0 2} engine-image-ingress-service-nodes:{1 2} engine-image-pach-process-service-nodes:{0 2} engine-timespace-face-feature-db-proxy-nodes:{1 2} engine-timespace-face-feature-db-worker-nodes:{0 2} engine-timespace-ped-feature-db-proxy-nodes:{1 2} engine-timespace-ped-feature-db-worker-nodes:{0 2} engine-video-crowd-oversea-process-service-nodes:{1 2} engine-video-face-process-service-nodes:{1 3} engine-video-headshoulder-process-service-nodes:{0 2} engine-video-pach-process-service-nodes:{1 2} tailing-detection-comparison-service-nodes:{1 2}] map[access-control-process-service-nosense-tasks:112 access-control-process-service-tasks:96 engine-alert-feature-db-dbs:101 engine-alert-feature-db-proxy-node-tps:1 engine-alert-feature-db-worker-node-capacity:2 engine-image-face-process-service-node-tps:160 engine-image-ingress-service-node-devices:96 engine-image-ingress-service-node-tps:160 engine-image-pach-process-service-node-tps:20 engine-struct-db-service-capacity:10 engine-struct-db-service-node-tps:5 engine-timespace-face-feature-db-proxy-node-tps:5 engine-timespace-face-feature-db-worker-node-capacity:40 engine-timespace-face-replica-factor:1 engine-timespace-ped-feature-db-proxy-node-tps:5 engine-timespace-ped-feature-db-worker-node-capacity:20 engine-timespace-ped-replica-factor:1 engine-video-crowd-oversea-process-service-node-streams:10 engine-video-face-process-service-node-streams:16 engine-video-headshoulder-process-service-node-streams:10 engine-video-pach-process-service-node-streams:10 iva_feature_encryption_key:DAC32A0B9E87C2D9 tailing-detection-comparison-service-tasks:128]}
	//fmt.Printf("nodeInfo:||%#v||\n\n", nodeInfo)
	// nodeInfo:||&license.CAStatus{Disable:false, IsActive:true, ActiveLimit:100, AloneTime:0, DongleTime:1595214814, Status:"alone", AuthID:"a3aa3f58-3bd1-4b04-bb93-8513fac42151", Product:"IVA-VIPER", DongleID:"1335129808", ExpiredAt:"99991231", Company:"sensetime_SC", FeatureIds:[]uint64{0x55f0}, Quotas:map[string]license.quotaLimit{"access-control-process-service-nodes":license.quotaLimit{Used:1, Total:2}, "engine-alert-feature-db-proxy-nodes":license.quotaLimit{Used:1, Total:2}, "engine-alert-feature-db-worker-nodes":license.quotaLimit{Used:1, Total:2}, "engine-image-face-process-service-nodes":license.quotaLimit{Used:0, Total:2}, "engine-image-ingress-service-nodes":license.quotaLimit{Used:1, Total:2}, "engine-image-pach-process-service-nodes":license.quotaLimit{Used:0, Total:2}, "engine-timespace-face-feature-db-proxy-nodes":license.quotaLimit{Used:1, Total:2}, "engine-timespace-face-feature-db-worker-nodes":license.quotaLimit{Used:0, Total:2}, "engine-timespace-ped-feature-db-proxy-nodes":license.quotaLimit{Used:1, Total:2}, "engine-timespace-ped-feature-db-worker-nodes":license.quotaLimit{Used:0, Total:2}, "engine-video-crowd-oversea-process-service-nodes":license.quotaLimit{Used:1, Total:2}, "engine-video-face-process-service-nodes":license.quotaLimit{Used:1, Total:3}, "engine-video-headshoulder-process-service-nodes":license.quotaLimit{Used:0, Total:2}, "engine-video-pach-process-service-nodes":license.quotaLimit{Used:1, Total:2}, "tailing-detection-comparison-service-nodes":license.quotaLimit{Used:1, Total:2}}, Consts:map[string]interface {}{"access-control-process-service-nosense-tasks":112, "access-control-process-service-tasks":96, "engine-alert-feature-db-dbs":101, "engine-alert-feature-db-proxy-node-tps":1, "engine-alert-feature-db-worker-node-capacity":2, "engine-image-face-process-service-node-tps":160, "engine-image-ingress-service-node-devices":96, "engine-image-ingress-service-node-tps":160, "engine-image-pach-process-service-node-tps":20, "engine-struct-db-service-capacity":10, "engine-struct-db-service-node-tps":5, "engine-timespace-face-feature-db-proxy-node-tps":5, "engine-timespace-face-feature-db-worker-node-capacity":40, "engine-timespace-face-replica-factor":1, "engine-timespace-ped-feature-db-proxy-node-tps":5, "engine-timespace-ped-feature-db-worker-node-capacity":20, "engine-timespace-ped-replica-factor":1, "engine-video-crowd-oversea-process-service-node-streams":10, "engine-video-face-process-service-node-streams":16, "engine-video-headshoulder-process-service-node-streams":10, "engine-video-pach-process-service-node-streams":10, "iva_feature_encryption_key":"DAC32A0B9E87C2D9", "tailing-detection-comparison-service-tasks":128}}||
	fmt.Printf("nodeInfo.Quotas:||%#v||\n\n", nodeInfo.Quotas)
	// nodeInfo.Quotas: map[string]license.quotaLimit{"access-control-process-service-nodes":license.quotaLimit{Used:1, Total:2}, "engine-alert-feature-db-proxy-nodes":license.quotaLimit{Used:1, Total:2}, "engine-alert-feature-db-worker-nodes":license.quotaLimit{Used:1, Total:2}, "engine-image-face-process-service-nodes":license.quotaLimit{Used:0, Total:2}, "engine-image-ingress-service-nodes":license.quotaLimit{Used:1, Total:2}, "engine-image-pach-process-service-nodes":license.quotaLimit{Used:0, Total:2}, "engine-timespace-face-feature-db-proxy-nodes":license.quotaLimit{Used:1, Total:2}, "engine-timespace-face-feature-db-worker-nodes":license.quotaLimit{Used:0, Total:2}, "engine-timespace-ped-feature-db-proxy-nodes":license.quotaLimit{Used:1, Total:2}, "engine-timespace-ped-feature-db-worker-nodes":license.quotaLimit{Used:0, Total:2}, "engine-video-crowd-oversea-process-service-nodes":license.quotaLimit{Used:1, Total:2}, "engine-video-face-process-service-nodes":license.quotaLimit{Used:1, Total:3}, "engine-video-headshoulder-process-service-nodes":license.quotaLimit{Used:0, Total:2}, "engine-video-pach-process-service-nodes":license.quotaLimit{Used:1, Total:2}, "tailing-detection-comparison-service-nodes":license.quotaLimit{Used:1, Total:2}}

	//time.Sleep(100 * time.Second)
	if err != nil {
		return nil, err
	}
	node := &nodeQuotas{
		AFD:  0,
		AC:   0,
		TD:   0,
		STFD: 0,
		TFD:  0,
		SDB: 0,
		IIS: 0,
	}
	//node.VPS = make(map[string]int32)
	//node.VPS = map[string]int32{"VPC": 0, "VPW": 0}  -- 原代码
	//                                              crowd      headshoulder
	node.VPS = map[string]int32{"VPC": 0, "VPW": 0, "VPCO":0, "VPHS":0}  // rock修改
	//var node = map[string]int32{"VPS":0,"VIS":0, "IPS":0, "IIS":0, "AFD":0, "TFD":0, "AC":0, "TD":0, "VPC":0 }
	//node := make(map[string]int32)
	for k, v := range nodeInfo.Quotas {
		switch k {
		// "engine-video-pach-process-service-nodes":license.quotaLimit{Used:1, Total:2}
		case "engine-video-pach-process-service-nodes":
			if v.Total != 0 {
				node.VPS["VPC"] = v.Total - 1 // Total大于1，则需要减1，这是版本的原因
			}
			if err != nil {
				return nil, err
			}

		case "engine-video-face-process-service-nodes":
			if v.Total != 0 {
				node.VPS["VPW"] = v.Total - 1
			}
			//config, err := s.GetInstanceConfig(string(k))
			if err != nil {
				return nil, err
			}

			// rock增加vps的 crowd headshoulder 部分
		case "engine-video-crowd-oversea-process-service-nodes":
			if v.Total != 0 {
				node.VPS["VPCO"] = v.Total - 1
			}
		case "engine-video-headshoulder-process-service-nodes":
			if v.Total != 0 {
				node.VPS["VPHS"] = v.Total - 1
			}

		case "access-control-process-service-nodes":
			if v.Total != 0 {
				node.AC = v.Total - 1
			}
			//config, err := s.GetInstanceConfig(string(k))
			if err != nil {
				return nil, err
			}

		case "engine-alert-feature-db-worker-nodes":
			if v.Total != 0 {
				node.AFD = v.Total - 1
			}
			//config, err := s.GetInstanceConfig(string(k))
			if err != nil {
				return nil, err
			}

		case "engine-timespace-ped-feature-db-worker-nodes":
			if v.Total != 0 {
				//node.TFD["TFDW"] = v.Total -1 + node.TFD["TFDW"]
				node.STFD = v.Total - 1
			}

		case "engine-timespace-face-feature-db-worker-nodes":
			if v.Total != 0 {
				node.TFD = v.Total - 1 + node.TFD
			}
			if err != nil {
				return nil, err
			}

		case "tailing-detection-comparison-service-nodes":
			if v.Total != 0 {
				node.TD = v.Total - 1 + node.TD
			}
			if err != nil {
				return nil, err
			}
		case "engine-struct-db-service-nodes":
			if v.Total != 0 {
				node.SDB = v.Total -1
			}
		case "engine-image-ingress-service-nodes":
			if v.Total != 0 {
				node.IIS = v.Total -1
			}

		}
	}
	return node, nil
}

func (s *SophonServer) GetLicenseHdInfo(ctx context.Context, req *api.GetLicenseHdInfoRequest) (*api.GetLicenseHdInfoResponse, error) {
	ctl, err := license.NewServiceCtl()
	if err != nil {
		return nil, err
	}
	caResp, err := ctl.HardwareInfo(license.ServerType(req.GetServerType()), 0)
	if err != nil {
		return nil, err
	}
	return &api.GetLicenseHdInfoResponse{
		Fingerprint: caResp.FingerPrint,
		C2V:         caResp.C2V,
	}, nil
}

func (s *SophonServer) ActiveLicense(ctx context.Context, req *api.ActiveLicenseRequest) (*api.ActiveLicenseResponse, error) {
	ctl, err := license.NewServiceCtl()  // 生成一个和license交互的客户端(里面有证书)
	if err != nil {
		return nil, err
	}
	var caResp *license.ActiveResponse
	fmt.Println("Rock: req.GetActiveType():", req.GetActiveType())  // Rock: req.GetActiveType(): ONLINE
	if req.GetActiveType() == api.LicenseActiveType_ONLINE {
		// license.ServerType = uint
		// req.ServerType = int32
		caResp, err = ctl.OnlineActivate(license.ServerType(req.ServerType), "activate")
		glog.Info("glog Rock: ---------------------> ActiveLicense2")  // 这里走到了
	} else {
		glog.Info("glog Rock: ---------------------> ActiveLicense3")  // 没有走到这
		caResp, err = ctl.OfflineActivate(license.ServerType(req.ServerType), req.V2C)
	}
	glog.Info("glog license.ServerType(req.ServerType): ", license.ServerType(req.ServerType))  // 0
	glog.Info("glog license.ServerType(req.GetServerType()): ", license.ServerType(req.GetServerType()))  // 0
	if err != nil {
		glog.Info("glog Rock: err :", err)
		// err :Post "https://private.ca.sensetime.com:8443/online": context deadline exceeded (Client.Timeout exceeded while awaiting headers)
		// 这里超时，说明license部分有问题了。可能是接口变了，也可能是license的状态变了。
		glog.Info("glog Rock: caResp :", caResp)
		return nil, err

	}
	resp := &api.ActiveLicenseResponse{
		Code:    caResp.StatusCode,
		Message: caResp.StatusMessage,
	}
	Quotas, err := s.AnalysisQuotas(license.ServerType(req.GetServerType()))
	fmt.Printf("use func get Quotas: %#v\n\n", Quotas)
	// use func get Quotas: &server.nodeQuotas{VPS:map[string]int32{"VPC":1, "VPW":2}, AFD:1, TFD:1, AC:1, TD:1, STFD:1, SDB:0, IIS:1}
	time.Sleep(5 * time.Second)

	// 初始化一下
	NebulaConfig := &nebulaConfig{}
	//                     01      1      21     3     41    51     6     71      8  // 后一位数是1，表示已经安装过了，sophon打印出来的
	NameList := []string{"VPS", "STFD", "AFD", "TFD", "AC", "TD", "SDB", "IIS", "IPS"}
	for i := 0; i < len(NameList); i++ {
		fmt.Printf("glog Rock: ---------------------> for i value: %#v\n", i)
		//time.Sleep(5 * time.Second)

		switch NameList[i] {
		case "VPS":
			// 获取配置文件
			vpsName, config, err := s.GetInstanceConfig("engine-video-process-service-nebula")
			//fmt.Printf("vps vpsName:||%#v||\n\n", vpsName)
			fmt.Printf("from helm client get config!!!!!!!------->:||%#v||\n\n", config)
			if err == nil {
				configValueW := ConfigMatch("worker", config, int(Quotas.VPS["VPW"]))
				//fmt.Printf("vps configValueW:||%#v||\n\n", configValueW)

				configValueP := ConfigMatch("pach_worker", configValueW, int(Quotas.VPS["VPC"]))
				//fmt.Printf("vps configValueP:||%#v||\n\n", configValueP)

				// crowd_oversea_worker
				configValueCO := ConfigMatch("crowd_oversea_worker", configValueP, int(Quotas.VPS["VPCO"]))
				//fmt.Printf("vps configValueP:||%#v||\n\n", configValueCO)
				// headshoulder_worker
				configValueHS := ConfigMatch("headshoulder_worker", configValueCO, int(Quotas.VPS["VPHS"]))
				fmt.Printf("vps configValueHS:||%#v||\n\n", configValueHS)

				err := s.SaveDeployInfo(NebulaConfig.VPS, vpsName,configValueHS)
				if err != nil {
					glog.Fatal(err)
				}
				//if configValueHS != "" {
				//	s.ReDeployInstance(vpsName, configValueHS)
				//}
				//fmt.Printf("%v time sleep 10 --------------------------------\n", vpsName)
				//time.Sleep(10 * time.Second)
			}
		case "STFD":
			stfdName, config, err := s.GetInstanceConfig("engine-struct-timespace-feature-db-nebula")
			if err == nil {
				configValueW := ConfigMatch("worker", config, int(Quotas.STFD))
				configValueC := ConfigMatch("coordinator", configValueW, int(Quotas.STFD))
				configValueR := ConfigMatch("reducer", configValueC, int(Quotas.STFD))
				err := s.SaveDeployInfo(NebulaConfig.STFD, stfdName,configValueR)
				if err != nil {
					glog.Fatal(err)
				}
				NebulaConfig.DelSvc = append(NebulaConfig.DelSvc, stfdName)
				//if configValueR != "" {
				//	s.ReDeployInstance(stfdName, configValueR)
				//}
				fmt.Printf("STFD configValueR:||%#v||\n\n", configValueR)
				//fmt.Printf("%v time sleep 10 --------------------------------\n", stfdName)
				//time.Sleep(10 * time.Second)
			}
		case "AFD":
			afdName, config, err := s.GetInstanceConfig("engine-static-feature-db-nebula")
			if err == nil {
				configValueW := ConfigMatch("worker", config, int(Quotas.AFD))
				configValueP := ConfigMatch("proxy", configValueW, int(Quotas.AFD))
				if configValueP != "" {
					s.ReDeployInstance(afdName, configValueP)
				}
				fmt.Printf("AFD configValueP:||%#v||\n\n", configValueP)
				//fmt.Printf("%v time sleep 10 --------------------------------\n", afdName)
				//time.Sleep(10 * time.Second)
			}
		case "TFD":
			tfdName, config, err := s.GetInstanceConfig("engine-timespace-feature-db-nebula")
			if err == nil {
				// configValueW := ConfigMatch("worker", config, int(Quotas.TFD))
				configValueW := ConfigMatch("worker", config, int(Quotas.TFD))
				configValueC := ConfigMatch("coordinator", configValueW, int(Quotas.TFD))
				configValueR := ConfigMatch("reducer", configValueC, int(Quotas.TFD))
				err := s.SaveDeployInfo(NebulaConfig.TFD, tfdName,configValueR)
				if err != nil {
					glog.Fatal(err)
				}
				NebulaConfig.DelSvc = append(NebulaConfig.DelSvc, tfdName)
				//if configValueR != "" {
				//	s.ReDeployInstance(tfdName, configValueR)
				//}
				fmt.Printf("TFD configValueR:||%#v||\n\n", configValueR)
				//fmt.Printf("%v time sleep 10 --------------------------------\n", tfdName)
				//time.Sleep(10 * time.Second)
			}
		case "AC":
			acName, config, err := s.GetInstanceConfig("access-control-process-nebula")
			if err == nil {
				configValueM := ConfigMatch("master", config, int(Quotas.AC))
				configValueW := ConfigMatch("worker", configValueM, int(Quotas.AC))
				if configValueW != "" {
					s.ReDeployInstance(acName, configValueW)
				}
				fmt.Printf("AC configValueW:||%#v||\n\n", configValueW)
				//fmt.Printf("%v time sleep 10 --------------------------------\n", acName)
				//time.Sleep(10 * time.Second)
			}
		case "TD":
			//tdName, config, err := s.GetInstanceConfig("engine-image-ingress-service-nebula")  // 之前传入的是iis，错了
			tdName, config, err := s.GetInstanceConfig("tailing-detection-comparison-service-nebula")
			if err == nil {
				configValueM := ConfigMatch("manager", config, int(Quotas.TD))
				configValueW := ConfigMatch("worker", configValueM, int(Quotas.TD))
				if configValueW != "" {
					s.ReDeployInstance(tdName, configValueW)
				}
				fmt.Printf("TD configValueW:||%#v||\n\n", configValueW)
				//fmt.Printf("%v time sleep 10 --------------------------------\n", tdName)
				//time.Sleep(10 * time.Second)
			}
		case "SDB":
			sdbName, config, err := s.GetInstanceConfig("engine-struct-db-nebula")
			if err == nil {
				configValueM := ConfigMatch("proxy", config, int(Quotas.SDB))
				configValueW := ConfigMatch("worker", configValueM, int(Quotas.SDB))
				if configValueW != "" {
					s.ReDeployInstance(sdbName, configValueW)
				}
				fmt.Printf("SDB configValueW:||%#v||\n\n", configValueW)
				//fmt.Printf("%v time sleep 10 --------------------------------\n", sdbName)
				//time.Sleep(10 * time.Second)
			}
		case "IIS":
			iisName, config, err := s.GetInstanceConfig("engine-image-ingress-service-nebula")
			if err == nil {
				configValueM := ConfigMatch("active_worker", config, int(Quotas.IIS))
				configValueW := ConfigMatch("master", configValueM, int(Quotas.IIS))
				if configValueW != "" {
					s.ReDeployInstance(iisName, configValueW)
				}
				fmt.Printf("IIS configValueW:||%#v||\n\n", configValueW)
				//fmt.Printf("%v time sleep 10 --------------------------------\n", iisName)
				//time.Sleep(10 * time.Second)
			}
		case "IPS":
			ipsName, config, err := s.GetInstanceConfig("engine-image-process-service-nebula")
			if err == nil {
				// ips为daemonset格式，所以不需要replicas调整
				err := s.SaveDeployInfo(NebulaConfig.IPS, ipsName,config)
				if err != nil {
					glog.Fatal(err)
				}
				NebulaConfig.DelSvc = append(NebulaConfig.DelSvc, ipsName)
				fmt.Printf("IIS configValueW:||%#v||\n\n", config)
				//fmt.Printf("%v time sleep 10 --------------------------------\n", iisName)
				//time.Sleep(10 * time.Second)
			}

		}
	}

	fmt.Printf("%#v\n\n\n", NebulaConfig)
	// 1.1获取当前机器显卡总数
	NvidiaCardInfo := nvidiaCard{}
	nvidiaTotalCmd := "/usr/bin/nvidia-smi -L | wc -l"
	nvidiaTotalStr, err := s.ExecLinuxCommand(nvidiaTotalCmd)
	if err != nil {
		glog.Fatal(err)
	}else{
		nvidiaTotalStr = strings.Replace(nvidiaTotalStr, "\n", "", -1)
		nvidiaTotalInt, err := strconv.Atoi(nvidiaTotalStr)
		if err != nil {
			glog.Fatal(err)
		}else {
			NvidiaCardInfo.CardTotal = nvidiaTotalInt
			fmt.Println(NvidiaCardInfo)  // {4 0 []}
		}

	}

	// 1.2删除所有stfd tfd ips服务，便于vps自动落卡
	for _,v := range NebulaConfig.DelSvc{
		fmt.Println("开始删除服务:",v)
		err := s.helmClient.DeleteInstance(v, true)
		if err != nil {
			glog.Fatal(err)
		}
	}

	// 1.3使用前面保存的override配置，安装vps服务
	fmt.Println("开始安装vps")
	err = s.helmClient.DeployInstance(NebulaConfig.VPS["name"],NebulaConfig.VPS["namespace"],NebulaConfig.VPS["chart"],NebulaConfig.VPS["config"],true)
	if err != nil {
		glog.Fatal(err)
	}
	// name ns chart version bool

	// 1.4检测vps占卡情况及可用的剩余卡

	// 1.5安装

	// update configmap
	licResp, err := ctl.GetClientLics(license.ServerType(req.ServerType))
	if err != nil || len(licResp.Licenses) == 0 {
		return resp, fmt.Errorf("get client license failed: %v", err)
	}
	if err := s.kubeClient.UpdateConfigmapWithLicense(licResp.Licenses[0]); err != nil {
		return resp, fmt.Errorf("update configmap failed: %v", err)
	}

	// restart pods
	if req.RestartPods {
		if err := s.kubeClient.RestartPodsWithLicense(60); err != nil {
			return resp, fmt.Errorf("restart pods failed: %v", err)
		}
	}

	return resp, nil
}

func (s *SophonServer) GetLicenseStatus(ctx context.Context, req *api.GetLicenseStatusRequest) (*api.GetLicenseStatusResponse, error) {
	ctl, err := license.NewServiceCtl()
	if err != nil {
		return nil, err
	}
	caResp, err := ctl.GetCAStatus(license.ServerType(req.GetServerType()))
	if err != nil {
		// try render dongle id from c2v.
		hdResp, hdErr := ctl.HardwareInfo(license.ServerType(req.GetServerType()), 0)
		if hdErr == nil {
			result := regexp.MustCompile(`<hasp id=["|']([0-9]+)["|']`).FindStringSubmatch(hdResp.C2V)
			if result[1] != "" {
				caResp = &license.CAStatus{
					DongleID: result[1],
				}
			}
		} else {
			return nil, err
		}
	}

	quotas := make(map[string]*api.LicenseQuotaLimit)
	for k, v := range caResp.Quotas {
		quotas[k] = &api.LicenseQuotaLimit{
			Used:  v.Used,
			Total: v.Total,
		}
	}
	consts := make(map[string]string)
	for k, v := range caResp.Consts {
		consts[k] = fmt.Sprint(v)
	}
	status := api.LicenseStatus{
		Disable:     caResp.Disable,
		IsActive:    caResp.IsActive,
		ActiveLimit: caResp.ActiveLimit,
		AloneTime:   caResp.AloneTime,
		DongleTime:  caResp.DongleTime,
		Status:      caResp.Status,
		AuthId:      caResp.AuthID,
		Product:     caResp.Product,
		DongleId:    caResp.DongleID,
		ExpiredAt:   caResp.ExpiredAt,
		Company:     caResp.Company,
		FeatureIds:  caResp.FeatureIds,
		Quotas:      quotas,
		Consts:      consts,
	}
	return &api.GetLicenseStatusResponse{Status: &status}, nil
}
