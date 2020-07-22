package main

import (
	"fmt"
	"github.com/golang/glog"
	"gopkg.in/yaml.v2"
	//"sigs.k8s.io/yaml"  // 这个不要用，不合适
)

func regexConfig(option string, config string, replicasNum int) {
	tmpMap := map[string]interface{}{}
	err := yaml.Unmarshal([]byte(config), tmpMap)
	if err != nil {
		glog.Fatal(err)
	}

	for k, v := range tmpMap {
		switch v2 := v.(type) {
		case  map[interface {}]interface {}:
			fmt.Println(k,v2["replicas"])
			if k == option{
				v2["replicas"] = replicasNum
				//tmpMap["crowd_oversea_worker"] = v2
			}
		case []interface {}:
			fmt.Println("[]interface {}+++++++>", v2)
		default:
			fmt.Printf("i dont know~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~,  %T\n", v2)
		}

	}

	fmt.Println(tmpMap)
	//var exp = `(?s)(?P<X>` + option + `:.*?replicas: )(?P<Y>.*?)\s`
	//optionRepli := regexp.MustCompile(exp)
	//optionStr := optionRepli.FindAllStringSubmatch(config, -1)
	//fmt.Println("no.1 test:", optionStr)
	//
	//if optionStr != nil {
	//	repl := regexp.MustCompile(`(replicas:)\s+(\d)`)
	//	replStr := repl.FindAllStringSubmatch(optionStr[0][0], -1)
	//	fmt.Println("no.2 test replStr: ", replStr) //  [[replicas: 2 replicas: 2]]
	//	replCount := regexp.MustCompile(`\d`)
	//	if replStr != nil {
	//		replconfig := replCount.ReplaceAllLiteralString(replStr[0][0], strconv.Itoa(replicasNum))
	//		fmt.Println("no.3 test replconfig:", replconfig)
	//		optionconfig := repl.ReplaceAllLiteralString(optionStr[0][0], replconfig)
	//		fmt.Println("no.4 test optionconfig:", optionconfig)
	//		optionConfigNew := optionRepli.ReplaceAllLiteralString(config, optionconfig)
	//		fmt.Println("no.5 test optionConfigNew:", optionConfigNew)
	//	}
	//	//optionlastStr := opeplis.ReplaceAllLiteralString(optionStr[0][0], strconv.Itoa(replicasNum))
	//} else {
	//	fmt.Println("Else : error ")
	//}
}

func main() {
	// 原来配置
	//config := `crowd_oversea_worker:\n  env:\n  - name: NVIDIA_VISIBLE_DEVICES\n    value: \"0\"\n  replicas: 1\n  repository: registry.sensenebula.io:5000/nebula-test/engine-video-crowd-oversea-process-service\nextraArgs:\n  customSelect: true\n  decodeInterval: 1\n  detectModelPath: models/kestrel_detect/KM_Hunter_SmallFace_Gray_4.10.29.model\n  featureVersion: 24901\nheadshoulder_worker:\n  env:\n  - name: NVIDIA_VISIBLE_DEVICES\n    value: \"3\"\n  replicas: 1\n  repository: registry.sensenebula.io:5000/nebula-test/engine-video-process-service\nhostAliases:\n- hostnames:\n  - private.ca.sensetime.com\n  ip: 10.9.244.167\n- hostnames:\n  - slave.private.ca.sensetime.com\n  ip: 127.0.0.1\nimage:\n  repository: registry.sensenebula.io:5000/nebula-test/engine-video-process-service\nkafka:\n  brokers:\n  - kafka-default-0.kafka-default.component.svc.cluster.local:9092\nmanager:\n  replicas: 1\n  repository: registry.sensenebula.io:5000/nebula-test/engine-video-crowd-oversea-process-service\npach_worker:\n  env:\n  - name: NVIDIA_VISIBLE_DEVICES\n    value: \"1\"\n  replicas: 1\n  repository: registry.sensenebula.io:5000/nebula-test/engine-video-process-service\nworker:\n  env:\n  - name: NVIDIA_VISIBLE_DEVICES\n    value: \"2\"\n  replicas: 1\n  repository: registry.sensenebula.io:5000/nebula-test/engine-video-process-service\n  service:\n    verbose: true\nzookeeper:\n  endpoints:\n  - zookeeper-default-0.component.svc.cluster.local:2181\n`

	// 这里改成不同的数字，方便查看方便
	config := "crowd_oversea_worker:\n  env:\n  - name: NVIDIA_VISIBLE_DEVICES\n    value: \"0\"\n  replicas: 5\n  repository: registry.sensenebula.io:5000/nebula-test/engine-video-crowd-oversea-process-service\nextraArgs:\n  customSelect: true\n  decodeInterval: 1\n  detectModelPath: models/kestrel_detect/KM_Hunter_SmallFace_Gray_4.10.29.model\n  featureVersion: 24901\nheadshoulder_worker:\n  env:\n  - name: NVIDIA_VISIBLE_DEVICES\n    value: \"3\"\n  replicas: 6\n  repository: registry.sensenebula.io:5000/nebula-test/engine-video-process-service\nhostAliases:\n- hostnames:\n  - private.ca.sensetime.com\n  ip: 10.9.244.167\n- hostnames:\n  - slave.private.ca.sensetime.com\n  ip: 127.0.0.1\nimage:\n  repository: registry.sensenebula.io:5000/nebula-test/engine-video-process-service\nkafka:\n  brokers:\n  - kafka-default-0.kafka-default.component.svc.cluster.local:9092\nmanager:\n  replicas: 7\n  repository: registry.sensenebula.io:5000/nebula-test/engine-video-crowd-oversea-process-service\npach_worker:\n  env:\n  - name: NVIDIA_VISIBLE_DEVICES\n    value: \"1\"\n  replicas: 8\n  repository: registry.sensenebula.io:5000/nebula-test/engine-video-process-service\nworker:\n  env:\n  - name: NVIDIA_VISIBLE_DEVICES\n    value: \"2\"\n  replicas: 9\n  repository: registry.sensenebula.io:5000/nebula-test/engine-video-process-service\n  service:\n    verbose: true\nzookeeper:\n  endpoints:\n  - zookeeper-default-0.component.svc.cluster.local:2181\n"
	replicasNum := 2
	option := "worker"
	//optionHeadShoulder := "headshoulder_worker"
	//optionCrowdOverseaWorker := "crowd_oversea_worker"
	//optionpachworker := "pach_worker"

	// headshoulder_worker 1
	// crowd_oversea_worker 1
	// pach_worker 1
	// worker 2 (engine-video-face-process-service-nodes)

	regexConfig(option, config, replicasNum)
}
