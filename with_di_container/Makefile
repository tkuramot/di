NAME = main
GEN = wire.go

all: run

# run
run: $(NAME)
	./$(NAME)

# build
$(NAME): $(GEN)
	go build -o $@ ./...

# generate wire_gen.go
$(GEN):
	wire ./...
