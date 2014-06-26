#!/usr/bin/env bash
echo "Testing the Ruby solution via stdin"
echo 
echo "Game test A, should output 0,1,NORTH"
cat ./data/game-a.txt | ./ruby/robot.rb 
echo 
echo "Game test B, should output 0,0,WEST"
cat ./data/game-b.txt | ./ruby/robot.rb 
echo 
echo "Game test C, should output 3,3,NORTH"
cat ./data/game-c.txt | ./ruby/robot.rb 
echo 
echo "Running unit tests"
/usr/bin/env ruby ./ruby/tests.rb 
