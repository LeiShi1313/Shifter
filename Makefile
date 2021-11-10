.ONESHELL:
.PHONY: statik all

all:
	$(MAKE) svelte
	go build -o Shifer ./main.go

svelte:
	cd web && yarn && yarn build && cd -
	statik -src=./web/build

clean:
	rm -rf ./statik ./web/build