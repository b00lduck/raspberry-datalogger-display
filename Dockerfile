FROM hypriot/rpi-golang
WORKDIR /gopath1.5/src/github.com/b00lduck/raspberry-datalogger-display
CMD ["raspberry-datalogger-display","/dev/fb1","/dev/input/touchscreen"]

ADD . /gopath1.5/src/github.com/b00lduck/raspberry-datalogger-display/
RUN go get
RUN go build
