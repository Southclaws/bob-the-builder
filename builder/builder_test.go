package builder

import (
	"fmt"
	"testing"
)

type testplayer struct {
	name  string
	roles []Role
}

func (t *testplayer) GetRoles() []Role {
	return t.roles
}
func (t *testplayer) GetID() string {
	return t.name
}
func (t *testplayer) String() string {
	return t.name
}

var players = []Player{
	&testplayer{"@SeveredSin", []Role{LeagueRoleTop, LeagueRoleJng}},
	&testplayer{"@Sweilous", []Role{LeagueRoleTop, LeagueRoleMid}},
	&testplayer{"@bis", []Role{LeagueRoleTop, LeagueRoleMid}},
	&testplayer{"@Sullies ⚔🌪", []Role{LeagueRoleTop, LeagueRoleMid, LeagueRoleJng, LeagueRoleSup, LeagueRoleAdc}},
	&testplayer{"@Virtx", []Role{LeagueRoleTop, LeagueRoleMid, LeagueRoleJng, LeagueRoleSup, LeagueRoleAdc}},
	&testplayer{"@Ben (THE HER0)", []Role{LeagueRoleTop, LeagueRoleJng, LeagueRoleAdc}},
	&testplayer{"@Southclaws", []Role{LeagueRoleSup, LeagueRoleAdc}},
	&testplayer{"@Tilly", []Role{LeagueRoleTop, LeagueRoleMid}},
	&testplayer{"@Pete", []Role{LeagueRoleTop, LeagueRoleMid}},
	&testplayer{"@The Heff", []Role{LeagueRoleTop, LeagueRoleMid, LeagueRoleJng, LeagueRoleSup, LeagueRoleAdc}},
}

var smallteam = []Player{
	&testplayer{"@SeveredSin", []Role{LeagueRoleTop, LeagueRoleJng}},
	&testplayer{"@Sweilous", []Role{LeagueRoleTop, LeagueRoleMid}},
	&testplayer{"@bis", []Role{LeagueRoleTop, LeagueRoleMid}},
	&testplayer{"@Sullies ⚔🌪", []Role{LeagueRoleTop, LeagueRoleMid, LeagueRoleJng, LeagueRoleSup, LeagueRoleAdc}},
	&testplayer{"@Virtx", []Role{LeagueRoleTop, LeagueRoleMid, LeagueRoleJng, LeagueRoleSup, LeagueRoleAdc}},
	&testplayer{"@Ben (THE HER0)", []Role{LeagueRoleTop, LeagueRoleJng, LeagueRoleAdc}},
	&testplayer{"@Southclaws", []Role{LeagueRoleSup, LeagueRoleAdc}},
}

func Test_bucketPlayers(t *testing.T) {
	result := bucketPlayers(players)
	for r, p := range result {
		fmt.Println(r, p)
	}
}

func Test_bucketPlayersSmol(t *testing.T) {
	result := bucketPlayers(smallteam)
	for r, p := range result {
		fmt.Println(r, p)
	}
}

func Test_removePlayerFromBucket(t *testing.T) {
	buckets := bucketPlayers(players)
	result := removePlayerFromBucket(buckets, "@Virtx")
	for r, p := range result {
		fmt.Println(r, p)
	}
}

func Test_countRoles(t *testing.T) {
	result := countRoles(players)
	fmt.Println(result)
}

func Test_sortRoles(t *testing.T) {
	roles := countRoles(players)
	result := sortRoles(roles)
	fmt.Println(result)
}

func Test_Build(t *testing.T) {
	b := RandomTeamBuilder{}
	printTeam(b.Build(players))
	printTeam(b.Build(players))
	printTeam(b.Build(players))
	printTeam(b.Build(players))
	printTeam(b.Build(players))
	printTeam(b.Build(players))
	printTeam(b.Build(players))
}

func Test_BuildSmol(t *testing.T) {
	b := RandomTeamBuilder{}
	tm, err := b.build(smallteam)
	if err != nil {
		fmt.Println(err)
	} else {
		printTeam(tm)
	}
}

func printTeam(teams *Teams) {
	fmt.Println("\n\nTeam 1:")
	for k, v := range teams[0] {
		fmt.Printf("%v: %v\n", k, v)
	}
	fmt.Println("Team 2:")
	for k, v := range teams[1] {
		fmt.Printf("%v: %v\n", k, v)
	}
}
