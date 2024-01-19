package main

import (
	// "bufio"
	"fmt"
	"os/exec"
	"strconv"
	"sync"
)

func main() {
	var prefix string
	fmt.Printf("Enter pinged ip prefix: ")
	fmt.Scan(&prefix)

	wg := new(sync.WaitGroup)
	iter_count := 1000
	host_count := 256
	for i := 0; i < host_count; i += 1 {
		ip := prefix + strconv.Itoa(i)

		wg.Add(1)
		go func() {
			for j := 0; j < iter_count; j += 1 {
				wg.Add(1)
				go func(j int) {
					fmt.Printf("Pinging %s, iteration: %d\n", ip, j)

					out, err := exec.Command("ping", "-w", "15", ip).Output()
					if err != nil {
						fmt.Printf("Pinging %s at iteration %d failed\n", ip, j)
					}
	                fmt.Printf("%s\n", out)
					// reader := bufio.NewReader(pipe)
					// var line string
					// line, err = reader.ReadString('\n')
					// for err == nil {
					// 	fmt.Println(line)
					// 	line, err = reader.ReadString('\n')
					// }
					wg.Done()
				}(j)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
