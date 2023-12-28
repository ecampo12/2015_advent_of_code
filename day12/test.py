import unittest
from AOC import *

class Test(unittest.TestCase):
    def test_part1(self):
        file = open("test_input.txt", "r")
        input = file.read().splitlines()
        file.close()
        results = [part1(line) for line in input]
        excepted = [6, 6, 3, 3, 0, 0, 0, 0]
        self.assertEqual(results, excepted)

    def test_part2(self):
        file = open("test_input2.txt", "r")
        input = file.read().splitlines()
        file.close()
        results = [part2(line) for line in input]
        excepted = [6, 4, 0, 6]
        self.assertEqual(results, excepted)
        
if __name__ == "__main__":
    unittest.main()
    
