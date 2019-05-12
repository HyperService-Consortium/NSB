
"""
do not use function sys.exit()
"""

import atexit

from py_nsbcli import *
from py_nsbcli.modules.admin import get_admin
import py_nsbcli

from option import Option


def check_glo_db_is_ok():
    global glo_db
    if glo_db.handler_num < 0:
        print("the leveldb at ./kvstore open failed")
        exit(1)


glo_db = py_nsbcli.LevelDB("./kvstore")
check_glo_db_is_ok()


# modules
admin = get_admin("")
cli = Client(admin)
cli.append_module("action", py_nsbcli.SystemAction(cli))
cli.append_module("token", py_nsbcli.SystemToken(cli))
cli.append_module("isc", py_nsbcli.ISC(cli))
kvdb = KVDB(glo_db)


@atexit.register
def atexit_close_global_db():
    global glo_db
    glo_db.close()
    print("gracefully stop")


if __name__ == '__main__':
    print("nsb-cli console")
    print("do not enter ^C(KeyBoardInterrupt)")
