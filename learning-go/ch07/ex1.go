package main

import (
	"io"
	"os"
	"sort"
)

type Team struct {
	Name    string
	Players []string
}

type League struct {
	Wins  map[string]int
	Teams map[string]Team
	Name  string
}

type Ranker interface {
	Ranking() []string
}

func RankPrinter(ranker Ranker, writer io.Writer) {
	results := ranker.Ranking()
	for _, v := range results {
		io.WriteString(writer, v)
		writer.Write([]byte("\n"))
	}
}

func (l *League) MatchResult(firstTeam string, firstScore int, secondTeam string, secondScore int) {
	if firstScore > secondScore {
		l.Wins[firstTeam] += 1
	} else if secondScore > firstScore {
		l.Wins[secondTeam] += 1
	}
}

func (l League) Ranking() []string {
	names := make([]string, 0, len(l.Teams))

	for k := range l.Teams {
		names = append(names, k)
	}

	sort.Slice(names, func(first, second int) bool {
		return l.Wins[names[first]] > l.Wins[names[second]]
	})

	return names
}

func main() {
	l := League{
		Name: "Big League",
		Teams: map[string]Team{
			"Italy": {
				Name:    "Italy",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"France": {Name: "France", Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"}},
			"India": {
				Name:    "India",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"Nigeria": {
				Name:    "Nigeria",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
		},
		Wins: map[string]int{},
	}
	l.MatchResult("Italy", 50, "France", 70)
	l.MatchResult("India", 85, "Nigeria", 80)
	l.MatchResult("Italy", 60, "India", 55)
	l.MatchResult("France", 100, "Nigeria", 110)
	l.MatchResult("Italy", 65, "Nigeria", 70)
	l.MatchResult("France", 95, "India", 80)

	RankPrinter(l, os.Stdout)
}
