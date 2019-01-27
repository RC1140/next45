# Next45 Project Mars
Demo Project For Next45

## Run instructions

### Requirements

* Make 
* Go >= 1.10

### Install Instructions

- clone the repo 
- cd $pathtorepo
- make install-deps
- make install

### Run Instructions

If your go bin is part of your path simply run

- `projectmars <path to datafile>`


e.g `projectmars $GOPATH/src/github.com/rc1140/next45/cmd/dat`

Note an example data file has been included in the project (next45/cmd/dat)

## Project structure

The project is split up into 2 packages, the library and the command line interface.
The library implements the basic requirements of the application such has handling 
the rovers movement and map setup. The cmd package handles IO and passing it to the library.

Each of the structs has their own test file which ensures that there is a clean relationship
between the tests and the functions they cover.
