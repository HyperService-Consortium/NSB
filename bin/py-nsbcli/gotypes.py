
import ctypes


# GoInt in C
GoInt8 = ctypes.c_int8
GoInt16 = ctypes.c_int16
GoInt32 = ctypes.c_int32
GoInt64 = ctypes.c_int64

# GoUint in C
GoUint8 = ctypes.c_uint8
GoUint16 = ctypes.c_uint16
GoUint32 = ctypes.c_uint32
GoUint64 = ctypes.c_uint64

GoInt = GoUint64
GoUInt = GoUint64

# the integer coresponds to the GolevelDB-array's index
GolevelDBptr = ctypes.c_int
GoWalletptr = ctypes.c_int


class GoBytes(object):
    # GoBytes in C
    Type = ctypes.c_void_p

    @staticmethod
    def frombytes(bytesarr: bytes):
        return ctypes.cast(ctypes.create_string_buffer(bytesarr, len(bytesarr)), ctypes.c_char_p)

    @staticmethod
    def convert(bytes_pointer, bytes_len=-1) -> bytes:
        return ctypes.string_at(bytes_pointer, bytes_len)


class GoString(object):
    # GoString in C

    Type = ctypes.POINTER(ctypes.c_char)

    @staticmethod
    def fromstr(pystr, enc):
        return ctypes.c_char_p(bytes(pystr.encode(enc)))

    @staticmethod
    def frombytes(pystr):
        return ctypes.c_char_p(pystr)

    @staticmethod
    def trans(pystr, enc='utf-8'):
        if isinstance(pystr, str):
            return GoString.fromstr(pystr, enc)
        elif isinstance(pystr, bytes):
            return GoString.frombytes(pystr)
        else:
            return 'error'


class GoStringSlice(object):
    # GoStringSlice in C

    Type = ctypes.POINTER(GoString.Type)

    @staticmethod
    def fromstrlist(strlist, enc):
        strs = []
        for hexstr in strlist:
            # const hexstr, so const strlist
            if hexstr[0:2] == '0x':
                hexstr = hexstr[2:]
            strptr = GoString.trans(hexstr, enc)
            if strptr == 'error':
                raise TypeError("hexstr-type in strlist needs str or bytes, but get", hexstr.__class__)
            strs.append(strptr)
        charparray = ctypes.c_char_p * len(strlist)
        charlist = charparray(*strs)
        return ctypes.cast(charlist, GoStringSlice.Type)
