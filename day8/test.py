import unittest
from AOC import *

class Test(unittest.TestCase):
    def test_part1(self):
        file = open("test_input.txt", "r")
        input = file.read().splitlines()
        file.close()
        self.assertEqual(part1(input), 12)

    def test_part2(self):
        file = open("test_input.txt", "r")
        input = file.read().splitlines()
        file.close()
        self.assertEqual(part2(input), 19)
        
if __name__ == "__main__":
    unittest.main()
    
