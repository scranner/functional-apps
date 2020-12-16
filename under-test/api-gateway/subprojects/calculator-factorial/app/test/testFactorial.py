import unittest, os, sys

sys.path.insert(0, os.path.abspath(os.path.join(os.path.dirname(__file__), '../main')))

from factorial import factorial


class Factorial(unittest.TestCase):

    def test_valid_returnsResult(self):
        self.assertEqual(factorial("4"), '{"result":"' + str(24) + '"}')

    def test_invalid_returnsNone(self):
        self.assertEqual(factorial("mfwerf"), None)

    def test_valid1_returnsResult(self):
        self.assertEqual(factorial(3), '{"result":"' + str(6) + '"}')

    def test_invalid1_returnsNone(self):
        self.assertEqual(factorial("1e1"), '{"result":"' + str(3628800) + '"}')


if __name__ == '__main__':
    unittest.main()
