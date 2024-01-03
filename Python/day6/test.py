import unittest
from AOC import *

class Test(unittest.TestCase):
    def test_part1(self):
        file = open("test_input.txt", "r")
        input = file.read().splitlines()
        file.close()
        x = parse_input(input)
        self.assertEqual(part1(x), 998996)

    def test_part2(self):
        file = open("test_input2.txt", "r")
        input = file.read().splitlines()
        file.close()
        x = parse_input(input)
        self.assertEqual(part2(x), 2000001)
        
if __name__ == "__main__":
    unittest.main()
    
