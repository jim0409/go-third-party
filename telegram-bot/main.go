package main

import "net/http"

/*
- https://api.telegram.org/bot5020139909:AAGJlFRhCGnr6PpiHjHg23sZWUUSEnX2ZZs/getUpdates
> 拿到所有訊息

- https://api.telegram.org/bot5020139909:AAGJlFRhCGnr6PpiHjHg23sZWUUSEnX2ZZs/sendMessage?chat_id=1068429272&parse_mode=Markdown&text=“watch you slowly bb”
> 發送訊息

*/

func main() {
	resp, err := http.Get("https://api.telegram.org/bot5020139909:AAGJlFRhCGnr6PpiHjHg23sZWUUSEnX2ZZs/sendMessage?chat_id=1068429272&parse_mode=Markdown&text='watch out'")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}
