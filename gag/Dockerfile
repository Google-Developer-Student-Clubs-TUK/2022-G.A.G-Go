FROM golang:alpine as builder

# Configure Go
# 경로 에러가 발생해 설정
RUN mkdir -p ${GOPATH}/src/app ${GOPATH}/bin

WORKDIR /go/src/app

# Git 에러가 발생해 설치
RUN apk add git

# Get github.com/cespare/reflex 가 gopath에 설치되도록 auto 사용
ENV GO111MODULE=auto
RUN go get github.com/cespare/reflex

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go mod tidy
RUN go build -o ./run .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

#Copy executable from builder
COPY --from=builder /go/src/app/run .

EXPOSE 8080
CMD ["./run"]