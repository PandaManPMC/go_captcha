package rotate

// CheckAngle is to the detection angle
func CheckAngle(angle, dAngle, padding int64) bool {
	minAngle := 360 - padding
	maxAngle := 360 + padding
	angle += dAngle

	return angle >= minAngle && angle <= maxAngle
}
