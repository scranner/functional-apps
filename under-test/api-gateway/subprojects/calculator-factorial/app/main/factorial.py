import math


def factorial(x):
    try:
        num = float(x)
        return '{"result":"' + str(math.factorial(num)) + '"}'
    except:
        return None
