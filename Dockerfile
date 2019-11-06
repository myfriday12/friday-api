##################################
# Build binary
##################################
FROM golang:1.13.1-alpine3.10 AS builder

RUN apk add --no-cache curl git

WORKDIR /go/src/friday-api
COPY . .

RUN apk --no-cache add tzdata

ENV TZ Asia/Kuala_Lumpur
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

ENV GO111MODULE=on
RUN cd api && CGO_ENABLED=0 go build -a -o fridayapp


##################################
# Build image
##################################
FROM alpine:3.10

RUN apk --no-cache add tzdata

ENV TZ Asia/Kuala_Lumpur
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

WORKDIR /go/src/friday-api
COPY . .
COPY --from=builder /go/src/friday-api/api/fridayapp /go/src/friday-api/api/
RUN chmod +x /go/src/friday-api/api/fridayapp
WORKDIR /go/src/friday-api/api
ENTRYPOINT ["./fridayapp"]