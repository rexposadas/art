ts := $(shell /bin/date "+%Y-%m-%d-%H-%M-%S")

build:
	go build --race -o /tmp/art

all: build
	/tmp/art circles regular -c 2 -f input/small.json
	/tmp/art circles gradient -c 2 -f input/small.json
	/tmp/art circles grid -c 2 -f input/small.json
	/tmp/art hole -c 2 -f input/small.json
	/tmp/art perls -c 2 -f input/small.json
	/tmp/art squares spiral -c 2 -f input/small.json
	/tmp/art random-shapes -c 2 -f input/small.json


for_sale: build
	/tmp/art circles regular -c 5 -f input/config.json
	/tmp/art circles gradient -c 5 -f input/config.json
	/tmp/art hole -c 5 -f input/config.json
	/tmp/art squares spiral -c 5 -f input/config.json
	/tmp/art random-shapes -c 5 -f input/config.json
	/tmp/art perls -c 5 -f input/config.json
	/tmp/art circles grid -c 5 -f input/config.json

gift: build
	/tmp/art gift -f input/list.csv

circles_grid: build
	/tmp/art circles grid -c 10 -f input/config.json

circles_gradient: build
	/tmp/art circles gradient -c 5 -f input/config.json

circles_loop: build
	/tmp/art circles loop -c 20 -f input/config.json

circles: build
#	/tmp/art circles regular -c 7 -f input/config.json
	/tmp/art circles

squares: build
	/tmp/art random-shapes

perls: build
	/tmp/art perls -c 10 -f input/config.json






samples: build
	/tmp/art samples


smoke: build
	/tmp/art smoke -c 10 -f input/config.json

junas: build
	/tmp/art junas -c 10 -f input/config.json

dot: build
	/tmp/art dot -c 5 -f input/config.json

domain: build
	/tmp/art domain -c 10 -f input/config.json

hole: build
	/tmp/art hole -c 10 -f input/config.json

julia: build
	/tmp/art julia -c 10 -f input/config.json

profile: build
	#rm output/*
	/tmp/art profile -f input/2018-profile.JPG
	open output/*

spiral: build
	rm output/*
	/tmp/art squares spiral -c 1 -f input/config.json
	open output/*

merge: build
	rm output/merged/*
	/tmp/art merge 	-f input/list.csv
	open output/merged/*


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