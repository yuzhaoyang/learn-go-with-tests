package v4

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"go", "java", "python"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(salutation)
		}()
	}
	wg.Wait()
}
