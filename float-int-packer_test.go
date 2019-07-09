package floatint

import (
	"fmt"
	"log"
	"testing"
)

func TestInt16(t *testing.T) {
	log.Println("\nTestInt16")
	vlu := []float64{1.0, 5.671, 8.188, 8.189, 8.191}
	for _, v := range vlu {
		retint16 := Pack(v)
		log.Printf("Encoded %f to %d", v, retint16)

		retvlu := Unpack(retint16)
		log.Printf("Decoded %d to %f", retint16, retvlu)

		if v > MaxSafeForInt16 {
			if retvlu != 0.0 {
				t.Errorf("Encode/Decode failed. exp: %f, got %f", v, retvlu)
			}
		} else {
			if float32(v) != retvlu {
				t.Errorf("Encode/Decode failed. Exp: %f, got %f", v, retvlu)
			}
		}
	}
}

func TestInt32(t *testing.T) {
	log.Println("\nTestInt32")
	vlu := []float64{1.0, 8.189, 16.36, 215341.273, 536870.875, 536870.876}
	for _, v := range vlu {
		retint32 := Pack32(v)
		fmt.Printf("Encoded %f to %d", v, retint32)

		retvlu := Unpack32(retint32)
		fmt.Printf("\tDecoded %d to %f", retint32, retvlu)

		if v > MaxSafeForInt32 {
			if retvlu != 0.0 {
				t.Errorf("Encode32/Decode32 failed. exp: %f, got %f", v, retvlu)
			}
		} else {
			if v != retvlu {
				t.Errorf("Encode32/Decode32 failed. exp: %f, got %f", v, retvlu)
			}
		}
		fmt.Println("")
	}
}
