ts := $(shell /bin/date "+%Y-%m-%d-%H-%M-%S")

build:
	go build -o /tmp/art

for_sale: build
	/tmp/art circles regular -c 5 -f input/basic.json
	/tmp/art circles gradient -c 5 -f input/basic.json
	/tmp/art hole -c 5 -f input/basic.json
	/tmp/art squares spiral -c 5 -f input/basic.json
	/tmp/art squares regular -c 5 -f input/basic.json
	/tmp/art perls -c 5 -f input/basic.json
	/tmp/art circles grid -c 5 -f input/basic.json

gift: build
	/tmp/art gift -f input/list.csv

circles_grid: build
	/tmp/art circles grid -c 10 -f input/basic.json

circles_gradient: build
	/tmp/art circles gradient -c 5 -f input/basic.json

circles_loop: build
	/tmp/art circles loop -c 20 -f input/basic.json

circles: build
#	/tmp/art circles regular -c 7 -f input/basic.json
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
	/tmp/art dot -c 5 -f input/basic.json

domain: build
	/tmp/art domain -c 10 -f input/basic.json

hole: build
	/tmp/art hole -c 10 -f input/basic.json

julia: build
	/tmp/art julia -c 10 -f input/basic.json

spiral: build
	#rm output/*.png
	/tmp/art squares spiral -c 5 -f input/basic.json
	open output/*.png

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