import unittest
from AOC import *

class Test(unittest.TestCase):
    def test_part1(self):
        total  = part1(">")
        self.assertEqual(total, 2)
        
        total  = part1("^>v<")
        self.assertEqual(total, 4)
        
        total  = part1("^v^v^v^v^v")
        self.assertEqual(total, 2)

    def test_part2(self):
        total  = part2("^v")
        self.assertEqual(total, 3)
        
        total  = part2("^>v<")
        self.assertEqual(total, 3)
        
        total  = part2("^v^v^v^v^v")
        self.assertEqual(total, 11)
        
if __name__ == "__main__":
    unittest.main()
    
