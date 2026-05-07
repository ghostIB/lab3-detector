package stats

var GlobalStats = make(map[string]int)

func IncrementProcessed(imageType string) {
	GlobalStats[imageType]++
}
