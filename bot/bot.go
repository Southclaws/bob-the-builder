package bot

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"

	"github.com/Southclaws/bob-the-builder/builder"
)

type Bot struct {
	sesh *discordgo.Session
	tb   builder.TeamBuilder
}

const (
	RoleTop = "Role: Top"
	RoleJng = "Role: Jng"
	RoleMid = "Role: Mid"
	RoleSup = "Role: Sup"
	RoleAdc = "Role: Adc"
)

func New(token string, tb builder.TeamBuilder) (bot *Bot, err error) {
	bot = &Bot{
		tb: tb,
	}

	bot.sesh, err = discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}
	bot.sesh.StateEnabled = true
	bot.sesh.State.TrackVoice = true

	bot.sesh.AddHandler(func(s *discordgo.Session, m *discordgo.Ready) {
		for _, g := range m.Guilds {
			toCreate := []string{
				RoleTop,
				RoleJng,
				RoleMid,
				RoleSup,
				RoleAdc,
			}
			guildroles, err := bot.sesh.GuildRoles(g.ID)
			if err != nil {
				return
			}

			for _, r := range guildroles {
				_, err := roleFromString(r.Name)
				if err == nil {
					n := 0
					for _, x := range toCreate {
						if x == r.Name {
							toCreate[n] = x
							n++
						}
					}
					toCreate = toCreate[:n]
				}
			}
			fmt.Println("Creating the following roles:", toCreate)

			for _, r := range toCreate {
				newrole, err := bot.sesh.GuildRoleCreate(g.ID)
				if err != nil {
					return
				}
				_, err = bot.sesh.GuildRoleEdit(g.ID, newrole.ID, r, 0xae4646, false, 0, false)
				if err != nil {
					return
				}
			}
		}
	})

	bot.sesh.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Content == "/teambuilder" {
			g, err := bot.sesh.State.Guild(m.GuildID) // bot.sesh.Guild(m.GuildID)
			if err != nil {
				return
			}

			rolemapping := bot.getLeagueRolesFromDiscordRoles(m.GuildID)

			var channel string
			for _, v := range g.VoiceStates {
				if v.UserID == m.Author.ID {
					channel = v.ChannelID
				}
			}

			if channel == "" {
				fmt.Println("Failed to get channel")
				return
			}

			users := []builder.Player{}
			for _, v := range g.VoiceStates {
				if v.ChannelID == channel {
					member, err := bot.sesh.GuildMember(m.GuildID, v.UserID)
					if err != nil {
						return
					}
					leagueroles := []builder.Role{}
					for _, r := range member.Roles {
						if leaguerole, ok := rolemapping[r]; ok {
							leagueroles = append(leagueroles, leaguerole)
						}
					}

					users = append(users, &Player{
						id:    v.UserID,
						roles: leagueroles,
					})
				}
			}

			if err := bot.buildTeam(m.ChannelID, users); err != nil {
				_, err := bot.sesh.ChannelMessageSend(m.ChannelID, "An error occurred:"+err.Error())
				if err != nil {
					fmt.Println("failed to send error:", err)
				}
			}
		}
	})

	return
}

func (b *Bot) Start(ctx context.Context) error {
	if err := b.sesh.Open(); err != nil {
		return err
	}
	defer b.sesh.Close()

	<-ctx.Done()

	return nil
}
