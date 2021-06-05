FROM centos:7

ADD ./server/grpc_server .

CMD ["./grpc_server"]