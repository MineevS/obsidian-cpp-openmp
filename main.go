package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"obsidian-cpp-openmp/util"
	//"/util"

	//"io/ioutil"
	//"flag"
)

const (
	VERSION = "0.04"
)

var (
	CppTemp       = filepath.Join(util.ExecuteDirectory, "temp.cpp")
	ExeTemp       = filepath.Join(util.ExecuteDirectory, "temp.exe")
	StringBuilder = strings.Builder{}
	CC            = ""
)

func init() {
	util.FileLogger(filepath.Join(util.ExecuteDirectory, "obsidian-cpp-openmp.log"))

	// detect c compiler

	if util.ExecuteCommandSilent("g++ --version") == 0 {
		//CC = "clang++ "
		CC = "g++ "
		return
	}

	if util.ExecuteCommandSilent("clang++ --version") == 0 {
		CC = "clang++ "
		return
	}

	println("No C++ compiler found")
	os.Exit(1)
}

func helpMessage() {
	fmt.Println("version : ", VERSION)
	fmt.Println("author : MineevS")
	fmt.Println("github : https://github.com/MineevS/obsidian-cpp-openmp")
}
func displayArgs() {
	for index, arg := range os.Args {
		println(index, arg)
	}
}

func Execute() {
	// TODO Read compiler flag and could choose compiler

	var temp_std string = "";
	var temp_file string = "";
	//var std string

	/// fmt.Println("Args_len: ", len(os.Args));
	/*
	for i := 0; i < len(os.Args); i++ {
		fmt.Println("index: ", i, "Arg_value: " ,os.Args[i]);
	}
	//*/

	if(strings.HasPrefix(os.Args[2], "-std=")){
		temp_std = os.Args[2];
	}
	/// fmt.Println("TEMP_STD: ", temp_std);

	//fmt.Println("IsLocal: ", strings.IsLocal(os.Args[4]));
	//fmt.Println("exist_file: ", os.path.exists(os.Args[4]));
	/// fmt.Println("exist_file: ", util.ExistPath(os.Args[4]));

	if(util.ExistPath(os.Args[4])){
		temp_file = os.Args[4];
	  fmt.Println("TEMP_FILE: ", temp_file);

		/*
		contents,_ := ioutil.ReadFile(temp_file)
		 println(string(contents))
		//*/

		CppTemp = temp_file;

		/// fmt.Println("ExeTemp: ", ExeTemp);

		command := CC + " " + temp_std + " " + CppTemp + " -fopenmp -o " + ExeTemp + " && " + ExeTemp
		print("\u200b")

		// fmt.Println("command: ", command);
		//*/
		util.ExecuteCommand(command)
		os.Remove(CppTemp)
		os.Remove(ExeTemp)
	} else {
			// fmt.Println("TEMP_FILE: ", temp_file);

			index := 2
			for index < len(os.Args) && !strings.HasPrefix(os.Args[index], "#") {
				index++
			}
			//std := flag.String("std", "", "")

			/// fmt.Println("index: ", index);
			/// fmt.Println("CppTemp: ", CppTemp);
			if index == len(os.Args) {
				/// fmt.Println("os.Args: ", os.Args);
				fmt.Println("Please disable the option \"Use main function\"")
				log.Fatal("out of args range")
				return
			}

			/*contents,_ := ioutil.ReadFile("")
				println(string(contents))
			//*/


			//fmt.Println("Далее!");
			// to prevent the previous program do not exit
			util.KillByName("temp.exe")
			/// fmt.Println("Temp: ", strings.Join(os.Args[index:], "\n"));
			util.CreatFile(CppTemp, strings.Join(os.Args[index:], "\n"))

			command := CC + " " + temp_std + " " + CppTemp + " -fopenmp -w -o " + ExeTemp + " && " + ExeTemp

			print("\u200b")

		  //*/
		  util.ExecuteCommand(command)
		  os.Remove(CppTemp)
		  os.Remove(ExeTemp)

		}

}

func main() {
	argc := len(os.Args)
	//displayArgs()
	switch argc {
	case 1:
		helpMessage()
	default:
		Execute()
	}
}
