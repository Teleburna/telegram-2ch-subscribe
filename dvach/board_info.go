package dvach

import (
	"encoding/json"
	"fmt"
	"sort"
)

const (
	ThreadUrl = "https://2ch.hk/%s/res/%s.html"
)

type ThreadInfo struct {
	ID        string `json:"num"`
	Subject   string `json:"subject"`
	Timestamp int64  `json:"timestamp"`
}

type ByTimestamp []ThreadInfo

func (threads ByTimestamp) Len() int           { return len(threads) }
func (threads ByTimestamp) Swap(i, j int)      { threads[i], threads[j] = threads[j], threads[i] }
func (threads ByTimestamp) Less(i, j int) bool { return threads[i].Timestamp < threads[j].Timestamp }

type BoardInfo struct {
	Board   string       `json:"board"`
	Threads []ThreadInfo `json:"threads"`
}

func NewBoardInfo(jsonBoardData []byte) *BoardInfo {
	boardInfo := BoardInfo{}
	json.Unmarshal(jsonBoardData, &boardInfo)

	return &boardInfo
}

func (boardInfo *BoardInfo) ThreadsAfter(timestamp int64) []ThreadInfo {
	threads := make([]ThreadInfo, 0)
	if boardInfo == nil || boardInfo.Threads == nil {
		return threads
	}
	for _, thread := range boardInfo.Threads {
		if thread.Timestamp > timestamp {
			threads = append(threads, thread)
		}
	}
	sort.Sort(ByTimestamp(threads))

	return threads
}

func (boardInfo *BoardInfo) ThreadUrl(id string) string {
	return fmt.Sprintf(ThreadUrl, boardInfo.Board, id)
}