import unittest

import robot

class TestTable(unittest.TestCase):

    def setUp(self):
        self.t = robot.Table(2,2)

    def test_boundaries(self):
        self.assertTrue(self.t.inBounds(0,0))
        self.assertTrue(self.t.inBounds(1,0))
        self.assertFalse(self.t.inBounds(1,3))


class TestRobot(unittest.TestCase):

    def setUp(self):
        self.r = robot.Robot(robot.Table(2,2))

    def test_movement(self):
        self.assertRaises(RuntimeError, self.r.move)
        self.assertRaises(RuntimeError, self.r.left)
        self.assertRaises(RuntimeError, self.r.right)
        self.assertRaises(RuntimeError, self.r.report)

        self.r.place(0,0,'NORTH')
        self.assertEqual(self.r.report(), (0,0,'NORTH'))

        self.r.left()
        self.assertEqual(self.r.report(), (0,0,'WEST'))

        self.r.right()
        self.assertEqual(self.r.report(), (0,0,'NORTH'))

        self.r.move()
        self.assertEqual(self.r.report(), (0,1,'NORTH'))

        self.assertRaises(RuntimeError, self.r.move)


class TestRobotController(unittest.TestCase):

    def setUp(self):
        self.c = robot.RobotController(robot.Robot(robot.Table(2,2)))

    def test_actions(self):
        self.assertEqual(None, self.c.action('Not a command'))

        self.assertEqual(None, self.c.action('MOVE'))
        self.assertEqual(None, self.c.action('LEFT'))
        self.assertEqual(None, self.c.action('RIGHT'))
        self.assertEqual(None, self.c.action('REPORT'))
 
if __name__ == '__main__':
    unittest.main()
