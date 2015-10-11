// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Generating random text: a Markov chain algorithm

Based on the program presented in the "Design and Implementation" chapter
of The Practice of Programming (Kernighan and Pike, Addison-Wesley 1999).
See also Computer Recreations, Scientific American 260, 122 - 125 (1989).

A Markov chain algorithm generates text by creating a statistical model of
potential textual suffixes for a given prefix. Consider this text:

	I am not a number! I am a free man!

Our Markov chain algorithm would arrange this text into this set of prefixes
and suffixes, or "chain": (This table assumes a prefix length of two words.)

	Prefix       Suffix

	"" ""        I
	"" I         am
	I am         a
	I am         not
	a free       man!
	am a         free
	am not       a
	a number!    I
	number! I    am
	not a        number!

To generate text using this table we select an initial prefix ("I am", for
example), choose one of the suffixes associated with that prefix at random
with probability determined by the input statistics ("a"),
and then create a new prefix by removing the first word from the prefix
and appending the suffix (making the new prefix is "am a"). Repeat this process
until we can't find any suffixes for the current prefix or we exceed the word
limit. (The word limit is necessary as the chain table may contain cycles.)

Our version of this program reads text from standard input, parsing it into a
Markov chain, and writes generated text to standard output.
The prefix and output lengths can be specified using the -prefix and -words
flags on the command-line.
*/
package main

import (
	"bufio"
	//"flag"
	"fmt"
	"io"
	//"math/rand"
	"os"
	"strings"
	//"time"
	"strconv"
)

// Prefix is a Markov chain prefix of one or more words.
type Prefix []string

// String returns the Prefix as a string (for use as a map key).
func (p Prefix) String() string {
	return strings.Join(p, " ")
}

// Shift removes the first word from the Prefix and appends the given word.
func (p Prefix) Shift(word string) {
	copy(p, p[1:])
	p[len(p)-1] = word
}

// Chain contains a map ("chain") of prefixes to a list of suffixes.
// A prefix is a string of prefixLen words joined with spaces.
// A suffix is a single word. A prefix can have multiple suffixes.
type Chain struct {
	chain     map[string]map[string]int
	prefixLen int
}

// NewChain returns a new Chain with prefixes of prefixLen words.
func NewChain(prefixLen int) *Chain {
	return &Chain{make(map[string]map[string]int), prefixLen}
}

// Build reads text from the provided Reader and
// parses it into prefixes and suffixes that are stored in Chain.
func (c *Chain) Build(r io.Reader) {
	br := bufio.NewReader(r)
	p := make(Prefix, c.prefixLen)
	p[0] = "' '"
	p[1] = "' '"
	for {
		var s string
		if _, err := fmt.Fscan(br, &s); err != nil {
			break
		}
		key := p.String()
		//fmt.Println(key)
		//c.chain[key] = append(c.chain[key], s)
		frequency,exists := c.chain[key]
		if !exists {
			frequency = make(map[string]int)
			c.chain[key] = frequency
		}
		c.chain[key][s]++
		//fmt.Println(c.chain[key][s])
		p.Shift(s)
	}
	//c.writetoFile(int(prefixLen),file1)
}
	// Function to write to a file

func (c *Chain) writetoFile(n int,file1 string) {
  outfile,err := os.Create(file1)
	if err!= nil {
		fmt.Println("Sorry: couldn't create the file!")
	}
	defer outfile.Close()
	c1 := n
	fmt.Fprintln(outfile,c1)
	for key1:= range c.chain {
		  fmt.Fprint(outfile,key1," ")
			for key2 := range c.chain[key1]{
				fmt.Fprint(outfile,key2," ")
				fmt.Fprint(outfile,c.chain[key1][key2]," ")
			}
	 fmt.Fprint(outfile,"\r\n")
	}
}

// Function to read from modelfile and generate the text

