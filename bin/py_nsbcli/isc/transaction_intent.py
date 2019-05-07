
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


class TransactionIntent(object):
    def __init__(self, fr, to, seq, amt, meta):
        self.fr = fr
        self.to = to
        self.seq = seq
        self.amt = amt
        self.meta = meta

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
    def seq(self):
        return self._seq

    @seq.setter
    def seq(self, val: int):
        self._seq = val

    @property
    def amt(self):
        return self._amt

    @amt.setter
    def amt(self, val: int):
        self._amt = val

    @property
    def meta(self):
        return self._meta

    @meta.setter
    def meta(self, val: bytes):
        self._meta = val

    def dict(self):
        return {
            'from': base64.b64encode(self._fr).decode() if self._fr is not None else None,
            'to': base64.b64encode(self._to).decode() if self._to is not None else None,
            'seq': self._seq,
            'amt': self._amt,
            'meta': base64.b64encode(self._meta).decode() if self._meta is not None else None
        }

    def json(self):
        return json.dumps(self.dict())


if __name__ == '__main__':
    pass

