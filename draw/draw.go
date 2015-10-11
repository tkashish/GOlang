package main
import (
       "fmt"
       "os"
       "strconv"
        "math/rand"
        "math"
       )
func main() {
//============================================PopSize=================================================//   
    max_t:= 100
    x_0 := 0.1
    r,err := strconv.ParseFloat(os.Args[1],64)
    if err != nil {
        fmt.Println("ERROR")
    }else{
    PopSize(r,x_0,max_t)
    }
//=========================================Rand walk===================================================//
    width:= 500.0
    height:= 500.0 
    stepsize, err3 := strconv.ParseFloat(os.Args[2],64)
    nos:= 1000
    seed:= 12345
    if err3!=nil {
        fmt.Println("Error: Please enter a number for Step Size ")
    }else if stepsize<0 {
        fmt.Println("Error: Value of stepsize is negative ",stepsize)
    }else{
    x := width/2
    y := height/2
    rand.Seed(int64(seed))

    pic := CreateNewCanvas(int(width),int(height))
    pic.SetStrokeColor(MakeColor(0,0,225))
    pic.SetLineWidth(1)
    pic.MoveTo(x,y)
    for i:=1;i<=nos;i++{
        // to get the new position
        x,y = randmove(x,y,stepsize,width,height)
        // to draw line to new position
        pic.LineTo(x,y)
    }
    dist1 := float32(distance(width/2,x,height/2,y))
    fmt.Println("Distance = ", dist1)
    pic.Stroke()
    pic.SaveToPNG("RandomWalk.png")
    }


//====================================================CA==================================================//

    // extract the rule:
    rule := os.Args[3]

    // if the rule is small, try to interpret it as an integer
    if len(rule) <= 3 {
        var good bool
        rule, good = numberToRule(rule)
        if !good {
            fmt.Println("Error: a rule should be 8 letters long or a number between 0 and 255")
            return
        }

    // if the rule is 8 characters, interpret it as a 0/1 string
    } else if len(rule) == 8 {
        if !checkRule(rule) {
            fmt.Println("Error: a rule should contain 0s and 1")
            return
        }
    // otherwise, it's a mistake
    } else {
        fmt.Println("Error: a rule should be 8 letters long or a number between 0 and 255")
        return
    }

    // givin the width
    width_ca:= 100

    // givin the number of steps
    nsteps:= 50

    // initialize the field
    start := make([]bool, width_ca)
    start[width_ca/ 2] = true
    // calling the funciton 
    CA(rule, start, nsteps)

//=====================================My Cool Picture=============================================//    
    // Function for Extra Credit Question
    batman()

}
func batman() {
    pic := CreateNewCanvas(int(400),int(400)) 
    pic.SetStrokeColor(MakeColor(0,0,0))
    pic.SetLineWidth(1)
    pic.MoveTo(0,200)
    step := 0.01 // Step size is kept small to get a smooth curve
    x := 0.0
    y := 200.0
    y1 := 0.0

    //Below are series of function to get the curves in the batman figure
    // 15 for loops have been used for 15 different curves in the figure
    // I have used for loop as the step size is small and has to be iterated to get x,y coordinates accoding to the function defination of the curve

// For the Left wing
        for i:=0 ; i<10000 ; i++ {

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
// for the Head
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
// For the right wing
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

// For the bottom Part of the picture    
    y1 = 0
    for i:=0 ; i<10000 ; i++ {
    y1 = y1 + step
    x = 400 - (y1*y1)/100
    y = 200 - y1 
    pic.LineTo(x, y)
    }
    pic.MoveTo(400,200)
    y1 = 0
    for i:=0 ; i<8360 ; i++ {
    y1 = y1 + step
    x = 400 - (y1*y1)/100
    y = 200 + y1 
    pic.LineTo(x, y)
    }  
    for i:=0 ; i<6500 ; i++ {
        x = x-step
        y = 283.0 - math.Sqrt(32.5*32.5 - (float64(i)*step - 32.5)*(float64(i)*step - 32.5))
        pic.LineTo(x, y)
    }
      
    for i:=0 ; i<6500 ; i++ {
        x = x - step
        y = 282.2 - (30*30 - (x - 235)*(x - 235))/25
        pic.LineTo(x, y)
    }  
    for i:=0 ; i<6500 ; i++ {
        x = x - step
        y = 294.89 - (35*35 - (x - 165)*(x - 165))/25
        pic.LineTo(x, y)
    }
    for i:=0 ; i<6500 ; i++ {
        x = x-step
        y = 281.6255 - math.Sqrt(32.5*32.5 - (float64(i)*step - 32.5)*(float64(i)*step - 32.5))
        pic.LineTo(x, y)
    }
    y1 = 0

    pic.MoveTo(0,200)
    
    for i:=0 ; i<8370 ; i++ {
    y1 = y1 + step
    x = (y1*y1*1.0007636649658262)/100
    y = 200 + y1 - 2.8806637550832
    pic.LineTo(x, y)
    }

    fmt.Println(x,y)
    pic.Fill()
    pic.SaveToPNG("MyCoolPicture.png")
}
//========================================================PopSize==============================================================//
func PopSize(r float64,x_0 float64 , max_t int ) {
    pic := CreateNewCanvas(500,100)
    pic.SetStrokeColor(MakeColor(0,0,225))
    pic.SetLineWidth(1)
    xt := x_0
    pic.LineTo(float64(0),100 - 100*xt) 

    for t:= 1; t<max_t; t++{
        pic.LineTo(float64(t*5),100 - 100*xt) 
        xt = r*xt*(1-xt)
        if xt<0 || xt >1{
            xt = 1.0
        }
    }

    pic.Stroke()
    pic.SaveToPNG("PopSize.png")
}

//=============================================================CA==============================================================//
// Write the current field out to the screen and calls the function draw to draw at each step
func printField(b Canvas,field []bool , j int) {
    i := 0  
    for _, v := range field {
        x := i*5
        y := j*5

        draw(b,v,float64(x),float64(y))  // calling the draw function to make a square of either yellow color(if v = " ") or black color (if v = "#")        
        if v {
            fmt.Print("#")
        } else {
            fmt.Print(" ")
        }
        i++
    }
    fmt.Println()
}

// access an item of the rule
func readRule(rule string, pos int) bool {
    if rule[pos] == '1' {
        return true
    } else {
        return false
    }
}

// return the 3 values of the field centered at i
func fieldAround(field []bool, i int) (bool,bool,bool) {
    var left, right bool
    if i > 0 {
        left = field[i-1]
    }
    if i+1 < len(field) {
        right = field[i+1]
    }
    return left, field[i], right
}

// figure out the universe a time t given field is at time t-1
func nextField(rule string, field []bool) []bool {
    out := make([]bool, len(field))
    for i := 0; i < len(field); i++ {
        l,m,r := fieldAround(field, i)
        if l && m && r {               //XXX
            out[i] = readRule(rule, 0)

        } else if l && m && !r {       //XX_
            out[i] = readRule(rule, 1)

        } else if l && !m && r {        //X_X
            out[i] = readRule(rule, 2)

        } else if l && !m && !r {        //X__
            out[i] = readRule(rule, 3)

        } else if !l && m && r {        //_XX
            out[i] = readRule(rule, 4)

        } else if !l && m && !r {        //_X_
            out[i] = readRule(rule, 5)

        } else if !l && !m && r {        //__X
            out[i] = readRule(rule, 6)

        } else if !l && !m && !r {        //___
            out[i] = readRule(rule, 7)
        }
    }
    return out
}

// run the CA for the given number of steps, starting from field
func CA(rule string, field []bool, nsteps int) {
    // for the required number of steps
    board := CreateNewCanvas(500, 255)  
    board.SetFillColor(MakeColor(0,0,0))
    i := 0
    for i = 0 ; i < nsteps; i++ {
        printField(board,field,i)
        field = nextField(rule, field)
    }
    printField(board,field,i)

    board.SaveToPNG("CA.png")    
}

// make sure the rule contians only 0s and 1s
func checkRule(r string) bool {
    for _, c := range r {
        if c != '1' && c != '0' {
            return false
        }
    }
    return true
}

// convert a number between 0 and 255 (encoded as a string) into a 0/1 binary
// string representing a rule; returns false if it fails
func numberToRule(n string) (string, bool) {
    nn, err := strconv.Atoi(n)
    if err != nil || nn < 0 || nn > 255 {
        return "", false
    }

    r := strconv.FormatInt(int64(nn), 2)
    for len(r) < 8 {
        r = "0" + r
    }
    return r, true
}
// this is the new function which draws a square of size 5 by 5 and fills it with color black (if its "#" ) and yellow (if its " ")
func draw(b Canvas,v bool , x float64,y float64) {
    if v {
    b.MoveTo(x,y)
    b.LineTo(x+5.0,y)
    b.LineTo(x+5.0,y+5.0)
    b.LineTo(x,y+5.0)
    b.LineTo(x,y)
    b.SetFillColor(MakeColor(0,0,0)) // black for " "
    b.Fill()
    }else{
    b.MoveTo(x,y)
    b.LineTo(x+5.0,y)
    b.LineTo(x+5.0,y+5.0)
    b.LineTo(x,y+5.0)
    b.LineTo(x,y)
    b.SetFillColor(MakeColor(255,255,0)) // Yellow for  "#"
    b.Fill()

    }
}
//=====================================================================random walk======================================================================//
// This function calculates distance betweeen e given points (x1,y1) & (x2,y2)
func distance(x1,x2,y1,y2 float64) float32 {
    a := math.Pow((x1-x2),2) + math.Pow((y1-y2),2)
    dist := math.Pow(a,0.5)
    return float32(dist)
}  


func randmove(x,y,stepsize,width,height float64) (float64,float64){

    delta:= 2*(math.Pi)*rand.Float64()
    //fmt.Println(delta)
    
    x1 := x+stepsize*math.Cos(delta)
    y1 := y+stepsize*math.Sin(delta)    
    
    if x1>width || x1<0 || y1 <0 || y1>height{
        x1,y1=randmove(x,y,stepsize,width,height)

        return x1,y1
    } else {
        return x1,y1
    
    }

} 