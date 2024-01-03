import unittest
from AOC import *

class Test(unittest.TestCase):
    def test_part1(self):
        file = open("test_input.txt", "r")
        input = file.read().splitlines()
        file.close()
        results = [part1(x)[0] for x in input]
        expected = [1, 2, 3, 4, 5, 6, 7, 8, 9]
        self.assertEqual(results, expected)

if __name__ == "__main__":
    unittest.main()
    
