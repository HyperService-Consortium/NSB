import json

try:
    import requests
except ImportError or ModuleNotFoundError as e:
    print(e)
    exit(1)

from config import HTTP_HEADERS, ENC
from py_nsbcli.modules.admin import Admin


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
        # for k, v in response.__dict__.items():
        #     print(k, v)
        return response.content

    def get_json(self, url, params=None):
        return json.loads(self.get(url, params), encoding=ENC)

    def abci_info(self):
        response = self.get_json(self._admin.abci_info_url)
        print(json.dumps(response, sort_keys=True, indent=4))

    def append_module(self, name: str, sub_module):
        setattr(self, name, sub_module)

