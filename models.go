package main

type Downtime struct {
	Item_name      string
	Start_datetime string
	End_datetime   string
	Duration       float64
	Downtime_type  string
	Downtime_text  string
	Comment        string
}

type Pstn struct {
	Item_name      string
	Start_datetime string
	Pressure       string
	Temperature    string
	Level          string
	Level_ph       string
	Oil_dens       string
}

type WaterVol struct {
	Name     string
	Date     string
	WaterVol string
}

type WellTest struct {
	Name             string // CHANGE TO ID
	Date             string
	CassingPressure1 float64
	CassingPressure2 float64
	CassingPressure3 float64
	CassingPressure4 float64
	PumpRPM          float64
	PumpCurrent      float64
	PumpFrequency    float64
	FlowPress        float64
	PumpEfficiency   float64
	WHPerss          float64
	WHTemp           float64
}

type Wells struct {
	WELLNAME      string
	RUS_FORMATION string
	PRODUCT       string
	BATTERY       string
	DOME          string
	PERIOD        string
	RUS_ASSET     string
	FACILITY      string
	RUS_LIFT_TYPE string
	PUMP_MODEL    string
	RUS_WELLTYPE  string
}
