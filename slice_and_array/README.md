# intro

在 go 裡面有兩種資料結構可以處理包含許多元素的清單資料, 分別是 Array 和 Slice:

- Array: 清單的長度是固定的(fixed length), 屬於原生型別(primitive type), 較少在程式中使用.
- Slice: 可以增加或減少清單的長度, 使用`[]`定義, 例如, []byte 是 byte slice, 指元素為 byte 的 slice;

e.g. []string 是 string slice, 指元素為 string 的 slice.

> 不論是 Array 或 Slice, 他們內部的元素都必須要有相同的資料型別(字串、數值、...)


在 Slice 中會包含

- Pointer to Array: 這個 pointer 會指向實際上在底層的 array
- Capacity: 從 slice 的第一個元素開始算起, 它底層 array 的元素數目
- Length: 該 slice 中的元素數目


# refer:
- https://pjchender.dev/golang/slice-and-array/
- https://coolshell.cn/articles/21128.html
