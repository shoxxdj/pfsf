package main 

import(
	"flag"
	"fmt"
	"os"
)

func main() {
	prefixFlag := flag.String("p","","Prefix value")
	suffixFlag := flag.String("s","","Suffix value")
	inputFlag  := flag.String("i","","input value (or Stdin)")

	flag.Usage = func (){
		fmt.Fprintf(os.Stderr,"PreFix SufFix a binary to add / remove prefix and/or suffix to an input string\n")
		flag.VisitAll(func(f *flag.Flag) {
			fmt.Fprintf(os.Stderr, "\t-%v: %v\n", f.Name, f.Usage) // f.Name, f.Value
		})
		fmt.Println("---- ---- ---- ---- Examples ---- ---- ---- ----")
		fmt.Println("|   pfsf -i website -p https:// -s .com        |")
		fmt.Println("|   echo 'website' | pfsf -p https:// -s .com  |")
		fmt.Println("---- ---- ---- ---- -------- ---- ---- ---- ----")
	}
	flag.Parse()

	var input string 
	var prefix string
	var suffix string 

	prefix = *prefixFlag
	suffix = *suffixFlag

	if len(*inputFlag) == 0 {
		for {
			_, err := fmt.Scanln(&input)
			fmt.Printf("%s%s%s\n",prefix,input,suffix)
			if err != nil {
				break
			}
		}
	}else{
		input = *inputFlag
		fmt.Printf("%s%s%s\n",prefix,input,suffix)
	}
}