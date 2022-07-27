FROM golang:1.18-alpine AS build
WORKDIR /src
COPY src ./src
COPY main.go .
COPY go.mod .
COPY .env .

RUN mkdir "../assets"
RUN touch ../assets/movies.csv

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o swagger .

FROM scratch AS runtime
COPY --from=build /src/swagger ./
COPY --from=build /src/.env ./

EXPOSE 8080/tcp
ENTRYPOINT ["./swagger"]