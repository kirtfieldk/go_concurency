package main
import(
	"fmt"
	"sync"
	"runtime"
)
// These method prints out whats up three timesssss
func printGreetings(){
	var wg sync.WaitGroup
	for _, salutations := range []string{"Hello", "Hreetings", "Whats up"}{
		wg.Add(1)
		go func(){
			defer wg.Done()
			fmt.Println(salutations)
		}()
	}
	wg.Wait()
}
// This prints out all three strings
func printGreetingsFixed(){
	var wg sync.WaitGroup
	for _, salutations := range []string{"Hello", "Hreetings", "Whats up"}{
		wg.Add(1)
		// For the anonomyse func give it a perameter 
		go func(str string){
			defer wg.Done()
			fmt.Println(str)
		}(salutations)
	}
	wg.Wait()
}
// Checking size of goroutines
func sizeOfGoroutines(){
	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}
	var c <- chan interface{}
	var wg sync.WaitGroup
	noop := func(){wg.Done(); <-c}
	const numGoroutines = 1e4
	wg.Add(numGoroutines)
	before := memConsumed()
	for i := numGoroutines; i > 0; i-- {
		go noop()
	}
	wg.Wait()
	after := memConsumed()
	fmt.Printf("Mem space for go routines %.3fkb", float64(after-before)/numGoroutines/1000)
}
func main(){
	printGreetings()
	printGreetingsFixed()
	sizeOfGoroutines()
}