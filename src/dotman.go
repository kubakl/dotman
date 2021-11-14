package src

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func AddLink(filename, linkname string) {
  db := connect()
  db.AutoMigrate(&Link{})
  link := fmt.Sprintf("%s/.dotfiles/%s", os.Getenv("HOME"), filepath.Base(filename))
  _, err := os.Stat(link)
  if err == nil {
    fmt.Println("A link to this file already exists.")
    os.Exit(1)
  }
  var exists []Link
  db.Raw("SELECT * FROM links WHERE link_name = ?", linkname).Scan(&exists)
  if len(exists) == 1 {
    fmt.Println("A dotfile with this name already exists.")
    os.Exit(1)
  }
  err = os.Link(filename, link)
  if err != nil {
    fmt.Println("Couldn't create the dotfile. Make sure you specified a good file name and that it's not a directory.") 
    os.Exit(1)
  }
  fullpath, _ := filepath.Abs(filename)
  db.Create(&Link{LinkName: linkname, OriginalPath: fullpath, LinkPath: link, CreationDate: time.Now()})
}

func ShowLinks() {
  db := connect()
  db.AutoMigrate(&Link{})
  var links []Link
  db.Raw("SELECT * FROM links").Scan(&links)
  if len(links) > 0  {
    fmt.Println("[")
    for i, v := range links {
      print_links(v, i + 1 == len(links))
    }
    fmt.Println("]")
  } else {
    fmt.Println("You don't have any dotfiles yet.")
  }
}

func RemoveLink(linkname string) {
  db := connect()
  db.AutoMigrate(&Link{})

  var exists []Link
  db.Raw("SELECT * FROM links WHERE link_name = ?", linkname).Scan(&exists)
  if len(exists) == 0 {
    fmt.Println("Dotfile with this name does not exist.")
    os.Exit(1)
  }
  db.Where("link_name = ?", linkname).Delete(&exists[0])
  os.Remove(exists[0].LinkPath)
}

func create_db_file() {
  dbfile := fmt.Sprintf("%s/.dotfiles.db", os.Getenv("HOME"))
  _, err := os.Stat(dbfile)
  if err != nil {
    if errors.Is(err, os.ErrNotExist) {
      os.Create(dbfile)
    }
  }
}

func make_dot_dir() {
  dotfiles := fmt.Sprintf("%s/.dotfiles", os.Getenv("HOME"))
  _, err := os.Stat(dotfiles)
  if err != nil {
    if errors.Is(err, os.ErrNotExist) {
      os.Mkdir(fmt.Sprintf("%s/.dotfiles", os.Getenv("HOME")), os.FileMode(int(0777)))
    }
  }
}

func connect() *gorm.DB {
  make_dot_dir()
  create_db_file()

  db, err := gorm.Open(sqlite.Open(fmt.Sprintf("%s/.dotfiles.db", os.Getenv("HOME"))))
  if err != nil {
    fmt.Println("Some error occured, please try again.")
    os.Exit(1)
  }
  return db
}

func print_links(link Link, last bool) {
  fmt.Println("  {")
  fmt.Printf("    Dotfile's name: %s\n", link.LinkName)
  fmt.Printf("    Original path: %s\n", link.OriginalPath)
  fmt.Printf("    Dotfile's path: %s\n", link.LinkPath)
  fmt.Printf("    Creation date: %s\n", format_date(link.CreationDate))
  if !last {
    fmt.Println("  },")
  } else {
    fmt.Println("  }")
  }
}

func format_date(date time.Time) string {
  return fmt.Sprintf("%s/%s/%s %s:%s",strconv.Itoa(date.Day()), strconv.Itoa(int(date.Month())), strconv.Itoa(date.Year()), strconv.Itoa(date.Local().Hour()), strconv.Itoa(date.Local().Minute())) 
}

type Link struct {
  LinkName string
  OriginalPath string
  LinkPath string
  CreationDate time.Time
}
