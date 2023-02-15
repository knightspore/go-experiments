# Bresenham's Line Algorithm Implemented in Go

An implementation of Bresenham's line algorithm written in go, based on the example described in the [Wikipedia Entry](https://en.wikipedia.org/wiki/Bresenham%27s_line_algorithm).

In addition to the core line drawing functionality, this package also includes:
- Methods for creating 2D and 3D coordinates in `primitives.go`
- A method for transforming 3D coordinates into 2D (barely) in `primitives.go`
- A method for creating polygons in `polygon.go`
- Support for rendering polygons to the terminal using [termbox-go](https://github.com/nsf/termbox-go)
- A method for loading `.obj` 3D model files into a Polygon (see example below)

Also provided is a method for loading a `.obj` 3d model file into a Polygon - an example of which is described below.

## Examples

To run the example, clone the directory and run `go mod tidy`, and then `go run .`. This will run the example in `main.go` using termbox. This will render the 3d model of a rock in your terminal.

## Contributing

Feel free to contribute if you would like to take a shot at improving upon, or adding to this little experiment.

Some Ideas:
- Improve the 3d -> 2d co-ordinate function
- Add support for different color models and color gradients to the rendered polygons
- Look at memory handling to print more complex models
- Add support for more advanced 3D file formats, such as STL or PLY\
- Expand to include circles or ellipses with Bresenham's Algorithm
