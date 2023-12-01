
WOLFMUD_PATH := $(patsubst %/github.com/einherij/wolfmud-web, %/code.wolfmud.org/WolfMUD.git, $(PWD))

run:
	go run ./cmd/webserver/main.go

update:
	git pull
	go mod download

update_wolfmud:
	cd $(WOLFMUD_PATH)
	git pull
