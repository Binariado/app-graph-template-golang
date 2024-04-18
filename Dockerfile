FROM golang:1.22

WORKDIR /app/app-graph
COPY . .

RUN go build -o app-graph

COPY . /app/app-graph

RUN chmod a+x /app/app-graph

EXPOSE $PORT

ENTRYPOINT ["go", "run", "/app/app-graph" ]
