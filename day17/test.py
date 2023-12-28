import unittest
from AOC import *

class Test(unittest.TestCase):
    def test_part1(self):
        file = open("test_input.txt", "r")
        input = file.read().splitlines()
        file.close()
        combos = part1(input)
        self.assertEqual(len(combos), 4)

    def test_part2(self):
        file = open("test_input.txt", "r")
        input = file.read().splitlines()
        file.close()
        min_len = part2(input)
        self.assertEqual(min_len, 3)
        
if __name__ == "__main__":
    unittest.main()
    
