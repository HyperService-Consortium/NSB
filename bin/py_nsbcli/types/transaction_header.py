
"""
type TransactionHeader struct {
    From            []byte        `json:"from"`
    ContractAddress []byte        `json:"to"`
    JsonParas       []byte        `json:"data"`
    Value           *math.Uint256 `json:"value"`
    Nonce           *math.Uint256 `json:"nonce"`
    Signature       []byte        `json:"signature"`
}
"""

import json
import base64

from py_nsbcli.util.util import randint256
from py_nsbcli.util.cast import transbytes, transint, MOD256
from py_nsbcli.types.wallet import Wallet
from io import BytesIO


class TransactionHeader(object):
    def __init__(self, fr: bytes = None, to: bytes = None, dat: bytes = None, value=None, sig: bytes = None):
        self.fr = fr
        self.to = to
        self.data = dat
        self.value = value
        self._nonce = randint256()
        self.signature = sig

    @property
    def fr(self):
        return self._fr

    @fr.setter
    def fr(self, val: bytes):
        self._fr = val

    @property
    def to(self):
        return self._to

    @to.setter
    def to(self, val: bytes):
        self._to = val

    @property
    def data(self):
        return self._data

    @data.setter
    def data(self, val: bytes):
        self._data = val

    @property
    def value(self):
        return self._value

    @value.setter
    def value(self, val):
        if val is None:
            self._value = b"\x00"
            return

        if isinstance(val, bytes):
            if len(val) > 64:
                raise ValueError("value is not a valid uint256 number")
            self._value = val
            return

        val = transint(val)

        if val < 0 or val >= MOD256:
            raise ValueError("value is not a valid uint256 number")

        self._value = transbytes(val, 32)

    @property
    def signature(self):
        return self._signature

    @signature.setter
    def signature(self, val: bytes):
        self._signature = val

    def dict(self):
        return {
            'from': base64.b64encode(self._fr).decode() if self._fr is not None else None,
            'to': base64.b64encode(self._to).decode() if self._to is not None else None,
            'data': base64.b64encode(self._data).decode() if self._data is not None else None,
            'value': base64.b64encode(self._value).decode() if self._value is not None else None,
            'nonce': base64.b64encode(self._nonce).decode() if self._nonce is not None else None,
            'signature': base64.b64encode(self._signature).decode()     if self._signature is not None else None
        }

    @property
    def body(self):
        bytes_buffer = BytesIO()
        bytes_buffer.write(self._fr if self._fr is not None else b"")
        bytes_buffer.write(self._to if self._to is not None else b"")
        bytes_buffer.write(self._data if self._data is not None else b"")
        bytes_buffer.write(self._value if self._value is not None else b"")
        bytes_buffer.write(self._nonce if self._nonce is not None else b"")
        return bytes_buffer.getvalue()

    def json(self):
        return json.dumps(self.dict())

    def sign(self, wlt: Wallet):
        self.signature = wlt.sign(self.body)


if __name__ == '__main__':
    tx_header = TransactionHeader(b"a", b"b", b"data", None)

    print(tx_header.body)
