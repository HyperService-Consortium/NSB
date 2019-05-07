
"""
do not use function sys.exit()
"""

import atexit

from py_nsbcli import *
from modules import *


def check_glo_db_is_ok():
    global glo_db
    if glo_db.handler_num < 0:
        print("the leveldb at ./kvstore open failed")
        exit(1)


glo_db = LevelDB("./kvstore")
check_glo_db_is_ok()


# modules
admin = Admin()
cli = Client(admin)
kvdb = KVDB(glo_db)


@atexit.register
def atexit_close_global_db():
    global glo_db
    glo_db.close()
    print("gracefully stop")


if __name__ == '__main__':
    print("nsb-cli console")
    print("do not enter ^C(KeyBoardInterrupt)")
