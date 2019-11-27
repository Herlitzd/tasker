FROM alpine:3.7 as Tasker

COPY tasker /tasker/
COPY ./examples /tasker/examples
WORKDIR /tasker

ENTRYPOINT ["./tasker", "./examples/basic.json", "Main"]