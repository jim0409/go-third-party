# intro
this design is motivate by instant message service

due to service is a read heavy system; one send msg related to group read

also, tons of group also cause write bottleneck with Mysql single node.

# feature
1. support horizontal scale for db nodes
2. mechanism for db nodes selected decision, local cache & hot reload
3. store decision, nodes mapping talbes, in mysql
4. 
