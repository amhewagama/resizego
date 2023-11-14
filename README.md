<h1>ResizeGo</h1>

<p><strong>Resize multiple JPG images with high quality and low space</strong></p>

<p>ResizeGo is a simple command-line tool written in Go that allows you to efficiently resize a collection of JPG images, optimizing for both high quality and reduced storage space. The tool supports concurrent processing for faster resizing and is designed to be easy to use.</p>

<h2>Features</h2>

<ul>
    <li>Resize multiple JPG images concurrently</li>
    <li>Optimize for high quality with adjustable JPEG compression</li>
    <li>Efficiently reduce file size for better storage utilization</li>
    <li>Specify custom output directories for resized images</li>
    <li>Command-line interface for easy integration into workflows</li>
</ul>

<h2>Usage</h2>

<ol>
    <li>Download the latest release or build from source.</li>
    <li>Run the executable with the source directory as a command-line argument:</li>
</ol>

<pre>./resize_images source_directory</pre>

<p>Resized images will be saved in separate output directories for high quality and low space versions.</p>

<h2>Options</h2>

<ul>
    <li><code>-quality</code>: Adjust JPEG compression quality (default is 80)</li>
    <li><code>-maxwidth</code> and <code>-maxheight</code>: Set maximum width and height for resizing</li>
    <li><code>-output</code>: Specify custom output directory names for different sizes</li>
</ul>

<h2>Requirements</h2>

<p>Go Lang</p>
> go get -u golang.org/x/image/draw </br>
> go get -u golang.org/x/image

<h2>Examples</h2>

<p>Resize images in the "travel" directory, saving large versions to "travel-lg" and small versions to "travel-sm":</p>

<pre>go run resize_images.go travel</pre>
<p>travel-lg, travel-sm directories will be created:</p>
<p>Resized images will be stored seperately.</p>
<p>1000 images may resize without memory overflow.</p>

<h2>Contributions</h2>

<p>Contributions are welcome! Feel free to open issues, submit pull requests, or suggest new features to make ResizeGo even more versatile.</p>

<h2>License</h2>

<p>This project is licensed under the <a href="LICENSE">MIT License</a>.</p>
