package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	)

/*===============================================================
 * Functions to manipulate a "field" of cells --- the main data
 * that must be managed by this program.
 *==============================================================*/
// The data stored in a single cell of a field
type Cell struct {
	kind  string
	score float64
	a int // this variable helps to get the color scheme mentioned in the paper of Nowark and May
		  // a = 1 if the previous kind of this struc was C and new kind is C
		  // a = 2 if the previous kind of this struc was D and new kind is C
		  // a = 3 if the previous kind of this struc was C and new kind is D
		  // a = 4 if the previous kind of this struc was D and new kind is D
}

// createField should create a new field of the ysize rows and xsize columns,
// so that field[r][c] gives the Cell at position (r,c).
func createField(rsize, csize int) [][]Cell {
	f := make([][]Cell, rsize)
	for i := range f {
		f[i] = make([]Cell, csize)
	}
	return f
}

// inField returns true iff (row,col) is a valid cell in the field
func inField(field [][]Cell, row, col int) bool {
	return row >= 0 && row < len(field) && col >= 0 && col < len(field[0])
}

// readFieldFromFile should open the given file and read the initial
// values for the field. The first line of the file will contain
// two space-separated integers saying how many rows and columns
// the field should have:
//    10 15
// each subsequent line will consist of a string of Cs and Ds, which
// are the initial strategies for the cells:
//    CCCCCCDDDCCCCCC
//
// If there is ever an error reading, this function should cause the
// program to quit immediately.
func readFieldFromFile(filename string) [][]Cell {
     file, err := os.Open(filename)
     if err != nil {
         fmt.Println("Error: something went wrong opening the file.")
         fmt.Println("Probably you gave the wrong filename.")
     }

     var lines []string = make([]string, 0)
     scanner := bufio.NewScanner(file)
     i := 0

     for i == 0 && scanner.Scan(){
        lines = append(lines, scanner.Text())
        i = 1
     }
     var a, b string
     fmt.Sscanf(lines[0], "%v %v", &a, &b)
     rsize,err1 := strconv.Atoi(a)
     csize,err2 := strconv.Atoi(b)

     if err1 !=nil || err2 != nil {
	 fmt.Println("ERROR: Something is not right")
     }
     lines = make([]string, 0)

     for scanner.Scan(){
         lines = append(lines, scanner.Text())  
     }
     
     field := createField(rsize,csize)
     for m:=0 ; m < rsize; m++ {
     	for n:=0; n < csize; n++ {
     		field[m][n].kind = string(lines[m][n])
     		if field[m][n].kind == "C"{
     			field[m][n].a = 1
     		}else if field[m][n].kind == "D"{
     			field[m][n].a = 4
     		}
      	}
     }
     
     return field
}

// drawField should draw a representation of the field on a canvas and save the
// canvas to a PNG file with a name given by the parameter filename.  Each cell
// in the field should be a 5-by-5 square, and cells of the "D" kind should be
// drawn red and cells of the "C" kind should be drawn blue.

// FOR FOUR COLOR SCHEME: 
// a = 1 if the previous kind of this struc was C and new kind is C, color = blue
// a = 2 if the previous kind of this struc was D and new kind is C, color = green
// a = 3 if the previous kind of this struc was C and new kind is D, color = yellow
// a = 4 if the previous kind of this struc was D and new kind is D, color = red
func drawField(field [][]Cell, filename string) {
	fmt.Println(len(field))
    length := len(field)
	width := len(field[1])
    b := CreateNewCanvas(5*(length), 5*(width))
    x:= 0.0
    y:= 0.0
    for m:=0; m < length; m++ {
    	for n:= 0; n< width; n++{
    		if field[m][n].a == 1  {
    			b = box(x,y,b)
			    b.SetFillColor(MakeColor(0,0,255)) // blue
			    b.Fill()
    		}else if field[m][n].a== 2{
    			b = box(x,y,b)
			    b.SetFillColor(MakeColor(0,255,0)) // green
			    b.Fill()
    		} else if field[m][n].a == 3{
    			b = box(x,y,b)
			    b.SetFillColor(MakeColor(255,255,0)) // yellow
			    b.Fill()
    		}else if field[m][n].a == 4 {
    			b = box(x,y,b)
			    b.SetFillColor(MakeColor(255,0,0)) // red
			    b.Fill()
    		}
    		x = x+5
    	}
    	x = x - float64(5*length)
    	y = y + 5
    }
    b.SaveToPNG(filename)

}

// This function creates a box of dimensions 5 by 5
// This function is called in the drawField function 
func box(x float64, y float64, b Canvas) Canvas{

	b.MoveTo(x,y)
    b.LineTo(x+5.0,y)
    b.LineTo(x+5.0,y+5.0)
    b.LineTo(x,y+5.0)
    b.LineTo(x,y)

    return b
}
// This function checks if the (r-1,c-1) corodinates lie inside the field or not
// If it lies outside the field, then update m,n,p,q
func check(field [][]Cell, r int, c int) (int, int, int, int){
	m := r-1
	n := c-1
	p := 3
	q := 3
    // if r-1, c-1 both <0  (Top right corner)
    // then update m = m+1, n = n+1, p & q == 2 (iterate only 2 times)
	if r-1 < 0 && c-1 < 0{
		m = m+1
		n = n+1
		p = 2
		q = 2
	}else if r-1 < 0 && c-1 >= 0{
        // This case is of left side of the field 
		m = m+1
		p = 2
	}else if c-1 < 0 && r-1 >= 0{
        // this is the case of top sied
		n = n+1
		q = 2
	}

	if c+1 > len(field[1]) -1 && r-1 < 0{
        // This is the case of Top right corner
		q = 2
		m = m+1
		p = 2
	}else if c+1 > len(field[1]) -1 && r-1 >= 0 {
        // This is the case of left side of the field
		q = 2
	}

	if r+1 > len(field)-1 && c+1 > len(field[1])-1{
        // this is the case of bottom left corner
		p = 2
		q = 2
	}else if r+1 > len(field)-1 && c+1 <= len(field[1])-1 {
        // This is the case of Bottom side of the field
		p = 2
	}else if c+1 > len(field[1])-1 && r+1 <= len(field) -1{
	   q = 2
	}else if r-1<0 && c+1 > len(field[1]) - 1{
	   // This is the case of bottom left corner
        p = 2
		n = n+1
		q = 2
	}

	return m,n,p,q
}

