import unittest
from AOC import *

class Test(unittest.TestCase):
    def test_part1(self):
        file = open("test_input.txt", "r")
        input = file.read().splitlines()
        file.close()
        times = part1(input)
        expected = [1120, 1056]
        self.assertEqual(times, expected)

    def test_part2(self):
        file = open("test_input.txt", "r")
        input = file.read().splitlines()
        file.close()
        points = part2(input)
        expected = [312, 689]
        self.assertEqual(points, expected)
        
if __name__ == "__main__":
    unittest.main()
    
