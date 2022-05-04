build:
	go build -o bin/ball-clock cmd/clock/main.go

ball-clock.prof: build	
	./bin/ball-clock -cpuprofile=ball-clock.prof

run:
	go run cmd/clock/main.go

clean:
	rm bin/ball-clock