func (c *Chain) readFromFile(lines []string,prefixLen int){

		//key1:=make(Prefix ,r)
	  for i:=1; i< len(lines);i++ {
      var items[]string = strings.Split(lines[i]," ")
		var words []string
		  for j:=0;j<prefixLen;j++ {
				words = append(words,items[j])
			}
			c.chain[strings.Join(words, " ")] = make(map[string]int)
	    for j:=prefixLen;j<len(items)-1;j=j+2 {
						freq64,_ := strconv.ParseInt(items[j+1], 0, 64)
						freq := int(freq64)
						c.chain[strings.Join(words, " ")][items[j]]=freq
			}
	  }
}


// Generate returns a string of at most n words generated from Chain.
// func (c *Chain) Generate(n int) string {
// 	p := make(Prefix, c.prefixLen)
// 	var words []string
// 	var next string
// 	var max = 0
// 	for i := 0; i < n; i++ {
// 		//choices := c.chain[p.String()]
// 		frequency := c.chain[p.String()]
// 		if len(frequency) == 0 {
// 			break
// 		}
// 		max=0
// 		for count := range frequency {
// 			if frequency[count] >= max {
// 				max = frequency[count]
// 				next = count
// 			}
// 		}
//     words = append(words, next)
// 	  p.Shift(next)
// 	}

// 	return strings.Join(words, " ")
// }
func (c *Chain) Generate(n int) string {
	fmt.Println(c.chain)
	p := make(Prefix, c.prefixLen)
	for y := 0; y < c.prefixLen; y++{
		p[y]= "' '"
	} 
	var words []string
	for i := 0; i < n; i++ {
		fmt.Println(p.String())

		choices := c.chain[p.String()]
		fmt.Println(choices)
		if len(choices) == 0 {
			break
		}
		max := 0
		var next string

		fmt.Println(choices)
		for s3 := range choices {
			if choices[s3] >= max {
				//fmt.Println(c.chain[p.String()][s3])
				max = choices[s3]
				next = s3
			}
		}
		fmt.Println(next)
		words = append(words, next)
		p.Shift(next)
		fmt.Println(words)
	}
	return strings.Join(words, " ")
}

func main() {
	// Register command-line arguments

  length := len(os.Args)
	command1 := os.Args[1]
	if command1!= "read" && command1!= "generate" {
		fmt.Println("Enter approriate command")
		return
	}
	if command1 == "read"{
		prefixLen,err1 := strconv.ParseInt(os.Args[2],0,64)
		if err1!=nil{
			fmt.Println("Enter correct prefix Length")
		}
		c := NewChain(int(prefixLen))
	  file1:= os.Args[3]
		for i:=4;i<=(length-1);i++{
		 infile1:= os.Args[i]
		 r1,err2:= os.Open(infile1)
		 if err2!=nil{
			fmt.Println("Sorry couldn't open the file")
		 }
	   c.Build(r1)
		 c.writetoFile(int(prefixLen),file1)
		}

	}
	if command1 == "generate"{
		//file1:= os.Args[3]
		n, _:= strconv.ParseInt(os.Args[3],0,64)
		//r2,err3:= os.Open(file1)
		//if err3!=nil{
		//fmt.Println("Sorry couldn't open the file")
		//}
		file, err := os.Open(os.Args[2])
						if err != nil{
							fmt.Println("Error: Something is wrong")
							os.Exit(3)
						}
		var lines[]string = make([]string,0)
		scanner := bufio.NewScanner(file)
		var r int
		for scanner.Scan(){
			lines = append(lines, scanner.Text())
		}
			fmt.Sscanf(lines[0], "%v", &r)
			prefixLen:= r
			c:=NewChain(prefixLen)
			c.readFromFile(lines,prefixLen)
			fmt.Println(c)
      		text:= c.Generate(int(n))
			fmt.Println(text)
	}
	//numWords := flag.Int("words", 100, "maximum number of words to print")
	//prefixLen := flag.Int("prefix", 2, "prefix length in words")

	//flag.Parse()                     // Parse command-line flags.
	//rand.Seed(time.Now().UnixNano()) // Seed the random number generator.

	//c := NewChain(*prefixLen)     // Initialize a new Chain.
	//c.Build(os.Stdin)             // Build chains from standard input.
	//c.writetoFile(*prefixLen)
	//text := c.Generate(*numWords) // Generate text.
	//fmt.Println(text)             // Write text to standard output.
}
