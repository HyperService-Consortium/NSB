
from ctypes import CDLL

try:
    from hexbytes import HexBytes
except ImportError or ModuleNotFoundError as e:
    print(e)
    exit(1)


from py_nsbcli.util.gotypes import (
    GoInt32,
    GoString,
    GoBytes,
    GolevelDBptr,
    GoWalletptr
)


import platform
from py_nsbcli.config import INCLUDE_PATH

ENC = "utf-8"

if platform.system() == "Windows":
    funcs = CDLL(INCLUDE_PATH + "/cwallet.dll")
elif platform.system() == "Darwin":
    funcs = CDLL(INCLUDE_PATH + "/cwallet_mac.dll")
elif platform.system() == "Linux":
    raise ImportError("can not import cwallet dynamic link library")
else:
    raise ImportError("no corresponding cwallet api on this platform")


funcs.CDLL_NewLevelDBHandler.argtype = GoString.Type
funcs.CDLL_NewLevelDBHandler.restype = GolevelDBptr

funcs.CDLL_CloseDB.argtype = GolevelDBptr
funcs.CDLL_CloseDB.restype = None

funcs.CDLL_PreCheckWallet.argtype = GoWalletptr
funcs.CDLL_PreCheckWallet.restype = GoInt32

funcs.CDLL_NewWalletHandlerFromDB.argtype = (GolevelDBptr, GoString.Type)
funcs.CDLL_NewWalletHandlerFromDB.restype = GoWalletptr

funcs.CDLL_NewWalletHandler.argtype = (GolevelDBptr, GoString.Type)
funcs.CDLL_NewWalletHandler.restype = GoWalletptr

funcs.CDLL_WalletAddress.argtype = (GoWalletptr, GoInt32)
funcs.CDLL_WalletAddress.restype = GoBytes.Type

funcs.CDLL_WalletSign.argtype = (GoWalletptr, GoInt32, GoBytes.Type, GoInt32)
funcs.CDLL_WalletSign.restype = GoBytes.Type

funcs.CDLL_WalletSignHash.argtype = (GoWalletptr, GoInt32, GoBytes.Type)
funcs.CDLL_WalletSignHash.restype = GoBytes.Type

funcs.CDLL_WalletVerifyByRaw.argtype = (GoWalletptr, GoInt32, GoBytes.Type, GoBytes.Type, GoInt32)
funcs.CDLL_WalletVerifyByRaw.restype = GoInt32

funcs.CDLL_WalletVerifyByHash.argtype = (GoWalletptr, GoInt32, GoBytes.Type, GoBytes.Type)
funcs.CDLL_WalletVerifyByHash.restype = GoInt32


class LevelDB:
    def __init__(self, path=None):
        self._handler_num = -1
        if path is not None:
            self.open(path)
        pass

    @property
    def handler_num(self):
        return self._handler_num

    def open(self, path):
        """
        NewLevelDBHandler(dbpath string) (handlerPtr int32)
        """
        self.close()
        self._handler_num = funcs.CDLL_NewLevelDBHandler(GoString.trans(path, ENC))
        return self._handler_num

    def close(self):
        """
        NewLevelDBHandler(dbpath string) (handlerPtr int32)
        """
        if self._handler_num > -1:
            funcs.CDLL_CloseDB(self._handler_num)

    @staticmethod
    def close_db(handler_num):
        funcs.CDLL_CloseDB(handler_num)


class Wallet:
    def __init__(self, db_handler, name):
        self._handler_num = -1
        self._name = name
        self.open(db_handler, name)

    @property
    def handler_num(self):
        return self._handler_num

    @property
    def name(self):
        return self._name

    def open(self, db_handler, name):
        if isinstance(db_handler, LevelDB):
            self._handler_num = funcs.CDLL_NewWalletHandlerFromDB(db_handler.handler_num, GoString.trans(name, ENC))
        elif isinstance(db_handler, int):
            self._handler_num = funcs.CDLL_NewWalletHandlerFromDB(db_handler, GoString.trans(name, ENC))
        else:
            self._handler_num = -1

    def address(self, idx=0):
        ptr = funcs.CDLL_WalletAddress(self._handler_num, idx)
        if ptr is None:
            return
        return GoBytes.convert(ptr, 32)

    @staticmethod
    def create(db_handler, name):
        wlt = Wallet(None, name)
        if isinstance(db_handler, LevelDB):
            wlt._handler_num = funcs.CDLL_NewWalletHandler(db_handler.handler_num, GoString.trans(name, ENC))
        elif isinstance(db_handler, int):
            wlt._handler_num = funcs.CDLL_NewWalletHandler(db_handler, GoString.trans(name, ENC))
        else:
            wlt._handler_num = -1
        return wlt

    def sign(self, msg: bytes) -> bytes or None:
        ptr = funcs.CDLL_WalletSign(self._handler_num, 0, GoBytes.frombytes(msg), len(msg))
        if ptr is None:
            return
        return GoBytes.convert(ptr, 64)

    def sign_hash(self, msg_hash: bytes) -> bytes or None:
        if len(msg_hash) != 64:
            raise ValueError("the length of SHA512 Hash(Bytes) must be 64")
        ptr = funcs.CDLL_WalletSign(self._handler_num, 0, GoBytes.frombytes(msg_hash))
        if ptr is None:
            return
        return GoBytes.convert(ptr, 64)

    def verify_by_raw(self, msg: bytes, signature: bytes) -> int:
        if len(signature) != 64:
            raise ValueError("the length of signature(Bytes) must be 64")
        print(msg, len(msg), GoBytes.convert(GoBytes.frombytes(signature), 64).hex())
        return funcs.CDLL_WalletVerifyByRaw(
            self._handler_num, 0,
            GoBytes.frombytes(signature),
            GoBytes.frombytes(msg),
            len(msg)
        )

    def verify_by_hash(self, msg_hash: bytes, signature: bytes) -> int:
        if len(msg_hash) != 64:
            raise ValueError("the length of SHA512 Hash(Bytes) must be 64")
        if len(signature) != 64:
            raise ValueError("the length of signature(Bytes) must be 64")
        return funcs.CDLL_WalletVerifyByHash(
            self._handler_num, 0,
            GoBytes.frombytes(signature),
            GoBytes.frombytes(msg_hash)
        )



if __name__ == '__main__':
    db = LevelDB("../kvstore")
    print(db.handler_num)
    test_wlt = Wallet(db, 'Alice')
    print(test_wlt.handler_num, test_wlt.address(0).hex())
    print(test_wlt.sign(b"\x10\x00").hex())
    aut = test_wlt.sign(b"\x10\x00")
    print(test_wlt.verify_by_raw(b"\x10\x00", aut))
    aut = test_wlt.sign(b"\x11\x00")
    print(test_wlt.verify_by_raw(b"\x10\x00", aut))

    test_wlt = Wallet(db, 'black_Alice')
    print(test_wlt.handler_num)
    print(test_wlt.sign(b"\x10\x00") is None)
    print(test_wlt.sign_hash(aut) is None)

    db.close()
