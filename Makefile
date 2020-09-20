all: build run move

run:
	@./bin/gp -status

build:
	@go build -o bin/gp && echo Build success || echo Build error

move:
	@cp bin/gp /usr/local/bin/ && echo Moved 'gp' bin to /usr/local/bin/ || echo Failed to move binary /usr/local/bin/
