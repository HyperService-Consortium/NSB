

class Contract(object):
    def __init__(self, name):
        self._name = name

    @property
    def name(self):
        return self._name

    def func(self, func_name):
        pass
