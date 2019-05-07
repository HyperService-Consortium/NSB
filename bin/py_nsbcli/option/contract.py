

import base64
import json

from py_nsbcli.types import TransactionHeader
from py_nsbcli.modules.contract import Contract
from py_nsbcli.util.cast import transbytes


class Option(Contract):
    def __init__(self, bind_cli):
        super().__init__(bind_cli)

    def create_option(self, wlt, owner, price, value):
        value = transbytes(price, 32)
        args_option = {
            "owner": base64.b64encode(owner).decode(),
            "strike_price": base64.b64encode(value).decode()
        }
        tx_header = TransactionHeader(wlt.address(0), None, json.dumps(args_option).encode(), value)
        tx_header.sign(wlt)
        return self.create_contract("option", tx_header)
