package robot
    
import (
    "fmt" 
    "errors"
    "strings"
    "strconv"
)
    

//A 2d table for a Robot to walk about on, has a maximum size x and y
type Table struct {
    X, Y int
}


/* Robot and methods. A Robot has a table and a place on that table along 
 * with a heading. It can move around the table and report it's location and 
 * current heading. 
 */

func NewRobot(t *Table) (*Robot) {
    /* Construct a new Robot, takes a table pointer for the robot to walk
     * around on
     */
    return &Robot {0, 0, "", t, false} 
}

//Directions a Robot can face
var ( directions = [4]string {"NORTH", "EAST", "SOUTH", "WEST"} )

type Robot struct {
    x, y int
    heading string
    table *Table
    placed bool
}

func (self *Robot) isOnTable(x int, y int) (bool) {
    /* Check to see if moving the robot to x, y would keep it 
     * on the table or not.
     */
    if x < self.table.X && x >= 0 {
        if y < self.table.Y && y >= 0 {
            return true
        }
    }
    return false
}

func (self *Robot) Place(x int, y int, heading string) error {
    /* Place the robot on it's table at a given x,y and facing in 
     * 'heading' direction.
     */
    if self.isOnTable(x, y) {
        self.x = x
        self.y = y
        self.heading = heading
        self.placed = true
        return nil
    }
    return errors.New("x, y coordinates are not on this robot's table")
}

func (self *Robot) Left() error {
    /* Turn the robot 90 degrees to the left. */
    if !self.placed {
        return errors.New("Robot must be placed first.")
    }
    //Get the current heading's index in the directions array
    currentIndex := findIndex(len(directions), 
        func(i int) bool { return directions[i] == self.heading })
    if currentIndex == 0 {
        self.heading = directions[len(directions)-1]
    } else {
        self.heading = directions[currentIndex-1]
    }
    return nil
}


func (self *Robot) Right() error {
    /* Turn the robot 90 degrees to the right. */
    if !self.placed {
        return errors.New("Robot must be placed first.")
    }
    //Get the current heading's index in the directions array
    currentIndex := findIndex(len(directions), 
        func(i int) bool { return directions[i] == self.heading })
    if currentIndex == len(directions)-1 {
        self.heading = directions[0]
    } else {
        self.heading = directions[currentIndex+1]
    }
    return nil
}

func (self *Robot) Move() error {
    /* Attempts to move the robot forward one space in the direction 
     * it is heading, will return an error if this would move it off 
     * the table.
     */
    
    x, y := self.x, self.y
    switch self.heading {
        case "NORTH":
            y = y+1
        case "EAST":
            x = x+1
        case "SOUTH":
            y = y-1
        case "WEST":
            x = x-1
    }

    if self.isOnTable(x, y) {
        self.x, self.y = x, y
        return nil
    } 
    return errors.New("Move invalid")
}

func (self *Robot) Report() {
    fmt.Printf("%d,%d,%s", self.x, self.y, self.heading)
}


/* Controller. Takes a Robot to control, then takes string input
 * which it parses looking for commands which it will pass on to
 * it's Robot.
 */

func NewController(r *Robot) *Controller {
    /* Controller constructor, takes a pointer to a robot to control */
    return &Controller {r}
}

type Controller struct {
    robot *Robot
}

func (self *Controller) DoCommand(s string) {
    /* DoCommand takes a string and attempts to match it with a valid robot
     * command:
     *
     * MOVE
     * LEFT
     * RIGHT
     * REPORT
     * PLACE X,Y,HEADING
     *
     * If a string matches none of these then it is ignored, as per 
     * requirements, any invalid commands pass silently.
     */

    s = strings.ToUpper(strings.Trim(s, " "))

    if s == "MOVE" {
        self.robot.Move()
    } else if s == "LEFT" {
        self.robot.Left()
    } else if s == "RIGHT" {
        self.robot.Right()
    } else if s == "REPORT" {
        self.robot.Report()
    } else if strings.HasPrefix(s, "PLACE") {
        bits := strings.Split(s, " ")
        if len(bits) == 2 {
            args := strings.Split(bits[1], ",")
            if len(args) == 3 {
                //First two should be ints
                x, _ := strconv.ParseInt(args[0], 10, 0)
                y, _ := strconv.ParseInt(args[1], 10, 0)
                self.robot.Place(int(x), int(y), args[2])
            }
        }

    }
}




/************** Helper funcs beyond this point ***************/

func findIndex(max int, check func(i int) bool) int {
    /* Return the first index where check func returns true, useful for
     * finding the first matching index in an array or slice, how is there 
     * nothing like this in Go already?
     */
    for i := 0; i < max; i++ {
        if check(i) {
            return i
        }
    }
    return -1
}



