package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)


type Program struct{
	programType string
	programName string

}
type Touch struct{
	programName string
	
}

func (touch Touch) touchCommand(){
	file,err := os.OpenFile(touch.programName,os.O_WRONLY|os.O_CREATE,0644)
	if err != nil{
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Printf("Your file %v was successfully created",touch.programName)

}


func (prog Program) generateProgram(){

	switch prog.programType{

		// }
	case  "c":
		fmt.Println("Case c activated")
		
		file,err := os.OpenFile(prog.programName,os.O_WRONLY|os.O_CREATE,0644)
		if err != nil{
			log.Fatal(err)


		}
		defer file.Close()
		data := "#include <stdio.h>\nvoid main(){\n\nprintf(\"Hello world\");\n\n}"
		_,err = file.WriteString(data)
		if err != nil{
			log.Fatal(err)
		}


		fmt.Printf("Your  %v program was successfully created",prog.programName)

	case "go":
		file,err := os.OpenFile(prog.programName,os.O_WRONLY|os.O_CREATE,0644)
		if err != nil{
			log.Fatal(err)


		}
		defer file.Close()
		data := "package main\nimport \"fmt\"\n\nfunc main(){\n\nfmt.println(\"Hello world\")\n\n}"
		_,err = file.WriteString(data)
		if err != nil{
			log.Fatal(err)
		}		
		fmt.Printf("Your  %v program was successfully created",prog.programName)


	}

}



func main() {
	if(len(os.Args) > 1){
	mainOption := os.Args[1]
	switch mainOption{
	case "generate":
		var language,filename string

		fmt.Println("Enter the Language:(C,GO)")
		fmt.Scanln(&language)
		if(language == "C" || language== "GO"){
			fmt.Println("Enter the filename please")
			fmt.Scanln(&filename)
			

			
	 	programConfig :=  Program{programType: strings.ToLower(language),programName: filename}
		programConfig.generateProgram()


		}else{
			fmt.Println("Language incorrect or not supported select another please");
		}


	


		
	
	
	case  "touch":
	filename :=  os.Args[2]
	touchedProgram := Touch{
		programName: filename,
	}

	touchedProgram.touchCommand()
default:
		fmt.Printf("Your command was not recognized please try again invoking correct commands <generate>  or <touch> filename")

	}


	}else{
		fmt.Println("Usage generate or touch <filename>")
	}


}

	
