import base64
import json

from config import ENC
from py_nsbcli.util.cast import transbytes
from py_nsbcli.modules.contract import Contract


class SystemToken(Contract):
    def __init__(self, bind_cli):
        super().__init__(bind_cli)

    def set_balance(self, wlt, value: int or bytes or str):

        value = transbytes(value, 32)

        if len(value) > 32:
            raise ValueError("value(uint256) overflow")

        data_set_balance = {
            "function_name": "setBalance",
            "args": base64.b64encode(json.dumps({
                "1": base64.b64encode(value).decode(),
            }).encode(ENC)).decode()
        }

        return self.exec_system_contract_method(
            wlt,
            b"systemCall\x19system.token\x18",
            json.dumps(data_set_balance).encode(ENC),
            0
        )
