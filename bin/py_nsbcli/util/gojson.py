
import json
import base64


class GoJsonEncoder(json.JSONEncoder):
    """
    Go Json Encoder
    +-------------------+---------------+
    | Python            | JSON          |
    +===================+===============+
    | dict              | object        |
    +-------------------+---------------+
    | list, tuple       | array         |
    +-------------------+---------------+
    | str               | string        |
    +-------------------+---------------+
    | int, float        | number        |
    +-------------------+---------------+
    | True              | true          |
    +-------------------+---------------+
    | False             | false         |
    +-------------------+---------------+
    | None              | null          |
    +-------------------+---------------+
    | bytes             | base64-string |
    +-------------------+---------------+
    """
    def default(self, obj):
        if isinstance(obj, bytes):
            return base64.b64encode(obj).decode('utf-8')
        # elif isinstance(obj, dict):
        #     return base64.b16decode(json.dumps(obj, cls=GoJsonEncoder).encode('utf-8')).decode('utf-8')
        return json.JSONEncoder.default(self, obj)


class GoJson:

    @staticmethod
    def dump(obj, fp, *args, **kwargs):
        if 'cls' not in kwargs:
            kwargs['cls'] = GoJsonEncoder
        json.dump(obj, fp, *args, **kwargs)

    @staticmethod
    def dumps(obj, *args, **kwargs):
        if 'cls' not in kwargs:
            kwargs['cls'] = GoJsonEncoder
        return json.dumps(obj, *args, **kwargs)

    @staticmethod
    def load(fp, *args, **kwargs):
        return json.load(fp, *args, **kwargs)

    @staticmethod
    def loads(s, *args, **kwargs):
        return json.loads(s, *args, **kwargs)


def action_hook(dct):
    if 'testbytes' in dct:
        dct['testbytes'] = base64.b64decode(dct['testbytes'])
        return dct
    return dct


if __name__ == "__main__":
    action = {
        "testbytes": b"12",
        "tees_ojb_bytes": {
            "testbytes": b"12"
        }
    }
    print(action)

    with open('test.json', 'w') as fp:
        GoJson.dump(action, fp)
        print(GoJson.dumps(action))

    with open('test.json', 'r') as fp:
        aut = GoJson.load(fp)
        print(GoJson.loads(GoJson.dumps(action), object_hook=action_hook))
