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
        if len(os.Args) > 3 {
          src.AddLink(os.Args[2], os.Args[3])
          os.Exit(0)
        }
        fmt.Println("You have to specify file's path and the link's name. Type <dotman --help> for more information about this command.")
        os.Exit(1)
        break
      case "show":
        src.ShowLinks()
        os.Exit(0)
        break
      case "remove":
        if len(os.Args) > 2 {
          src.RemoveLink(os.Args[2])
          os.Exit(0)
        }
        fmt.Println("You have to specify the link's name. Type <dotman --help> for more information about this command.")
        os.Exit(1)
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
