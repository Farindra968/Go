// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
	"time"
)

// Post represents a shared resource that multiple goroutines may access.
// The mutex protects the 'view' field from race conditions.
type Post struct {
	view int
	mu   sync.Mutex
}

// inc increases the view count safely.
// We lock the mutex before changing shared data and unlock it after.
func (p *Post) inc() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.view++
}

// readView returns the current number of views safely.
func (p *Post) readView() int {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.view
}

// addViews increases the view by a given number.
func (p *Post) addViews(n int) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.view += n
}

func main() {
	fmt.Println("Example 1: Single goroutine")
	post := Post{view: 0}
	post.inc()
	fmt.Println("View count after one increment:", post.readView())

	fmt.Println("\nExample 2: Multiple goroutines update same value safely")
	var wg sync.WaitGroup
	sharedPost := Post{view: 0}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sharedPost.inc()
		}()
	}

	wg.Wait()
	fmt.Println("Final view count:", sharedPost.readView())

	fmt.Println("\nExample 3: Update with a larger step")
	article := Post{view: 50}
	article.addViews(25)
	fmt.Println("View count after adding 25:", article.readView())

	fmt.Println("\nExample 4: Shared resource with different goroutines and delays")
	var counter Post
	var workerWG sync.WaitGroup

	for i := 0; i < 5; i++ {
		workerWG.Add(1)
		go func(id int) {
			defer workerWG.Done()
			for j := 0; j < 3; j++ {
				time.Sleep(10 * time.Millisecond)
				counter.inc()
				fmt.Printf("Goroutine %d: current value = %d\n", id, counter.readView())
			}
		}(i)
	}

	workerWG.Wait()
	fmt.Println("Final value after all goroutines finished:", counter.readView())

	fmt.Println("\nExample 5: Mutex prevents data races")
	// Without a mutex, multiple goroutines writing to the same variable at the same time
	// can cause a race condition and produce incorrect results.
	// With a mutex, each goroutine runs safely one at a time.
	var safeWG sync.WaitGroup
	result := Post{view: 0}

	for i := 0; i < 10; i++ {
		safeWG.Add(1)
		go func() {
			defer safeWG.Done()
			for j := 0; j < 100; j++ {
				result.inc()
			}
		}()
	}

	safeWG.Wait()
	fmt.Println("Safe total after 10 goroutines * 100 increments:", result.readView())
}
