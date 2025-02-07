package practice

// ### Example #1: Concurrent File Download
//
// Implement a function that, given an array of URLs and an existing download function,
// downloads all the data from the urls in parallel, merges the results into a single dictionary of {url:data}
// and then returns the dictionary.
//
// Result: they are separate JSON payloads w/o overlapping keys. i.e., the download result of each URL is a number
// and we want to return the total sum at the end.

type ActionableUrl func() error

func GetUrl() {
	mu.Unlock()
	defer mu.Lock()
	// act on the url here and get the json result
}
