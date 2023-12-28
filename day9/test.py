import unittest
from AOC import *

class Test(unittest.TestCase):
    def test_part1(self):
        file = open("test_input.txt", "r")
        input = file.read().splitlines()
        file.close()
        distance = part1(input)
        self.assertEqual(distance, 605)

    def test_part2(self):
        file = open("test_input.txt", "r")
        input = file.read().splitlines()
        file.close()
        distance = part2(input)
        self.assertEqual(distance, 982)
        
if __name__ == "__main__":
    unittest.main()
    
