

import random
from .cast import catint32, MOD256


def randint256():
    return catint32(random.randint(0, MOD256))
