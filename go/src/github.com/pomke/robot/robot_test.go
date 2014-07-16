package robot

import "testing"

func TestController(t *testing.T) {
    /* Test weather controller accepts a command, does not check if 
     * moves are valid for the robot.
     */

    r := NewRobot(&Table {5, 5})
    c := NewController(r)
    
    checkCommandValid := func(command string) {
        e := c.DoCommand(command)
        if e != nil {
            t.Error("Could not handle valid command " + command)
        }
    }

    checkCommandInvalid := func(command string) {
        e := c.DoCommand(command)
        if e == nil {
            t.Error("Accepted invalid command " + command)
        }
    }
   
    checkCommandValid("PLACE 0,0,NORTH")
    checkCommandValid("LEFT")
    checkCommandValid("RIGHT")
    checkCommandValid("MOVE")
    checkCommandInvalid("I like skunks")
    checkCommandInvalid("Noodles make me giggle")
    checkCommandInvalid("I am the very model of a modern major general")
    
}

/* Tests relating to robots and their movement */

func TestRobotPlace(t *testing.T) {
    r := NewRobot(&Table {2,2})
    r.Place(0, 0, "NORTH")
    x, y, heading := r.Report()
    if x != 0 || y != 0 || heading != "NORTH" {
        t.Error("Failed to place robot")
    }
}

func TestRobotLeft(t *testing.T) {
    r := NewRobot(&Table {2,2})
    r.Place(0, 0, "NORTH")
    r.Left()
    x, y, heading := r.Report()
    if x != 0 || y != 0 || heading != "WEST" {
        t.Error("Failed to turn robot left")
    }
}

func TestRobotRight(t *testing.T) {
    r := NewRobot(&Table {2,2})
    r.Place(0, 0, "NORTH")
    r.Right()
    x, y, heading := r.Report()
    if x != 0 || y != 0 || heading != "EAST" {
        t.Error("Failed to turn robot right")
    }
}

func TestRobotMove(t *testing.T) {
    r := NewRobot(&Table {2,2})
    r.Place(0, 0, "NORTH")
    r.Move()
    x, y, heading := r.Report()
    if x != 0 || y != 1 || heading != "NORTH" {
        t.Error("Failed to move robot")
    }
}

func TestRobotMoveOutOfBounds(t *testing.T) {
    r := NewRobot(&Table {2,2})
    r.Place(0, 0, "NORTH")
    r.Left()
    r.Move()
    x, y, heading := r.Report()
    if x != 0 || y != 0 || heading != "WEST" {
        t.Error("Moved off the table")
    }
}



