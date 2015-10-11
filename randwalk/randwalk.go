package main
import (
	"fmt"
	"os"
	"strconv"
	"math/rand"
	"math"
)
func main() {
 	if len(os.Args) != 6 {
		fmt.Println("Error: Please enter all the values : WIDTH HEIGHT STEPSIZE NUMBER-OF-STEPS SEED")
	}else{
	width, err1 := strconv.ParseFloat(os.Args[1],64)
	height, err2 := strconv.ParseFloat(os.Args[2],64)
	stepsize, err3 := strconv.ParseFloat(os.Args[3],64)
	nos, err4 := strconv.Atoi(os.Args[4])
	seed, err5 := strconv.Atoi(os.Args[5])
    if err1!= nil {
		fmt.Println(" Error: Please enter a number for width")
	}else if width <0 {
		fmt.Println("Error: Value of width is negative ",width)
	}else if err2!=nil {
		fmt.Println("Error: Please enter a number for height ")
	}else if height<0 {
		fmt.Println("Error: Value of height is negative ",height)
	}else if err3!=nil {
		fmt.Println("Error: Please enter a number for Step Size ")
	}else if stepsize<0 {
		fmt.Println("Error: Value of stepsize is negative ",stepsize)
	}else if err4!=nil {
		fmt.Println("Error: Please enter an integer for number of steps")
	}else if nos<0 {
		fmt.Println("Error: Value of Number of Steps is negative ",nos)
	}else if err5!=nil {
		fmt.Println("Error: Please enter a number for seed")
	}else{
	x := width/2
	y := height/2
	fmt.Println(x,y)
	rand.Seed(int64(seed))

    pic := CreateNewCanvas(int(width),int(height))
    pic.SetStrokeColor(MakeColor(0,0,225))
    pic.SetLineWidth(1)
    pic.MoveTo(x,y)
	for i:=1;i<=nos;i++{
		x,y = randmove(x,y,stepsize,width,height)
		pic.LineTo(x,y)
		fmt.Println(x,y)
	}
	dist1 := float32(distance(width/2,x,height/2,y))
	fmt.Println("Distance = ", dist1)

    pic.Stroke()
    pic.SaveToPNG("randwalk.png")
	}
}
}

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