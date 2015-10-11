package main

import (
    "fmt"
    "os"
    "strconv"
)


// Write the current field out to the screen
func printField(b Canvas,field []bool , j int) {
    i := 0  
    for _, v := range field {
        x := i*5
        y := j*5

        draw(b,v,float64(x),float64(y))        
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

// run the CA for the given number of steps, staring from field
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

    board.SaveToPNG("ca.png")    
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

func main() {
    // parse the command line
    if len(os.Args) != 4 {
        fmt.Println("Error: should be ca rule width nsteps")
        return
    }

    // extract the rule:
    rule := os.Args[1]

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

    // get the width
    width, err := strconv.Atoi(os.Args[2])
    if err != nil || width <= 0 {
        fmt.Println("Error: width should be a positive integer")
        return
    }

    // get the number of steps
    nsteps, err := strconv.Atoi(os.Args[3])
    if err != nil || nsteps <= 0 {
        fmt.Println("Error: number of steps should be a positive integer")
        return
    }

    // initialize the field
    start := make([]bool, width)
    start[width / 2] = true

    CA(rule, start, nsteps)
}

func draw(b Canvas,v bool , x float64,y float64) {
    if v {
    b.MoveTo(x,y)
    b.LineTo(x+5.0,y)
    b.LineTo(x+5.0,y+5.0)
    b.LineTo(x,y+5.0)
    b.LineTo(x,y)
    b.SetFillColor(MakeColor(0,0,0))
    b.Fill()
    }else{
    b.MoveTo(x,y)
    b.LineTo(x+5.0,y)
    b.LineTo(x+5.0,y+5.0)
    b.LineTo(x,y+5.0)
    b.LineTo(x,y)
    b.SetFillColor(MakeColor(255,255,0))
    b.Fill()

    }
}
