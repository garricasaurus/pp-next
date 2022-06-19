FROM golang:alpine AS builder

WORKDIR /pp-next

COPY /go.mod ./
COPY /go.sum ./
COPY /*.go ./
COPY /config ./config
COPY /consts ./consts
COPY /viewmodel ./viewmodel
COPY /store ./store
COPY /controller ./controller
COPY /model ./model

RUN go build


FROM alpine:latest

WORKDIR /pp-next

COPY --from=builder /pp-next/ppnext ./ppnext
COPY /assets ./assets
COPY /templates ./templates

CMD [ "./ppnext" ]
