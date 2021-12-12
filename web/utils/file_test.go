package utils

import (
	"fmt"
	"strings"
	"testing"
)

func TestString2File(t *testing.T) {

	for i := 0; i < 1000; i++ {
		Bytes2File([]byte(`[INFO ] 2018-06-23 00:00:04.158 [pool-7-thread-1] [com.jianxiaoli.spotmix.transfer.slave.EsDataTransfer$TransferTask] - liuheniu_json no message is available for consumption after the specified interval!
[INFO ] 2018-06-23 00:00:04.164 [pool-7-thread-2] [com.jianxiaoli.spotmix.transfer.slave.EsDataTransfer$TransferTask] - liuheniu_json no message is available for consumption after the specified interval!
[INFO ] 2018-06-23 00:00:04.403 [pool-10-thread-1] [com.jianxiaoli.spotmix.transfer.slave.EsDataTransfer$TransferTask] - liuheniu1_json no message is available for consumption after the specified interval!
[INFO ] 2018-06-23 00:00:04.417 [pool-10-thread-2] [com.jianxiaoli.spotmix.transfer.slave.EsDataTransfer$TransferTask] - liuheniu1_json no message is available for consumption after the specified interval!
`),
			"/home/tqyin/a.log")
	}
}
func TestReadFileBytes(t *testing.T) {
	data, _ := ReadFileBytes("/home/tqyin/business_log.log.2018-10-31", 1024*1024)
	fmt.Println(strings.Contains(string(data), "\n"))
}
