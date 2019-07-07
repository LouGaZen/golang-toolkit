package conv

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBoolValue(t *testing.T) {
	cases := []struct {
		value  interface{}
		result bool
		hasErr bool
	}{
		{nil, false, true},
		{true, true, false},
		{false, false, false},
		{"abc", false, true},
	}

	Convey("TestBoolValue", t, func() {
		for i, c := range cases {
			Convey(fmt.Sprintf("case %v", i), func() {
				res, err := BoolValue(c.value)
				if c.hasErr {
					So(err, ShouldBeError)
				} else {
					So(err, ShouldBeNil)
				}
				So(res, ShouldEqual, c.result)
			})
		}
	})
}

func TestBoolDefault(t *testing.T) {
	cases := []struct {
		value        interface{}
		defaultValue bool
		result       bool
	}{
		{nil, false, false},
		{nil, true, true},
		{true, false, true},
		{false, true, false},
		{"abc", true, true},
	}

	Convey("TestBoolDefault", t, func() {
		for i, c := range cases {
			Convey(fmt.Sprintf("case %v", i), func() {
				So(BoolDefault(c.value, c.defaultValue), ShouldEqual, c.result)
			})
		}
	})
}

func TestBool(t *testing.T) {
	cases := []struct {
		value  interface{}
		result bool
	}{
		{nil, false},
		{true, true},
		{false, false},
		{"abc", false},
	}

	Convey("TestBool", t, func() {
		for i, c := range cases {
			Convey(fmt.Sprintf("case %v", i), func() {
				So(Bool(c.value), ShouldEqual, c.result)
			})
		}
	})
}

func TestBoolForce(t *testing.T) {
	channel := make(chan int, 1)
	channel <- 1

	cases := []struct {
		value  interface{}
		result bool
	}{
		{nil, false},
		{true, true},
		{false, false},

		{int(0), false},
		{int(-1), true},
		{int8(0), false},
		{int8(-1), true},
		{int16(0), false},
		{int16(-1), true},
		{int32(0), false},
		{int32(-1), true},
		{int64(0), false},
		{int64(-1), true},

		{uint(0), false},
		{uint(1), true},
		{uint8(0), false},
		{uint8(1), true},
		{uint16(0), false},
		{uint16(1), true},
		{uint32(0), false},
		{uint32(1), true},
		{uint64(0), false},
		{uint64(1), true},

		{float32(0), false},
		{float32(1.1), true},
		{float64(0), false},
		{float64(-1.1), true},

		{[]byte{}, false},
		{[]byte{1}, true},
		{make(chan bool, 1), false},
		{channel, true},
		{map[int]int{}, false},
		{map[int]int{1: 1}, true},
		{"", false},
		{"abc", true},
		{func() {}, true},
	}

	Convey("TestBoolForce", t, func() {
		for i, c := range cases {
			Convey(fmt.Sprintf("case %v", i), func() {
				So(BoolForce(c.value), ShouldEqual, c.result)
			})
		}
	})
}

func TestIntValue(t *testing.T) {
	cases := []struct {
		value  interface{}
		result int64
		hasErr bool
	}{
		{nil, 0, true},
		{true, 1, false},
		{false, 0, false},
		{int(-123), -123, false},
		{int8(-123), -123, false},
		{int16(-123), -123, false},
		{int32(-123), -123, false},
		{int64(-123), -123, false},
		{uint(123), 123, false},
		{uint8(123), 123, false},
		{uint16(123), 123, false},
		{uint32(123), 123, false},
		{uint64(123), 123, false},
		{float32(123.45), 123, false},
		{float64(-123.45), -123, false},
		{"123", 123, false},
		{"-123", -123, false},
		{"-123.45", 0, true},
		{"abc", 0, true},
		{func() {}, 0, true},
	}

	Convey("TestIntValue", t, func() {
		for i, c := range cases {
			Convey(fmt.Sprintf("case %v", i), func() {
				res, err := IntValue(c.value)
				if c.hasErr {
					So(err, ShouldBeError)
				} else {
					So(err, ShouldBeNil)
				}
				So(res, ShouldEqual, c.result)
			})
		}
	})
}

func TestIntDefault(t *testing.T) {
	cases := []struct {
		value        interface{}
		defaultValue int64
		result       int64
	}{
		{nil, 1, 1},
		{true, 0, 1},
		{false, 1, 0},
		{int(-123), 1, -123},
		{int8(-123), 1, -123},
		{int16(-123), 1, -123},
		{int32(-123), 1, -123},
		{int64(-123), 1, -123},
		{uint(123), 1, 123},
		{uint8(123), 1, 123},
		{uint16(123), 1, 123},
		{uint32(123), 1, 123},
		{uint64(123), 1, 123},
		{float32(123.45), 1, 123},
		{float64(-123.45), 1, -123},
		{"123", 1, 123},
		{"-123", 1, -123},
		{"-123.45", 1, 1},
		{"abc", 1, 1},
		{func() {}, 1, 1},
	}

	Convey("TestIntDefault", t, func() {
		for i, c := range cases {
			Convey(fmt.Sprintf("case %v", i), func() {
				So(IntDefault(c.value, c.defaultValue), ShouldEqual, c.result)
			})
		}
	})
}

