import unittest
from AOC import *

class Test(unittest.TestCase):
    def test_part1(self):
        floor = part1("(())")
        self.assertEqual(floor, 0)
        floor = part1("()()")
        self.assertEqual(floor, 0)
        
        floor = part1("(((")
        self.assertEqual(floor, 3)
        floor = part1("(()(()(")
        self.assertEqual(floor, 3)
        floor = part1("))(((((")
        self.assertEqual(floor, 3)
        
        floor = part1("())")
        self.assertEqual(floor, -1)
        floor = part1("))(")
        self.assertEqual(floor, -1)
        
        floor = part1(")))")
        self.assertEqual(floor, -3)
        floor = part1(")())())")
        self.assertEqual(floor, -3)

    def test_part2(self):
        index = part2(")")
        self.assertEqual(index, 1)
        index = part2("()())")
        self.assertEqual(index, 5)
        
if __name__ == "__main__":
    unittest.main()
    
