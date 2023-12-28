import unittest
from AOC import *

class Test(unittest.TestCase):
    def test_part1(self):
        file = open("test_input.txt", "r")
        input = file.read().splitlines()
        file.close()
        result = [look_and_say(line) for line in input]
        expected = ["11", "21", "1211", "111221", "312211"]
        self.assertEqual(result, expected)
        
if __name__ == "__main__":
    unittest.main()
    
