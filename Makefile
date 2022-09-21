ts := $(shell /bin/date "+%Y-%m-%d-%H-%M-%S")

build:
	go build -o art

all: build
	make gift
	make circle

merge: build
	./art merge -f input/list.csv

gift: build
	./art gift -f input/list.csv

try: build
	./art try
	#open output/merged/*

circle: build
	./art circles -c 7 -f input/circle.json
	./art circles -c 7 -f input/circle-spring.json



store:
	mv output output-$(ts)
	mkdir output
	mkdir output/gift
	mkdir output/merged
	mkdir output/generated
	mkdir output/generated/circle


samples: build
	./art samples

clean:
	@rm output/*.png &
	@rm output/generated/*.png &
	@rm output/generated/circles/*.png &
	@rm output/gift/*.png &
	@rm output/merged/*.png &