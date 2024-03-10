package tunda

type Kerjaan struct {
	Data         interface{} `json:"data"`
	TimeDuration int         `json:"time_duration"`
}

type KerjaanID string

const (
	KEY_LIST_OF_JOBS string = "tunda-daftar-kerjaan"
	JOB_PREFIX       string = "tunda-kerjaan-"
)