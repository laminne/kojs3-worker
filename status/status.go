package status

type RunningStatus struct {
	TaskID string       // 問題ID
	Status []TestStatus // テストケースごと
}

type TestStatus struct {
	TestID string // テストケースのID
	ExitStatus int // 実行時のステータスコード
	Duration int // 実行秒数
	Status string // WA RE などのステータス
	Memory int64 // メモリ使用量
}
