// Package robot implements a robot simulation
package robot

import (
	"fmt"
)

const (
	N = 0
	E = 1
	S = 2
	W = 3
)

func (d Dir) String() string {
	return []string{"N", "E", "S", "W"}[d]
}

func (r *Step2Robot) String() string {
	return fmt.Sprintf("Direction %s, position %v", r.Dir, r.Pos)
}

func (r Rect) contains(p Pos) bool {
	return (p.Northing >= r.Min.Northing &&
		p.Northing <= r.Max.Northing &&
		p.Easting >= r.Min.Easting &&
		p.Easting <= r.Max.Easting)
}

// Action is the type for a robot action in step 2
type Action func(*Step2Robot)

// Action3 is the type for a robot action in step 3
type Action3 struct {
	Name    string
	Actions []Action
}

/* STEP 3 */

type state struct {
	finished bool
	robot    *Step2Robot
}

// StartRobot3 starts a robot in step 3
func StartRobot3(name, script string, outAction chan<- Action3, log chan string) {
	var actions []Action

	commandsCh := make(chan Command)
	actionsCh := make(chan Action)
	go StartRobot(commandsCh, actionsCh)

	for _, code := range script {
		commandsCh <- Command(code)
		if action := <-actionsCh; action != nil {
			actions = append(actions, action)
		} else {
			log <- fmt.Sprintf("Invalid command to robot %s: %s", name, string(code))
			break
		}
	}
	close(commandsCh)

	outAction <- Action3{name, actions}

}

// Room3 simulates the room in step 3
func Room3(extent Rect, robots []Step3Robot, action chan Action3, report chan []Step3Robot, log chan string) {
	robotStates := make(map[string]state)

	for i, r := range robots {
		_, nameTaken := robotStates[r.Name]

		switch {
		case r.Name == "":
			log <- "Found a robot without a name"
		case nameTaken:
			log <- fmt.Sprintf("Name %s is already taken", r.Name)
		case !extent.contains(r.Pos):
			log <- fmt.Sprintf("Robot placed outside of the room: %s", r.Name)
		case robotOnPosition(r.Pos, robotStates):
			log <- fmt.Sprintf("Cannot place %s on occupied position %v", r.Name, r.Pos)
		default:
			robotStates[r.Name] = state{finished: false, robot: &(robots[i].Step2Robot)}
			continue
		}
		report <- robots
		return
	}

	for !allRobotsFinished(robotStates) {
		act := <-action
		state, ok := robotStates[act.Name]

		if !ok {
			log <- fmt.Sprintf("Robot %s doesn't exist", act.Name)
			break
		}

		robot := state.robot

		for _, action := range act.Actions {
			oldPos := robot.Pos
			action(robot)

			switch {
			case !extent.contains(robot.Pos):
				log <- fmt.Sprintf("Robot %s hit a wall", act.Name)
				robot.Pos = oldPos
			case collisionAt(robot.Pos, robotStates):
				log <- fmt.Sprintf("Robot %s hit another robot", act.Name)
				robot.Pos = oldPos
			}
		}
		state.finished = true
		robotStates[act.Name] = state
	}
	report <- robots
}

func allRobotsFinished(states map[string]state) bool {
	for _, state := range states {
		if state.finished == false {
			return false
		}
	}

	return true
}

func robotOnPosition(pos Pos, states map[string]state) bool {
	for _, state := range states {
		if state.robot.Pos == pos {
			return true
		}
	}
	return false
}

func collisionAt(p Pos, states map[string]state) bool {
	var foundRobot bool

	for _, state := range states {
		if state.robot.Pos == p {
			if foundRobot {
				return true
			} else {
				foundRobot = true
			}
		}
	}

	return false
}

/* STEP 2 */

// StartRobot starts a robot in step 2
func StartRobot(cmdChan <-chan Command, actionChan chan<- Action) {
	for c := range cmdChan {
		switch c {
		case 'R':
			actionChan <- (*Step2Robot).right
		case 'L':
			actionChan <- (*Step2Robot).left
		case 'A':
			actionChan <- (*Step2Robot).advance
		default:
			actionChan <- nil
		}
	}
	close(actionChan)
}

// Room simulates the room in step 2
func Room(extent Rect, robot Step2Robot, act <-chan Action, rep chan Step2Robot) {
	for a := range act {
		oldPos := robot.Pos
		a(&robot)
		if !extent.contains(robot.Pos) {
			robot.Pos = oldPos
		}
	}
	rep <- robot
}

func (r *Step2Robot) advance() {
	switch r.Dir {
	case N:
		r.Northing++
	case E:
		r.Easting++
	case S:
		r.Northing--
	case W:
		r.Easting--
	}
}
func (r *Step2Robot) right() {
	r.Dir = (r.Dir + 1) % 4
}
func (r *Step2Robot) left() {
	r.Dir = Dir((uint(r.Dir) - 1) % 4)
}

/* STEP 1 */

// Advance moves the robot in the direction it's facing in step 1
func Advance() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Y++
	case E:
		Step1Robot.X++
	case S:
		Step1Robot.Y--
	case W:
		Step1Robot.X--
	}
}

// Right turns the robot to the right in step 1
func Right() {
	Step1Robot.Dir = (Step1Robot.Dir + 1) % 4
}

// Left turns the robot to the left in step 1
func Left() {
	Step1Robot.Dir = Dir((uint(Step1Robot.Dir) - 1) % 4)
}