/*===============================================================
 * Functions to simulate the spatial games
 *==============================================================*/

// play a game between a cell of type "me" and a cell of type "them" (both me
// and them should be either "C" or "D"). This returns the reward that "me"
// gets when playing against them.

func gameBetween(me, them string, b float64) float64 {
	if me == "C" && them == "C" {
		return 1
	} else if me == "C" && them == "D" {
		return 0
	} else if me == "D" && them == "C" {
		return b
	} else if me == "D" && them == "D" {
		return 0
	} else {
		fmt.Println("type ==", me, them)
		panic("This shouldn't happen")
	}
}

// updateScores goes through every cell, and plays the Prisoner's dilema game
// with each of it's in-field nieghbors (including itself). It updates the
// score of each cell to be the sum of that cell's winnings from the game.
func updateScores(field [][]Cell, b float64){

    r := 0
    c := 0
    for r<len(field){
    	c = 0
    	for c<len(field[1]){
    		field[r][c].score = calculatescore(field,r,c,b)
    		c = c+1
    	}
    	r = r+1
    }
}
// this function calculates and returns the score at a particular cell of the field (r,c)
func calculatescore(field [][]Cell, r int,c int, b float64) float64{

	score := 0.0
    // m,n is set to the r-1 and c-1, which is then checked and updated if it is outside the field 
    // using the check function 
	m:= r-1
	n:= c-1
    // p,q is the number of times the loop m,n have to be run 
    // initial value of p,q have been set to 3,3, considering the (r,c) is inside the field 
    // (or not on the boundary of the field )
	// IF the (r,c) is at the boundary of the field then the check function will update p,q accordingly
    p:=3
	q:=3
	m,n,p,q = check (field,r,c)
	x := m
	y := n

		for m < x+p{
			n = y
			for n < y+q {
				score = score + gameBetween(field[r][c].kind,field[m][n].kind,b)
				n = n +1
			}
			m = m+1
		}

	return score
}
// updateStrategies create a new field by going through every cell (r,c), and
// looking at each of the cells in its neighborhood (including itself) and the
// setting the kind of cell (r,c) in the new field to be the kind of the
// neighbor with the largest score
func updateStrategies(field [][]Cell) [][]Cell {
    // x1 is a new field of same dimensions as "field"
    var x1 [][]Cell
    x1 = createField(len(field),len(field[1]))

    for r:=0; r<len(field); r++{
    	for c:= 0; c<len(field[1]); c++{
    		x1[r][c].kind = maxscore(field,r,c)
    		x1[r][c].a = updateA(x1[r][c].kind, field[r][c].kind)
    	}
    }
	return x1 
}
// This function updates the value of "a" of that cell (a is defined in the struct)
func updateA(x1 string, field string) int {

	if field == "C" && x1== "C" {
		return 1
	} else if field == "C" && x1 == "D" {
		return 2
	} else if field == "D" && x1 == "C" {
		return 3
	} else if field == "D" && x1 == "D" {
		return 4
	}
	return 0
}
// this function returns the "kind" of the cell which is around cell[r][c] and has the maximum score
func maxscore(field [][]Cell, r int,c int) string{
	maxstring := field[r][c].kind
	maxscore := field[r][c].score
	m, n, p, q := check(field, r, c)
	x := m
	y := n
	for m < x + p  {
		n = y
		for n < y + q {
			if field[m][n].score > maxscore {
				maxscore = field[m][n].score
				maxstring = field[m][n].kind
			}else{

			}
			n = n+1
		}
		m = m+1
	}

	return maxstring
}

// evolve takes an intial field and evolves it for nsteps according to the game
// rule. At each step, it should call "updateScores()" and the updateStrategies
func evolve(field [][]Cell, nsteps int, b float64) [][]Cell {
	updateScores(field,b)
	for i := 0; i < nsteps; i++ {
		field = updateStrategies(field)
		updateScores(field,b)
	}
	return field
}

// Implements a Spatial Games version of prisoner's dilemma. The command-line
// usage is:
//     ./spatial field_file b nsteps
// where 'field_file' is the file continaing the initial arrangment of cells, b
// is the reward for defecting against a cooperator, and nsteps is the number
// of rounds to update stategies.
//
func main() {
	// parse the command line
	if len(os.Args) != 4 {
		fmt.Println("Error: should spatial field_file b nsteps")
		return
	}

	fieldFile := os.Args[1]

	b, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil || b <= 0 {
		fmt.Println("Error: bad b parameter.")
		return
	}

	nsteps, err := strconv.Atoi(os.Args[3])
	if err != nil || nsteps < 0 {
		fmt.Println("Error: bad number of steps.")
		return
	}

    // read the field
	field := readFieldFromFile(fieldFile)
	//fmt.Println(field)
    fmt.Println("Field dimensions are:", len(field), "by", len(field[3]))
    // evolve the field for nsteps and write it as a PNG
	field = evolve(field, nsteps, b)
	drawField(field, "Prisoners.png")
}
