FROM alpine

ADD ./server /server
CMD ["/server", "-logtostderr"]
