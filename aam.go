package nmea

const (
	// TypeAAM type of AAM sentence for Waypoint Arrival Alarm
	TypeAAM = "AAM"
)

// AAM - Waypoint Arrival Alarm
// This sentence is generated by some units to indicate the status of arrival (entering the arrival circle, or passing
// the perpendicular of the course line) at the destination waypoint (source: GPSD).
// https://gpsd.gitlab.io/gpsd/NMEA.html#_aam_waypoint_arrival_alarm
//
// Format: $--AAM,A,A,x.x,N,c--c*hh<CR><LF>
// Example: $GPAAM,A,A,0.10,N,WPTNME*43
type AAM struct {
	BaseSentence
	// StatusArrivalCircleEntered is warning of arrival to waypoint circle
	// * A = Arrival Circle Entered
	// * V = not entered
	StatusArrivalCircleEntered string

	// StatusPerpendicularPassed is warning for perpendicular passing of waypoint
	// * A = Perpendicular passed at waypoint
	// * V = not passed
	StatusPerpendicularPassed string

	// ArrivalCircleRadius is radius for arrival circle
	ArrivalCircleRadius float64

	// ArrivalCircleRadiusUnit is unit for arrival circle radius
	ArrivalCircleRadiusUnit string

	// DestinationWaypointID is destination waypoint ID
	DestinationWaypointID string
}

// newAAM constructor
func newAAM(s BaseSentence) (AAM, error) {
	p := NewParser(s)
	p.AssertType(TypeAAM)
	return AAM{
		BaseSentence:               s,
		StatusArrivalCircleEntered: p.EnumString(0, "arrival circle entered status", WPStatusArrivalCircleEnteredA, WPStatusArrivalCircleEnteredV),
		StatusPerpendicularPassed:  p.EnumString(1, "perpendicularly passed status", WPStatusPerpendicularPassedA, WPStatusPerpendicularPassedV),
		ArrivalCircleRadius:        p.Float64(2, "arrival circle radius"),
		ArrivalCircleRadiusUnit:    p.EnumString(3, "arrival circle radius units", DistanceUnitKilometre, DistanceUnitNauticalMile, DistanceUnitStatuteMile, DistanceUnitMetre),
		DestinationWaypointID:      p.String(4, "destination waypoint ID"),
	}, p.Err()
}
