from py_nsbcli import LevelDB, Wallet


class KVDB(object):
    def __init__(self, dbhandler: LevelDB):
        self._dbhandler = dbhandler

    def set(self, db):
        if isinstance(db, str):
            self._dbhandler.close()
            self._dbhandler = LevelDB(db)
            if self._dbhandler.handler_num < 0:
                raise Exception("Open failed")
        elif isinstance(db, LevelDB):
            if db.handler_num < 0:
                raise Exception("invalid db ptr")
            self._dbhandler = db

    def load_wallet(self, name) -> Wallet:
        x = Wallet(self._dbhandler, name)
        if x.handler_num < 0:
            raise Exception("create failed")
        return x

    def create_wallet(self, name) -> Wallet:
        x = Wallet.create(self._dbhandler, name)
        if x.handler_num < 0:
            raise Exception("create failed")
        return x
