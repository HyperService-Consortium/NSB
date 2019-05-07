import base64
import json

from config import ENC
from py_nsbcli.system_action.action import Action
from py_nsbcli.modules.contract import Contract


class SystemAction(Contract):
    def __init__(self, bind_cli):
        super().__init__(bind_cli)

    def add_action(self, wlt, action: Action):
        data_add_action = {
            "function_name": "addAction",
            "args": base64.b64encode(action.json().encode(ENC)).decode()
        }

        return self.exec_system_contract_method(
            wlt,
            b"systemCall\x19system.action\x18",
            json.dumps(data_add_action).encode(ENC),
            0
        )

    def get_action(self, wlt, isc_address: bytes or str, tid: int, aid: int):
        if isinstance(isc_address, str):
            isc_address = bytes.fromhex(isc_address)
        data_add_action = {
            "function_name": "getAction",
            "args": base64.b64encode(json.dumps({
                "1": base64.b64encode(isc_address).decode(),
                "2": tid,
                "3": aid
            }).encode(ENC)).decode()
        }

        return self.exec_system_contract_method(
            wlt,
            b"systemCall\x19system.action\x18",
            json.dumps(data_add_action).encode(ENC),
            0
        )
