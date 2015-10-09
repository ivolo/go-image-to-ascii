
# go-image-to-ascii

Convert image to ascii.

```go
import "github.com/ivolo/go-image-to-ascii"
// ..

f, err := os.Open("screenshot.png")
img, err := png.Decode(f)
fmt.Print(ascii.Convert(img))
```

![image](https://cloud.githubusercontent.com/assets/658544/10387777/1334b9c2-6e19-11e5-9671-1248fe2b541f.png)
