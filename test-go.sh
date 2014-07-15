
#!/usr/bin/env bash
echo "Building Go.."
./go/build.sh
echo "Testing the Golang solution via stdin"
echo 
echo "Game test A, should output 0,1,NORTH"
cat ./data/game-a.txt | ./go/bin/game
echo 
echo "Game test B, should output 0,0,WEST"
cat ./data/game-b.txt | ./go/bin/game
echo 
echo "Game test C, should output 3,3,NORTH"
cat ./data/game-c.txt | ./go/bin/game
echo 
#echo "Running unit tests"
#/usr/bin/env python ./python/tests
