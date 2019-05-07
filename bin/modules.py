
try:
    import requests
except ImportError or ModuleNotFoundError as e:
    print(e)
    exit(1)

import json, base64

from config import HTTP_HEADERS, ENC
from io import BytesIO
from py_nsbcli.cast import transbytes
from py_nsbcli import TransactionHeader, Action, Wallet, LevelDB



class Admin(object):
    def __init__(self):
        self.rpc_host = None

        self._abci_info_url = None
        self._abci_query_url = None
        self._block_url = None
        self._block_result_url = None
        self._block_chain_url = None
        self._broadcast_tx_async_url = None
        self._broadcast_tx_commit_url = None
        self._broadcast_tx_sync_url = None
        self._commit_url = None
        self._consensus_params_url = None
        self._dump_consensus_url = None
        self._genesis_url = None
        self._health_url = None
        self._net_info_url = None
        self._num_unconfirmed_txs_url = None
        self._status_url = None
        self._subscrible_url = None
        self._tx_url = None
        self._tx_search_url = None
        self._unconfirmed_txs_url = None
        self._unsubscribe_url = None
        self._unsubscribe_all_url = None
        self._validatos_url = None

    @property
    def abci_info_url(self):
        return self._abci_info_url

    @property
    def abci_query_url(self):
        return self._abci_query_url

    @property
    def broadcast_tx_commit_url(self):
        return self._broadcast_tx_commit_url

    def set_rpc_host(self, host_name):
        self.rpc_host = host_name

        self._abci_info_url = self.rpc_host + "/abci_info"
        self._abci_query_url = self.rpc_host + "/abci_query"
        self._broadcast_tx_commit_url = self.rpc_host + "/broadcast_tx_commit"

        print(host_name)


class Client(object):
    def __init__(self, bind_admin: Admin):
        self._admin = bind_admin
        self.http_header = HTTP_HEADERS

    @property
    def admin(self):
        return self._admin

    def set_http_header(self, new_header):
        self.http_header = new_header

    def get(self, url, params=None):
        response = requests.get(
            url,
            headers=self.http_header,
            params=params
        )
        if response.status_code != 200:
            raise Exception(response)
        for k, v in response.__dict__.items():
            print(k, v)
        return response.content

    def get_json(self, url, params=None):
        return json.loads(self.get(url, params), encoding=ENC)

    @staticmethod
    def generate_contract_tx(transaction_type: bytes, contract_type: bytes, tx_header: bytes):
        bytes_buffer = BytesIO()

        bytes_buffer.write(transaction_type)
        bytes_buffer.write(b"\x19")
        bytes_buffer.write(contract_type)
        bytes_buffer.write(b"\x18")
        bytes_buffer.write(tx_header)

        return "0x" + bytes_buffer.getvalue().hex()

    def abci_info(self):
        response = self.get_json(self._admin.abci_info_url)
        print(json.dumps(response, sort_keys=True, indent=4))

    def broadcast_tx_commit(self, tx: str):
        response = self.get_json(self._admin.broadcast_tx_commit_url, params={"tx": tx})
        print(response)

    def create_contract(self, name: str or bytes, tx_header: TransactionHeader):
        if isinstance(name, str):
            name = name.encode(ENC)

        self.broadcast_tx_commit(tx=Client.generate_contract_tx(
            b"createContract",
            name,
            tx_header.json().encode('utf-8')
        ))

    def exec_contract_method(self, name: str or bytes, tx_header: TransactionHeader):
        if isinstance(name, str):
            name = name.encode(ENC)

        self.broadcast_tx_commit(tx=Client.generate_contract_tx(
            b"sendTransaction",
            name,
            tx_header.json().encode('utf-8')
        ))

    def exec_system_contract_method(self, wlt: Wallet, prec_name: bytes, args: bytes, value: int):

        tx_header = TransactionHeader(wlt.address(0), None, args, value)
        tx_header.sign(wlt)

        bytes_buffer = BytesIO()
        bytes_buffer.write(prec_name)
        bytes_buffer.write(tx_header.json().encode('utf-8'))

        self.broadcast_tx_commit(tx="0x" + bytes_buffer.getvalue().hex())

    def create_isc(self, wlt, owners, funds, value, intents):
        owners = [base64.b64encode(owner).decode() for owner in owners]
        args_isc = {
            "isc_owners": owners,
            "required_funds": funds,
            "ves_signature": base64.b64encode(b"").decode(),
            "transactionIntents": [base64.b64encode(intent.json()).decode() for intent in intents]
        }
        tx_header = TransactionHeader(wlt.address(0), None, json.dumps(args_isc).encode(ENC), value)
        tx_header.sign(wlt)
        return self.create_contract("isc", tx_header)

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
        x = Wallet.create(self._dbhandler, name)
        if x.handler_num < 0:
            raise Exception("create failed")
        return x
