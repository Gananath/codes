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


extern crate num;
extern crate image;

use image::{ImageBuffer, Pixel, Rgb};

struct Fractal{
    width:i32,
    height:i32,
    max_iter:i32
}

fn _print_type_of<T>(_: &T) {
    println!("{}", std::any::type_name::<T>())
}

fn main() {
    let frac = Fractal{
        width:300,
        height:300,
        max_iter:30
    };
    let mb = 0; // 0 for mandelbrot and 1 for julia set
    let zoom = 1.0 as f64; // zoom parameter
    let mut img = ImageBuffer::new(frac.width as u32, frac.height as u32);
    let mut c_re:f64;
    let mut c_im:f64;
    let mut z:num::complex::Complex<f64>; 
    for x in 0..frac.width{
        for y in 0..frac.height{
            if mb==0{
                // Mandelbrot Set
                c_re = (1.0/zoom)*normalize(x as f64,0.0, frac.width as f64,-2.5,2.5);
                c_im = (1.0/zoom)*normalize(y as f64,0.0, frac.height as f64,-2.5,2.5);
                z = num::complex::Complex::new(0.0, 0.0);
            }else{
                				// Julia Set
				// wikipedia.org/wiki/Julia_set
				c_re = -0.8; 
				c_im =  0.156;
                let z_re = (1.0/zoom)*normalize(x as f64,0.0, frac.width as f64,-2.5,2.5);
                let z_im = (1.0/zoom)*normalize(y as f64,0.0, frac.height as f64,-2.5,2.5);
                z = num::complex::Complex::new(z_re, z_im);
            }
            let c = num::complex::Complex::new(c_re, c_im);
            
            let n = mandelbrot_julia(c,z,frac.max_iter);
            
            let pixel;
            if n==frac.max_iter{
                pixel = Rgb::from_channels(0, 0, 0, 0);
            }else{
                let col =255 - (n  * 255 / frac.max_iter);
                let colr = col as u8;
                pixel = Rgb::from_channels(colr, colr, colr, 255);
            }
            

            img.put_pixel(x as u32, y as u32, pixel);
           
           
        }
    }
    img.save("fractal.png");
}


fn normalize(x:f64,x_min:f64,x_max:f64,y_min:f64,y_max:f64) ->f64{
	// https://stats.stackexchange.com/a/281165
	(((y_max-y_min)*((x-x_min)/x_max-x_min))+y_min)
}

fn mandelbrot_julia(mut c:num::complex::Complex<f64>,mut z:num::complex::Complex<f64>,max_iter:i32)->i32{
    for i in 0..max_iter{
        z = z*z+c;
        if z.norm()>2.0{
            return i;
        }
    }
    return max_iter;
    
}


