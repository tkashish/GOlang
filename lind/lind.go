package main

import (
    "fmt"
    "strconv"
    "os"
    "math"
)


/* Goal:

repeatedly apply all the substitutions in the rules in parallel.

*/
func getRhsFor(char string, lhs, rhs []string) (string, bool) {
	for i, l := range lhs {
		if l == char {
			return rhs[i], true
		}
	}
	return "", false
}

// applyRules will take string "start" and apply each of the given rules in
// parallel, returning the new string.
func applyRules(lhs, rhs []string, start string) string {
    var out string
    for _, char := range start {
        charRhs, existed := getRhsFor(string(char), lhs, rhs)
        if existed {
            out = out + charRhs
        } else {
            out = out + string(char)
        }
    }
    return out
}


/*
    A and B: draw line forward in the direction you’re facing
-: turn right by 60°
+: turn left by 60°
*/
func drawLindenmayer(s string) {
    const w = 10000
    const h = 10000
    pic := CreateNewCanvas(w, h)
    pic.SetStrokeColor(MakeColor(255, 0, 255))
    pic.SetLineWidth(1)
    x,y := 0.0, 0.9*h
    dir := 0.0
    step := 2.0

    pic.MoveTo(x,y)

    for _, c := range s {
        if c == 'A' || c == 'B' {
            x = x + step * math.Cos(dir)
            y = y - step * math.Sin(dir)
            pic.LineTo(x, y)

        } else if c == '+' {
            // turn left
            dir = dir + math.Pi / 3.0
        } else if c == '-' {
            // trun right
            dir = dir - math.Pi / 3.0
        } else {
            panic("Wow, somethings really wrong.")
        }
    }
    pic.Stroke()

    pic.SaveToPNG("Lind.png")
}

func main() {
    lhs := []string{"A", "B"}
    rhs := []string{"B-A-B", "A+B+A"}

    fmt.Println(lhs, rhs)

    if len(os.Args) < 2 {
        fmt.Println("Wrong.")
        return
    }

    steps, err := strconv.Atoi(os.Args[1])
    if err != nil {
        fmt.Println("Wrong.")
        return
    }

    var current string = "A"
    for i := 0; i < steps; i++ {
        current = applyRules(lhs, rhs, current)
    }

    drawLindenmayer(current)
}
