package main

import "fmt"

func main() {
	// http server
	insecureAddr := ""
	insecurePort := ""

	//httpServer := ""
	serveHttp := insecureAddr != "" && insecurePort != ""
	fmt.Println(serveHttp)
}