func TestInt(t *testing.T) {
	cases := []struct {
		value  interface{}
		result int64
	}{
		{nil, 0},
		{true, 1},
		{false, 0},
		{int(-123), -123},
		{int8(-123), -123},
		{int16(-123), -123},
		{int32(-123), -123},
		{int64(-123), -123},
		{uint(123), 123},
		{uint8(123), 123},
		{uint16(123), 123},
		{uint32(123), 123},
		{uint64(123), 123},
		{float32(123.45), 123},
		{float64(-123.45), -123},
		{"123", 123},
		{"-123", -123},
		{"-123.45", 0},
		{"abc", 0},
		{func() {}, 0},
	}

	Convey("TestInt", t, func() {
		for i, c := range cases {
			Convey(fmt.Sprintf("case %v", i), func() {
				So(Int(c.value), ShouldEqual, c.result)
			})
		}
	})
}

func TestUintValue(t *testing.T) {
	cases := []struct {
		value  interface{}
		result uint64
		hasErr bool
	}{
		{nil, 0, true},
		{true, 1, false},
		{false, 0, false},
		{int(123), 123, false},
		{int8(123), 123, false},
		{int16(123), 123, false},
		{int32(123), 123, false},
		{int64(123), 123, false},
		{uint(123), 123, false},
		{uint8(123), 123, false},
		{uint16(123), 123, false},
		{uint32(123), 123, false},
		{uint64(123), 123, false},
		{float32(123.45), 123, false},
		{float64(123.45), 123, false},
		{"123", 123, false},
		{"123.45", 0, true},
		{"abc", 0, true},
		{func() {}, 0, true},
	}

	Convey("TestUintValue", t, func() {
		for i, c := range cases {
			Convey(fmt.Sprintf("case %v", i), func() {
				Convey(fmt.Sprintf("case %v", i), func() {
					res, err := UintValue(c.value)
					if c.hasErr {
						So(err, ShouldBeError)
					} else {
						So(err, ShouldBeNil)
					}
					So(res, ShouldEqual, c.result)
				})
			})
		}
	})
}

func TestUintDefault(t *testing.T) {
	cases := []struct {
		value        interface{}
		defaultValue uint64
		result       uint64
	}{
		{nil, 1, 1},
		{true, 0, 1},
		{false, 1, 0},
		{int(123), 1, 123},
		{int8(123), 1, 123},
		{int16(123), 1, 123},
		{int32(123), 1, 123},
		{int64(123), 1, 123},
		{uint(123), 1, 123},
		{uint8(123), 1, 123},
		{uint16(123), 1, 123},
		{uint32(123), 1, 123},
		{uint64(123), 1, 123},
		{float32(123.45), 1, 123},
		{float64(123.45), 1, 123},
		{"123", 1, 123},
		{"abc", 1, 1},
		{func() {}, 1, 1},
	}

	Convey("TestUintDefault", t, func() {
		for i, c := range cases {
			Convey(fmt.Sprintf("case %v", i), func() {
				So(UintDefault(c.value, c.defaultValue), ShouldEqual, c.result)
			})
		}
	})
}

func TestUint(t *testing.T) {
	cases := []struct {
		value  interface{}
		result uint64
	}{
		{nil, 0},
		{true, 1},
		{false, 0},
		{int(123), 123},
		{int8(123), 123},
		{int16(123), 123},
		{int32(123), 123},
		{int64(123), 123},
		{uint(123), 123},
		{uint8(123), 123},
		{uint16(123), 123},
		{uint32(123), 123},
		{uint64(123), 123},
		{float32(123.45), 123},
		{float64(123.45), 123},
		{"123", 123},
		{"abc", 0},
		{func() {}, 0},
	}

	Convey("TestUint", t, func() {
		for i, c := range cases {
			Convey(fmt.Sprintf("case %v", i), func() {
				So(Uint(c.value), ShouldEqual, c.result)
			})
		}
	})
}

func TestFloatValue(t *testing.T) {
	cases := []struct {
		value  interface{}
		result float64
		hasErr bool
	}{
		{nil, 0, true},
		{true, 1, false},
		{false, 0, false},
		{int(-123), -123, false},
		{int8(-123), -123, false},
		{int16(-123), -123, false},
		{int32(-123), -123, false},
		{int64(-123), -123, false},
		{uint(123), 123, false},
		{uint8(123), 123, false},
		{uint16(123), 123, false},
		{uint32(123), 123, false},
		{uint64(123), 123, false},
		{float32(123.45), float64(float32(123.45)), false},
		{float64(-123.45), -123.45, false},
		{"123", 123, false},
		{"-123.45", -123.45, false},
		{"abc", 0, true},
		{func() {}, 0, true},
	}

	Convey("TestFloatValue", t, func() {
		for i, c := range cases {
			Convey(fmt.Sprintf("case %v", i), func() {
				res, err := FloatValue(c.value)
				if c.hasErr {
					So(err, ShouldBeError)
				} else {
					So(err, ShouldBeNil)
				}
				So(res, ShouldEqual, c.result)
			})
		}
	})
}

