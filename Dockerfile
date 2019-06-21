FROM golang
COPY ./go-mux-api /go-mux-api
RUN chmod +x /go-mux-api
ENTRYPOINT ["/go-mux-api"]
