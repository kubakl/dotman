package main

import (
  "os"
  "fmt"
  src "github.com/kubakl/dotman/src"
)

func main() {
  if len(os.Args) > 1 {
    switch os.Args[1] {
      case "add":
        src.AddLink(os.Args[2], os.Args[3])
        break
      case "show":
        src.ShowLinks()
        os.Exit(0)
        break
      case "remove":
        src.RemoveLink()
        os.Exit(0)
        break
      case "--help", "-h":
        fmt.Println("Showing command details.")
        os.Exit(0)
      default:
        fmt.Printf("Couldn't recognize option %q. Type <dotman --help> for more information about this command.\n", os.Args[1])
        os.Exit(1)
    }
  } else {
    fmt.Println("Wrong command usage. You have to pass at least one subcommand. Type <dotman --help> for more information about this command.")
    os.Exit(1)
  }
}