func TestFloatDefault(t *testing.T) {
	cases := []struct {
		value        interface{}
		defaultValue float64
		result       float64
	}{
		{nil, 1, 1},
		{true, 0, 1},
		{false, 1, 0},
		{int(-123), 1, -123},
		{int8(-123), 1, -123},
		{int16(-123), 1, -123},
		{int32(-123), 1, -123},
		{int64(-123), 1, -123},
		{uint(123), 1, 123},
		{uint8(123), 1, 123},
		{uint16(123), 1, 123},
		{uint32(123), 1, 123},
		{uint64(123), 1, 123},
		{float32(123.45), 1, float64(float32(123.45))},
		{float64(-123.45), 1, -123.45},
		{"123", 1, 123},
		{"-123.45", 1, -123.45},
		{"abc", 1, 1},
		{func() {}, 1, 1},
	}

	Convey("TestFloatDefault", t, func() {
		for i, c := range cases {
			Convey(fmt.Sprintf("case %v", i), func() {
				So(FloatDefault(c.value, c.defaultValue), ShouldEqual, c.result)
			})
		}
	})
}

func TestFloat(t *testing.T) {
	cases := []struct {
		value  interface{}
		result float64
	}{
		{nil, 0},
		{true, 1},
		{false, 0},
		{int(-123), -123},
		{int8(-123), -123},
		{int16(-123), -123},
		{int32(-123), -123},
		{int64(-123), -123},
		{uint(123), 123},
		{uint8(123), 123},
		{uint16(123), 123},
		{uint32(123), 123},
		{uint64(123), 123},
		{float32(123.45), float64(float32(123.45))},
		{float64(-123.45), -123.45},
		{"123", 123},
		{"-123.45", -123.45},
		{"abc", 0},
		{func() {}, 0},
	}

	Convey("TestFloat", t, func() {
		for i, c := range cases {
			Convey(fmt.Sprintf("case %v", i), func() {
				So(Float(c.value), ShouldEqual, c.result)
			})
		}
	})
}

func TestStringValue(t *testing.T) {
	cases := []struct {
		value  interface{}
		result string
		hasErr bool
	}{
		{nil, "", true},
		{"", "", false},
		{"abc", "abc", false},
		{0, "", true},
	}

	Convey("TestStringValue", t, func() {
		for i, c := range cases {
			Convey(fmt.Sprintf("case %v", i), func() {
				res, err := StringValue(c.value)
				if c.hasErr {
					So(err, ShouldBeError)
				} else {
					So(err, ShouldBeNil)
				}
				So(res, ShouldEqual, c.result)
			})
		}
	})
}

func TestStringDefault(t *testing.T) {
	cases := []struct {
		value        interface{}
		defaultValue string
		result       string
	}{
		{nil, "abc", "abc"},
		{"", "abc", ""},
		{"abc", "def", "abc"},
		{0, "abc", "abc"},
	}

	Convey("TestStringDefault", t, func() {
		for i, c := range cases {
			Convey(fmt.Sprintf("case %v", i), func() {
				So(StringDefault(c.value, c.defaultValue), ShouldEqual, c.result)
			})
		}
	})
}

func TestString(t *testing.T) {
	cases := []struct {
		value  interface{}
		result string
	}{
		{nil, ""},
		{"", ""},
		{"abc", "abc"},
		{0, ""},
	}

	Convey("TestString", t, func() {
		for i, c := range cases {
			Convey(fmt.Sprintf("case %v", i), func() {
				So(String(c.value), ShouldEqual, c.result)
			})
		}
	})
}

func TestStringForce(t *testing.T) {
	cases := []struct {
		value  interface{}
		result string
	}{
		{nil, ""},
		{"", ""},
		{"abc", "abc"},
		{-123.45, "-123.45"},
		{[]float64{123.45, -67.89}, fmt.Sprintf("%v", []float64{123.45, -67.89})},
		{map[int]string{1: "1", 2: "2"}, fmt.Sprintf("%v", map[int]string{1: "1", 2: "2"})},
	}

	Convey("TestStringForce", t, func() {
		for i, c := range cases {
			Convey(fmt.Sprintf("case %v", i), func() {
				So(StringForce(c.value), ShouldEqual, c.result)
			})
		}
	})
}

func TestUnsafeBytesToString(t *testing.T) {
	Convey("TestUnsafeBytesToString", t, func() {
		So(UnsafeBytesToString([]byte("abc")), ShouldEqual, "abc")
	})
}

func TestUnsafeStringToBytes(t *testing.T) {
	Convey("TestUnsafeStringToBytes", t, func() {
		So(UnsafeStringToBytes("abc"), ShouldResemble, []byte("abc"))
	})
}
