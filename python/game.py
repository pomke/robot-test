#!/usr/bin/env python
import sys

from robot import Table, Robot, RobotController

#Create a robot and give him a 5x5 table to play on
astro = Robot(Table(5, 5))

#Create a controller and give it a robot to control
controller = RobotController(astro)

#Read commands from stdin and feed them to our controller
for line in sys.stdin:
    controller.action(line)
