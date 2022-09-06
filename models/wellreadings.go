package models

type WellReadings struct {
	AssetID          int64
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

type WaterVol struct {
	AssetID  int64
	Name     string
	Date     string
	WaterVol string
}
