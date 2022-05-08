FROM golang:1.17 as builder
WORKDIR /build
COPY . .
ENV CGO_ENABLED=0
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
ENV GOOS=linux 
ENV GOARCH=amd64
RUN  go build -o httpserver .


FROM scratch
COPY --from=builder /build/httpserver /
EXPOSE 80
ENTRYPOINT ["/httpserver"]