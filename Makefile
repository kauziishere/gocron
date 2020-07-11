all:
	@echo "\tgocron bash utility\n"
	@echo "\tAssumes pre-installed golang"
	@echo "\tdeps:\t\tsetup dependancies for package"
	@echo "\tbuild:\t\tbuild gocron"
	@echo "\tinstall:\tinstall gocron"

deps:
	@echo "Install dependancies"
	go get github.com/rs/zerolog/log

build: deps
	@echo "Building gocron"
	go build

install: build
	@echo "Install gocron"
	go install
	@echo "Installation completed for gocron"
	export gocron=$GOPATH'/bin/gocron'
