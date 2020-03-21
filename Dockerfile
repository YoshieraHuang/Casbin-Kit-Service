FROM golang:alpine AS builder
RUN apk --no-cache add build-base
COPY . /code/casbinsvc
RUN cd /code/casbinsvc/cmd/casbinsvc \
    && GOOS=linux CGO_ENABLE=1 GOARCH=amd64 go build .

FROM alpine:latest
RUN apk --no-cache add tzdata ca-certificates libc6-compat libgcc libstdc++
COPY --from=builder /code/casbinsvc/cmd/casbinsvc/casbinsvc /app/casbinsvc
COPY --from=builder /code/casbinsvc/cmd/casbinsvc/model/ /app/model/
CMD [ "/app/casbinsvc" ]