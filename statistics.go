package goutil

import (
	"container/ring"
	"sync"
	"time"

	"github.com/jinzhu/now"
)

type Item struct {
	Type    string
	SubType string
	Key     string
}

type Info struct {
	Time time.Time
	Data map[string]uint64
}
type Statistics struct {
	data       map[string]map[string]*ring.Ring
	chanItem   chan Item
	mutex      *sync.RWMutex
	dataLength int
	timeFormat string
	timeFunc   func() time.Time
}

func NewStatistics(itemLength int, dataLength int, timeFormat string, timeFunc func() time.Time) *Statistics {
	chanItem := make(chan Item, itemLength)

	staitsitcs := &Statistics{
		data:       map[string]map[string]*ring.Ring{},
		chanItem:   chanItem,
		mutex:      &sync.RWMutex{},
		dataLength: dataLength,
		timeFormat: timeFormat,
		timeFunc:   timeFunc,
	}

	go staitsitcs.consume()

	return staitsitcs
}

func (s *Statistics) AddItem(item Item) {
	s.chanItem <- item
}

func (s *Statistics) consume() {
	for item := range s.chanItem {
		typeName := item.Type
		subType := item.SubType
		key := item.Key
		s.mutex.Lock()
		ringStruct := s.getRing(typeName, subType)
		if ringStruct.Value == nil {
			ringStruct.Value = Info{
				Time: s.timeFunc(),
				Data: map[string]uint64{},
			}
		}

		if ringStruct.Value.(Info).Time.Equal(s.timeFunc()) {
			ringStruct.Value.(Info).Data[key]++
			s.mutex.Unlock()
			continue
		}

		ringStruct = ringStruct.Prev()
		s.data[typeName][subType] = ringStruct
		ringStruct.Value = Info{
			Time: s.timeFunc(),
			Data: map[string]uint64{key: 1},
		}
		s.mutex.Unlock()
	}
}

func (s *Statistics) getRing(typeName string, subType string) *ring.Ring {
	_, ok := s.data[typeName]
	if !ok {
		s.data[typeName] = make(map[string]*ring.Ring)
	}

	_, ok = s.data[typeName][subType]
	if !ok {
		s.data[typeName][subType] = ring.New(s.dataLength)
	}

	return s.data[typeName][subType]
}

func (s *Statistics) GetInfoByType(typeStr string) interface{} {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	_, ok := s.data[typeStr]
	if !ok {
		return nil
	}

	ret := make(map[string][]interface{})

	for subType, _ := range s.data[typeStr] {
		temp := []interface{}{}

		ringStruct, _ := s.data[typeStr][subType]
		if ringStruct == nil {
			continue
		}

		ringStruct.Do(func(v interface{}) {
			if v == nil {
				return
			}
			tempInfo := map[string]interface{}{}
			tempInfo["time"] = v.(Info).Time.Format("2006-01-02 15:04:05")
			for k, v := range v.(Info).Data {
				tempInfo[k] = v
			}

			temp = append(temp, tempInfo)
		})

		ret[subType] = temp
	}
	return ret
}

func (s *Statistics) GetInfoByTypeAndSubType(typeStr string, subType string) interface{} {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	_, ok := s.data[typeStr]
	if !ok {
		return nil
	}

	_, ok = s.data[typeStr][subType]
	if !ok {
		return nil
	}

	ret := []interface{}{}

	ringStruct, _ := s.data[typeStr][subType]
	if ringStruct == nil {
		return nil
	}

	ringStruct.Do(func(v interface{}) {
		if v == nil {
			return
		}
		tempInfo := map[string]interface{}{}
		tempInfo["time"] = v.(Info).Time.Format("2006-01-02 15:04:05")
		for k, v := range v.(Info).Data {
			tempInfo[k] = v
		}

		ret = append(ret, tempInfo)
	})

	return ret
}

func getTime() time.Time {
	return now.BeginningOfMinute()
}
