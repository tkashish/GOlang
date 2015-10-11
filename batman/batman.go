package main
import(
   "fmt"
    "math"
    )
func main() {
	pic := CreateNewCanvas(int(400),int(400))
    pic.SetStrokeColor(MakeColor(0,0,0))
    pic.SetLineWidth(1)
    pic.MoveTo(0,200)
    step := 0.01
    x := 0.0
    y := 200.0
    y1 := 0.0
    for i:=0 ; i<10000 ; i++ {
    //left wing
    y1 = y1 + step
    x = (y1*y1)/100
    y = 200 - y1 
    pic.LineTo(x, y)
	}

    for i:=0 ; i<6000 ; i++ {
        x = x + step
        y = 100 + (50*50 - ((x-150)*(x-150)))/50
    pic.LineTo(x,y)
    }

    for i:=0 ; i<1000 ; i++ {
        x = x + step
        y = 148 - (x - 160)*58/10
    pic.LineTo(x,y)
    }
    for i:=0 ; i<1000 ; i++ {
        x = x + step
        y = 90 + (x - 170)*58/30
    pic.LineTo(x,y)
    }

    for i:=0 ; i<4000 ; i++ {
        x = x + step
    pic.LineTo(x,y)
    }

    for i:=0 ; i<1000 ; i++ {
        x = x + step
        y = 109.333- (x - 220)*58/30
    pic.LineTo(x,y)
    }

    for i:=0 ; i<1000 ; i++ {
        x = x + step
        y = 90 + (x - 230)*58/10
    pic.LineTo(x,y)
    }

    for i:=0 ; i<6000 ; i++ {
        x = x + step
        y = 148 +(10*10 - ((x-250)*(x-250)))/50
    pic.LineTo(x,y)
    }

    pic.MoveTo(400,200)
    y1 = 0
    for i:=0 ; i<10000 ; i++ {
    //left wing
    y1 = y1 + step
    x = 400 - (y1*y1)/100
    y = 200 - y1 
    pic.LineTo(x, y)
    }
    pic.MoveTo(400,200)
    // lower half    
    y1 = 0
    for i:=0 ; i<8360 ; i++ {
    //left wing
    y1 = y1 + step
    x = 400 - (y1*y1)/100
    y = 200 + y1 
    pic.LineTo(x, y)
    }

//  330.1103999999902 283.6000000000059   
    for i:=0 ; i<6500 ; i++ {
        x = x-step
        y = 283.0 - math.Sqrt(32.5*32.5 - (float64(i)*step - 32.5)*(float64(i)*step - 32.5))
        pic.LineTo(x, y)
    }
    //265.1104000000493 282.1938362449227
      
    for i:=0 ; i<6500 ; i++ {
        x = x - step
        y = 282.2 - (30*30 - (x - 235)*(x - 235))/25
        pic.LineTo(x, y)
    }  
    //200.11040000010843 294.8913675260973
    for i:=0 ; i<6500 ; i++ {
        x = x - step
        y = 294.89 - (35*35 - (x - 165)*(x - 165))/25
        pic.LineTo(x, y)
    }
    //135.11040000016754 281.62552752599936
    for i:=0 ; i<6500 ; i++ {
        x = x-step
        y = 281.6255 - math.Sqrt(32.5*32.5 - (float64(i)*step - 32.5)*(float64(i)*step - 32.5))
        pic.LineTo(x, y)
    }
    //70.1104000001444 280.8193362449227
    pic.MoveTo(0,200)
    // lower half    
    y1 = 0
    for i:=0 ; i<8370 ; i++ {
    //left wing
    y1 = y1 + step
    x = (y1*y1)/100
    y = 200 + y1 
    pic.LineTo(x, y)
    }
    fmt.Println(x,y)
    
	pic.Stroke()
    pic.SaveToPNG("batman.png")
}