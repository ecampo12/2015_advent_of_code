import unittest
from AOC import *

class Test(unittest.TestCase):
    def test_passwords_validity(self):
        file = open("test_input.txt", "r")
        input = file.read().splitlines()
        file.close()
        result = [password_validity(line) for line in input]
        expected = [False, False, False]
        self.assertEqual(result, expected)
    
    def test_next_password(self):
        file = open("test_next.txt", "r")
        input = file.read().splitlines()
        file.close()
        result = [next_password(line) for line in input]
        expected = ["xy", "xz", "yb"]
        self.assertEqual(result, expected)
    
    def test_part1(self):
        file = open("test_part1.txt", "r")
        input = file.read().splitlines()
        file.close()
        result = [part1(line) for line in input]
        expected = ["abcdffaa", "ghjaabcc"]
        self.assertEqual(result, expected)
        
if __name__ == "__main__":
    unittest.main()
    
