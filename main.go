package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"rpgmaniac-discord-bot/config"
	"rpgmaniac-discord-bot/dice"
	"rpgmaniac-discord-bot/structs"

	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	Token       string
	CurrentGame structs.Game
)

func init() {

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	err := config.ReadConfig()
	CurrentGame = structs.DetermineGame(config.Game)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)

	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Set status and current game
	dg.Identify.Presence.Game.Name = string(CurrentGame)
	dg.Identify.Presence.Game.Type = discordgo.ActivityTypeGame
	dg.Identify.Presence.Status = string(discordgo.StatusDoNotDisturb)

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	channelName, _ := channelName(s, m)

	// If the message is "ping" reply with "Pong!"
	if dice.IsDiceMessage(m.Content) && channelName == config.Channel {
		s.ChannelMessageSend(m.ChannelID, dice.DescribeDices(m.Content))
	}

}

// TODO: refactor to separate package
func channelName(s *discordgo.Session, m *discordgo.MessageCreate) (string, error) {
	chName, err := s.State.Channel(m.ChannelID)

	if err != nil {
		chName, err = s.Channel(m.ChannelID)
	}
	return chName.Name, err
}
