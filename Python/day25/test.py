import unittest
from AOC import *

class Test(unittest.TestCase):
    def test_coord_pos(self):
        coords = [(1, 1), (2, 1), (1, 2), (1, 6), (4, 3), (4, 1), (3, 3), (3, 4), (2, 2)]
        expected = [1, 2, 3, 21, 18, 7, 13, 19, 5]
        results = [get_triangle_code(coord) for coord in coords]
        self.assertEqual(results, expected)
        
    def test_part1(self):
        end_coord = (6, 6)
        code = part1(end_coord)
        self.assertEqual(code, 27995004)
        
if __name__ == "__main__":
    unittest.main()
    
