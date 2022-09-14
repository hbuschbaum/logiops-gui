package logiops

type Pid int

type LogiData struct {
	Ignore  []Pid
	Devices []LogiDevice
}

type LogiDevice struct {
	Name string
	//Buttons []LogiButton
	//Dpi []int
	//Smartshift LogiSmartshift
	//Hiresscroll LogiHiresscroll
	//Thumbwheel LogiThumbwheel
}

type LogiButton struct {
	Cid    int
	Action LogiAction
}

type LogiSmartshift struct {
	On               bool
	Threshold        int
	DefaultThreshold int
}

type LogiHiresscroll struct {
	Hires, Invert, Target bool
}

type LogiThumbwheel struct {
	Divert, Invert bool
	Left, Right    LogiActionGesture
}

type LogiAction interface {
	getInfo() string
}

type LogiActionGesture struct {
	Gestures []LogiGestures
}

type LogiGestures struct {
	Direction string
	Threshold int
	Mode      string
	Action    LogiActionKeypress
}

func (LogiActionGesture) getInfo() string {
	return ""
}

type LogiActionKeypress struct {
}

func (LogiActionKeypress) getInfo() string {
	return ""
}
