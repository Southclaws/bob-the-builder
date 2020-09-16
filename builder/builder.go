package builder

import (
	"errors"
	"math/rand"
	"sort"
)

type Role int

const (
	LeagueRoleTop Role = iota
	LeagueRoleJng Role = iota
	LeagueRoleMid Role = iota
	LeagueRoleSup Role = iota
	LeagueRoleAdc Role = iota
)

func createBucketsPlayers() map[Role][]Player {
	return map[Role][]Player{
		LeagueRoleTop: {},
		LeagueRoleJng: {},
		LeagueRoleMid: {},
		LeagueRoleSup: {},
		LeagueRoleAdc: {},
	}
}
func createBucketsCounts() map[Role]int {
	return map[Role]int{
		LeagueRoleTop: 0,
		LeagueRoleJng: 0,
		LeagueRoleMid: 0,
		LeagueRoleSup: 0,
		LeagueRoleAdc: 0,
	}
}

func (r Role) String() string {
	switch r {
	case LeagueRoleTop:
		return "Top"
	case LeagueRoleJng:
		return "Jng"
	case LeagueRoleMid:
		return "Mid"
	case LeagueRoleSup:
		return "Sup"
	case LeagueRoleAdc:
		return "Adc"
	default:
		return "Unknown role"
	}
}

type Player interface {
	GetID() string
	GetRoles() []Role
}

type Lobby interface {
	GetPlayers() []Player
}

type TeamBuilder interface {
	Build([]Player) *Teams
}

type Team map[Role]Player

type Teams [2]Team

// rando team builder

type RandomTeamBuilder struct {
}

// create a bucket for each role and place players in buckets (duplicates allowed)
// create a list of roles, sorted smallest first by number of players in each
// for each role, starting with smallest:
// - pick two players from the role's bucket at random
// - remove those two players from all other buckets

func (b *RandomTeamBuilder) Build(players []Player) *Teams {
	var teams *Teams
	var err error

	for i := 0; i < 50; i++ {
		teams, err = b.build(players)
		if err == nil {
			break
		}
	}
	return teams
}

func (b *RandomTeamBuilder) build(players []Player) (*Teams, error) {
	buckets := bucketPlayers(players)
	counts := countRoles(players)
	sorted := sortRoles(counts)

	red := make(Team)
	blu := make(Team)
	var redplayer Player
	var bluplayer Player

	var err error
	for _, role := range sorted {
		buckets, redplayer, err = pickAndRemove(buckets, role)
		if err != nil {
			continue
		}
		red[role] = redplayer

		buckets, bluplayer, err = pickAndRemove(buckets, role)
		if err != nil {
			continue
		}
		blu[role] = bluplayer
	}

	return &Teams{red, blu}, nil

}

func bucketPlayers(players []Player) map[Role][]Player {
	buckets := createBucketsPlayers()
	for _, p := range players {
		roles := p.GetRoles()
		for _, r := range roles {
			buckets[r] = append(buckets[r], p)
		}
	}
	return buckets
}

func removePlayerFromBucket(buckets map[Role][]Player, id string) map[Role][]Player {
	newBuckets := createBucketsPlayers()
	for role, players := range buckets {
		for _, i := range players {
			if i.GetID() == id {
				continue
			}
			newBuckets[role] = append(newBuckets[role], i)
		}
	}
	return newBuckets
}

func countRoles(players []Player) map[Role]int {
	counts := createBucketsCounts()
	for _, i := range players {
		for _, r := range i.GetRoles() {
			counts[r]++
		}
	}
	return counts
}

type rolecount struct {
	r Role
	c int
}
type rolecounts []rolecount

func (a rolecounts) Len() int           { return len(a) }
func (a rolecounts) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a rolecounts) Less(i, j int) bool { return a[i].c < a[j].c }

func sortRoles(roles map[Role]int) []Role {
	var counts rolecounts
	for r, c := range roles {
		counts = append(counts, rolecount{r, c})
	}
	sort.Sort(counts)
	sorted := make([]Role, 5)
	for i, r := range counts {
		sorted[i] = r.r
	}
	return sorted
}

func pickAndRemove(buckets map[Role][]Player, role Role) (map[Role][]Player, Player, error) {
	players := buckets[role]
	n := len(players)
	if n == 0 {
		return buckets, nil, errors.New("players pool empty")
	}

	idx := rand.Intn(n)
	player := players[idx]
	newbuckets := removePlayerFromBucket(buckets, player.GetID())
	return newbuckets, player, nil
}
