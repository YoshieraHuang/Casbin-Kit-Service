FROM alpine:latest

RUN mkdir /app
WORKDIR /app

COPY casbinsvc /app/
COPY model/ /app/model/

ENTRYPOINT [ "./casbinsvc" ]