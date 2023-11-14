# resizego
REsize JPG high quality and law space

#install Go

#install packages
go get -u golang.org/x/image/draw
go get -u golang.org/x/image

#run app
source directory will be in same path (eg: 'trvel')

go run resize_images.go travel

#results

travel-lg, travel-sm directories will be created.
Resized images will be stored seperately. 1000 images may resize without trash.

