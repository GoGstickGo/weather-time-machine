package views

type Data struct {
	Alert *Alert
	Yield interface{}
}

type Alert struct {
	Level   string
	Message string
}

const (
	AlertLvlError   = "danger"
	AlertLvlWarning = "warning"
	AlertLvlInfo    = "info"
	AlertLvlSuccess = "success"

	AlertGenericMsg = "Something went wrong. Please refresh the page " +
		"and contact us if the problem persists."
)
