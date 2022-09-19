# genRand

this utils can convert image to asci matrix.

The algorithm uses decomposition into pixel brightness, and ASCIBrightnessMap.

# Using 

`genrang  -flags value pathToImage`

flags:
- `-comp`: Compression cooficent
  - -1: shrink to window size
  - 0: no compres
  - n>1: compress n times 
  
- `-color`: Color shift(exsposition)
  - shift num color map

- `-h\-w`: Height/widht crop resolution
  - croped to this res
 
