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
		"team01",
		"team02",
		"team03",
		"team04",
		"team05",
		"team06",
		"team07",
		"team08",
		"team09",
		"team10",
		"team11",
		"team12",
		"team13",
		"team14",
		"team15",
		"team16",
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
		fmt.Printf("%02d番目の発表チームは %v (チーム%02d)\n", i+1, t.name, t.order)
	}
}

// const { SHA3 } = require('sha3');

// const arr5 = [
//   'Tk',
//   '土鍋',
//   'sou7',
//   'matumoto1234',
//   'マヤミト',
//   'すず',
// ];

// const arr10 = [
//   'Endered',
//   '朱雀匠',
//   'eraser5th',
//   'tashiro',
//   'しんぶんぶん',
// ];

// const arr = [arr5, arr10];

// for (let q = 0; q < 2; q += 1) {
//   const hashArr = [];
//   const teamObj = {};
//   const hash = new SHA3(512);

//   for (let i = 0; i < arr[q].length; i += 1) {
//     hash.update(arr[q][i]);
//     hashArr.push(hash.digest('hex').toString());
//     teamObj[hash.digest('hex')] = arr[q][i];
//     hash.reset();
//   }

//   hashArr.sort();

//   for (let i = 0; i < hashArr.length; i += 1) {
//     if (q === 0) {
//       console.log(`学生枠5分枠発表順${i + 1}: ${teamObj[hashArr[i]]}`);
//     } else {
//       console.log(`学生枠10分枠発表順${i + 1}: ${teamObj[hashArr[i]]}`);
//     }
//   }
// }
