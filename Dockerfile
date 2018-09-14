FROM golang

WORKDIR /go

COPY ./cards .

# RUN go build

CMD ["go", "run", "main.go", "deck.go" ]

