import unittest
from AOC import *

class Test(unittest.TestCase):
    def test_part1(self):
        file = open("test_input.txt", "r")
        input = file.read().splitlines()
        file.close()
        values = {
            'd': 72,
            'e': 507,
            'f': 492,
            'g': 114,
            'h': 65412,
            'i': 65079,
            'x': 123,
            'y': 456
        }
        for key, value in part1(input).items():
            self.assertEqual(value, values[key])
        
if __name__ == "__main__":
    unittest.main()
    
