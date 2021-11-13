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
        fmt.Println("Dotman is a tool for managing dotfiles.")
        fmt.Println("\nUsage:\n\n\tdotman <command> [arguments]")
        fmt.Println("\nThe commands are:")
        fmt.Println("\n\tadd <file's path> <link's name>  add a specified file to .dotfiles directory and dotfiles database.")
        fmt.Println("\tremove <link's name>             remove the specified link from .dotfiles directory and dotfiles database.")
        fmt.Println("\tshaw                             shaw all the links from the dotfiles database.")
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
