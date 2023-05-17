package slice_and_array

import (
	"fmt"
	"reflect"
	"strings"
)

/*
在 go 裡面有兩種資料結構可以處理包含許多元素的清單資料, 分別是 Array 和 Slice:
- Array: 清單的長度是固定的(fixed length), 屬於原生型別(primitive type), 較少在程式中使用.
- Slice: 可以增加或減少清單的長度, 使用`[]`定義, 例如, []byte 是 byte slice, 指元素為 byte 的 slice;
[]string 是 string slice, 指元素為 string 的 slice.

不論是 Array 或 Slice, 他們內部的元素都必須要有相同的資料型別(字串、數值、...)
*/

// slice 型別: []T
// slice := make([]T, len, cap)
func BasicSliceInfo() {
	// 方式ㄧ: 建立一個帶有資料的 string slice, 適合用在知道 slice 裡面的元素有哪些時
	peopleWithElement := []string{"Araon", "Jim", "Bob", "ken"}
	_ = peopleWithElement

	// 方式二: 透過 make 可以建立`空 slice`, 適合用會對 slice 中特定位置元素進行操作時
	peopleWithKnownNum := make([]string, 4)
	_ = peopleWithKnownNum

	// 方式三: 空的 slice, 一般會搭配 append 使用
	var peopleEmptySlice []string
	_ = peopleEmptySlice

	// 方式四: 大概知道需要多少元素時使用, 搭配 append 使用
	peopleWithCap := make([]string, 0, 5) // len=0, cap=5, []
	_ = peopleWithCap

	/*
		實際上, 當 Go 在建立 slice 時, 它內部會建立兩個不同的資料結構, 分別是 slice 和 array
		- slice 的 zero value 是 `nill`, 而 `nil` 的 slice 其 `len` 和 `cap` 都是 `0`
	*/
}

/*
在 Slice 中會包含

- Pointer to Array: 這個 pointer 會指向實際上在底層的 array
- Capacity: 從 slice 的第一個元素開始算起, 它底層 array 的元素數目
- Length: 該 slice 中的元素數目

調用 slice 中的元素時, 會先指向該 slice, 而該 slice 的 pointer to head 會在指到底層的 array
*/

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// slice 的 `length` 和 `capacity` 可以變更:
func CapacityAndLength() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s) // len=6 cap=6 [2 3 5 7 11 13]

	s = s[:0]
	printSlice(s) // len=0 cap=6 []

	s = s[:4]
	printSlice(s) // len=4 cap=6 [2 3 5 7]

	s = s[2:]
	printSlice(s) // len=2 cap=4 [5 7]
}

// 使用 make 建立 slice 時，可以指定該 slice 的 length 和 capacity
func MakeTheSlice() {
	// make(T, length, capacity)
	a := make([]int, 5) // len(a)=5, cap(a)=5, [0 0 0 0 0]
	_ = a

	// 建立特定 capacity 的 slice
	b := make([]int, 0, 5) // len=0 cap=5, []
	b = b[:cap(b)]         // len=5 cap=5, [0 0 0 0 0]
	b = b[1:]              // len=4 cap=4, [0 0 0 0]
	_ = b
}

// 若 length 的數量不足時，將無法將元素放入 slice 中，這時候可以使用 append 來擴展 slice
// slice 的 length 不足時無法填入元素
func SliceLengthIsNotEnough() {
	scores := make([]int, 0, 10)
	fmt.Println(len(scores), cap(scores)) // 0, 10
	//scores[7] = 9033  // 無法填入元素，因為 scores 的 length 不足

	// 這時候可使用 append 來擴展 slice
	scores = append(scores, 5)
	fmt.Println(scores)
	fmt.Println(len(scores), cap(scores)) // 1, 10

	// 但要達到原本的目的，需要使用切割 slice
	scores = scores[0:8]
	fmt.Println(len(scores), cap(scores)) // 8, 10
	scores[7] = 9033
	fmt.Println(scores) // [5 0 0 0 0 0 0 9033]
}

