# goutil

[![Build Status](https://api.travis-ci.org/keyboard1989/goutil.svg?branch=master)](https://travis-ci.org/keyboard1989/goutil)

## document site

[http://goutil.readthedocs.io/](http://goutil.readthedocs.io/)

## Description

goutil is just a collection of go common used method

## serval import struct or method

### Statistics

Statistics is used to stat realtime data in a web server or a backgroud service  
It contain a main type and a sub type, and  the third is the key string you want to calculate  

eg: Type:API  SubType: User Key:get  

```
statis := NewStatistics(10000, 10, "2006-01-02 15:04:05", MockGetTime)
item := Item{Type:"API", SubType:"User", Key:"get"}
statis.AddItem(item)

//Get a data value
data := statis.GetInfoByTypeAndSubType("API", "User")
// data format
// map["time"] = "2017-01-01 12:12:00"
// map["get] = 123
```

### GetDayRange

GetDayRange is used to get a range of start and end time, see the comment for usage   

