# GO IMAGE PROCCESS
## What the purpose of this package?
The purpose of this package is to _image resizing_, _compressing_ processes and _adding to S3 Bucket_.

## How can use this package?
### First of all install package and import
```cmd
go get "github.com/dahaiyiyim/imgkit"
```
```go
import imgkit "github.com/dahaiyiyim/imgkit"
```
### Resizing and compressing image with using ImgCompress() method.
```go
base64Str := "image_base64_string"
imageFile, _ := imgkit.ImgCompress(240, 240, base64Str)
```
### Create config with using New() method.
```go
config := imgkit.New(&imgkit.S3Config{
S3_ACCESS_KEY: "YOUR_ACCESS_KEY",
S3_SECRET_KEY: "YOUR_SECRET_KEY",
S3_REGION:     "YOUR_REGION",
S3_BUCKET:     "YOUR_BUCKET",
S3_OBJECT_KEY: "YOUR_OBJECT_KEY",
})
```
### Add to S3 with using AddS3() method.
```go
_ = config.AddS3(imageFile, "image.jpg")
```