/*
Zero Value
*/

// slice 的 zero value 是 nil
// 也就是當一個 slice 的 length 和 capacity 都是 0，並且沒有底層的 array 時
func ZeroSliceIsNil() {
	var s []int
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s) // len=0 cap=0 []

	if s == nil {
		fmt.Println("nil!") // nil!
	}
}

// slice of integer, boolean, struct
func TypeOfSlice() {
	// integer slice
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	// boolean slice
	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	// struct slice
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

// 建立多維度的 slice（slices of slices）
func MultiDimSlice() {
	arr := make([][]int, 3)
	fmt.Println(arr) // [[] [] []]

	// 賦值
	arr[0] = []int{1}
	arr[1] = []int{2}
	arr[2] = []int{3}
	fmt.Println(arr) // [[1] [2] [3]]
}

// 建立多維度的陣列並同時賦值
func MultiDimSliceWithValue() {
	arr := [][]int{
		[]int{1},
		[]int{2},
	}
	fmt.Println(arr) // [[1] [2]]
}

// 更複雜的 ...
func ComplexitySlice() {
	board := [][]string{
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
	}
	fmt.Println(board) // [[- - -] [- - -] [- - -]]

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	fmt.Println(board) // [[X - X] [O - X] [- - O]]

	//X - X
	//O - X
	//- - O
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

// slice 的操作與疊代
func SliceOperationAndIter() {
	// 定義內部元素型別為字串的陣列
	fruits := []string{"apple", "banana", "grape", "orange"}

	// 取得 slice 的 length 和 capacity
	fmt.Println(len(fruits)) // length, 4
	fmt.Println(cap(fruits)) // capacity, 4

	// append 為陣列添加元素（不會改變原陣列）
	fruits = append(fruits, "strawberry")

	// range syntax: fruits[start:end] 可以截取陣列，從 "start" 開始到 "end-1" 的元素
	fmt.Println(fruits[0:2]) //  [apple, banana]
	fmt.Println(fruits[:3])  //  [apple, banana, grape]
	fmt.Println(fruits[2:])  //  [grape, orange, strawberry]

	// 透過 range cards 可以疊代陣列元素
	for _, fruit := range fruits {
		fmt.Println(fruit)
	}

	/*
	   在疊代陣列時會使用到 for i, card := range cards{…}，
	   其中如果 i 在疊代時沒有被使用到，但有宣告它時，程式會噴錯;
	   因此如果不需要用到這個變數時，需把它命名為 _。
	*/
}

// 切割(:)
/*
須要留意的是，在 golang 中使用`:`來切割 slice 時, 並不會複製原本的 slice 中的資料.
而是建立一個新的 slice,

import!!!
但實際上還是指稱到相同位址的底層 array(並沒有複製新的)，因此還是會改到原本的元素!
但實際上還是指稱到相同位址的底層 array(並沒有複製新的)，因此還是會改到原本的元素!
但實際上還是指稱到相同位址的底層 array(並沒有複製新的)，因此還是會改到原本的元素!
*/

func SliceSplit() {
	scores := []int{1, 2, 3, 4, 5}
	newSlice := scores[2:4]
	fmt.Println(newSlice) // 3, 4
	newSlice[0] = 999     // 把原本 scores 中 index 值為 3 的元素改成 999
	fmt.Println(scores)   // 1, 2, 999, 4, 5
}

// 另外，使用 : 時，不能超過它本身的 capacity，否則會導致 runtime panic：
func PanicWhileSplitting() {
	s := []int{1, 2, 3, 4, 5}

	// over the slice capacity will cause runtime panic
	overSliceCap := s[:6]
	fmt.Println(overSliceCap)
}

/*
由於無法直接透過`:`來擴充 slice 的 capacity,
因此若有需要擴充原 slice 的 capacity 時,

!!!
需要透過 append 或 copy 的方式來擴充原本 slice 的 capacity。

*/

// func append(s []T, x ...T) []T
func AppendSlice() {
	var s []int
	printSlice(s) // len=0 cap=0 []

	// 一次添加多個元素
	s = append(s, 2, 3)
	printSlice(s) //len=2 cap=2 [2 3]

	// 一次添加一個元素
	s = append(s, 4)
	printSlice(s) // len=3 cap=4 [2 3 4]
	// cap 的數量再擴展時會是 1, 2, 4, 8, 16, .... 。

	// 使用 append(a, b...) 就可以把 b 這個 slice append 到 a 這個 slice 中：
	a := []int{1, 2}
	b := []int{3, 4, 5}

	a = append(a, b...) // 等同於 append(s, b[0], b[1], b[2])
	fmt.Println(a)      // [1 2 3 4 5]
}

/*
- copy
func copy(dst, src []T) int
回傳的值是複製了多少元素（the number of elements copied）
*/
func cloneScores() {
	scores := []int{1, 2, 3, 4, 5}

	// STEP 1：建立空 slice 且長度為 4
	cloneScores := make([]int, 4)

	// STEP 2：使用 copy 將前 scores 中的前三個元素複製到 cloneScores 中
	copy(cloneScores, scores[:len(scores)-2])
	fmt.Println(cloneScores) // [1,2,3, 0]
}

/*
cloneScores 的元素數量如果「多於」被複製進去的元素時，會用 zero value 去補。例如，當
cloneScores 的長度是 4，但只複製 3 個元素進去時，最後位置多出來的元素會補 zero value。

cloneScores 的元素數量如果「少於」被複製進去的元素時，超過的元素不會被複製進去。例如，當
cloneScores 的長度是 1，但卻複製了 3 個元素進去時，只會有 1 個元素被複製進去。
*/

// Growing slices: https://blog.golang.org/slices-intro
func GrowingSlices() {
	// 使用 copy 可以用來擴展擴展 slice 的 capacity。在沒有 copy 時，寫法會像這樣：
	s := []byte{'1', 'a', 'c'}

	t := make([]byte, len(s), (cap(s)+1)*2) // +1 in case cap(s) == 0
	for i := range s {
		t[i] = s[i]
	}
	s = t
}

func CopySlice() {
	s := []byte{'1', 'a', 'c'}
	t := make([]byte, len(s), (cap(s)+1)*2)
	copy(t, s)
	s = t
}

func RangeSlice() {
	//	如果只需要 index：for i := range pow
	// 如果只需要 value：for _, value := range pwd
	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

// 把 slice 當成 function 的參數時
/*
所以當我們在 func 中把 slice 當成參數傳入時, 它仍然還是使用 pass by value 的方式.
但因為它複製的是 slice 本身, 而這個 slice 卻還是指向在同一個記憶體位址的 Array.
這也就是為什麼, 在函式中當我們沒有使用 Pointer 的方式去修改這個 slice 時, 但它卻還是會改動到 array 內的元素:

實際上 slice 並不會保存任何資料, 它只是描述底層的 array, 當我們修改 slice 中的元素時,
實際上是修改底層 array 的元素.

不同的 slice 只要共享相同底層的 array 時，就會看到相同對應的變化。
*/

func changeSliceItem(words []string) {
	words[0] = "Hi"
}

func UpdateSlice() {
	names := []string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	a := names[0:2]
	b := names[1:3]

	fmt.Println(a, b) // [John Paul] [Paul George

	b[0] = "XXX"       // a 和 b 這兩個 slice 參照到的是底層相同的 array
	fmt.Println(a, b)  // [John XXX] [XXX George]
	fmt.Println(names) // [John XXX George Ringo]

	words := []string{"Hello", "every", "one"}
	fmt.Println(words) // [Hello every one]

	changeSliceItem(words)
	fmt.Println(words) // [Hi every one]
}

// Array 建立: 在 Array 中陣列的元素是固定的
// 陣列型別：[n]T
// 陣列的元素只能有 10 個，且都是數值
func BuildArray() {
	// 先定義再賦值
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a) // [Hello World]

	var people [4]int               // len=4 cap=4, [0,0,0,0]
	people = [4]int{10, 20, 30, 40} // len=4 cap=4, [10,20,30,40]
	fmt.Println(people)

	// 定義且同時賦值
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes) // [2 3 5 7 11 13]

	// 使用 ... 根據元素數目建立 Array : 使用 [...]T{} 可以根據元素的數目自動建立陣列：
	// 沒有使用 ...，建立出來的會是 slice
	arr := []string{"North", "East", "South", "West"}
	fmt.Println(reflect.TypeOf(arr).Kind(), len(arr)) // slice 4

	// 使用 ...，建立出來的會是 array
	arrWithDots := [...]string{"North", "East", "South", "West"}
	fmt.Println(reflect.TypeOf(arrWithDots).Kind(), len(arrWithDots)) // array 4
}

// Filter: https://gosamples.dev/generics-filter/
// 利用`泛型`實作 filter & includes
func filter[T any](slice []T, f func(T) bool) []T {
	var filtered []T

	for _, element := range slice {
		if f(element) {
			filtered = append(filtered, element)
		}
	}

	return filtered
}

func includes[T any](slice []T, f func(T) bool) bool {
	for _, element := range slice {
		if f(element) {
			return true
		}
	}
	return false
}

// Remove: 移除元素
func RemoveLastElem() {
	scores := []int{1, 2, 3, 4, 5}
	removeLastIndex := scores[:len(scores)-1]

	fmt.Println(removeLastIndex) // [1,2,3,4]
}

func RemoveParticularElem(source []int, index int) []int {
	// STEP 1：取得最後一個元素的 index
	lastIndex := len(source) - 1

	// STEP 2：把要移除的元素換到最後一個位置
	source[index], source[lastIndex] = source[lastIndex], source[index]

	// STEP 3：除了最後一個位置的元素其他回傳出去
	return source[:lastIndex]
}

func DemoRemoveParticularElem() {
	scores := []int{1, 2, 3, 4, 5}
	scores = RemoveParticularElem(scores, 2)
	fmt.Println(scores) // [1, 2, 5, 4]
}

// Implicit memory aliasing in for loop : https://stackoverflow.com/questions/62446118/implicit-memory-aliasing-in-for-loop
func ImplicitMemoryAliasingInForLoop() {
	var versions []int
	createWorkerFor := func(i *int) int { return *i }

	for i, v := range versions {
		_ = v                                // 不要使用 res := createWorkerFor(&v)
		res := createWorkerFor(&versions[i]) // 使用 versions[i]
		// ...
		_ = res
	}
}

// preallocating (prealloc) 的問題: https://stackoverflow.com/questions/59734706/how-to-resolve-consider-preallocating-prealloc-lint
/*
雖然 append 會自動幫我們擴展 slice 的 capacity.
但如果可能的話, 在建立 slice 時預先定義好 slice 的 length 和 capacity 將可以避免不必要的記憶體分派（memory allocations）

每次 array 的 capacity 改變時，就會進行一次 memory allocation。
*/

func PreAllocatingSlice() {
	// 方法一: 使用 append, 搭配 length: 0, 並設定 capacity 定好:
	contacts := make([]model.Contact, 0, len(patientContacts))
	for i := range patientContacts {
		contact := toInternalContactFromPCC(p, &patientContacts[i], residentID)
		contacts = append(contacts, contact)
	}

	// 方法二(maybe better): 不使用 append, 直接設定 len 和 cap 並透過 index 的方式把值給進去:
	contacts := make([]model.Contact, len(patientContacts))
	for i := range patientContacts {
		contact := toInternalContactFromPCC(p, &patientContacts[i], residentID)
		contacts[i] = contact
	}
}
