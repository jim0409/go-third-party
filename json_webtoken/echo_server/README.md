# intro

工作上有時候需要做JWT，提供API調度，方便即時回饋~

# test with `ab`
> ab -p post.json -T application/json -c 100 -n 10000 http://127.0.0.1:8000/echo

```log
Concurrency Level:      100
Time taken for tests:   1.385 seconds
Complete requests:      10000
Failed requests:        0
Total transferred:      3010000 bytes
Total body sent:        2540000
HTML transferred:       1830000 bytes
Requests per second:    7222.65 [#/sec] (mean)
Time per request:       13.845 [ms] (mean)
Time per request:       0.138 [ms] (mean, across all concurrent requests)
Transfer rate:          2123.06 [Kbytes/sec] received
                        1791.55 kb/s sent
                        3914.62 kb/s total
```