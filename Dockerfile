FROM golang:1.15
ENV INSTALL_PATH /go/src/stack-overflow-extension-backend
RUN mkdir $INSTALL_PATH
WORKDIR $INSTALL_PATH
RUN go get
RUN go build -o /app/main .
CMD ["/app/main"]