
from ctypes import CDLL

from gotypes import (
    GoInt32,
    GoString,
    GoBytes,
    GolevelDBptr,
    GoWalletptr
)

from config import INCLUDE_PATH

ENC = "utf-8"

funcs = CDLL(INCLUDE_PATH + "/cwallet.dll")

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

funcs.CDLL_WalletSign.argtype = (GoWalletptr, GoInt32, GoBytes.Type, GoInt32)
funcs.CDLL_WalletSign.restype = GoBytes.Type

funcs.CDLL_WalletSignHash.argtype = (GoWalletptr, GoInt32, GoBytes.Type)
funcs.CDLL_WalletSignHash.restype = GoBytes.Type

funcs.CDLL_WalletVerifyByRaw.argtype = (GoWalletptr, GoInt32, GoBytes.Type, GoBytes.Type, GoInt32)
funcs.CDLL_WalletVerifyByRaw.restype = GoInt32

funcs.CDLL_WalletVerifyByHash.argtype = (GoWalletptr, GoInt32, GoBytes.Type, GoBytes.Type)
funcs.CDLL_WalletVerifyByHash.restype = GoInt32


class LevelDB:
    def __init__(self):
        self.handler_num = 0

    def new_level_db_handler(self, path):
        pass


class Wallet:
    pass
