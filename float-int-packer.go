package floatint

//MaxSafeForInt16 is maximum float value that can be safely packed into int16
const MaxSafeForInt16 float64 = 8.189

//MaxSafeForInt32 is maximum float value that can be safely packed into int32
const MaxSafeForInt32 float64 = 536870.875

/*
 * Using 16 byte (int16) to store small floating point numbers
 *   2 bytes used for decimal places (2^2 = 4 decimal places)
 *      can store up to 4 decimal places
 *      but encoding upto 3 decimal places only
 *
 *   remaing 13 bytes used for value, sign bit not used
 *
 *   Max value can be packed is 8.189
 *
 */

//Pack convert given float value to int16
//If given float value is > MaxSafeForInt16 then it will return 0
func Pack(value float64) int16 {
	if value > MaxSafeForInt16 { //max value that can be safely encoded tp int16
		return 0
	}

	// get up to 3 decimal places
	decplaces := 3
	//val := math.Floor(value)

	for i := 0; i < decplaces; i++ {
		//log.Printf("Value: %f, Val: %f   ok: %v\n", value, val, ok)
		value *= 10.0
		//cnt++
	}
	return int16(((decplaces << 13) + int(value)))
}

//Unpack convert given int16 back to float value
func Unpack(value int16) float32 {
	decplaces := 3 //value >> 12
	//log.Printf("Cnt: %d\n", decplaces)

	result := float32(value & 0x1fff)
	//log.Printf("result: %f\n", result)

	for ok := true; ok; ok = (decplaces > 0) {
		result /= 10.0
		decplaces--
	}
	return result
}

/*
 * Using 32 byte (int) to store small floating point numbers
 *   2 bytes used for decimal places (2^2 = 4 decimal places)
 *      can store up to 4 decimal places
 *      but encoding upto 3 decimal places only
 *
 *   remaing 29 bytes used for value, sign bit not used
 *
 *   Max value can be safely packed is 536870.875
 *
 */

//Pack32 convert given float value to int32
//If given float value is > MaxSafeForInt32 then it will return 0
func Pack32(value float64) int {
	if value > MaxSafeForInt32 { //max value that can be safely encoded to int32
		return 0
	}
	// get up to 3 decimal places
	decplaces := 3
	for i := 0; i < decplaces; i++ {
		value *= 10.0
	}
	return int(((decplaces << 29) + int(value)))
}

//Unpack32 convert given int32 back to float value
func Unpack32(value int) float64 {
	decplaces := 3
	result := float64(value & 0x1fffffff)
	for ok := true; ok; ok = (decplaces > 0) {
		result /= 10.0
		decplaces--
	}
	return result
}
