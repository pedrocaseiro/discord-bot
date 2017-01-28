package main

import (
  "fmt"
  "flag"

  // "github.com/bwmarrin/discordgo"
)

var token string

// init will run before main

func init() {
  flag.StringVar(&token, "t", "", "Bot Token")
  flag.Parse()
}

func main() {
  if token == "" {
    fmt.Println("No token provided. Please run: duskbot -t <bot token>")
    return
  }
}

