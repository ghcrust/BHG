package main

import (
	"fmt"
	"net"
	"sort"
	//"sync"
)

var url = "scanme.nmap.org"

//var wg sync.WaitGroup

func main() {
	ports := make(chan int, 20)
	results := make(chan int)
	var openports []int

	for worker := 0; worker < cap(ports); worker++ {
		go ScannerWorker(ports, results)
	}
	go func() {
		for port := 0; port < 1024; port++ {
			//wg.Add(1)
			ports <- port
		}
	}()
	for i := 0; i < 1024; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}
	close(ports)
	close(results)

	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("port %d is open\n", port)
	}
	//wg.Wait()

}

func ScannerWorker(ports, results chan int) {
	for port := range ports {
		address := fmt.Sprintf("%s:%d", url, port)
		if conn, err := net.Dial("tcp", address); err == nil {
			results <- port
			conn.Close()
			continue
		}
		results <- 0
		//wg.Done()
	}
}
