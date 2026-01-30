// Package snowflake provides a very simple Twitter snowflake generator and parser.
package utils

import (
	"encoding/binary"
	"math/rand"
	"net"
	"strconv"
	"sync"
	"time"
)

const (
	workerIdBits   = uint(10)                         // 机器id所占的位数
	sequenceBits   = uint(12)                         // 序列所占的位数
	workerIdMax    = int64(-1 ^ (-1 << workerIdBits)) // 支持的最大机器id数量
	sequenceMask   = int64(-1 ^ (-1 << sequenceBits)) // 毫秒内存列的最大值
	workerIdShift  = sequenceBits                     // 机器id左移位数
	timestampShift = sequenceBits + workerIdBits      // 时间戳左移位数 22位
)

type businessID int // 业务id
const (
	businessIDUserID         businessID = 1
	businessIDWalletID       businessID = 2
	businessIDCardID         businessID = 3
	businessID3dsID          businessID = 4
	businessIDReapTXID       businessID = 5
	businessIDCallbackTaskID businessID = 6
	businessGeneralID        businessID = 7
)

// 开始时间截数组
var twepoch = [6]int64{
	1290000000000,
	1614000000000,
	1638000000000,
	1234000000000,
	1602000000000,
	1601000000000,
}

// 定义结构体struct
type SnowFlake struct {
	sync.Mutex
	timestamp  int64
	workerId   int64
	sequence   int64
	businessId businessID
}

// 实例化snowflake
func NewSnowFlake(workerId int64, businessId businessID) *SnowFlake {
	// 判定workerid是否在合理范围内
	if workerId < 0 || workerId > workerIdMax {
		workerId = 1
	}

	return &SnowFlake{
		timestamp:  0,
		workerId:   workerId,
		sequence:   0,
		businessId: businessId,
	}
}

func (s *SnowFlake) Generate() uint64 {
	s.Lock()
	defer s.Unlock()

	now := time.Now().UnixMilli()
	// 如果是微妙时间戳、且处于同一秒
	if s.timestamp == now {
		// 确保sequence不会溢出,最大值为4095、保证1ms内最多生成4096个ID值
		s.sequence = (s.sequence + 1) & sequenceMask
		if s.sequence == 0 {
			// 如果当前时间小于等于当前时间戳
			for now <= s.timestamp {
				// 死循环等待下一个毫秒值
				now = time.Now().UnixMilli()
			}
		}
	} else {
		s.sequence = 0
	}
	// todo: 时钟回拨判断、如果获取到的时间戳比上一个小,则存在时钟回拨状态,需要抛出异常
	s.timestamp = now

	e := twepoch[int(s.businessId)%len(twepoch)]
	// 时间戳左移22位 | workeid左移12位 | sequence兜底
	r := (now-e)<<timestampShift | (s.workerId << workerIdShift) | (s.sequence)
	return uint64(r)
}

func (s *SnowFlake) GenerateWithPrefix(prefix uint64) (uint64, error) {
	var now int64
	func() {
		s.Lock()
		defer s.Unlock()
		now = time.Now().UnixMilli()
		// 如果是微妙时间戳、且处于同一秒
		if s.timestamp == now {
			// 确保sequence不会溢出,最大值为4095、保证1ms内最多生成4096个ID值
			s.sequence = (s.sequence + 1) & sequenceMask
			if s.sequence == 0 {
				// 如果当前时间小于等于当前时间戳
				for now <= s.timestamp {
					// 死循环等待下一个毫秒值
					now = time.Now().UnixMilli()
				}
			}
		} else {
			s.sequence = 0
		}
	}()

	// todo: 时钟回拨判断、如果获取到的时间戳比上一个小,则存在时钟回拨状态,需要抛出异常
	s.timestamp = now

	e := twepoch[int(s.businessId)%len(twepoch)]
	// 时间戳左移22位 | workeid左移12位 | sequence兜底
	r := (now-e)<<timestampShift | (s.workerId << workerIdShift) | (s.sequence)
	generated := uint64(r)
	len := len(strconv.FormatUint(prefix, 10))
	return strconv.ParseUint(strconv.FormatUint(prefix, 10)+strconv.FormatUint(generated, 10)[:19-len], 10, 64)
}

func getNodeId() int64 {

	interfaces, err := net.Interfaces()
	if err != nil {
		return int64(rand.Int63() % workerIdMax)
	}
	var out string
	for _, inter := range interfaces {
		out = out + inter.HardwareAddr.String()
	}
	return int64(binary.BigEndian.Uint64(Md5([]byte(out)))) % workerIdMax

}

var SnowFlakeWalletID = NewSnowFlake(getNodeId(), businessIDWalletID)
var SnowFlakeUserID = NewSnowFlake(getNodeId(), businessIDUserID)
var SnowFlakeCardID = NewSnowFlake(getNodeId(), businessIDCardID)
var SnowFlake3dsID = NewSnowFlake(getNodeId(), businessID3dsID)
var SnowFlakeReapTXID = NewSnowFlake(getNodeId(), businessIDReapTXID)
var SnowFlakeCallbackTaskID = NewSnowFlake(getNodeId(), businessIDCallbackTaskID)
var SnowFlakeID = NewSnowFlake(getNodeId(), businessGeneralID)
