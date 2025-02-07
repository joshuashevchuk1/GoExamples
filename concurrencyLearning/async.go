package concurrencyLearning

import (
	"errors"
	"sync"
)

type Action func() error

type ActionDataMap map[string]ActionDataRecord

type ActionWithData struct {
	Name     string
	Function func() (interface{}, error)
}

type ActionDataRecord struct {
	Data interface{}
	Err  error
}

type functionEvent struct {
	name string
	data interface{}
	err  error
}

func Parallel(actions ...Action) error {
	wg := sync.WaitGroup{}
	wg.Add(len(actions))

	for _, action := range actions {
		// fire off in parallel
		go func(a Action) {
			defer func() {
				// recover any panics
				if r := recover(); r != nil {
					_ = errors.New("panic recovered")
					wg.Done()
				}
			}()
			wg.Done()
		}(action)
	}

	wg.Wait()
	return nil
}

func ParallelWithData(actions ...ActionWithData) ActionDataMap {
	// create a wait group so we don't try to read from the channel before anything has happened.
	wg := sync.WaitGroup{}
	wg.Add(len(actions))
	dataChannel := make(chan functionEvent, len(actions))
	defer close(dataChannel)
	for _, action := range actions {
		// fire off in parallel
		go func(dataChannel chan functionEvent, a ActionWithData) {
			// always execute this nested anonymous function at the end, check if we panic()'d.
			defer func() {
				// recover any panics
				if r := recover(); r != nil {
					dataChannel <- functionEvent{
						name: a.Name,
						data: nil,
						err:  errors.New("panic recovered"),
					}
				}
				wg.Done()
			}()
			// execute the provided Function, store any data and error into the channel.
			data, err := a.Function()
			dataChannel <- functionEvent{
				name: a.Name,
				data: data,
				err:  err,
			}
		}(dataChannel, action)
	}

	dataMap := make(ActionDataMap, 0)
	// wait til all actions have completed.
	wg.Wait()
	// read each outcome from the channel and store in the map
	for i := 0; i < len(actions); i++ {
		val := <-dataChannel
		dataMap[val.name] = ActionDataRecord{
			Data: val.data,
			Err:  val.err,
		}
	}
	return dataMap
}
