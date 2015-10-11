package main
import (
		"fmt"
		"strconv"
		"os"
)

type board struct {
	b [][]int
}

func (B *board) NumCols() int{
	return len(B.b[1])
}

func (B *board) NumRows() int{
	return len(B.b)
} 

func (B *board) Cell(r,c int) int{
	return B.b[r][c]
} 

func (B *board) Set(r,c, value int) {
	B.b[r][c] = value
} 

// this function checks if the cell(r,c) lies inside the board
func (B *board) Contains(r,c int) bool{
	if r > B.NumRows() - 1 || r < 0{
		return false
	}else if c > B.NumCols() - 1 || c < 0{
		return false
	}else{
		return true
	}
}
// this function creats and return the board using the input from the user
func CreateBoard() *board {
	var board board

	size,err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("ERROR")
	} 
	pile,err1 := strconv.Atoi(os.Args[2])
	if err1 != nil {
		fmt.Println("ERROR")
	} 

	length := size
	width := size
	board.b = make([][]int, length)
	for i:=0; i<length; i++{
		board.b[i] = make([]int, width)
	}
	board.b[(length)/2][(width)/2] = pile
	return &board
}

// This function checks if the value in all the cell of the board is less than 4
// by going through all the cells
func (B *board) IsConverged() bool {
	for i:=0; i<len(B.b); i++{
		for j:=0; j<len(B.b); j++{
			if B.b[i][j] > 4 || B.b[i][j] == 4 {
				return false
			} 
		}
	}
	return true
}
// This function takes in the coordinates of the cell and checks if the value of that cell is less than 4 or not
// If not then it dicrease its value until its value is less than 4 and it increments the value of the adjacent cells acordingly 
// After that, it check the value of the adjacent cells and topples them if their value of grater than 3 by recursively calling itself.
// This makes sure that value of the cells adjacent to cell(r,c) is less than 4 
func (B *board) Topple(r,c int) {
	value := B.Cell(r,c)
	value1 := value / 4
	value2 := value%4
	if value > 4 || value == 4{
		//B.Topple(r,c)
		B.Set(r,c,value2)
		if B.Contains(r,c-1){ // check the cell on the north of cell(r,c)
			B.Set(r,c-1,B.Cell(r,c-1) + value1)
			n := B.Cell(r,c-1)
			if n > 4 || n == 4 {
				B.Topple(r,c-1)
			}	
		}
		if B.Contains(r,c+1){ // check the cell on the south of cell(r,c)
			B.Set(r,c+1,B.Cell(r,c+1) + value1)
			s := B.Cell(r,c+1)
			if s > 4 || s == 4 {
				B.Topple(r,c+1)
			}
		}
		if B.Contains(r-1,c){ // check the cell on the west of cell(r,c)
			B.Set(r-1,c,B.Cell(r-1,c) + value1)
			w := B.Cell(r-1,c)
			if w > 4 || w == 4 {
				B.Topple(r-1,c)
			}
		}
		if B.Contains(r+1,c){ // check the cell on the east of cell(r,c)
			B.Set(r+1,c,B.Cell(r+1,c) + value1)
			e := B.Cell(r+1,c)
			if e > 4 || e == 4 {
				B.Topple(r+1,c)
			}
		}
	}else{
	return
	}
}

// this function calls the function Topple if the value of the cell (size/2, size/2) is more than or equal to 4
func ComputeSteadyState(b *board) {
	r := (len(b.b))/2
	c := (len(b.b[1]))/2
	value := b.Cell(r,c)
	if value > 4 || value == 4{
		b.Topple(r,c)
	}
	
}
// this function Draws board 
func (B *board) DrawBoard(){
	length := len(B.b)
	width := len(B.b)
    b := CreateNewCanvas(length, width)
    x:= 0.0
    y:= 0.0
    b.SetFillColor(MakeColor(0,0,0))
    b.Clear()
    for m:=0; m < length; m++ {
    	for n:= 0; n< width; n++{
    		if B.b[m][n] == 1{
    			b = box(x,y,b)
			    b.SetFillColor(MakeColor(85,85,85))
			    b.FillStroke()
    		} else if B.b[m][n] == 2{
    			b = box(x,y,b)
			    b.SetFillColor(MakeColor(170, 170, 170))
			    b.FillStroke()
    		}else if B.b[m][n] == 3 {
    			b = box(x,y,b)
			    b.SetFillColor(MakeColor(255,255,255))
			    b.FillStroke()
    		}
    		x = x+1.0
    	}
    	x = x - float64(1*length)
    	y = y + 1.0
    }
    b.SaveToPNG("board.png")
}
// this funciton draws a box of size 1x1
func box(x float64, y float64, b Canvas) Canvas{

	b.MoveTo(x,y)
    b.LineTo(x+1.0,y)
    b.LineTo(x+1.0,y+1.0)
    b.LineTo(x,y+1.0)
    b.LineTo(x,y)
    b.SetLineWidth(0.0)
    return b
}

func main() {
	board := CreateBoard()	
	ComputeSteadyState(board)
	if board.IsConverged() {
		fmt.Println("The Program converged")
	}else{

		fmt.Println("The Program did not converge")
	}
	board.DrawBoard()
}