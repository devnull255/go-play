package main
//Runs a command passed at the command line
import (
   "fmt"
   "log"
   "os"
   "os/exec"
)

func info(msg string) {
   log.Println("INFO:",msg)
}

func main() {
  var nargs = len(os.Args)
  var msg string
  //var cmd string
  msg = fmt.Sprintf("Number of arguments: %d",nargs)
  //log.Println(msg)
  info(msg)
  if nargs > 1 {
     args := os.Args[1:]     
     msg = fmt.Sprintf(" Running command: %q\n",args)
     //log.Println(msg)
     info(msg)
     out, err := exec.Command(args[0],args[1:]...).Output()
     if err != nil {
        log.Fatal(err)
     }
     fmt.Printf("%s\n",out)
     //log.Println("Finished executing command.")
     info("Finished executing command.")

   } else {
     fmt.Println("Usage: shell_cmd.go <command [cmdarg...]>")
   }
}
