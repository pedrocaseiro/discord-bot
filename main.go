package main

import (
  "fmt"
  "flag"

  "github.com/bwmarrin/discordgo"
)

var token string
var buffer = make([][]byte, 0)

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

  // Create a new Discord session using the provided bot token
  discordGo, err := discordgo.New(token)
  if err != nil {
    fmt.Println("Error creating Discord session: ", err)
    return
  }

  err = discordGo.Open()
  if err != nil {
    fmt.Println("Error opening discord session: ", err)
    return
  }

  fmt.Println("Bot is now running. Press CTRL-C to exit.")
  <-make(chan struct{})
  return
}

