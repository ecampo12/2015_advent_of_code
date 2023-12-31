import unittest
from AOC import *

class Test(unittest.TestCase):
    def test_part1(self):
        file = open("test_input.txt", "r")
        input = file.read().splitlines()
        file.close()
        ans = part1(input)
        self.assertEqual(ans['a'], 2)

    def test_part2(self):
        file = open("test_input.txt", "r")
        input = file.read().splitlines()
        file.close()
        ans = part2(input)
        self.assertEqual(ans['a'], 7)
        
if __name__ == "__main__":
    unittest.main()
    
