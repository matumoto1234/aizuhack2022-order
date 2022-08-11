package main

import (
	"crypto/sha256"
	"fmt"
	"sort"
)

type Team struct {
	name  string
	order int
	hash  [sha256.Size]byte
}

func NewTeam(name string, order int, hash [sha256.Size]byte) Team {
	return Team{
		name:  name,
		order: order,
		hash:  hash,
	}
}

func main() {
	teamNames := [16]string{
		"edamame",
		"L&P",
		"一夜城",
		"チーム04",
		"LIoT",
		"ティラノがやってきたぞ！",
		"スーパーマイマイ",
		"チーズケーキ",
		"くぁwせdrftgyふじこlp",
		"🍆 or 🐊",
		"ぼっこりカマンベール",
		"コロンブス",
		"musicA",
		"分速3分メートル",
		"チームこねこね",
		"酪王",
	}

	hashs := make([][sha256.Size]byte, 16)

	for i, name := range teamNames {
		hash := sha256.Sum256([]byte(name))
		hashs[i] = hash
	}

	teams := make([]Team, 16)
	for i := range teams {
		teams[i] = NewTeam(teamNames[i], i+1, hashs[i])
	}

	sort.Slice(teams, func(i, j int) bool {
		hashStrI := fmt.Sprintf("%x", hashs[i])
		hashStrJ := fmt.Sprintf("%x", hashs[j])
		return hashStrI < hashStrJ
	})

	for i, t := range teams {
		fmt.Printf("%02d番目の発表チームは チーム%02d: %s \n", i+1, t.order, t.name)
	}
}
