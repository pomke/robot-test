require_relative "robot"
require "test/unit"

class TestTable < Test::Unit::TestCase

    # test that a table knows it's boundaries
    def test_boundaries
        t = Table.new 2,2
        assert t.inBounds? 1,1
        assert t.inBounds? 0,1
        assert !t.inBounds?(2,1)
    end

end


class TestRobot < Test::Unit::TestCase

    def test_movement
        r = Robot.new Table.new 2,2

        #should raise on actions before a place
        assert_raise(RuntimeError) { r.move }
        assert_raise(RuntimeError) { r.left }
        assert_raise(RuntimeError) { r.right }
        assert_raise(RuntimeError) { r.report }

        r.place 0,0,"NORTH"
        #check we are on the table
        assert_equal [0,0,"NORTH"], r.report

        r.left
        assert_equal [0,0,"WEST"], r.report

        r.right
        assert_equal [0,0,"NORTH"], r.report

        r.move
        assert_equal [0,1,"NORTH"], r.report
        
        #perform an illegal move (off the table)
        assert_raise(RuntimeError) { r.move }
    end

end

class TestController < Test::Unit::TestCase

    def test_actions
        c = RobotController.new Robot.new Table.new 2,2

        #accepts bogus commands
        assert_nothing_raised(RuntimeError) { c.action "not a command"}

        #accepts invalid command orders
        assert_nothing_raised(RuntimeError) { c.action "MOVE"}
        assert_nothing_raised(RuntimeError) { c.action "LEFT"}
        assert_nothing_raised(RuntimeError) { c.action "RIGHT"}
        assert_nothing_raised(RuntimeError) { c.action "REPORT"}

    end

end
 
