<b># resizego</b>
Resize JPG high quality and low space

<b># install Go</b>

<b># install packages</b>

> go get -u golang.org/x/image/draw </br>
> go get -u golang.org/x/image


<b># run app</b>

<i>source directory will be in same path (eg: 'trvel')</i>
> go run resize_images.go travel


<b># results</b>

travel-lg, travel-sm directories will be created.</br>
Resized images will be stored seperately. </br>
1000 images may resize without memory overflow.</br>


