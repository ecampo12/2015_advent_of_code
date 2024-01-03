import unittest
from AOC import *

class Test(unittest.TestCase):
    def test_part1(self):
        file = open("test_input.txt", "r")
        input = file.read().splitlines()
        file.close()
        self.assertEqual(part1(input), 330)
     
if __name__ == "__main__":
    unittest.main()
    
