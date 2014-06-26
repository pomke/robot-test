def mkFloat(x, y):
    """Make a floating point number from two ints."""
    return float('%s.%s'% (x, y))

def mkCoords(f):
    """Break a floating point number into x, y."""
    return map(lambda point: int(point), str(f).split('.'))


class Table(object):
    """I am a Table, a simple 2d grid with no special behavior.
    I can tell you if a coordinate is within my boundary however."""

    def __init__(self, x=5, y=5):
        """Create a table with dimensions x and y."""
        self.squares = []
        for i in range(x):
            for ii in range(y):
                self.squares.append(mkFloat(i,ii))

    def inBounds(self, x, y):
        """Check if given coordinate is on the table."""
        return mkFloat(x,y) in self.squares


class Robot(object):
    """Hi, I'm a Robot that can safely navigate a Table, I require
    a Table to walk on and can perform particular movements."""

    #possible directions a Robot can face
    directions = ('NORTH', 'EAST', 'SOUTH', 'WEST')

    def __init__(self, table):
        self.table = table;
        self.pos = self.heading = None

    def place(self, x, y, heading):
        """Set me at a given position on the table. Raise an error
        if the placement is invalid or the heading incorrect."""
        if heading not in self.directions:
            raise ValueError('Invalid Heading', heading)

        if not self.table.inBounds(x, y):
            raise RuntimeError('Move Invalid')
        #Set our position and heading
        self.pos = mkFloat(x, y)
        self.heading = heading

    def left(self):
        """Turn 90 degrees to the left."""
        if self.pos is None:
            raise RuntimeError('Not on the table yet')
        self.heading = self.directions[self.directions.index(self.heading)-1]

    def right(self):
        """Turn 90 degrees to the right."""
        if self.pos is None:
            raise RuntimeError('Not on the table yet')
        i = self.directions.index(self.heading) +1
        if i > len(self.directions)-1:
            i = 0
        self.heading = self.directions[i]

    def move(self):
        """Move me forward one place in the direction I am facing,
        raise an error if the placement is invalid."""
        if self.pos is None:
            raise RuntimeError('Not on the table yet')

        #calculate where we'd end up if we move
        x, y = mkCoords(self.pos) 
        if self.heading == 'NORTH':
            y += 1
        if self.heading == 'EAST':
            x += 1
        if self.heading == 'SOUTH':
            y -= 1
        if self.heading == 'WEST':
            x -= 1

        if self.table.inBounds(x, y):
            #we're still on the table, lets move!
            self.pos = mkFloat(x, y)
        else:
            #we'd fall off the table, lets stay put.
            raise RuntimeError('Move invalid')

    def report(self):
        """Return a three-tuple of x, y and heading, raise an error if
        we're not yet on the table."""
        if self.pos is None:
            raise RuntimeError('Not on the table yet')
        x, y = mkCoords(self.pos)
        return (x, y, self.heading)


class RobotController(object):
    """A robot controller parses actions and controls it's robot, handling
    errors from user-input or from the robot."""

    def __init__(self, robot):
        self.robot = robot

    def action(self, string):
        """Takes a string and attempts to find an action in it for it's
        robot to perform, silently fails if actions are invalid."""
        inp = string.strip().upper()
        if inp == 'MOVE':
            try:
                self.robot.move()
            except Exception, e:
                pass #silently

        if inp == 'LEFT':
            try:
                self.robot.left()
            except Exception, e:
                pass #silently

        if inp == 'RIGHT':
            try:
                self.robot.right()
            except Exception, e:
                pass #silently

        if inp == 'REPORT':
            try:
                print '%d,%d,%s' % self.robot.report()
            except Exception, e:
                pass #silently

        if inp.startswith('PLACE'):
            try:
                self.robot.place(*inp.split()[-1].split(','))
            except Exception, e:
                pass #silently
