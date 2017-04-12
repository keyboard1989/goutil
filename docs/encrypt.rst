Encrypt
=======

Md5
------

.. note:: Md5(str string) string

.. sourcecode:: go

    //Usage
    str := "abc"
    strMd5 := goutil.Md5(str)
    
    //resust:900150983cd24fb0d6963f7d28e17f72


Base64Encode
--------------

.. note:: Base64Encode(str string) string

.. sourcecode:: go

    //Usage
    str := abc
    strBase64 := goutil.Base64Encode(str)

    //result:YWJj


Base64Decode
-------------

.. note:: Base64Decode(str string) ([]byte, error)

.. sourcecode:: go

    //Usage
    str := YWJj
    strBase64Decode, err := gotuil(str)
    if err != nil {
        fmt.Println(strBase64Decode)
    }

    //result:: abc