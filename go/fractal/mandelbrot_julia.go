/*
 * Author: Gananath R (2020)
 * 
 * A fractal generating code for mandelbrot and julia set
 * 
 * Julia set is similar to mandelbrot except the c is a constant complex equation and z iterates from different location in a complex plane
 * 
 * complex plane{x=real;y=imaginary}
 */


package main

import (
    "os"
    "image"
    "image/color"
    "image/png"
    "math/cmplx"
    //"fmt"
)


func main(){

    width := 300
    height := 300
    max_iter := 30
    
    var img = image.NewRGBA(image.Rect(0, 0, width, height))

    for x := 0; x < width; x++ {
        for y := 0; y < height; y++ {
            
            // Mandelbrot Set
            c_re := normalize(float64(x),0,float64(width),-2,2)
            c_im := normalize(float64(y),0,float64(height),-2,2)
            z := complex(0,0)
            
            // Julia Set
            // Uncomment below lines to u
			      //c_re := -0.8 
			      //c_im :=  0.156
            //z_re := normalize(float64(x),0,float64(width),-2,2)
            //z_im := normalize(float64(y),0,float64(height),-2,2)
            //z := complex(z_re,z_im)
            
            c := complex(c_re,c_im)
            n:= mandelbrot_julia(c,z,max_iter)
            if n == max_iter{
                img.Set(x, y, color.Black)                
            }else{
				col :=uint8(255 - int(n * 255 / max_iter))
				colr := color.RGBA{col, col, col, 255}
                img.Set(x, y, colr)
            }
            
        }
    }
	// Encode as PNG.
	f, _ := os.Create("image.png")
	png.Encode(f, img)
 
}

func mandelbrot_julia(c complex128,z complex128,max_iter int)int{
    for i:=0;i<max_iter;i++{
        z = z*z+c
        if cmplx.Abs(z)>2{
            return i
        }
    }
    return max_iter
    
}

func normalize(x float64,x_min float64,x_max float64,y_min float64,y_max float64) float64{
	// https://stats.stackexchange.com/a/281165
	return (y_max-y_min)*((x-x_min)/x_max-x_min)+y_min
}

