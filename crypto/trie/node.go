package trie
import (
	//"fmt"
	//"bytes"
	"errors"
	"encoding/json"
	_ "encoding/hex"
	// "github.com/tendermint/tendermint/rlp"
	dbm "github.com/tendermint/tendermint/libs/db"
)

type HashFunc func(...[]byte) []byte
type TrieNode struct {
	Nodes [][]byte `json:"nodes"`
}

type MerklePatricaTree struct {
	RootHash MerkleHash
	Hasher HashFunc
	DBer dbm.DB
}

func NewTrieNode(jsonedNode []byte) (*TrieNode, error) {
	var node *TrieNode = nil
	err := json.Unmarshal(jsonedNode, *node);
	if err != nil {
		// log it after
		return nil, err
	}
	return node, nil
}
func NewTrieNodeByBytess(dat ...[]byte) (*TrieNode) {
	var node TrieNode
	node.Nodes = append(node.Nodes, dat...)
	return &node
}


func (trnode *TrieNode) Extend() {
	for idx := 0; idx < 16; idx++ {
		
	}
}

func getNode(db dbm.DB, toGet []byte) (*TrieNode, error) {
	return NewTrieNode(db.Get(toGet))
}

func (mpt *MerklePatricaTree) Insert(key []byte, value []byte) (MerkleHash, error) {
	if mpt.RootHash == nil {// mpt is nil
		hashnode, err := json.Marshal(NewTrieNodeByBytess(key, value))
		if err != nil {
			return nil, err
		}
		hashBytes := mpt.Hasher(hashnode)
		mpt.DBer.Set(hashBytes, value)
		return hashBytes, nil
	} else {
		hashnode, serialNode, err := mpt.RootHash, new(TrieNode), errors.New("...")
		 for {
			serialNode, err = getNode(mpt.DBer, hashnode)
			if err != nil {
				return nil, err
			}
			if serialNode == nil {
				hashnode, err = json.Marshal(NewTrieNodeByBytess(key, value))
				if err != nil {
					return nil, err
				}
				hashBytes := mpt.Hasher(hashnode)
				mpt.DBer.Set(hashBytes, value)
				return hashBytes, nil
			} else if len(serialNode.Nodes) == 2 {
				if len(serialNode.Nodes[0]) < len(key) {
					if len(serialNode.Nodes[0]) == 1 {
						serialNode.Extend()
						newNode := NewTrieNodeByBytess(key[1:], value)
						hashnode, err = json.Marshal(newNode)
						if err != nil {
							return nil, err
						}
						return mpt.Hasher(hashnode), nil
					}
				}
			} else if len(serialNode.Nodes) == 17 {
				hashnode = serialNode.Nodes[key[0]]
				key = key[1:]
			} else {
				return nil, errors.New("un recognized node type");
			}
		}
	}
}
