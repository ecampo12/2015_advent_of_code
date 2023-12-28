import unittest
from AOC import *

class Test(unittest.TestCase):
    def test_part1(self):
        file = open("test_input.txt", "r")
        input = file.read().splitlines()
        file.close()
        dims = parse_input(input)
        total  = part1(dims)
        self.assertEqual(total, 101)

    def test_part2(self):
        file = open("test_input.txt", "r")
        input = file.read().splitlines()
        file.close()
        dims = parse_input(input)
        total  = part2(dims)
        self.assertEqual(total, 48)
        
if __name__ == "__main__":
    unittest.main()
    
