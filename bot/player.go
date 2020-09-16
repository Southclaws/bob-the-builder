package bot

import (
	"errors"

	"github.com/Southclaws/bob-the-builder/builder"
)

type Player struct {
	id    string
	roles []builder.Role
}

func (p *Player) GetID() string {
	return p.id
}

func (p *Player) GetRoles() []builder.Role {
	return p.roles
}

func roleFromString(name string) (builder.Role, error) {
	switch name {
	case RoleTop:
		return builder.LeagueRoleTop, nil
	case RoleJng:
		return builder.LeagueRoleJng, nil
	case RoleMid:
		return builder.LeagueRoleMid, nil
	case RoleSup:
		return builder.LeagueRoleSup, nil
	case RoleAdc:
		return builder.LeagueRoleAdc, nil
	default:
		return builder.LeagueRoleTop, errors.New("unknown role")
	}
}

func (b *Bot) getLeagueRolesFromDiscordRoles(guild string) map[string]builder.Role {
	ids := map[string]builder.Role{}
	roles, err := b.sesh.GuildRoles(guild)
	if err != nil {
		return nil
	}
	for _, r := range roles {
		role, err := roleFromString(r.Name)
		if err == nil {
			ids[r.ID] = role
		}
	}
	return ids
}
