package logiops

type Pid int

type LogiData struct {
	Ignore  []Pid
	Devices []LogiDevice
}

type LogiDevice struct {
	Name string
	//Buttons []LogiButton
	Dpi        int
	Smartshift *LogiSmartshift
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

type LogiAction struct {
	ActionTyoe string
	Action     interface{}
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

type LogiActionKeypress struct {
}
