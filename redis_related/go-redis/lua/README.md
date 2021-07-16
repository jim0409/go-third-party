# intro
使用luascript來做redis上的操作

# quick setup env
> docker run --rm --name redis-lab -p 6379:6379 -d redis

# Lua 簡介
Lua 是一個小巧的腳本語言。是巴西里約熱內盧天主教大學（Pontifical Catholic University of Rio de Janeiro）里的一個研究小組，由Roberto Ierusalimschy、Waldemar Celes 和 Luiz Henrique de Figueiredo所組成並於1993年開發。 其設計目的是為了嵌入應用程式中，從而為應用程式提供靈活的擴展和定製功能。Lua由標準C編寫而成，幾乎在所有作業系統和平台上都可以編譯，運行。Lua並沒有提供強大的庫，這是由它的定位決定的。所以Lua不適合作為開發獨立應用程式的語言。Lua 有一個同時進行的JIT項目，提供在特定平台上的即時編譯功能。

Lua腳本可以很容易的被C/C++ 代碼調用，也可以反過來調用C/C++的函數，這使得Lua在應用程式中可以被廣泛應用。不僅僅作為擴展腳本，也可以作為普通的配置文件，代替XML,ini等文件格式，並且更容易理解和維護。 Lua由標準C編寫而成，代碼簡潔優美，幾乎在所有作業系統和平台上都可以編譯，運行。一個完整的Lua解釋器不過200k，在目前所有腳本引擎中，Lua的速度是最快的。這一切都決定了Lua是作為嵌入式腳本的最佳選擇。


# 使用 Lua 腳本的好處
1、減少網絡開銷：可以將多個請求通過腳本的形式一次發送，減少網絡時延和請求次數。

2、原子性的操作：Redis會將整個腳本作為一個整體執行，中間不會被其他命令插入。因此在編寫腳本的過程中無需擔心會出現競態條件，無需使用事務。

3、代碼復用：客戶端發送的腳步會永久存在redis中，這樣，其他客戶端可以復用這一腳本來完成相同的邏輯。

4、速度快：見 與其它語言的性能比較, 還有一個 JIT編譯器可以顯著地提高多數任務的性能; 對於那些仍然對性能不滿意的人, 可以把關鍵部分使用C實現, 然後與其集成, 這樣還可以享受其它方面的好處。

5、可以移植：只要是有ANSI C 編譯器的平台都可以編譯，你可以看到它可以在幾乎所有的平台上運行:從 Windows 到Linux，同樣Mac平台也沒問題, 再到移動平台、遊戲主機，甚至瀏覽器也可以完美使用 (翻譯成JavaScript).

6、源碼小巧：20000行C代碼，可以編譯進182K的可執行文件，加載快，運行快。


# Redis 和 Lua 整合詳解
1. 調用Lua腳本的語法：
```
$ redis-cli --eval path/to/redis.lua KEYS[1] KEYS[2] , ARGV[1] ARGV[2] ...
--eval，告訴redis-cli讀取並運行後面的lua腳本
path/to/redis.lua，是lua腳本的位置
KEYS[1] KEYS[2]，是要操作的鍵，可以指定多個，在lua腳本中通過KEYS[1], KEYS[2]獲取
ARGV[1] ARGV[2]，參數，在lua腳本中通過ARGV[1], ARGV[2]獲取。
```
> 注意： KEYS和ARGV中間的 ',' 兩邊的空格，不能省略。

- redis支持大部分Lua標準庫
```
Base 提供一些基礎函數
String 提供用於字符串操作的函數
Table 提供用於表操作的函數
Math 提供數學計算函數
Debug 提供用於調試的函數
```

2. 在腳本中調用 Redis 命令
在腳本中可以使用redis.call函數調用Redis命令
```
redis.call('set', 'foo', 'bar')
```
```
local value=redis.call('get', 'foo')
return value 
-- value 的值為 bar
```

`redis.call` 函數的返回值就是 Redis 命令的執行結果

Redis命令的返回值有5種類型，redis.call函數會將這5種類型的回覆轉換成對應的Lua的數據類型，具體的對應規則如下（空結果比較特殊，其對應Lua的false）

- redis返回值類型和Lua數據類型轉換規則
```
redis 返回值類型/ Lua 數據類型
整數回復/ 數字類型
字符串回復/ 字符串類型
多行字符串回復/ table類型(數組形式)
狀態回復/ table類型(只有一個 ok 字段儲存狀態信息)
錯誤回復/ table類型(只有一個 err 字段儲存錯誤信息)
```

