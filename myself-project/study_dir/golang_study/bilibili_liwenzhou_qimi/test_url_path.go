package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

var skipLogPath = []string{"/health", "/swagger/index.html", "/swagger/swagger-ui.css",
	"/swagger/swagger-ui-standalone-preset.js", "/swagger/swagger-ui-bundle.js", "/swagger/swagger-ui.css.map",
	"/swagger/doc.json", "/swagger/swagger-ui-standalone-preset.js.map", "/swagger/swagger-ui-bundle.js.map",
	"/swagger/favicon-32x32.png", "/swagger/favicon-16x16.png"}

func main() {
	AccessLog(skipLogPath...)
}

func AccessLog(notlogged ...string) {
	var skip map[string]struct{}

	if length := len(notlogged); length > 0 {
		skip = make(map[string]struct{}, length)

		for _, path := range notlogged {
			skip[path] = struct{}{} // skip["/health"]=struct{}{}  skip["/swagger/index.html"]=struct{}{}
		}
	}
	logrus.Infof("skip path info: %v", skip)
	// map[/health:{} /swagger/doc.json:{} /swagger/favicon-16x16.png:{} /swagger/favicon-32x32.png:{} /swagger/index.html:{} /swagger/swagger-u
	//i-bundle.js:{} /swagger/swagger-ui-bundle.js.map:{} /swagger/swagger-ui-standalone-preset.js:{} /swagger/swagger-ui-standalone-preset.js.map:{} /swagger/swagger-ui.css:{} /swagger/swagger-ui.css.map:{}]

	ac2(skip)
}

func ac2(skip map[string]struct{}) {
	// Start timer
	start := time.Now()
	//path := c.Request.URL.Path    // 获取请求的 URL 路径
	//raw := c.Request.URL.RawQuery // 获取请求的参数
	path := "/health"
	raw := "age=18"

	logrus.Infof("Rock: access_log.go line:25 , raw: [%v], path:[%v]", raw, path)
	// time="2020-09-11T16:11:19+08:00" level=info msg="Rock: access_log.go line:25 , raw: [], path:[/health]"
	// # curl http://10.151.3.84:32500/swagger/index.html?a=1
	// time="2020-09-11T16:15:13+08:00" level=info msg="Rock: access_log.go line:25 , raw: [a=1], path:[/swagger/index.html]"

	// Process request
	//c.Next()
	logrus.Infof("skip: %v", skip)
	logrus.Infof("skip[/health]: %v", skip["/health"])

	// Log only when path is not being skipped
	_, ok := skip[path]
	logrus.Infof("ok:%v", ok)
	if ok {
		//if _, ok := skip[path]; !ok {
		// Stop timer
		end := time.Now()
		latency := end.Sub(start)
		fmt.Println(latency)

		//clientIP := c.ClientIP()
		//method := c.Request.Method
		//statusCode := c.Writer.Status()

		clientIP := "10.20.30.40"
		method := "Get"
		statusCode := 666

		if raw != "" {
			path = path + "?" + raw
		}
		// INFO[2019-03-28 17:16:47] ip: 127.0.0.1       latency: 4.9982ms   code: 400   method: POST     path: /v1/apps
		logrus.Infof("ip: %-15s latency: %-10v code: %-5d method: %-8s path: %s \n", clientIP, latency, statusCode, method, path)
	}
}
