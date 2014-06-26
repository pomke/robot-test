#!/usr/bin/env bash
echo "Testing the Python solution via stdin"
echo 
echo "Game test A, should output 0,1,NORTH"
cat ./data/game-a.txt | ./python/game.py 
echo 
echo "Game test B, should output 0,0,WEST"
cat ./data/game-b.txt | ./python/game.py 
echo 
echo "Game test C, should output 3,3,NORTH"
cat ./data/game-c.txt | ./python/game.py 
echo 
echo "Running unit tests"
/usr/bin/env python ./python/tests.py 
