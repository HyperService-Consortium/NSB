
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
