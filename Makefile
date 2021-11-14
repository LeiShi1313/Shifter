.ONESHELL:
.PHONY: statik all

all:
	$(MAKE) frontend
	go build -o Shifer ./main.go

frontend:
	cd frontend && yarn && yarn build && cd -
	statik -src=./web/build

clean:
	rm -rf ./statik ./frontend/build