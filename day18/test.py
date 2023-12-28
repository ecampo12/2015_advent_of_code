import unittest
from AOC import *

class Test(unittest.TestCase):
    def test_part1(self):
        file = open("test_input.txt", "r")
        input = file.read().splitlines()
        file.close()
        on = part1(input)
        self.assertEqual(on, 4)

    def test_part2(self):
        file = open("test_input.txt", "r")
        input = file.read().splitlines()
        file.close()
        on = part2(input)
        self.assertEqual(on, 17)
        
if __name__ == "__main__":
    unittest.main()
    
