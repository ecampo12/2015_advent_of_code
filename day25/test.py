import unittest
from AOC import *

class Test(unittest.TestCase):
    def test_part1(self):
        start = 20151125
        start_coord = (1, 1)
        end_coord = (6, 6)
        code = part1(start, start_coord, end_coord)
        self.assertEqual(code, 27995004)
        
if __name__ == "__main__":
    unittest.main()
    
