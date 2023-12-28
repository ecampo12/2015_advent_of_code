import unittest
from AOC import *

class Test(unittest.TestCase):
    def test_part1(self):
        x = part1("abcdef")
        self.assertEqual(x, 609043)
        
        x = part1("pqrstuv")
        self.assertEqual(x, 1048970)
        
if __name__ == "__main__":
    unittest.main()
    
