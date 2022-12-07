# genRand

this utils can convert image to asci matrix.

The algorithm uses decomposition into pixel brightness, and ASCIBrightnessMap.


## "Install"

1. Add your file build to `$PATH` 
2. Rename or create symlink to `gen-rand`


# Using 

`gen-rand  -flags value path_to_image`

flags:
- `-comp`: Compression cooficent
  - -1: shrink to window size
  - 0: no compres
  - n>1: compress n times 
  
- `-color`: Color shift(exsposition)
  - shift num color map

- `-h\-w`: Height/widht crop resolution
  - croped to this res

# Anim_sript

`./play_anim.sh 'file_path/url'`

Required:
	
- `ffmpeg`, `magick`, `wget` 
- Instaling gen-rand app

Script save cashe to /tmp/genRand

## Configure script

configure only ENV)):

- `SLEEP_TIME` sleep time between frame

- `COLOR_SHIFT` as well as `-color` flag in gen-rand app

