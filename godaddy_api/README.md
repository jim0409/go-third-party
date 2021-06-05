# intro
use godaddy api to retrive dns domain & handle account extra behavior

# quick start
### 1. Create a file named `godaddy_ac.csv`
run go process with specific csv file, format with
```csv
id,name,customer_id,key,secret
1,jim,280076842,e5N8Yzzn7hBo_Fxe7rVEZwvyTfMW4ztQmJG,PDnDKUFDZjefBUUyTeWGkb
```

### 2. processing with `go run`
> go run main.go apis.go tools.go -file "godaddy_ac.csv"

# build with cli
```bash
#!/bin/bash

# mac_os
GO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.gitcommitnum=`git rev-parse --short=6 HEAD`" .

# linux_os
GO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X main.gitcommitnum=`git rev-parse --short=6 HEAD`" .
```

# refer:
- https://developer.godaddy.com/getstarted
