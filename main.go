package main

import "time"

func main() {
	applyThemeology()

	applyTicker := time.NewTicker(5 * time.Minute)

	func() {
		for {
			select {
			case <-applyTicker.C:
				applyThemeology()
			}
		}
	}()
}
