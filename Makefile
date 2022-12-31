GC=go

all:
	$(GC) build -o bin/buzzer

dev:
	$(GC) run -tags dev main.go