## Led-Matrix-UI  

WebUI to manage HUB75 LED Matrix Panels with the hzeller/rpi-rgb-led-matrix library  

Build with:  
  * Go-Gin
  * Bootstrap
  * htmx

## Usage  

Run as ***root*** to get access to the `/dev` GPIO  
Configure it through the app flags  
They follow the hzeller library  
```  
sudo ./matrix-led-ui-bin -h
```

## Build  

You'll need to install [Zig](https://ziglang.org/) and Golang from your distribution package manager.  

Clone the two repos:  
 * zaggash/go-rpi-rgb-led-matrix
 * zaggash/led-matrix-ui

Move to `./go-rpi-rgb-led-matrix/lib/rpi-rgb-led-matrix/lib` and run Make  
Then go to the `led-matrix-ui` folder and run build.sh  

```
WORKDIR=$(pwd)
git clone https://github.com/zaggash/go-rpi-rgb-led-matrix.git
git clone https://github.com/zaggash/led-matrix-ui.git
cd $WORKDIR/go-rpi-rgb-led-matrix/lib/rpi-rgb-led-matrix/lib/
make \
    CC="zig cc -target arm-linux-gnueabihf -march=arm1176jz_s -mfpu=vfp -mfloat-abi=hard" \
    CXX="zig c++ -target arm-linux-gnueabihf -march=arm1176jzf_s -mfpu=vfp -mfloat-abi=hard"

cd $WORKDIR/led-matrix-ui/
./build.sh
```

TODO
 * Split build action with matrix to build more rpi device

__Please feel free to report any issues or suggestions__
