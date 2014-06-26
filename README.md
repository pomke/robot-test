Robot Test
==========

Requires Bash, Ruby 2.x and Python 2.7+

* /data contains the 3 test runs in text files
* /ruby contains a ruby implementation
* /python contains a python implementation

The files in /data can be piped to stdin of ./python/game.py and ./ruby/robot.rb

running ./test-python.sh  and ./test-ruby.sh will exercise the robots with each 
of the test run files and also run the unit-tests for each of the implementations 
respectively.
