# resizego
REsize JPG high quality and law space

#install Go

#install packages
<code>
go get -u golang.org/x/image/draw
go get -u golang.org/x/image
</code>

#run app
<code>
source directory will be in same path (eg: 'trvel')

go run resize_images.go travel
</code>

#results
<code>
travel-lg, travel-sm directories will be created.
Resized images will be stored seperately. 1000 images may resize without trash.
</code>
