# fastimage

[![GoDoc](https://godoc.org/github.com/rubenfonseca/fastimage?status.png)](https://godoc.org/github.com/rubenfonseca/fastimage) [![Build Status](https://travis-ci.org/rubenfonseca/fastimage.svg?branch=master)](http://travis-ci.org/rubenfonseca/fastimage)

by Ruben Fonseca (@rubenfonseca)

Golang implementation of [fastimage](https://pypi.python.org/pypi/fastimage/0.2.1).
Finds the size or type of an image given its uri by fetching as little as needed.

## How?

fastimage parses the image data as it is downlaoded. As soon as it finds out
the size and type of the image, it stops the download.

## Install

    $ go get github.com/rubenfonseca/fastimage

