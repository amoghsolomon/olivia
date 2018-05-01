package core

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

type message struct {
    ID  string
    ChannelID string   
}
var (
	botPrefix string = "olivia "
	messages = make(map[string]chan message)
	responses = make(map[string] string)
)

func messageHandler(session *discordgo.Session, content string, contentWithMentionsReplaced string, authorID string, authorUsername string, authorDiscriminator string, channelID string, messageID string, updateMessage bool) {
	if strings.HasPrefix(content, botPrefix) {
		session.ChannelTyping(channelID) // Send a typing event

		cmdMsg := strings.Replace(content, botPrefix, "", 1)

		cmd := strings.Split(cmdMsg, " ")

		switch cmd[0] {

			case "help":

				session.ChannelMessageSend(channelID, "```Olivia alpha. I am a bot created by Amogh. You can add custom responses to a specific word."+
					"To add a custom response. "+
					"wcustom=insert command:insert response```")
				session.ChannelMessageSend(channelID,"```Example: wcustom=idk:idk works!"+
					  "You:olivia idk Olivia:idk works!```")

			case "insult":

				mention := strings.Split(contentWithMentionsReplaced, "@")
				if (len(mention) >= 2){
					session.ChannelMessageSend(channelID, mention[1] +" "+ oliviainsult())
				} else {
					session.ChannelMessageSend(channelID, oliviainsult())
				}

			case "add":

				content = strings.Replace(content, "olivia add", "", 1)
				content = strings.TrimSpace(content)
				if strings.Contains(content, "=") == true{
				var stuff []string = strings.Split(content, "=")
				var cmd string = strings.TrimSpace(stuff[0])
				var res string = strings.TrimSpace(stuff[1])
				var status string = oliviasend(authorUsername, cmd, res)
				session.ChannelMessageSend(channelID, status)
				} else {
					session.ChannelMessageSend(channelID, "Assign response to command")
				}

			default:
				replacer := strings.NewReplacer(",", "", ".", "", ";", "")
				content = replacer.Replace(content)
				words := strings.Fields(content)
				if len(words) >= 2 {
					words = append(words[:0], words[1:]...)
					for _, word := range words {
						word = strings.ToLower(word)
						var fetchval string = oliviafetch(word)
						if fetchval != "Command not assigned" {
							session.ChannelMessageSend(channelID, fetchval)
						} else {
							session.ChannelMessageSend(channelID, "Command not assigned")
						}
					}
				}
		}
	}
}

func MessageUpdate(s *discordgo.Session, m *discordgo.MessageUpdate) {
	if m.Content == "" {
		return //No need to continue if there's no message
	}
	if (m.Author.ID == s.State.User.ID || m.Author.ID == "" || m.Author.Username == "") {
		return //Don't want the bot to reply to itself or to thin air
	}
	if m.ChannelID == "" {
		return //Where did this message even come from!?
	}
	contentWithMentionsReplaced := m.ContentWithMentionsReplaced()
	doesMessageExist := false
	for _, v := range messages {
		for obj := range v {
			if (obj.ChannelID == m.ChannelID && obj.ID == m.ID) {
				doesMessageExist = true
				break
			}
		}
		if doesMessageExist {
			break
		} else {
			return
		}
	}
	go messageHandler(s, m.Content, contentWithMentionsReplaced, m.Author.ID, m.Author.Username, m.Author.Discriminator, m.ChannelID, m.ID, true)
}
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == "" {
		return //No need to continue if there's no message
	}
	if (m.Author.ID == s.State.User.ID || m.Author.ID == "" || m.Author.Username == "") {
		return //Don't want the bot to reply to itself or to thin air
	}
	if m.ChannelID == "" {
		return //Where did this message even come from!?
	}
	contentWithMentionsReplaced := m.ContentWithMentionsReplaced()
	if messages[m.ChannelID] == nil {
		messages[m.ChannelID] = make(chan message, 0)
	}
	go func() {
		messages[m.ChannelID] <- message{ID:m.ID, ChannelID:m.ChannelID}
	}()
	go messageHandler(s, m.Content, contentWithMentionsReplaced, m.Author.ID, m.Author.Username, m.Author.Discriminator, m.ChannelID, m.ID, false)
}

func guildDetails(channelID string, s *discordgo.Session) (*discordgo.Guild, error) {
	channelInGuild, err := s.State.Channel(channelID)
	if err != nil {
		return nil, err
	}
	guildDetails, err := s.State.Guild(channelInGuild.GuildID)
	if err != nil {
		return nil, err
	}
	return guildDetails, nil
}
func channelDetails(channelID string, s *discordgo.Session) (*discordgo.Channel, error) {
	channelInGuild, err := s.State.Channel(channelID)
	if err != nil {
		return nil, err
	}
	return channelInGuild, nil
}