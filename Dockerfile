FROM golang
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go install github.com/mitranim/gow@latest
EXPOSE 3011
CMD ["gow", "run", "main.go"]