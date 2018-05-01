package main

import (
	"./core"	
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"encoding/json"
	"io/ioutil"
	"github.com/bwmarrin/discordgo"
)

var (
	Token string
	config *configStruct
)

type configStruct struct {
	Token     string `json:"Token"`
}

func ReadConfig() error {
	fmt.Println("Reading config file...")
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	Token = config.Token
	return nil
}

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	var err error
	var dg *discordgo.Session
	err = ReadConfig()
	if err != nil {
		fmt.Println(err)
	}
	dg, err = discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(core.MessageCreate)
	dg.AddHandler(core.MessageUpdate)
	err = dg.Open()

	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Connected to Database!!")
	fmt.Println(`               			  
   ____  ___       _      
  / __ \/ (_)   __(_)___ _
 / / / / / / | / / / __ '/
/ /_/ / / /| |/ / / /_/ / 
\____/_/_/ |___/_/\__,_/   
	   `)
	fmt.Println("Olivia Alpha V4 is now running!")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}