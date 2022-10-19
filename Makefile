ts := $(shell /bin/date "+%Y-%m-%d-%H-%M-%S")

build:
	go build -o /tmp/art

all: build
	make gift
	make circle

merge: build
	/tmp/art merge -f input/list.csv

gift: build
	/tmp/art gift -f input/list.csv

try: build
	/tmp/art try
	#open output/merged/*

circles_grid: build
	/tmp/art circles grid -c 10 -f input/circle.json

circles_gradient: build
	/tmp/art circles gradient -c 20 -f input/circle.json

circles_loop: build
	/tmp/art circles loop -c 20 -f input/circle.json

circles: build
#	/tmp/art circles regular -c 7 -f input/circle.json
	/tmp/art circles

squares: build
	/tmp/art squares regular

perls: build
	/tmp/art perls -c 10 -f input/basic.json



samples: build
	/tmp/art samples


smoke: build
	/tmp/art smoke -c 10 -f input/basic.json

junas: build
	/tmp/art junas -c 10 -f input/basic.json

dot: build
	/tmp/art dot -c 10 -f input/basic.json

domain: build
	/tmp/art domain -c 10 -f input/basic.json

hole: build
	/tmp/art hole -c 10 -f input/basic.json

julia: build
	/tmp/art julia -c 10 -f input/basic.json

spiral_square: build
	/tmp/art squares spiral -c 20 -f input/basic.json


store:
	mv output output-$(ts)
	mkdir output
	mkdir output/gift
	mkdir output/merged
	mkdir output/generated
	mkdir output/generated/circle


clean:
	@rm output/*.png &
	@rm output/generated/*.png &
	@rm output/generated/circles/*.png &
	@rm output/gift/*.png &
	@rm output/merged/*.png &