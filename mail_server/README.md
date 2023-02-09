## intro

send mail via smtp server

## pre required
1. go env required
> https://go.dev/dl/
2. install related package
> go mod vendor
3. copy a config.ini & filled value
> cp config.ini.tmpl config.ini

## quick start
1. build
> go build -o app .
2. run
> ./app


## testing
### required:
1. authcode: for authentication
2. id: for render msg index
3. mail: sending address
### option:
1. sub: subject
 
```bash
curl "http://127.0.0.1:8000/msg/send" -d '{"authcode":"test", "id":1, "sub": "good_subject", "mail": "berserker.01.tw@gmail.com"}'

# lack sub
curl "http://127.0.0.1:8000/msg/send" -d '{"authcode":"test", "id":1, "mail": "berserker.01.tw@gmail.com"}'

```

## config google smtp server settings
- https://www.webdesigntooler.com/google-smtp-send-mail

## refer:
- https://gist.github.com/jpillora/cb46d183eca0710d909a

