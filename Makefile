build:
	go build -o bin/ball-clock cmd/clock/main.go

run:
	go run cmd/clock/main.go

clean:
	rm bin/ball-clock