redis還提供了redis.pcall函數，功能與redis.call相同

唯一的區別是當命令執行出錯時，redis.pcall會記錄錯誤並繼續執行，而redis.call會直接返回錯誤，不會繼續執行

在腳本中可以使用return語句將值返回給客戶端，如果沒有執行return語句則默認返回nil

- Lua 數據類型和 redis 返回值類型轉換規則
```
Lua 數據類型/ redis 返回值類型
數字類型/ 整數回復(Lua的數字類型會被自動轉換成整數)
字符串類型/ 字符串回復
table類型(數組形式)/ 多行字符串回復
table類型(只有一個 ok 字段存儲狀態信息)/ 狀態回復
table類型(只有一個 err 字段存儲錯誤信息)/ 錯誤回復
```

3. 腳本相關命令
EVAL 語法: eval script numkeys key [key ...] arg [arg ...]
通過 key 和 arg 這兩類參數向腳本傳遞數據，他們的值在腳本中分別使用 KEYS 和 ARGV 兩個表類型的全局變量訪問

script: 是 lua 腳本

numkeys: 表示有幾個 key，分別是 KEYS[1], KEYS[2] ... 

```
note: EVAL 命令依據參數 numkeys 來將其後面的所有參數分別存入腳本中 KEYS 和 ARGV 兩個table類型的全局變量
當腳本不需要任何參數時，也不能省略這個參數(設為0)

192.168.127.128:6379>eval "return redis.call('set',KEYS[1],ARGV[1])" 1 name liulei
OK
192.168.127.128:6379>get name "liulei"
```

4. EVALSHA命令

在腳本比較長的情況下，如果每次調用腳本都需要將整個腳本傳給Redis會占用較多的帶寬

為了解決這個問題，Redis提供了EVALSHA命令，允許開發者通過腳本內容的SHA1摘要來執行腳本

該命令的用法和EVAL一樣，只不過是將腳本內容替換成腳本內容的SHA1摘要。

Redis在執行EVAL命令時會計算腳本的SHA1摘要並記錄在腳本緩存中

執行EVALSHA命令時Redis會根據提供的摘要從腳本緩存中查找對應的腳本內容

如果找到了則執行腳本，否則會返回錯誤："NOSCRIPT No matching script. Please use EVAL."

在程序中使用EVALSHA命令的一般流程如下。

1）、先計算腳本的SHA1摘要，並使用EVALSHA命令執行腳本。
2）、獲得返回值，如果返回「NOSCRIPT」錯誤則使用EVAL命令重新執行腳本。

```
SCRIPTLOAD "lua-script" 將腳本加入緩存，但不執行， 返回：腳本的SHA1摘要

SCRIPT EXISTS lua-script-sha1 判斷腳本是否已被緩存
```

5. SCRIPT FLUSH（該命令不區分大小寫）

清空腳本緩存，redis將腳本的SHA1摘要加入到腳本緩存後會永久保留，不會刪除，但可以手動使用SCRIPT FLUSH命令情況腳本緩存。
```
192.168.127.128:6379>script flush
OK
192.168.127.128:6379>SCRIPT FLUSH
OK
```

6. SCRIPT KILL（該命令不區分大小寫）

強制終止當前腳本的執行

如果當前執行的腳步對redis的數據進行了寫操作，則SCRIPT KILL命令不會終止腳本的運行，以防止腳本只執行了一部分

腳本中的所有命令，要麼都執行，要麼都不執行。
```
192.168.127.128:6379>script kill
(error)NOTBUSY No scripts in execution right now
192.168.127.128:6379>SCRIPT KILL
(error)NOTBUSY No scripts in execution right now
//這是當前沒有腳本在執行，所以提示該錯誤
```

7. lua-time-limit 5000（redis.conf配置文件中）

為了防止某個腳本執行時間過長導致Redis無法提供服務（比如陷入死循環）

Redis提供了lua-time-limit參數限制腳本的最長運行時間，默認為5秒鐘

當腳本運行時間超過這一限制後，Redis將開始接受其他命令但不會執行（以確保腳本的原子性，因為此時腳本並沒有被終止），而是會返回「BUSY」錯誤。



# refer:
- https://www.jishuwen.com/d/2ZKf/zh-tw
- https://kknews.cc/zh-tw/code/py82bbz.html