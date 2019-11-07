
import sys

# Minimum makefile impl github.com/Myriad-Dreamin/pymake
made = set()
__magic_module = sys.modules[__name__]


def consume_makefile(method):
    i = str(method)
    print(i, made)
    if i in made:
        made.update(i)
    else:
        return method


# object requirement
def cr(attr: str):
    # raise if not exists
    def __get_object_attr(obj, *_): return getattr(obj, attr)
    return __get_object_attr


# object requirements
def crs(*attrs): return map(cr, attrs)


# for module makefile:
# return function in module scope
def find(preq):
    return __magic_module.__dict__[preq] if preq in __magic_module.__dict__ else None


def require(*targets):
    def c(f):
        def wrapper(*args, **kwargs):
            for target in targets:
                target = (getattr(target, '__name__', None) == '__get_object_attr' and target(*args, **kwargs)) or target
                target = find(target) if isinstance(target, str) else target
                consume_makefile(target)(*args, **kwargs) if callable(target) else None
            return f(*args, **kwargs)
        return wrapper
    return c


# for shadow everything
class Makefile:

    @classmethod
    @require('muf', *crs('start2'))
    def start(cls, *_):
        print("here2")

    @classmethod
    def start2(cls, *_):
        print("here3")

    @classmethod
    @require('muf', *crs('start', 'start2'))
    def all(cls, *_):
        print("here")


def muf(*_):
    print("QAQ")


print(len(sys.argv), sys.argv)

Makefile().all()
