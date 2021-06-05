# intro
- nsq is a queue manager like kafka-zookeeper

# different with kafka
- https://blog.csdn.net/weixin_41571449/article/details/79117892
- https://blog.csdn.net/lucyxu107/article/details/88675400
```
kafka-zookeeper: topic/partition & offset to persist data
nsq-nsqd : topic/channel & memory cache without persist data
``` 

# no log was record, need to use systemd or other tools to record the stdout/stderr
- https://github.com/nsqio/nsq/issues/977#issuecomment-354047112
```
There is currently no feature to log to file directly, logs are always written to stderr of the process (which is the terminal, if you run it in a terminal). You can use a service manager like systemd, upstart, docker etc. to have the stderr of the process written to rotating log files. (Unless you use MS Windows. That is more-or-less a bug.)
```

# Inflight phase mechanism in NSQ
- FIN: Finish a message, success
- REQ: Re-queue a message, fail to deal the msg and gonna to retry
- TOUCH: Reset the timeout for an in-flight message, need to reset the event source timeout setting
```
In inflight status, NSQ can make sure message can be `comsumed` at least once

After message was sent to client, `msg` & `msg timeout` would be stored into pqueue first

If client receive the msg, response would be `FIN` / `REQ` or `TOUCH`

For timeout msg, msg would stored in queue and resend to clients

That being said, NSQD only guarantee that every msg would be sent `at least once` whereas `exactly once` reliablity need to be performed by clients.
```

# About nsqd lookup
- https://nsq.io/components/nsqlookupd.html
```
nsqlookupd is the daemon that manages topology information. Clients query nsqlookupd to discover nsqd producers for a specific topic and nsqd nodes broadcasts topic and channel information.

There are two interfaces: A TCP interface which is used by nsqd for broadcasts and an HTTP interface for clients to perform discovery and administrative actions.
```


# refer:
- http://tleyden.github.io/blog/2014/11/12/an-example-of-using-nsq-from-go/
- https://nsq.io/overview/design.html
- https://github.com/nsqio/nsq/issues/977
- https://www.youtube.com/watch?v=GCOvuCKe5zA
- https://swanspouse.github.io/2018/12/10/nsq-message-in-flight/
- https://blog.csdn.net/sd653159/article/details/83624661
- https://nsq.io/components/nsqlookupd.html
