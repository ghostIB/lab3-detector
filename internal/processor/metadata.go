package processor

import (
	"fmt"
	"regexp"
	"sync"
	"time"
)

const maxCacheSize = 500

var (
	LeakCache = make(map[string][]byte)
	mu        sync.Mutex
	re        = regexp.MustCompile(`^image_data_\d+_timestamp_\d+$`)
)

func RunWorkerPool(count int) {
	for i := 0; i < count; i++ {
		go func(id int) {
			for {
				processImage(id)
				time.Sleep(10 * time.Millisecond)
			}
		}(i)
	}
	select {} // block forever
}

func processImage(workerID int) {
	data := fmt.Sprintf("image_data_%d_timestamp_%d", workerID, time.Now().UnixNano())
	matched := re.MatchString(data)

	if matched {
		key := fmt.Sprintf("key_%d", time.Now().UnixNano())

		mu.Lock()
		LeakCache[key] = make([]byte, 1024*10)
		mu.Unlock()
	}
}
