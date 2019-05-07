

import base64
import json

from py_nsbcli.types import TransactionHeader
from py_nsbcli.modules.contract import Contract


class ISC(Contract):
    def __init__(self, bind_cli):
        super().__init__(bind_cli)

    def create_isc(self, wlt, owners, funds, value, intents):
        owners = [base64.b64encode(owner).decode() for owner in owners]
        args_isc = {
            "isc_owners": owners,
            "required_funds": funds,
            "ves_signature": base64.b64encode(b"").decode(),
            "transactionIntents": [base64.b64encode(intent.json()).decode() for intent in intents]
        }
        tx_header = TransactionHeader(wlt.address(0), None, json.dumps(args_isc).encode(), value)
        tx_header.sign(wlt)
        return self.create_contract("isc", tx_header)
