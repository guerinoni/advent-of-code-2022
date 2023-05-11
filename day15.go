package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"
)

//go:embed input/day15
var d15 string

func day15(input string, yLineToMonitor int) (int, int) {
	lines := strings.Split(input, "\n")
	sensors := make(map[point]point)
	for _, line := range lines {
		var sensor point
		var beacon point
		_, err := fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensor.x, &sensor.y, &beacon.x, &beacon.y)
		if err != nil {
			panic(err)
		}
		sensors[sensor] = beacon
	}
	return d15p1(sensors, yLineToMonitor), d15p2(sensors)
}

// manhattanDistance returns the manhattan distance between two points
func manhattanDistance(p1, p2 point) int {
	x := p1.x - p2.x
	y := p1.y - p2.y
	d := math.Abs(float64(x)) + math.Abs(float64(y))
	return int(d)
}

func d15p1(sensors map[point]point, yLineToMonitor int) int {
	lineToMonitor := make(map[int]bool)
	for sensor, beacon := range sensors {
		distanceFromBeacon := manhattanDistance(sensor, beacon)
		distanceFromY := math.Abs(float64(sensor.y - yLineToMonitor))
		impactOnY := distanceFromBeacon - int(distanceFromY)
		for i := 0; i <= impactOnY; i++ {
			lineToMonitor[sensor.x+i] = true
			lineToMonitor[sensor.x-i] = true
		}

		if beacon.y == yLineToMonitor {
			delete(lineToMonitor, beacon.x)
		}
	}
	return len(lineToMonitor)
}

func d15p2(sensors map[point]point) int {
	const limit = 4000000
	for sensor, beacon := range sensors {
		distanceFromBeacon := manhattanDistance(sensor, beacon)
		// add points just out of each direction of the sensor reach ("draw" a diamond)
		for i := distanceFromBeacon; i >= 0; i-- {
			var points []point
			points = append(points, point{sensor.x + (distanceFromBeacon - i), sensor.y + (distanceFromBeacon - (distanceFromBeacon - i)) + 1}) // bottom
			points = append(points, point{sensor.x - (distanceFromBeacon - i), sensor.y - (distanceFromBeacon - (distanceFromBeacon - i)) - 1}) // top
			points = append(points, point{sensor.x + (distanceFromBeacon - (distanceFromBeacon - i)) + 1, sensor.y + (distanceFromBeacon - i)}) // right
			points = append(points, point{sensor.x - (distanceFromBeacon - (distanceFromBeacon - i)) - 1, sensor.y - (distanceFromBeacon - i)}) // left

			// check the four points generated
			for _, p := range points {
				if !checkPointInRange(sensors, p, limit) {
					return p.x*limit + p.y
				}
			}
		}
	}

	return 0
}

func checkPointInRange(sensors map[point]point, point point, limit int) bool {
	// if current border point outside the range
	if point.x > limit || point.x < 0 || point.y > limit || point.y < 0 {
		return true
	}

	// scan each sensor
	for sen, bea := range sensors {

		// distance from border point to sensor
		thisdist := manhattanDistance(sen, point)

		// distance from sensor to his beacon
		sendist := manhattanDistance(sen, bea)

		// if sensor to beacon distance is greater than border point to beacon then the current border point is in sensor range
		if thisdist <= sendist {
			return true
		}
	}

	// point not in reach of any sensor
	return false
}
