package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"net/url"
)

type K8sConfig struct {
	Clusters       []Clusters  `yaml:"clusters"`
	//Contexts       []Contexts  `json:"contexts"`
	//CurrentContext string      `json:"current-context"`
	//Kind           string      `json:"kind"`
	//Preferences    Preferences `json:"preferences"`
	//Users          []Users     `json:"users"`
}
type Cluster struct {
	CertificateAuthorityData string `yaml:"certificate-authority-data"`
	Server                   string `yaml:"server"`
}
type Clusters struct {
	Cluster Cluster `yaml:"cluster"`
	Name    string  `yaml:"name"`
}
type Context struct {
	Cluster string `yaml:"cluster"`
	User    string `yaml:"user"`
}
//type Contexts struct {
//	Context Context `json:"context"`
//	Name    string  `json:"name"`
//}
//type Preferences struct {
//}
//
//type User struct {
//	ClientCertificateData string `json:"client-certificate-data"`
//	ClientKeyData         string `json:"client-key-data"`
//}
//type Users struct {
//	Name string `json:"name"`
//	User User   `json:"user"`
//}

func main() {
	adminConf := "apiVersion: v1\nclusters:\n- cluster:\n    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUN5akNDQWJLZ0F3SUJBZ0lCQURBTkJna3Foa2lHOXcwQkFRc0ZBREFWTVJNd0VRWURWUVFERXdwcmRXSmwKY201bGRHVnpNQ0FYRFRJeE1ETXdPREE1TlRBeE5sb1lEekl4TWpFd01qRXlNRGsxTURFMldqQVZNUk13RVFZRApWUVFERXdwcmRXSmxjbTVsZEdWek1JSUJJakFOQmdrcWhraUc5dzBCQVFFRkFBT0NBUThBTUlJQkNnS0NBUUVBCjZTdmNubytSN1hHMlYrMDg2MW52MW13U05HSEpob0FJOTNyRzk1Z2xzeGROQzQ3dHc2UHJsRHVJNllhcDEwZ0UKanN1MWQ0MFpnbzFZc2pmWjVCRTVENFhDY2dHeWx0THlhOERVNzh4VG1JZlRxbGpWUGFVNGNDdFpnaHUvenlLNQpWbFNRMXRDdjhyS3k2aVZUb0g2QXB4OWJHQ3pITFJBNW1hOWhpMGNhZjR0M2tqc0Y4U3d4Y1ZMeHZDT1RpcWNuCmQ1cVhzYmhmOUJqZ2RJWFkyM1g0KzA1VlVBNW5ZR1BEdTJ3bTlYenVhcllOaTl0VG53R3ZwWUVNaWRPamZyb2wKMDhiS1pQcmduTHdGeHJOVk5lM2VuWXRSTlQ5cDZHSHZUdW16QlVDcVd0aE10MWhJd0ZBZjhxM25pdkhYeGYxWQo4WWpWdlgybTNjRDB6c1ZTbDBwS3p3SURBUUFCb3lNd0lUQU9CZ05WSFE4QkFmOEVCQU1DQXFRd0R3WURWUjBUCkFRSC9CQVV3QXdFQi96QU5CZ2txaGtpRzl3MEJBUXNGQUFPQ0FRRUFGNHhxbXN2M1BWTU90SEgxbnM3QVNmc3QKaWdhbHUyMm00WG1iYVhjV294QTlVQTYwY1NSOVRySnRXY1dTcTRwYk9vRVJwZE5rZXRhNnRydnRGMFNaN2ZCaApCVFlKc0l4dklOb2xNdlgwZnBtK0U2a1NJbDZNZzlyVTdaenNwbnlTQXJpTk04WkhSY2ZFbDBhNGlqSkRjQi8zCkxEUDZmelNRSWRLcDVYRGJlTHFNTk14eHpkUEdGUVdRdWdNVjQyWXFjdVZBcmFZVVB3LzZBVEhJaWg1YlZVQzIKRjhmWmZtQ3IybzNUR05rTFd0TEJob1lFTG43MlArRVVRMkNEeFpnS2dVcVAzNVJlYkZJZkV0SlZ5Y1k0aTVpRQpWZUwzaVMwMEFVRVdCUWJHNEhHNW0yNytNS0t5MktBcnlBR3hpdjllLzkwSk8vbEZNSStOOW1qMUh3Vzl4UT09Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K\n    server: https://127.0.0.1:16443\n  name: nebula-cluster\ncontexts:\n- context:\n    cluster: nebula-cluster\n    user: kubernetes-admin\n  name: kubernetes-admin@nebula-cluster\ncurrent-context: kubernetes-admin@nebula-cluster\nkind: Config\npreferences: {}\nusers:\n- name: kubernetes-admin\n  user:\n    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUM5RENDQWR5Z0F3SUJBZ0lJSUZRZTN5RGduM1V3RFFZSktvWklodmNOQVFFTEJRQXdGVEVUTUJFR0ExVUUKQXhNS2EzVmlaWEp1WlhSbGN6QWdGdzB5TVRBek1EZ3dPVFV3TVRaYUdBOHlNVEl4TURJeE1qQTVOVEF4TjFvdwpOREVYTUJVR0ExVUVDaE1PYzNsemRHVnRPbTFoYzNSbGNuTXhHVEFYQmdOVkJBTVRFR3QxWW1WeWJtVjBaWE10CllXUnRhVzR3Z2dFaU1BMEdDU3FHU0liM0RRRUJBUVVBQTRJQkR3QXdnZ0VLQW9JQkFRQzJwemwxbzg2WlVubHEKM2oyK1YyWjdXMzNpd3JPWnpMNVBYRmRqbVY1UWR0TS9Rbk5RYU5icy8yK05RVWl2azV6bXNjUWdkMmIyMDNWQwplai9OZCtsRUVUQnptQUl5czZhS3lDRTdIYXEzNVArZGQ0SDh0WG1KSUJXbGlOQWNUSnRBWTZVRm5NcE14ckM3CmRzdHpSVUo2aE5WRDZxZ3FtdFhxS0FYSUJYczU0WjhXcjExSkNWdGk1Wlk1dklKdVZtZ3p2NDBxZ29IQzQ1MDEKaCtOTnNiSmpvWHh0WkNFK2JlUWNpVEhoaHl4c3A5OHNKQU9kUFJpSkhhemc1Yk9QOXl2cDl0V1hEejJrU3oweQozRzhNQzh1VmhJSHRvK0FnM2V3QUtJbUN5TlFMWEtvQmlYVW5ncW5FMFMxVkhndGFPZ0F0Y3ZKMWpUUDJYNC9hCm5VWEt4N29IQWdNQkFBR2pKekFsTUE0R0ExVWREd0VCL3dRRUF3SUZvREFUQmdOVkhTVUVEREFLQmdnckJnRUYKQlFjREFqQU5CZ2txaGtpRzl3MEJBUXNGQUFPQ0FRRUFkWFh0TnhPZU9SRlMzWWtERHppcFA2cElJZHpzbTIrbwoyT2IxQ0F3QXBvbnRVSW52YTVQTHRmaisxK0VHREZzTzZVOFJNWnhudkFxVzVyWlFyd2hqdlpFWlM4MWJNRlVFCml5RHAzaXBHNWM3ZEd0S2RlN1dqbnFaL3FzUm1HTVBWNTRVbG1FOENxQlgvYStNam9PaE9HTlB0QmRUcm8zZDQKNXkrdktsUW1kaHZWU01jcXl0akxkdmRGb3VMMEV4ZDdvSzJQaFpEbWh1elRab0pqOXI2M1lybFhLQVhoUTRmWgpQZXNEYWdKRUt2elloa2N5QXM5ZSs2dkdKLzVUempVcDVBRDhSRVdFN1ZSKzNReFBXZHcvcHA5US9wQSsvdnZtCjgybUVlOWlkTmczT2F0N3NKakJ0djc2OE81cGwzN2d5TTlBT0I2WmNTSGtMSXVCZmlPRkdmQT09Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K\n    client-key-data: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcFFJQkFBS0NBUUVBdHFjNWRhUE9tVko1YXQ0OXZsZG1lMXQ5NHNLem1jeStUMXhYWTVsZVVIYlRQMEp6ClVHalc3UDl2alVGSXI1T2M1ckhFSUhkbTl0TjFRbm8velhmcFJCRXdjNWdDTXJPbWlzZ2hPeDJxdCtUL25YZUIKL0xWNWlTQVZwWWpRSEV5YlFHT2xCWnpLVE1hd3UzYkxjMFZDZW9UVlErcW9LcHJWNmlnRnlBVjdPZUdmRnE5ZApTUWxiWXVXV09ieUNibFpvTTcrTktvS0J3dU9kTllmalRiR3lZNkY4YldRaFBtM2tISWt4NFljc2JLZmZMQ1FECm5UMFlpUjJzNE9XemovY3I2ZmJWbHc4OXBFczlNdHh2REF2TGxZU0I3YVBnSU4zc0FDaUpnc2pVQzF5cUFZbDEKSjRLcHhORXRWUjRMV2pvQUxYTHlkWTB6OWwrUDJwMUZ5c2U2QndJREFRQUJBb0lCQUNtYmhGSlJBMUlYNElMQQpwNDRCZU4zbTVKRlFTdnRoRVlVZ2d1TzBYekN6TU1IN1VDdGtCbm1zWW12cUVEVzJ5WW9mdW82dXhHeTdsTDJVCi9ZMi9vQXFhWTlBaE0zYjdSNWZicVA0L2J0RS9RQjlOeHZXYkhWbG9UcVlMdVdTSTZ1RExFaDlxTFFjNlI2NGMKakpId21Ba3kweVlLZ1U1YlFSbDdEbjRYQ1pBUWU5bGRnYlA4UU5GcmN3RG4wS29WVUlwK0lCcHFjNy9GcW96QQo1V2FTeDBCMndmZDJMSkRqVFFnZTV2Um9ka0Q5RXF1MzV6YnR5eEFYL29TOWEvT0tTRXUrR2ZMWFpLU0JHVkxFCkNFN052dkYyb0w5YTgzYkgzNnhTS3dMQnoyZGFVN01qQ1dTb2JHT0l2dENPMkdaY3A3bnloV1dlZVgyT3RRSHkKQ3daTTRFa0NnWUVBMk00SS91U0RVUEhGNUpJUjgrOEVxWE1YU1BCcWZVTDE3NWdLYldvZ2VabkNWdkVRSUpnaQp3MjV0U21kZzVMYzNocEhXckRuMzN3dHMyVXhWbFdXaFhiVWwvRm40LzZTeHoyOWhpMVJoN01Vak9NR0dGdnVkCjNvR3FPaWtKRDFOTE5za3Z1Z0w4QkJtbjdrVnN3bjlTWmV1TjB4WlV4WHlKSnhxQ0kwMWdUZFVDZ1lFQTE2eWQKbXVHODRnZXh3SCszQTJtN0JhbURjSjBHL0tqSTJ6US9CY2tXZzlMRmJ6ZXpPVmJaanhaNHl2QVZ6bzAwVmFQRwpaLzFndXRiS05rRG1VMFhhZXlpdmUvU1VlZW9yUWNPN0xPeWgzZDRGR0hBLzl1YUk1L1JFZHZJN0R3QXBpcmpLCjFEdzExQjl3TXJXc0Z1ckQ5OFYzVWE1cFRRM2pkMVQvOEZPWmFtc0NnWUVBMFhHNW53U0RGNTk1M2VaL0VYakcKdUN4SFFsOS9nS00vazhiOXk5WEtxYTZ3Myt6aDk1c0JlbXNqaWMxMC9YOUNWTzV1eE5MN2NxNUc2V1dsQ0Q2MgpOU0RiSVg2SjNlM3RHTWd4STdqUklUc1JXN01aSnZyTStEZXJpNlE0N0dVN05DSUh5VnM0dTMxTmpoSGpqOElRCmlBV1hiV1VWWC9OL1RCZC93VHJ1M2tFQ2dZRUFnVHVYdk1UakQzV1kxMFl2L2NXUElWZHZ3VHoyZ05NS2QvOFkKZkhrQUdzRmpPTDloajB3eFZRSWNJMDQxWXUxTm5MdGtHYy9EOUhEYi9pSnBZU0NNU0J3YSt1V1ZTbkE4dDAxMwpqUFhHRUZJSEw1dWpXR2pQUkY0ajREcEpsVFY5cnZnSDRhQ3FpSTdHQmp6Qi80RnhKN2lWUk1hRHBuNVovSmxjCk96OTh3RDBDZ1lFQXVuSkswaXNDZER0RlV6NCtLSFZ5SGVvbXNDY1VkYlo5RXpGdGl3eC9lbDMxNzdremptb28KdHRBbjQ4dWY2QjA3bGdjWEdOTUt1d0NlblBEbW05Wm9CNC9XeGovem9UR1J6SzNzWFN2MTR5UDVZOW9oN0E2Nwo0d2xieExFWlp1Q3hjMGlaRkVzNktQMWtwVWl1SGVaem1FZkpjN3RrWmtPTWhBNDVGMW9uWis4PQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo=\n"
	clusterAddr, err := GetClusterHostFromConfig(adminConf)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("clusterAddr: ", clusterAddr)
	}
}

func GetClusterHostFromConfig(config string) (string, error) {
	var k8sConfig K8sConfig
	if err := yaml.Unmarshal([]byte(config), &k8sConfig); err != nil {
		return "", err
	}

	var serverAddr string
	if len(k8sConfig.Clusters) >= 1{
		cluster := k8sConfig.Clusters[0]
		serverAddr = cluster.Cluster.Server
	}

	var Url *url.URL
	var err error
	if serverAddr != "" {
		Url, err = url.Parse(serverAddr)
		if err != nil {
			return "", err
		}
	}
	fmt.Printf("%#v\n", Url)
	fmt.Printf("%#v\n",Url.Hostname())
	fmt.Printf("%T\n",Url.Hostname())
	return Url.Hostname(), nil

}
