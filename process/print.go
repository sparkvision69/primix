package process

import (
	"os"
)

// Print function for printing messages in a custom way with color and font style
func Print(message interface{}) {
	// Custom format with timestamp (for example)
	var messageStr string

	// ANSI escape codes for color and style
	// Bold, Blue text for "PROCESS", and Italic style
	processPrefix :=  "\033[1;34;3mPROCESS\033[0m " // Blue, Bold, and Italic text

	// Handle different types manually
	switch v := message.(type) {
	case string:
		messageStr = v
	case int:
		messageStr = intToString(v)
	case int32:
		messageStr = intToString(int(v))
	case int64:
		messageStr = intToString(int(v))
	case float32:
		messageStr = floatToString(float64(v)) // Convert float32 to float64
	case float64:
		messageStr = floatToString(v)
	default:
		messageStr = "Unsupported type"
	}

	// Custom formatted output with color and style
	output := []byte(processPrefix + messageStr + "\n")
	_, _ = os.Stdout.Write(output)
}

// Manually convert an int to a string
func intToString(i int) string {
	if i == 0 {
		return "0"
	}

	var result string
	isNegative := i < 0
	if isNegative {
		i = -i
	}

	for i > 0 {
		result = string(rune(i%10+'0')) + result
		i = i / 10
	}

	if isNegative {
		result = "-" + result
	}

	return result
}

// Manually convert a float to a string (basic version for float32 and float64)
func floatToString(f float64) string {
	if f == 0 {
		return "0"
	}

	var result string
	if f < 0 {
		f = -f
		result = "-"
	}

	// Convert the integer part
	intPart := int(f)
	result += intToString(intPart)

	// Convert the fractional part
	fraction := f - float64(intPart)
	if fraction > 0 {
		result += "."
		for fraction > 0 && len(result) < 15 { // Limit precision
			fraction *= 10
			result += intToString(int(fraction))[0:1]
			fraction -= float64(int(fraction))
		}
	}

	return result
}
