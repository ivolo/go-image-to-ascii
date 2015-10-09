package ascii

import (
  "bytes"
  "image"
  "strconv"
  "github.com/ivolo/go-x256"
  "github.com/mgutz/ansi"
)

// thanks to IonicaBizau/image-to-ascii for algorithm
func Convert(img image.Image) (str string) {
  var buffer bytes.Buffer
  size := img.Bounds().Max;
  pixels := ".,:;i1tfLCG08@"
  precision := (255*4) / (len(pixels)-1)

  // ansi color end constant
  reset := ansi.ColorCode("reset")

  for y := 0; y < size.Y; y += 1 {
    for x := 0; x < size.X; x += 1 {
      r, g, b, a := img.At(x, y).RGBA()
      sum := r >> 8 + g >> 8 + b >> 8 + a >> 8
      pixel := pixels[int(sum) / precision]
      // find the closest ansi color code
      code := x256.ClosestCode(uint8(r >> 8), uint8(g >> 8), uint8(b >> 8))
      /// get the ansi color code
      color := ansi.ColorCode(strconv.Itoa(code))
      
      // write the pixel
      buffer.WriteString(color)
      buffer.WriteString(string(pixel))
      buffer.WriteString(reset)
    }
    buffer.WriteString("\n")
  }

  return buffer.String()
}