Art is a CLI which makes it easy to generated 2D art in Go.  

# Quick start

To start creating art, build the app and run it:


`go build -o ./art && ./art circles`

That creates a directory called "output" with art you just generated. 

If you are on a macbook you can open the files with `open output/*`. A sample circle image is shown below.

![](images/samples-669a073e.png)


# Using more options 

You can view the options you can have by: 

`./art`


# Other calls you can try 

1. `./art

# using the config file
Config files lets you alter the looks of your art. in the input folder you can see this config: 

```json
{
  "type": "circle",
  "out": {
    "prefix": "circle",
    "dir": "output"
  },
  "canvas": {
    "width": 300,
    "height": 300
  },
  "colors": {
    "scheme": [
      {"r":111, "g": 84, "b": 140, "a":  255},
      {"r":92, "g": 204, "b": 206, "a":  255},
      {"r":178, "g": 162, "b": 150, "a":  255}
    ]
  }
}
```

You can set the canvas width and color schemes. 

To make use of this config file you can make a call like so: 

`./art circles grid -c 10 -f input/circle.json`

That will produce 10 images with the circle config. A sample image would be: 

![](images/circle-300_x_300.png)



# Features not implemented but planned

1. Each image is created with an associated config file. This config file shows what parameters were used when that image was generated. This is helpful if you want to recreate a similar image. (not implemented)



