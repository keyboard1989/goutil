Time
======

GetCurrentTimeStr
-------------------

.. note:: GetCurrentTimeStr() string

.. sourcecode:: go

    //Usage
    t := goutil.GetCurrentTimeStr()
    fmt.Println(t)

    //res: 2017-01-01 12:12:12


GetDayRange
--------------

.. note:: GetDayRange(start string, end string, format string) ([]string, error)

.. sourcecode:: go

    //Usage
    start := "2017-01-01"
    end := "2017-01-31"
    format := "2006-01-02"

    dayRange , err := goutil(start,end,format)
    fmt.Println(dayRange)

    //result ["2017-01-01","2017-01-02",...,"2017-01-29","2017-01-30"]