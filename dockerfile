FROM ubuntu:18.04
COPY ./server-linux /server-linux
ENTRYPOINT /server-linux