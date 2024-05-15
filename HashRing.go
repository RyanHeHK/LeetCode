package main

import (
	"fmt"
	"hash/crc32"
	"sort"
)

type HashRing struct {
	replicas   int               // 虚拟节点的复制因子
	hashRing   []uint32          // 哈希环，存储节点的哈希值
	nodeToHash map[uint32]string // 哈希值到节点的映射
}

func NewHashRing(replicas int) *HashRing {
	return &HashRing{
		replicas:   replicas,
		hashRing:   []uint32{},
		nodeToHash: make(map[uint32]string),
	}
}

func (hr *HashRing) AddNode(node string) {
	for i := 0; i < hr.replicas; i++ {
		replicaKey := hr.getReplicaKey(node, i)
		hashValue := hr.hash(replicaKey)
		hr.hashRing = append(hr.hashRing, hashValue)
		hr.nodeToHash[hashValue] = node
	}

	sort.Slice(hr.hashRing, func(i, j int) bool {
		return hr.hashRing[i] < hr.hashRing[j]
	})
}

func (hr *HashRing) RemoveNode(node string) {
	for i := 0; i < hr.replicas; i++ {
		replicaKey := hr.getReplicaKey(node, i)
		hashValue := hr.hash(replicaKey)
		delete(hr.nodeToHash, hashValue)

		index := -1
		for j, val := range hr.hashRing {
			if val == hashValue {
				index = j
				break
			}
		}
		if index != -1 {
			hr.hashRing = append(hr.hashRing[:index], hr.hashRing[index+1:]...)
		}
	}
}

func (hr *HashRing) GetNode(key string) string {
	if len(hr.hashRing) == 0 {
		return ""
	}

	hashValue := hr.hash(key)
	index := hr.findIndex(hashValue)
	return hr.nodeToHash[hr.hashRing[index]]
}

func (hr *HashRing) findIndex(hashValue uint32) int {
	lo := 0
	hi := len(hr.hashRing) - 1

	for lo < hi {
		mid := (lo + hi) / 2
		if hr.hashRing[mid] < hashValue {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	if lo < len(hr.hashRing) {
		return lo
	}
	return 0
}

func (hr *HashRing) getReplicaKey(node string, replicaIndex int) string {
	return fmt.Sprintf("%s-%d", node, replicaIndex)
}

func (hr *HashRing) hash(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key))
}

//func main() {
//	nodes := []string{"Node1", "Node2", "Node3"}
//	hashRing := NewHashRing(3)
//
//	for _, node := range nodes {
//		hashRing.AddNode(node)
//	}
//
//	// 添加新节点
//	hashRing.AddNode("Node4")
//
//	// 移除节点
//	hashRing.RemoveNode("Node2")
//
//	// 查找节点
//	key := "data_key"
//	node := hashRing.GetNode(key)
//	fmt.Printf("Data with key '%s' is stored on node: %s\n", key, node)
//}
