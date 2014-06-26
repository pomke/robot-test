#!/usr/bin/env ruby

# Turn x,y into floating point
def mkFloat x, y
    return Float "#{x}.#{y}"
end

# Turn a float into x,y
def mkCoords f
    return (String f).split('.').map {|i| Integer i }
end

# I am a Table, I hold a 2d grid and can tell you if given coordinates
# are within my boundaries.
class Table

    # initialize a table with an array of floating point coordinates
    def initialize x=5, y=5
        @squares = []
        (0..x-1).each do |i|
            (0..y-1).each do |ii|
                @squares.push mkFloat i, ii
            end
        end
    end

    # Check if given coordinate is on the table
    def inBounds? x, y
        return @squares.include? mkFloat x, y
    end

end

# I am a Robot that can safely navigate a Table, I require
# a Table to walk on and provide an API for movment.
class Robot

    # possible directions I can face.
    DIRECTIONS = ["NORTH", "EAST", "SOUTH", "WEST"] 

    def initialize table
        @table = table
        @pos = @heading = nil
    end

    # Set me at a given position on the table
    def place x, y, heading
        if !DIRECTIONS.include? heading then
            raise "Invalid Heading"
        end

        if !@table.inBounds? x, y then
            raise "Move Invalid"
        end

        #set position and heading
        @pos = mkFloat x, y
        @heading = heading
    end

    # Turn 90 degrees to the left
    def left
        if !@pos then
            raise "Not on the table yet"
        end
        @heading = DIRECTIONS[DIRECTIONS.index(@heading) -1]
    end

    # Turn 90 degrees to the right
    def right
        if !@pos then
            raise "Not on the table yet"
        end
        i = DIRECTIONS.index(@heading) +1
        if i > DIRECTIONS.length() -1 then
            i = 0
        end
        @heading = DIRECTIONS[i]
    end

    # Move forward one place in the direction I am facing.
    def move
        if !@pos then
            raise "Not on the table yet"
        end
        
        # calculate where we'd end up if we move
        x, y = mkCoords @pos
        if @heading == "NORTH" then
            y+=1
        elsif @heading == "EAST" then
            x+=1
        elsif @heading == "SOUTH" then
            y-=1
        elsif @heading == "WEST" then
            x-=1
        end

        if @table.inBounds? x, y then
            @pos = mkFloat x, y
        else
            raise "Move Invalid"
        end
    end

    # Return x, y and heading based on my current state.
    def report
        if !@pos then
            raise "Not on the table yet"
        end
        x, y = mkCoords @pos
        return [x, y, @heading]
    end
end


# I control a Robot and feed it actions which I parse from string input.
# I silently pass on any errors from input or Robot.
class RobotController

    def initialize robot
        @robot = robot
    end

    # Takes a string and attempts to find an action for my Robot to perform.
    def action string
        string = string.strip.upcase
        if string == "MOVE" then
            begin
                @robot.move
            rescue RuntimeError
                #pass silently
            end
        elsif string == "LEFT" then
            begin
                @robot.left
            rescue RuntimeError
                #pass silently
            end
        elsif string == "RIGHT" then
            begin
                @robot.right
            rescue RuntimeError
                #pass silently
            end
        elsif string == "REPORT" then
            begin
                x, y, heading = @robot.report
                puts "#{x},#{y},#{heading}"
            rescue RuntimeError
                #pass silently
            end
        elsif string.start_with? "PLACE" then
            begin
                @robot.place *string.split[-1].split(',')
            rescue RuntimeError
                #pass silently
            end
        end
    end
end

if __FILE__ == $0
    #Create a robot and give him a 5x5 table to play on
    atlas = Robot.new Table.new 5,5

    #Create a controller and give it a robot to control
    controller = RobotController.new atlas

    #Read commands from stdin and feed them to our controller
    ARGF.each_line do |line|
        controller.action line
    end
end

