import unittest
from AOC import *

class Test(unittest.TestCase):
    def test_part1(self):
        file = open("test_input.txt", "r")
        input = file.read().splitlines()
        file.close()
        mol = part1(input)
        self.assertEqual(mol, 4)

if __name__ == "__main__":
    unittest.main()
    
