/*
 * Author: Gananath R (2020)
 * 
 * A fractal generating code for mandelbrot and julia set
 * 
 * f(z) = z^2+c
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
    "fmt"
)

var mb,w,h,mi,zoom int
type fractal struct{
	width int
	height int
	max_iter int
}

func init(){
	fmt.Print("Enter 0 for mandelbrot 1 for julia: ")
	fmt.Scanln(&mb)
	fmt.Print("Enter Width: ")
	fmt.Scanln(&w)
	fmt.Print("Enter Height: ")
	fmt.Scanln(&h)
	fmt.Print("Enter Iteration: ")
	fmt.Scanln(&mi)
	fmt.Print("Enter zoom: ")
	fmt.Scanln(&zoom)
}

func main(){
	frac := fractal{w,h,mi}
    var img = image.NewRGBA(image.Rect(0, 0, frac.width, frac.height))
	var c_re,c_im float64
	var z complex128
    for x := 0; x < frac.width; x++ {
        for y := 0; y < frac.height; y++ {
            if mb == 0{
				c_re = float64(1/zoom)*normalize(float64(x),0,float64(frac.width),-2,2)
				c_im = float64(1/zoom)*normalize(float64(y),0,float64(frac.height),-2,2)
				z = complex(0,0)
			}else if mb==1{
            // Julia Set
            // Uncomment below lines to u
				c_re = -0.8 
				c_im =  0.156
				z_re := float64(1/zoom)*normalize(float64(x),0,float64(frac.width),-2,2)
				z_im := float64(1/zoom)*normalize(float64(y),0,float64(frac.height),-2,2)
				z = complex(z_re,z_im)
			}else{
				fmt.Println("Something went wrong...exiting")
			}
            c := complex(c_re,c_im)
            n:= mandelbrot_julia(c,z,frac.max_iter)
            if n == frac.max_iter{
                img.Set(x, y, color.Black)

            }else{
				col :=uint8(255 - int(n * 255 / frac.max_iter))
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

