import unittest
from AOC import *

class Test(unittest.TestCase):
    def test_part1(self):
        file = open("test_input.txt", "r")
        input = file.read().splitlines()
        file.close()
        total = part1(input)
        self.assertEqual(total, 62842880)

    def test_part2(self):
        file = open("test_input.txt", "r")
        input = file.read().splitlines()
        file.close()
        total = part2(input)
        self.assertEqual(total, 57600000)
        
if __name__ == "__main__":
    unittest.main()
    
