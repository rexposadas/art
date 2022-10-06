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

circles_grid: build
	./art circles grid -c 10 -f input/circle.json

circles_gradient: build
	./art circles gradient -c 20 -f input/circle.json

circles_loop: build
	./art circles loop -c 20 -f input/circle.json

circles: build
	./art circles regular -c 7 -f input/circle.json

squares: build
	./art squares regular

perls: build
	./art perls -c 10 -f input/basic.json



samples: build
	./art samples -c 10 -f input/basic.json


smoke: build
	./art smoke -c 10 -f input/basic.json

junas: build
	./art junas -c 10 -f input/basic.json

dot: build
	./art dot -c 10 -f input/basic.json

domain: build
	./art domain -c 10 -f input/basic.json

hole: build
	./art hole -c 10 -f input/basic.json

julia: build
	./art julia -c 10 -f input/basic.json

spiral_square: build
	./art squares spiral -c 20 -f input/basic.json


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