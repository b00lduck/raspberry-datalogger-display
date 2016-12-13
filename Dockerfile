FROM rem/rpi-golang-1.7:latest

WORKDIR /gopath/src/github.com/b00lduck/raspberry-datalogger-display
ENTRYPOINT ["raspberry-datalogger-display","/dev/fb1","/dev/input/touchscreen"]

ADD . /gopath/src/github.com/b00lduck/raspberry-datalogger-display/
RUN go get
RUN go build
