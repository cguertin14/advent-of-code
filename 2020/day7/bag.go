package day7

import (
	"strconv"
	"strings"

	"github.com/cguertin14/advent-of-code-2020/utils"
)

func parseBags(reader utils.Reader) (containers map[string]map[string]int, containedBy map[string][]string) {
	containers = make(map[string]map[string]int)
	containedBy = make(map[string][]string)

	for _, s := range reader.Lines {
		split := strings.Split(s[:len(s)-1], " bags contain ")
		if split[1] == "no other bags" {
			continue
		}

		container := split[0]
		contained := strings.Split(split[1], ", ")
		containers[container] = make(map[string]int)

		for _, b := range contained {
			b = strings.TrimSuffix(b, " bags")
			b = strings.TrimSuffix(b, " bag")

			subContained := strings.SplitN(b, " ", 2)
			qty, _ := strconv.Atoi(subContained[0])
			content := subContained[1]

			containedBy[content] = append(containedBy[content], container)
			containers[container][content] = qty
		}
	}

	return
}

func countShinyGoldContainers(containers map[string]map[string]int, containedBy map[string][]string) (shinyGoldContainers int) {
	newlyFound := []string{"shiny gold"}
	found := map[string]bool{
		"shiny gold": true,
	}

	for {
		nextRound := make([]string, 0)
		for _, b := range newlyFound {
			foundContainers := containedBy[b]
			for _, c := range foundContainers {
				if !found[c] {
					shinyGoldContainers++
					found[c] = true
					nextRound = append(nextRound, c)
				}
			}
		}

		if len(nextRound) == 0 {
			break
		}
		newlyFound = nextRound
	}

	return
}

func countShinyGoldContained(containers map[string]map[string]int) (shinyGoldContained int) {
	bagsFound := make(map[string]int)
	newlyFound := map[string]int{
		"shiny gold": 1,
	}

	for {
		nextRound := make(map[string]int)
		for k, v := range newlyFound {
			contained := containers[k]
			for c, count := range contained {
				bagsFound[c] += count * v
				nextRound[c] += count * v
			}
		}

		if len(nextRound) == 0 {
			break
		}
		newlyFound = nextRound
	}

	for _, b := range bagsFound {
		shinyGoldContained += b
	}

	return
}
