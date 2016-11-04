FROM alpine:latest

EXPOSE 80
ENV SM_IMAGE cloudmore/ext-auth-srv

ADD bin bin

CMD ["bin/ext-auth-srv"]
