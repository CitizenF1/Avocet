package main

type DowntimeUpdate struct {
	AssetID        int64
	Item_name      string
	Start_datetime string
	End_datetime   string
	Duration       float64
	Downtime_type  string
	Downtime_text  string
	Comment        string
}

type Pstn struct {
	AssetID        int64
	Item_name      string
	Start_datetime string
	Pressure       string
	Temperature    string
	Level          string
	Level_ph       string
	Oil_dens       string
}

type WaterVol struct {
	AssetID  int64
	Name     string
	Date     string
	WaterVol string
}

type WellTest struct {
	ID               int64 // CHANGE TO ID
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
