
"""
type TransactionIntent struct {
    Fr   []byte `json:"from"`
    To   []byte `json:"to"`
    Seq  uint   `json:"seq"`
    Amt  uint   `json:"amt"`
    Meta []byte `json:"meta"`
}
"""

import json, base64

try:
    from hexbytes import HexBytes
except ImportError or ModuleNotFoundError as e:
    print(e)
    exit(1)


class Action(object):
    def __init__(self, isc_address, tid, aid, action_type, content, signature):
        self.addr = isc_address
        self.tid = tid
        self.aid = aid
        self.aty = action_type
        self.content = content
        self.signature = signature

    @property
    def addr(self):
        return self._addr

    @addr.setter
    def addr(self, val: bytes or str):
        if isinstance(val, bytes):
            self._addr = val
        else:
            self._addr = bytes.fromhex(val)

    @property
    def tid(self):
        return self._tid

    @tid.setter
    def tid(self, val: int):
        self._tid = val

    @property
    def aid(self):
        return self._aid

    @aid.setter
    def aid(self, val: int):
        self._aid = val

    @property
    def aty(self):
        return self._aty

    @aty.setter
    def aty(self, val: int):
        self._aty = val

    @property
    def content(self):
        return self._content

    @content.setter
    def content(self, val: bytes):
        self._content = val

    @property
    def signature(self):
        return self._signature

    @signature.setter
    def signature(self, val: bytes):
        self._signature = val

    def dict(self):
        return {
            '1': base64.b64encode(self._addr).decode() if self._addr is not None else None,
            '2': self._tid,
            '3': self._aid,
            '4': self._aty,
            '5': base64.b64encode(self._content).decode() if self._content is not None else None,
            '6': base64.b64encode(self._signature).decode() if self._signature is not None else None
        }

    def json(self):
        return json.dumps(self.dict())


if __name__ == '__main__':
    pass

