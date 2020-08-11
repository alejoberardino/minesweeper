FROM golang:1.14

WORKDIR /go/src/minesweeper

COPY . .

# TODO: This is not optmized, at all
RUN ./scripts/setup.sh 

CMD ["gin", "minesweeper", "--port 80"]

