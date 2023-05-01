package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input/day16
var d16 string

type Valve struct {
	Name     string
	FlowRate int
	Tunnels  []string
}

func day16(input string) (int, int) {
	lines := strings.Split(input, "\n")
	valves := make(map[string]Valve)
	shortest := map[string]int{}
	distance := 1

	for _, line := range lines {
		if line == "" {
			continue
		}
		var name string
		var flowRate int
		var tunnels []string
		_, err := fmt.Sscanf(line, "Valve %s has flow rate=%d", &name, &flowRate)
		if err != nil {
			fmt.Println(err)
			return 0, 0
		}

		if !strings.Contains(line, "tunnels lead to valves ") {
			line = strings.Replace(line, "tunnel leads to valve ", "tunnels lead to valves ", 1)
		}

		t := strings.Split(line, "tunnels lead to valves ")
		tt := strings.Split(t[1], ", ")
		for _, ttt := range tt {
			tunnels = append(tunnels, ttt)
			shortest[name+ttt] = distance
			shortest[ttt+name] = distance
		}

		valves[name] = Valve{Name: name, FlowRate: flowRate, Tunnels: tunnels}
	}

	// create all combinations of paths
	canContinue := true
	for canContinue {
		canContinue = false
		distance++
		for k := range valves {
			for k2, val2 := range valves {
				if _, ok := shortest[k+k2]; ok {
					continue
				}

				for _, tunnel := range val2.Tunnels {
					if dist, ok := shortest[k+tunnel]; ok && dist < distance {
						canContinue = true
						shortest[k+k2] = distance
						shortest[k2+k] = distance
						break
					}
				}
			}
		}
	}

	toVisit := map[string]bool{}
	for name, val := range valves {
		if val.FlowRate > 0 {
			toVisit[name] = false
		}
	}

	minutes := 30
	minutesElephant := 4

	return d16p1(valves["AA"], minutes+1, shortest, valves, toVisit),
		d16p2(valves["AA"], valves["AA"], minutes-minutesElephant, minutes-minutesElephant+1, shortest, valves, toVisit)
}

func d16p1(v Valve, time int, shortest map[string]int, valves map[string]Valve, toVisit map[string]bool) int {
	time--
	if time <= 0 {
		return 0
	}
	max := 0
	for k, visit := range toVisit {
		if visit {
			continue
		}
		toVisit[k] = true
		release := d16p1(valves[k], time-shortest[v.Name+k], shortest, valves, toVisit)
		if release > max {
			max = release
		}
		toVisit[k] = false
	}
	return max + v.FlowRate*time
}

func d16p2(v, v2 Valve, time, time2 int, shortest map[string]int, valves map[string]Valve, toVisit map[string]bool) int {
	time2--
	if time2 <= 0 {
		return 0
	}
	max := 0
	for k, visit := range toVisit {
		if visit == true {
			continue
		}
		toVisit[k] = true
		release := d16p2(v2, valves[k], time2, time-shortest[v.Name+k], shortest, valves, toVisit)
		if release > max {
			max = release
		}
		toVisit[k] = false
	}
	return max + v2.FlowRate*time2
}
