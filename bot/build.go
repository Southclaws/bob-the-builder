package bot

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/Southclaws/bob-the-builder/builder"
)

func (b *Bot) buildTeam(channel string, players []builder.Player) error {
	team := b.tb.Build(players)

	if team == nil {
		_, err := b.sesh.ChannelMessageSendEmbed(channel, &discordgo.MessageEmbed{
			Title:       "Team builder result",
			Type:        discordgo.EmbedTypeRich,
			Description: "Could not build a team! Current player pool needs a more diverse set of roles.",
		})
		return err
	}

	_, err := b.sesh.ChannelMessageSendEmbed(channel, &discordgo.MessageEmbed{
		Title: "Team builder result",
		Type:  discordgo.EmbedTypeRich,
		Fields: []*discordgo.MessageEmbedField{
			{Name: "Team 1", Value: format(team[0]), Inline: true},
			{Name: "Team 2", Value: format(team[1]), Inline: true},
		},
	})
	return err
}

func format(t builder.Team) string {
	if len(t) == 0 {
		return "(empty)"
	}
	s := strings.Builder{}
	for k, v := range t {
		s.WriteString(fmt.Sprintf("> `%s`: <@%s>\n", k, v.GetID()))
	}
	return s.String()
}
