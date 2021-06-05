# 介紹
unixsocket 是一個限定單台上的特殊 socket 傳輸協議，傳輸上較 tcp socket 快速。

使用`UNIX Domain Socket`的過程和網路`socket`十分相似，也要先呼叫`socket()`建立一個FD，

`"address family"`指定為`"AF_UNIX"`，

`"type"`可以選擇`"SOCK_DGRAM"`或`"SOCK_STREAM"`，

`"protocol"`引數仍然指定為`0`


# 實現
`UNIX Domain Socket`與網路`socket`程式設計最明顯的不同在於地址格式不同，用struct `sockaddr_un`表示，

網路程式設計中的`socket`地址是`IP`地址加埠號。

而`UNIX Domain Socket`的地址是一個`socket`型別的檔案在檔案系統中的路徑，

該socket檔案由`bind()`呼叫建立;(如果`bind()`時檔案已存在，則返回錯誤)。



# refer:
- https://www.itread01.com/content/1544790071.html

