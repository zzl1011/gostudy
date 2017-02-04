/**
 * author:zhaozhilu
 * goroutine
 *
 */
package main

import "time"

func main() {
	go func() {
		for {
			time.Sleep(time.Second)
		}
	}()

	select {} //block here for ever
}
