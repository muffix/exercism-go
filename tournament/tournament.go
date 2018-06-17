// Package tournament contains functions for a tournament
package tournament

import (
	"bytes"
	"fmt"
	"io"
	"sort"
	"strings"
)

// stats is the structure that records statistics for a team
type stats struct {
	name   string
	played int
	won    int
	lost   int
	drawn  int
	points int
}

// table is an ordered slice of team statistics. implements sort.Interface
type table []*stats

// tournamentStats is an unordered collection of team statistics with fast lookup by name
type tournamentStats map[string]*stats

const tableRowFormatString = "%-31s| %2d | %2d | %2d | %2d | %2d\n"

func (t table) Len() int {
	return len(t)
}

func (t table) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t table) Less(i, j int) bool {
	if t[i].points == t[j].points {
		return t[i].name < t[j].name
	}
	return t[i].points > t[j].points
}

// Tally reads tornament results and writes a table to the given writer
func Tally(r io.Reader, w io.Writer) error {
	statistics := &tournamentStats{}

	var inputBuffer bytes.Buffer
	_, err := inputBuffer.ReadFrom(r)

	if err != nil {
		return err
	}

	lines := strings.Split(inputBuffer.String(), "\n")

	for _, line := range lines {
		if line = cleanLine(line); line != "" {
			err := statistics.processLine(line)
			if err != nil {
				return err
			}
		}
	}

	statistics.writeTable(w)

	return nil
}

// cleanLine removes all irrelevant characters from a line (comments, leading/trailing whitespace)
func cleanLine(line string) string {
	line = strings.TrimSpace(line)
	if strings.HasPrefix(line, "#") {
		return ""
	}
	return line
}

// processLine splits a single line of input and starts processing it
func (statistics tournamentStats) processLine(line string) error {
	match := strings.Split(line, ";")

	if len(match) != 3 {
		return fmt.Errorf("Malformed line")
	}

	err := statistics.recordMatch(match[0], match[1], match[2])

	if err != nil {
		return err
	}

	return nil
}

// getTeamStats returns the statistics for a team of the given name
func (statistics tournamentStats) getTeamStats(teamName string) *stats {
	team, ok := statistics[teamName]

	if !ok {
		team = &stats{name: teamName}
		statistics[teamName] = team
	}

	return team
}

// recordMatch updates the statistics for the teams based on the result
func (statistics tournamentStats) recordMatch(teamNameA, teamNameB, result string) error {
	teamA := statistics.getTeamStats(teamNameA)
	teamB := statistics.getTeamStats(teamNameB)

	switch result {
	case "win":
		teamA.recordWin()
		teamB.recordLoss()
	case "loss":
		teamB.recordWin()
		teamA.recordLoss()
	case "draw":
		teamA.recordDraw()
		teamB.recordDraw()
	default:
		return fmt.Errorf("Malformed line")
	}

	return nil
}

// recordWin updates the staristics for a winning team
func (team *stats) recordWin() {
	team.played++
	team.won++
	team.points += 3
}

// recordLoss updates the staristics for a losing team
func (team *stats) recordLoss() {
	team.played++
	team.lost++
}

// recordDraw updates the staristics for a draw
func (team *stats) recordDraw() {
	team.played++
	team.drawn++
	team.points++
}

// writeHeader writes the table header to an io.Writer
func writeHeader(w io.Writer) {
	fmt.Fprintf(w, "%-31s| %2s | %2s | %2s | %2s | %2s\n", "Team", "MP", "W", "D", "L", "P")
}

// writeTable writes the table rows to an io.Writer
func (statistics tournamentStats) writeTable(w io.Writer) {
	var t table

	for _, teamStats := range statistics {
		t = append(t, teamStats)
	}

	sort.Sort(t)

	writeHeader(w)

	for _, teamStats := range t {
		fmt.Fprintf(w, tableRowFormatString, teamStats.name, teamStats.played, teamStats.won, teamStats.drawn, teamStats.lost, teamStats.points)
	}
}
