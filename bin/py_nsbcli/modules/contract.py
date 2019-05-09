from io import BytesIO

from config import ENC
from py_nsbcli import TransactionHeader, Wallet
from py_nsbcli.modules import Client
from py_nsbcli.util import GoJson



class Contract(object):
    def __init__(self, cli: Client):
        self.cli = cli

    def broadcast_tx_commit(self, tx: str):
        return self.cli.get_json(self.cli.admin.broadcast_tx_commit_url, params={"tx": tx})

    @staticmethod
    def generate_contract_tx(transaction_type: bytes, contract_type: bytes, tx_header: bytes):
        bytes_buffer = BytesIO()

        bytes_buffer.write(transaction_type)
        bytes_buffer.write(b"\x19")
        bytes_buffer.write(contract_type)
        bytes_buffer.write(b"\x18")
        bytes_buffer.write(tx_header)

        return "0x" + bytes_buffer.getvalue().hex()

    def create_contract(self, contract_name: str or bytes, tx_header: TransactionHeader):
        if isinstance(contract_name, str):
            contract_name = contract_name.encode(ENC)

        self.broadcast_tx_commit(tx=Contract.generate_contract_tx(
            b"createContract",
            contract_name,
            tx_header.json().encode('utf-8')
        ))

    def exec_contract_method(self, contract_name: str or bytes, tx_header: TransactionHeader):
        if isinstance(contract_name, str):
            contract_name = contract_name.encode(ENC)

        self.broadcast_tx_commit(tx=Contract.generate_contract_tx(
            b"sendTransaction",
            contract_name,
            tx_header.json().encode('utf-8')
        ))

    def contract_invoke_helper(
        self,
        contract_name: str or bytes,
        contract_address, function_name, args: list, value,
        sig_or_wallet: Wallet or bytes,
        fr=None, encoder_helper=None, default_helper=None
    ):
        data = GoJson.dumps({
            "function_name": function_name,
            "args": dict(((str(idx + 1), arg) for idx, arg in enumerate(args)))
        }, cls=encoder_helper, default_helper=default_helper).encode(ENC)

        if isinstance(sig_or_wallet, Wallet):
            if fr is None:
                fr = sig_or_wallet.address(0)

            tx_header = TransactionHeader(fr, contract_address, data, value)
            tx_header.sign(sig_or_wallet)
        else:
            tx_header = TransactionHeader(fr, contract_address, data, value, sig_or_wallet)
        if fr is None:
            raise ValueError("from address is necessary")

        return self.exec_contract_method(contract_name, tx_header)

    def exec_system_contract_method(self, wlt: Wallet, prec_name: bytes, args: bytes, value: int):

        tx_header = TransactionHeader(wlt.address(0), None, args, value)
        tx_header.sign(wlt)

        bytes_buffer = BytesIO()
        bytes_buffer.write(prec_name)
        bytes_buffer.write(tx_header.json().encode('utf-8'))

        self.broadcast_tx_commit(tx="0x" + bytes_buffer.getvalue().hex())
