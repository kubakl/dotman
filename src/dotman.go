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
    fmt.Println("A link with this name already exists.")
    os.Exit(1)
  }
  err = os.Link(filename, link)
  if err != nil {
    fmt.Println("Couldn't create the link. Make sure you specified a good file/directory name.") 
    os.Exit(1)
  }
  db.Create(&Link{LinkName: linkname, OriginalPath: filename, LinkPath: link, CreationDate: time.Now()})
}

func ShowLinks() {
  db := connect()
  db.AutoMigrate(&Link{})
  var links []Link
  db.Raw("SELECT * FROM links").Scan(&links)
  fmt.Println("[")
  for i, v := range links {
    print_links(v, i + 1 == len(links))
  }
  fmt.Println("]")
}

func RemoveLink() {
  db := connect()
  db.AutoMigrate(&Link{})
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
  fmt.Printf("    Link name: %s\n", link.LinkName)
  fmt.Printf("    Original path: %s\n", link.OriginalPath)
  fmt.Printf("    Link's path: %s\n", link.LinkPath)
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
