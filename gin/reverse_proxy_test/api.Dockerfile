FROM centos:7

ADD api_server .

CMD ["./api_server"]