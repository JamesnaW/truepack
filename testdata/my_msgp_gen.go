package testdata

// NOTE: THIS FILE WAS PRODUCED BY THE
// TRUEPACK CODE GENERATION TOOL (github.com/glycerine/truepack)
// DO NOT EDIT

import (
	"github.com/glycerine/truepack/msgp"
)

// MSGPfieldsNotEmpty supports omitempty tags
func (z *A) MSGPfieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 6
	}
	var fieldsInUse uint32 = 6
	isempty[0] = (len(z.Name) == 0) // string, omitempty
	if isempty[0] {
		fieldsInUse--
	}
	isempty[1] = (z.Bday.IsZero()) // time.Time, omitempty
	if isempty[1] {
		fieldsInUse--
	}
	isempty[2] = (len(z.Phone) == 0) // string, omitempty
	if isempty[2] {
		fieldsInUse--
	}
	isempty[3] = (z.Sibs == 0) // number, omitempty
	if isempty[3] {
		fieldsInUse--
	}
	isempty[4] = (z.GPA == 0) // number, omitempty
	if isempty[4] {
		fieldsInUse--
	}
	isempty[5] = (!z.Friend) // bool, omitempty
	if isempty[5] {
		fieldsInUse--
	}

	return fieldsInUse
}

// MSGPMarshalMsg implements msgp.Marshaler
func (z *A) MSGPMarshalMsg(b []byte) (o []byte, err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.MSGPMsgsize())

	// honor the omitempty tags
	var empty [6]bool
	fieldsInUse := z.fieldsNotEmpty(empty[:])
	o = msgp.AppendMapHeader(o, fieldsInUse)

	if !empty[0] {
		// string "name_zid00_str"
		o = append(o, 0xae, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x73, 0x74, 0x72)
		o = msgp.AppendString(o, z.Name)
	}

	if !empty[1] {
		// string "Bday_zid01_tim"
		o = append(o, 0xae, 0x42, 0x64, 0x61, 0x79, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x31, 0x5f, 0x74, 0x69, 0x6d)
		o = msgp.AppendTime(o, z.Bday)
	}

	if !empty[2] {
		// string "phone_zid02_str"
		o = append(o, 0xaf, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x32, 0x5f, 0x73, 0x74, 0x72)
		o = msgp.AppendString(o, z.Phone)
	}

	if !empty[3] {
		// string "Sibs_zid03_int"
		o = append(o, 0xae, 0x53, 0x69, 0x62, 0x73, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x33, 0x5f, 0x69, 0x6e, 0x74)
		o = msgp.AppendInt(o, z.Sibs)
	}

	if !empty[4] {
		// string "GPA_zid04_f64"
		o = append(o, 0xad, 0x47, 0x50, 0x41, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x34, 0x5f, 0x66, 0x36, 0x34)
		o = msgp.AppendFloat64(o, z.GPA)
	}

	if !empty[5] {
		// string "Friend_zid05_boo"
		o = append(o, 0xb0, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x35, 0x5f, 0x62, 0x6f, 0x6f)
		o = msgp.AppendBool(o, z.Friend)
	}

	return
}

// MSGPUnmarshalMsg implements msgp.Unmarshaler
func (z *A) MSGPUnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.MSGPUnmarshalMsgWithCfg(bts, nil)
}
func (z *A) MSGPUnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields0zgensym_58b5c2649cecee38_1 = 6

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields0zgensym_58b5c2649cecee38_1 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields0zgensym_58b5c2649cecee38_1, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft0zgensym_58b5c2649cecee38_1 := totalEncodedFields0zgensym_58b5c2649cecee38_1
	missingFieldsLeft0zgensym_58b5c2649cecee38_1 := maxFields0zgensym_58b5c2649cecee38_1 - totalEncodedFields0zgensym_58b5c2649cecee38_1

	var nextMiss0zgensym_58b5c2649cecee38_1 int32 = -1
	var found0zgensym_58b5c2649cecee38_1 [maxFields0zgensym_58b5c2649cecee38_1]bool
	var curField0zgensym_58b5c2649cecee38_1 string

doneWithStruct0zgensym_58b5c2649cecee38_1:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zgensym_58b5c2649cecee38_1 > 0 || missingFieldsLeft0zgensym_58b5c2649cecee38_1 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zgensym_58b5c2649cecee38_1, missingFieldsLeft0zgensym_58b5c2649cecee38_1, msgp.ShowFound(found0zgensym_58b5c2649cecee38_1[:]), unmarshalMsgFieldOrder0zgensym_58b5c2649cecee38_1)
		if encodedFieldsLeft0zgensym_58b5c2649cecee38_1 > 0 {
			encodedFieldsLeft0zgensym_58b5c2649cecee38_1--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField0zgensym_58b5c2649cecee38_1 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zgensym_58b5c2649cecee38_1 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss0zgensym_58b5c2649cecee38_1 = 0
			}
			for nextMiss0zgensym_58b5c2649cecee38_1 < maxFields0zgensym_58b5c2649cecee38_1 && (found0zgensym_58b5c2649cecee38_1[nextMiss0zgensym_58b5c2649cecee38_1] || unmarshalMsgFieldSkip0zgensym_58b5c2649cecee38_1[nextMiss0zgensym_58b5c2649cecee38_1]) {
				nextMiss0zgensym_58b5c2649cecee38_1++
			}
			if nextMiss0zgensym_58b5c2649cecee38_1 == maxFields0zgensym_58b5c2649cecee38_1 {
				// filled all the empty fields!
				break doneWithStruct0zgensym_58b5c2649cecee38_1
			}
			missingFieldsLeft0zgensym_58b5c2649cecee38_1--
			curField0zgensym_58b5c2649cecee38_1 = unmarshalMsgFieldOrder0zgensym_58b5c2649cecee38_1[nextMiss0zgensym_58b5c2649cecee38_1]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zgensym_58b5c2649cecee38_1)
		switch curField0zgensym_58b5c2649cecee38_1 {
		// -- templateUnmarshalMsg ends here --

		case "name_zid00_str":
			found0zgensym_58b5c2649cecee38_1[0] = true
			z.Name, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Bday_zid01_tim":
			found0zgensym_58b5c2649cecee38_1[1] = true
			z.Bday, bts, err = nbs.ReadTimeBytes(bts)

			if err != nil {
				return
			}
		case "phone_zid02_str":
			found0zgensym_58b5c2649cecee38_1[2] = true
			z.Phone, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Sibs_zid03_int":
			found0zgensym_58b5c2649cecee38_1[3] = true
			z.Sibs, bts, err = nbs.ReadIntBytes(bts)

			if err != nil {
				return
			}
		case "GPA_zid04_f64":
			found0zgensym_58b5c2649cecee38_1[4] = true
			z.GPA, bts, err = nbs.ReadFloat64Bytes(bts)

			if err != nil {
				return
			}
		case "Friend_zid05_boo":
			found0zgensym_58b5c2649cecee38_1[5] = true
			z.Friend, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	if nextMiss0zgensym_58b5c2649cecee38_1 != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	if p, ok := interface{}(z).(msgp.PostLoad); ok {
		p.PostLoadHook()
	}

	return
}

// fields of A
var unmarshalMsgFieldOrder0zgensym_58b5c2649cecee38_1 = []string{"name_zid00_str", "Bday_zid01_tim", "phone_zid02_str", "Sibs_zid03_int", "GPA_zid04_f64", "Friend_zid05_boo"}

var unmarshalMsgFieldSkip0zgensym_58b5c2649cecee38_1 = []bool{false, false, false, false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *A) MSGPMsgsize() (s int) {
	s = 1 + 15 + msgp.StringPrefixSize + len(z.Name) + 15 + msgp.TimeSize + 16 + msgp.StringPrefixSize + len(z.Phone) + 15 + msgp.IntSize + 14 + msgp.Float64Size + 17 + msgp.BoolSize
	return
}

// MSGPfieldsNotEmpty supports omitempty tags
func (z *Big) MSGPfieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 5
	}
	var fieldsInUse uint32 = 5
	isempty[0] = (len(z.Slice) == 0) // string, omitempty
	if isempty[0] {
		fieldsInUse--
	}
	isempty[1] = (len(z.Transform) == 0) // string, omitempty
	if isempty[1] {
		fieldsInUse--
	}
	isempty[2] = (z.Myptr == nil) // pointer, omitempty
	if isempty[2] {
		fieldsInUse--
	}
	isempty[3] = (len(z.Myarray) == 0) // string, omitempty
	if isempty[3] {
		fieldsInUse--
	}
	isempty[4] = (len(z.MySlice) == 0) // string, omitempty
	if isempty[4] {
		fieldsInUse--
	}

	return fieldsInUse
}

// MSGPMarshalMsg implements msgp.Marshaler
func (z *Big) MSGPMarshalMsg(b []byte) (o []byte, err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.MSGPMsgsize())

	// honor the omitempty tags
	var empty [5]bool
	fieldsInUse := z.fieldsNotEmpty(empty[:])
	o = msgp.AppendMapHeader(o, fieldsInUse)

	if !empty[0] {
		// string "Slice_zid00_slc"
		o = append(o, 0xaf, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x73, 0x6c, 0x63)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Slice)))
		for zgensym_58b5c2649cecee38_2 := range z.Slice {
			o, err = z.Slice[zgensym_58b5c2649cecee38_2].MSGPMarshalMsg(o)
			if err != nil {
				return
			}
		}
	}

	if !empty[1] {
		// string "Transform_zid01_map"
		o = append(o, 0xb3, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x31, 0x5f, 0x6d, 0x61, 0x70)
		o = msgp.AppendMapHeader(o, uint32(len(z.Transform)))
		for zgensym_58b5c2649cecee38_3, zgensym_58b5c2649cecee38_4 := range z.Transform {
			o = msgp.AppendInt(o, zgensym_58b5c2649cecee38_3)
			if zgensym_58b5c2649cecee38_4 == nil {
				o = msgp.AppendNil(o)
			} else {
				o, err = zgensym_58b5c2649cecee38_4.MSGPMarshalMsg(o)
				if err != nil {
					return
				}
			}
		}
	}

	if !empty[2] {
		// string "Myptr_zid02_ptr"
		o = append(o, 0xaf, 0x4d, 0x79, 0x70, 0x74, 0x72, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x32, 0x5f, 0x70, 0x74, 0x72)
		if z.Myptr == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Myptr.MSGPMarshalMsg(o)
			if err != nil {
				return
			}
		}
	}

	if !empty[3] {
		// string "Myarray_zid03_ary"
		o = append(o, 0xb1, 0x4d, 0x79, 0x61, 0x72, 0x72, 0x61, 0x79, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x33, 0x5f, 0x61, 0x72, 0x79)
		o = msgp.AppendArrayHeader(o, 3)
		for zgensym_58b5c2649cecee38_5 := range z.Myarray {
			o = msgp.AppendString(o, z.Myarray[zgensym_58b5c2649cecee38_5])
		}
	}

	if !empty[4] {
		// string "MySlice_zid04_slc"
		o = append(o, 0xb1, 0x4d, 0x79, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x34, 0x5f, 0x73, 0x6c, 0x63)
		o = msgp.AppendArrayHeader(o, uint32(len(z.MySlice)))
		for zgensym_58b5c2649cecee38_6 := range z.MySlice {
			o = msgp.AppendString(o, z.MySlice[zgensym_58b5c2649cecee38_6])
		}
	}

	return
}

// MSGPUnmarshalMsg implements msgp.Unmarshaler
func (z *Big) MSGPUnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.MSGPUnmarshalMsgWithCfg(bts, nil)
}
func (z *Big) MSGPUnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields7zgensym_58b5c2649cecee38_8 = 5

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields7zgensym_58b5c2649cecee38_8 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields7zgensym_58b5c2649cecee38_8, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft7zgensym_58b5c2649cecee38_8 := totalEncodedFields7zgensym_58b5c2649cecee38_8
	missingFieldsLeft7zgensym_58b5c2649cecee38_8 := maxFields7zgensym_58b5c2649cecee38_8 - totalEncodedFields7zgensym_58b5c2649cecee38_8

	var nextMiss7zgensym_58b5c2649cecee38_8 int32 = -1
	var found7zgensym_58b5c2649cecee38_8 [maxFields7zgensym_58b5c2649cecee38_8]bool
	var curField7zgensym_58b5c2649cecee38_8 string

doneWithStruct7zgensym_58b5c2649cecee38_8:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft7zgensym_58b5c2649cecee38_8 > 0 || missingFieldsLeft7zgensym_58b5c2649cecee38_8 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft7zgensym_58b5c2649cecee38_8, missingFieldsLeft7zgensym_58b5c2649cecee38_8, msgp.ShowFound(found7zgensym_58b5c2649cecee38_8[:]), unmarshalMsgFieldOrder7zgensym_58b5c2649cecee38_8)
		if encodedFieldsLeft7zgensym_58b5c2649cecee38_8 > 0 {
			encodedFieldsLeft7zgensym_58b5c2649cecee38_8--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField7zgensym_58b5c2649cecee38_8 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss7zgensym_58b5c2649cecee38_8 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss7zgensym_58b5c2649cecee38_8 = 0
			}
			for nextMiss7zgensym_58b5c2649cecee38_8 < maxFields7zgensym_58b5c2649cecee38_8 && (found7zgensym_58b5c2649cecee38_8[nextMiss7zgensym_58b5c2649cecee38_8] || unmarshalMsgFieldSkip7zgensym_58b5c2649cecee38_8[nextMiss7zgensym_58b5c2649cecee38_8]) {
				nextMiss7zgensym_58b5c2649cecee38_8++
			}
			if nextMiss7zgensym_58b5c2649cecee38_8 == maxFields7zgensym_58b5c2649cecee38_8 {
				// filled all the empty fields!
				break doneWithStruct7zgensym_58b5c2649cecee38_8
			}
			missingFieldsLeft7zgensym_58b5c2649cecee38_8--
			curField7zgensym_58b5c2649cecee38_8 = unmarshalMsgFieldOrder7zgensym_58b5c2649cecee38_8[nextMiss7zgensym_58b5c2649cecee38_8]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField7zgensym_58b5c2649cecee38_8)
		switch curField7zgensym_58b5c2649cecee38_8 {
		// -- templateUnmarshalMsg ends here --

		case "Slice_zid00_slc":
			found7zgensym_58b5c2649cecee38_8[0] = true
			if nbs.AlwaysNil {
				(z.Slice) = (z.Slice)[:0]
			} else {

				var zgensym_58b5c2649cecee38_9 uint32
				zgensym_58b5c2649cecee38_9, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Slice) >= int(zgensym_58b5c2649cecee38_9) {
					z.Slice = (z.Slice)[:zgensym_58b5c2649cecee38_9]
				} else {
					z.Slice = make([]S2, zgensym_58b5c2649cecee38_9)
				}
				for zgensym_58b5c2649cecee38_2 := range z.Slice {
					bts, err = z.Slice[zgensym_58b5c2649cecee38_2].MSGPUnmarshalMsg(bts)
					if err != nil {
						return
					}
					if err != nil {
						return
					}
				}
			}
		case "Transform_zid01_map":
			found7zgensym_58b5c2649cecee38_8[1] = true
			if nbs.AlwaysNil {
				if len(z.Transform) > 0 {
					for key, _ := range z.Transform {
						delete(z.Transform, key)
					}
				}

			} else {

				var zgensym_58b5c2649cecee38_10 uint32
				zgensym_58b5c2649cecee38_10, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Transform == nil && zgensym_58b5c2649cecee38_10 > 0 {
					z.Transform = make(map[int]*S2, zgensym_58b5c2649cecee38_10)
				} else if len(z.Transform) > 0 {
					for key, _ := range z.Transform {
						delete(z.Transform, key)
					}
				}
				for zgensym_58b5c2649cecee38_10 > 0 {
					var zgensym_58b5c2649cecee38_3 int
					var zgensym_58b5c2649cecee38_4 *S2
					zgensym_58b5c2649cecee38_10--
					zgensym_58b5c2649cecee38_3, bts, err = nbs.ReadIntBytes(bts)
					if err != nil {
						return
					}
					if nbs.AlwaysNil {
						if zgensym_58b5c2649cecee38_4 != nil {
							zgensym_58b5c2649cecee38_4.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != zgensym_58b5c2649cecee38_4 {
								zgensym_58b5c2649cecee38_4.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if zgensym_58b5c2649cecee38_4 == nil {
								zgensym_58b5c2649cecee38_4 = new(S2)
							}
							bts, err = zgensym_58b5c2649cecee38_4.MSGPUnmarshalMsg(bts)
							if err != nil {
								return
							}
							if err != nil {
								return
							}
						}
					}
					z.Transform[zgensym_58b5c2649cecee38_3] = zgensym_58b5c2649cecee38_4
				}
			}
		case "Myptr_zid02_ptr":
			found7zgensym_58b5c2649cecee38_8[2] = true
			if nbs.AlwaysNil {
				if z.Myptr != nil {
					z.Myptr.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
				}
			} else {
				// not nbs.AlwaysNil
				if msgp.IsNil(bts) {
					bts = bts[1:]
					if nil != z.Myptr {
						z.Myptr.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
					}
				} else {
					// not nbs.AlwaysNil and not IsNil(bts): have something to read

					if z.Myptr == nil {
						z.Myptr = new(S2)
					}
					bts, err = z.Myptr.MSGPUnmarshalMsg(bts)
					if err != nil {
						return
					}
					if err != nil {
						return
					}
				}
			}
		case "Myarray_zid03_ary":
			found7zgensym_58b5c2649cecee38_8[3] = true
			var zgensym_58b5c2649cecee38_11 uint32
			zgensym_58b5c2649cecee38_11, bts, err = nbs.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if !nbs.IsNil(bts) && zgensym_58b5c2649cecee38_11 != 3 {
				err = msgp.ArrayError{Wanted: 3, Got: zgensym_58b5c2649cecee38_11}
				return
			}
			for zgensym_58b5c2649cecee38_5 := range z.Myarray {
				z.Myarray[zgensym_58b5c2649cecee38_5], bts, err = nbs.ReadStringBytes(bts)

				if err != nil {
					return
				}
			}
		case "MySlice_zid04_slc":
			found7zgensym_58b5c2649cecee38_8[4] = true
			if nbs.AlwaysNil {
				(z.MySlice) = (z.MySlice)[:0]
			} else {

				var zgensym_58b5c2649cecee38_12 uint32
				zgensym_58b5c2649cecee38_12, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.MySlice) >= int(zgensym_58b5c2649cecee38_12) {
					z.MySlice = (z.MySlice)[:zgensym_58b5c2649cecee38_12]
				} else {
					z.MySlice = make([]string, zgensym_58b5c2649cecee38_12)
				}
				for zgensym_58b5c2649cecee38_6 := range z.MySlice {
					z.MySlice[zgensym_58b5c2649cecee38_6], bts, err = nbs.ReadStringBytes(bts)

					if err != nil {
						return
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	if nextMiss7zgensym_58b5c2649cecee38_8 != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	if p, ok := interface{}(z).(msgp.PostLoad); ok {
		p.PostLoadHook()
	}

	return
}

// fields of Big
var unmarshalMsgFieldOrder7zgensym_58b5c2649cecee38_8 = []string{"Slice_zid00_slc", "Transform_zid01_map", "Myptr_zid02_ptr", "Myarray_zid03_ary", "MySlice_zid04_slc"}

var unmarshalMsgFieldSkip7zgensym_58b5c2649cecee38_8 = []bool{false, false, false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Big) MSGPMsgsize() (s int) {
	s = 1 + 16 + msgp.ArrayHeaderSize
	for zgensym_58b5c2649cecee38_2 := range z.Slice {
		s += z.Slice[zgensym_58b5c2649cecee38_2].MSGPMsgsize()
	}
	s += 20 + msgp.MapHeaderSize
	if z.Transform != nil {
		for zgensym_58b5c2649cecee38_3, zgensym_58b5c2649cecee38_4 := range z.Transform {
			_ = zgensym_58b5c2649cecee38_4
			_ = zgensym_58b5c2649cecee38_3
			s += msgp.IntSize
			if zgensym_58b5c2649cecee38_4 == nil {
				s += msgp.NilSize
			} else {
				s += zgensym_58b5c2649cecee38_4.MSGPMsgsize()
			}
		}
	}
	s += 16
	if z.Myptr == nil {
		s += msgp.NilSize
	} else {
		s += z.Myptr.MSGPMsgsize()
	}
	s += 18 + msgp.ArrayHeaderSize
	for zgensym_58b5c2649cecee38_5 := range z.Myarray {
		s += msgp.StringPrefixSize + len(z.Myarray[zgensym_58b5c2649cecee38_5])
	}
	s += 18 + msgp.ArrayHeaderSize
	for zgensym_58b5c2649cecee38_6 := range z.MySlice {
		s += msgp.StringPrefixSize + len(z.MySlice[zgensym_58b5c2649cecee38_6])
	}
	return
}

// MSGPfieldsNotEmpty supports omitempty tags
func (z *S2) MSGPfieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 7
	}
	var fieldsInUse uint32 = 7
	isempty[1] = (len(z.B) == 0) // string, omitempty
	if isempty[1] {
		fieldsInUse--
	}
	isempty[2] = (len(z.R) == 0) // string, omitempty
	if isempty[2] {
		fieldsInUse--
	}
	isempty[3] = (z.P == 0) // number, omitempty
	if isempty[3] {
		fieldsInUse--
	}
	isempty[4] = (z.Q == 0) // number, omitempty
	if isempty[4] {
		fieldsInUse--
	}
	isempty[5] = (z.T == 0) // number, omitempty
	if isempty[5] {
		fieldsInUse--
	}
	isempty[6] = (len(z.Arr) == 0) // string, omitempty
	if isempty[6] {
		fieldsInUse--
	}
	isempty[7] = (z.MyTree == nil) // pointer, omitempty
	if isempty[7] {
		fieldsInUse--
	}

	return fieldsInUse
}

// MSGPMarshalMsg implements msgp.Marshaler
func (z *S2) MSGPMarshalMsg(b []byte) (o []byte, err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.MSGPMsgsize())

	// honor the omitempty tags
	var empty [8]bool
	fieldsInUse := z.fieldsNotEmpty(empty[:])
	o = msgp.AppendMapHeader(o, fieldsInUse)

	if !empty[1] {
		// string "beta_zid01_str"
		o = append(o, 0xae, 0x62, 0x65, 0x74, 0x61, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x31, 0x5f, 0x73, 0x74, 0x72)
		o = msgp.AppendString(o, z.B)
	}

	if !empty[2] {
		// string "ralph_zid02_map"
		o = append(o, 0xaf, 0x72, 0x61, 0x6c, 0x70, 0x68, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x32, 0x5f, 0x6d, 0x61, 0x70)
		o = msgp.AppendMapHeader(o, uint32(len(z.R)))
		for zgensym_58b5c2649cecee38_13, zgensym_58b5c2649cecee38_14 := range z.R {
			o = msgp.AppendString(o, zgensym_58b5c2649cecee38_13)
			o = msgp.AppendUint8(o, zgensym_58b5c2649cecee38_14)
		}
	}

	if !empty[3] {
		// string "P_zid03_u16"
		o = append(o, 0xab, 0x50, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x33, 0x5f, 0x75, 0x31, 0x36)
		o = msgp.AppendUint16(o, z.P)
	}

	if !empty[4] {
		// string "Q_zid04_u32"
		o = append(o, 0xab, 0x51, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x34, 0x5f, 0x75, 0x33, 0x32)
		o = msgp.AppendUint32(o, z.Q)
	}

	if !empty[5] {
		// string "T_zid05_i64"
		o = append(o, 0xab, 0x54, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x35, 0x5f, 0x69, 0x36, 0x34)
		o = msgp.AppendInt64(o, z.T)
	}

	if !empty[6] {
		// string "arr_zid06_ary"
		o = append(o, 0xad, 0x61, 0x72, 0x72, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x36, 0x5f, 0x61, 0x72, 0x79)
		o = msgp.AppendArrayHeader(o, 6)
		for zgensym_58b5c2649cecee38_15 := range z.Arr {
			o = msgp.AppendFloat64(o, z.Arr[zgensym_58b5c2649cecee38_15])
		}
	}

	if !empty[7] {
		// string "MyTree_zid07_ptr"
		o = append(o, 0xb0, 0x4d, 0x79, 0x54, 0x72, 0x65, 0x65, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x37, 0x5f, 0x70, 0x74, 0x72)
		if z.MyTree == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.MyTree.MSGPMarshalMsg(o)
			if err != nil {
				return
			}
		}
	}

	return
}

// MSGPUnmarshalMsg implements msgp.Unmarshaler
func (z *S2) MSGPUnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.MSGPUnmarshalMsgWithCfg(bts, nil)
}
func (z *S2) MSGPUnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields16zgensym_58b5c2649cecee38_17 = 8

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields16zgensym_58b5c2649cecee38_17 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields16zgensym_58b5c2649cecee38_17, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft16zgensym_58b5c2649cecee38_17 := totalEncodedFields16zgensym_58b5c2649cecee38_17
	missingFieldsLeft16zgensym_58b5c2649cecee38_17 := maxFields16zgensym_58b5c2649cecee38_17 - totalEncodedFields16zgensym_58b5c2649cecee38_17

	var nextMiss16zgensym_58b5c2649cecee38_17 int32 = -1
	var found16zgensym_58b5c2649cecee38_17 [maxFields16zgensym_58b5c2649cecee38_17]bool
	var curField16zgensym_58b5c2649cecee38_17 string

doneWithStruct16zgensym_58b5c2649cecee38_17:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft16zgensym_58b5c2649cecee38_17 > 0 || missingFieldsLeft16zgensym_58b5c2649cecee38_17 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft16zgensym_58b5c2649cecee38_17, missingFieldsLeft16zgensym_58b5c2649cecee38_17, msgp.ShowFound(found16zgensym_58b5c2649cecee38_17[:]), unmarshalMsgFieldOrder16zgensym_58b5c2649cecee38_17)
		if encodedFieldsLeft16zgensym_58b5c2649cecee38_17 > 0 {
			encodedFieldsLeft16zgensym_58b5c2649cecee38_17--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField16zgensym_58b5c2649cecee38_17 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss16zgensym_58b5c2649cecee38_17 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss16zgensym_58b5c2649cecee38_17 = 0
			}
			for nextMiss16zgensym_58b5c2649cecee38_17 < maxFields16zgensym_58b5c2649cecee38_17 && (found16zgensym_58b5c2649cecee38_17[nextMiss16zgensym_58b5c2649cecee38_17] || unmarshalMsgFieldSkip16zgensym_58b5c2649cecee38_17[nextMiss16zgensym_58b5c2649cecee38_17]) {
				nextMiss16zgensym_58b5c2649cecee38_17++
			}
			if nextMiss16zgensym_58b5c2649cecee38_17 == maxFields16zgensym_58b5c2649cecee38_17 {
				// filled all the empty fields!
				break doneWithStruct16zgensym_58b5c2649cecee38_17
			}
			missingFieldsLeft16zgensym_58b5c2649cecee38_17--
			curField16zgensym_58b5c2649cecee38_17 = unmarshalMsgFieldOrder16zgensym_58b5c2649cecee38_17[nextMiss16zgensym_58b5c2649cecee38_17]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField16zgensym_58b5c2649cecee38_17)
		switch curField16zgensym_58b5c2649cecee38_17 {
		// -- templateUnmarshalMsg ends here --

		case "beta_zid01_str":
			found16zgensym_58b5c2649cecee38_17[1] = true
			z.B, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "ralph_zid02_map":
			found16zgensym_58b5c2649cecee38_17[2] = true
			if nbs.AlwaysNil {
				if len(z.R) > 0 {
					for key, _ := range z.R {
						delete(z.R, key)
					}
				}

			} else {

				var zgensym_58b5c2649cecee38_18 uint32
				zgensym_58b5c2649cecee38_18, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.R == nil && zgensym_58b5c2649cecee38_18 > 0 {
					z.R = make(map[string]uint8, zgensym_58b5c2649cecee38_18)
				} else if len(z.R) > 0 {
					for key, _ := range z.R {
						delete(z.R, key)
					}
				}
				for zgensym_58b5c2649cecee38_18 > 0 {
					var zgensym_58b5c2649cecee38_13 string
					var zgensym_58b5c2649cecee38_14 uint8
					zgensym_58b5c2649cecee38_18--
					zgensym_58b5c2649cecee38_13, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zgensym_58b5c2649cecee38_14, bts, err = nbs.ReadUint8Bytes(bts)

					if err != nil {
						return
					}
					z.R[zgensym_58b5c2649cecee38_13] = zgensym_58b5c2649cecee38_14
				}
			}
		case "P_zid03_u16":
			found16zgensym_58b5c2649cecee38_17[3] = true
			z.P, bts, err = nbs.ReadUint16Bytes(bts)

			if err != nil {
				return
			}
		case "Q_zid04_u32":
			found16zgensym_58b5c2649cecee38_17[4] = true
			z.Q, bts, err = nbs.ReadUint32Bytes(bts)

			if err != nil {
				return
			}
		case "T_zid05_i64":
			found16zgensym_58b5c2649cecee38_17[5] = true
			z.T, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				return
			}
		case "arr_zid06_ary":
			found16zgensym_58b5c2649cecee38_17[6] = true
			var zgensym_58b5c2649cecee38_19 uint32
			zgensym_58b5c2649cecee38_19, bts, err = nbs.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if !nbs.IsNil(bts) && zgensym_58b5c2649cecee38_19 != 6 {
				err = msgp.ArrayError{Wanted: 6, Got: zgensym_58b5c2649cecee38_19}
				return
			}
			for zgensym_58b5c2649cecee38_15 := range z.Arr {
				z.Arr[zgensym_58b5c2649cecee38_15], bts, err = nbs.ReadFloat64Bytes(bts)

				if err != nil {
					return
				}
			}
		case "MyTree_zid07_ptr":
			found16zgensym_58b5c2649cecee38_17[7] = true
			if nbs.AlwaysNil {
				if z.MyTree != nil {
					z.MyTree.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
				}
			} else {
				// not nbs.AlwaysNil
				if msgp.IsNil(bts) {
					bts = bts[1:]
					if nil != z.MyTree {
						z.MyTree.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
					}
				} else {
					// not nbs.AlwaysNil and not IsNil(bts): have something to read

					if z.MyTree == nil {
						z.MyTree = new(Tree)
					}
					bts, err = z.MyTree.MSGPUnmarshalMsg(bts)
					if err != nil {
						return
					}
					if err != nil {
						return
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	if nextMiss16zgensym_58b5c2649cecee38_17 != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	if p, ok := interface{}(z).(msgp.PostLoad); ok {
		p.PostLoadHook()
	}

	return
}

// fields of S2
var unmarshalMsgFieldOrder16zgensym_58b5c2649cecee38_17 = []string{"", "beta_zid01_str", "ralph_zid02_map", "P_zid03_u16", "Q_zid04_u32", "T_zid05_i64", "arr_zid06_ary", "MyTree_zid07_ptr"}

var unmarshalMsgFieldSkip16zgensym_58b5c2649cecee38_17 = []bool{true, false, false, false, false, false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *S2) MSGPMsgsize() (s int) {
	s = 1 + 15 + msgp.StringPrefixSize + len(z.B) + 16 + msgp.MapHeaderSize
	if z.R != nil {
		for zgensym_58b5c2649cecee38_13, zgensym_58b5c2649cecee38_14 := range z.R {
			_ = zgensym_58b5c2649cecee38_14
			_ = zgensym_58b5c2649cecee38_13
			s += msgp.StringPrefixSize + len(zgensym_58b5c2649cecee38_13) + msgp.Uint8Size
		}
	}
	s += 12 + msgp.Uint16Size + 12 + msgp.Uint32Size + 12 + msgp.Int64Size + 14 + msgp.ArrayHeaderSize + (6 * (msgp.Float64Size)) + 17
	if z.MyTree == nil {
		s += msgp.NilSize
	} else {
		s += z.MyTree.MSGPMsgsize()
	}
	return
}

// MSGPfieldsNotEmpty supports omitempty tags
func (z Sys) MSGPfieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 1
	}
	var fieldsInUse uint32 = 1
	isempty[0] = false
	if isempty[0] {
		fieldsInUse--
	}

	return fieldsInUse
}

// MSGPMarshalMsg implements msgp.Marshaler
func (z Sys) MSGPMarshalMsg(b []byte) (o []byte, err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.MSGPMsgsize())

	// honor the omitempty tags
	var empty [1]bool
	fieldsInUse := z.fieldsNotEmpty(empty[:])
	o = msgp.AppendMapHeader(o, fieldsInUse)

	if !empty[0] {
		// string "F_zid00_ifc"
		o = append(o, 0xab, 0x46, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x69, 0x66, 0x63)
		o, err = msgp.AppendIntf(o, z.F)
		if err != nil {
			return
		}
	}

	return
}

// MSGPUnmarshalMsg implements msgp.Unmarshaler
func (z *Sys) MSGPUnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.MSGPUnmarshalMsgWithCfg(bts, nil)
}
func (z *Sys) MSGPUnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields20zgensym_58b5c2649cecee38_21 = 1

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields20zgensym_58b5c2649cecee38_21 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields20zgensym_58b5c2649cecee38_21, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft20zgensym_58b5c2649cecee38_21 := totalEncodedFields20zgensym_58b5c2649cecee38_21
	missingFieldsLeft20zgensym_58b5c2649cecee38_21 := maxFields20zgensym_58b5c2649cecee38_21 - totalEncodedFields20zgensym_58b5c2649cecee38_21

	var nextMiss20zgensym_58b5c2649cecee38_21 int32 = -1
	var found20zgensym_58b5c2649cecee38_21 [maxFields20zgensym_58b5c2649cecee38_21]bool
	var curField20zgensym_58b5c2649cecee38_21 string

doneWithStruct20zgensym_58b5c2649cecee38_21:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft20zgensym_58b5c2649cecee38_21 > 0 || missingFieldsLeft20zgensym_58b5c2649cecee38_21 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft20zgensym_58b5c2649cecee38_21, missingFieldsLeft20zgensym_58b5c2649cecee38_21, msgp.ShowFound(found20zgensym_58b5c2649cecee38_21[:]), unmarshalMsgFieldOrder20zgensym_58b5c2649cecee38_21)
		if encodedFieldsLeft20zgensym_58b5c2649cecee38_21 > 0 {
			encodedFieldsLeft20zgensym_58b5c2649cecee38_21--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField20zgensym_58b5c2649cecee38_21 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss20zgensym_58b5c2649cecee38_21 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss20zgensym_58b5c2649cecee38_21 = 0
			}
			for nextMiss20zgensym_58b5c2649cecee38_21 < maxFields20zgensym_58b5c2649cecee38_21 && (found20zgensym_58b5c2649cecee38_21[nextMiss20zgensym_58b5c2649cecee38_21] || unmarshalMsgFieldSkip20zgensym_58b5c2649cecee38_21[nextMiss20zgensym_58b5c2649cecee38_21]) {
				nextMiss20zgensym_58b5c2649cecee38_21++
			}
			if nextMiss20zgensym_58b5c2649cecee38_21 == maxFields20zgensym_58b5c2649cecee38_21 {
				// filled all the empty fields!
				break doneWithStruct20zgensym_58b5c2649cecee38_21
			}
			missingFieldsLeft20zgensym_58b5c2649cecee38_21--
			curField20zgensym_58b5c2649cecee38_21 = unmarshalMsgFieldOrder20zgensym_58b5c2649cecee38_21[nextMiss20zgensym_58b5c2649cecee38_21]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField20zgensym_58b5c2649cecee38_21)
		switch curField20zgensym_58b5c2649cecee38_21 {
		// -- templateUnmarshalMsg ends here --

		case "F_zid00_ifc":
			found20zgensym_58b5c2649cecee38_21[0] = true
			z.F, bts, err = nbs.ReadIntfBytes(bts)

			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	if nextMiss20zgensym_58b5c2649cecee38_21 != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	if p, ok := interface{}(z).(msgp.PostLoad); ok {
		p.PostLoadHook()
	}

	return
}

// fields of Sys
var unmarshalMsgFieldOrder20zgensym_58b5c2649cecee38_21 = []string{"F_zid00_ifc"}

var unmarshalMsgFieldSkip20zgensym_58b5c2649cecee38_21 = []bool{false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Sys) MSGPMsgsize() (s int) {
	s = 1 + 12 + msgp.GuessSize(z.F)
	return
}

// MSGPfieldsNotEmpty supports omitempty tags
func (z *Tree) MSGPfieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 3
	}
	var fieldsInUse uint32 = 3
	isempty[0] = (len(z.Chld) == 0) // string, omitempty
	if isempty[0] {
		fieldsInUse--
	}
	isempty[1] = (len(z.Str) == 0) // string, omitempty
	if isempty[1] {
		fieldsInUse--
	}
	isempty[2] = (z.Par == nil) // pointer, omitempty
	if isempty[2] {
		fieldsInUse--
	}

	return fieldsInUse
}

// MSGPMarshalMsg implements msgp.Marshaler
func (z *Tree) MSGPMarshalMsg(b []byte) (o []byte, err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.MSGPMsgsize())

	// honor the omitempty tags
	var empty [3]bool
	fieldsInUse := z.fieldsNotEmpty(empty[:])
	o = msgp.AppendMapHeader(o, fieldsInUse)

	if !empty[0] {
		// string "Chld_zid00_slc"
		o = append(o, 0xae, 0x43, 0x68, 0x6c, 0x64, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x73, 0x6c, 0x63)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Chld)))
		for zgensym_58b5c2649cecee38_22 := range z.Chld {
			o, err = z.Chld[zgensym_58b5c2649cecee38_22].MSGPMarshalMsg(o)
			if err != nil {
				return
			}
		}
	}

	if !empty[1] {
		// string "Str_zid01_str"
		o = append(o, 0xad, 0x53, 0x74, 0x72, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x31, 0x5f, 0x73, 0x74, 0x72)
		o = msgp.AppendString(o, z.Str)
	}

	if !empty[2] {
		// string "Par_zid02_ptr"
		o = append(o, 0xad, 0x50, 0x61, 0x72, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x32, 0x5f, 0x70, 0x74, 0x72)
		if z.Par == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Par.MSGPMarshalMsg(o)
			if err != nil {
				return
			}
		}
	}

	return
}

// MSGPUnmarshalMsg implements msgp.Unmarshaler
func (z *Tree) MSGPUnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.MSGPUnmarshalMsgWithCfg(bts, nil)
}
func (z *Tree) MSGPUnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields23zgensym_58b5c2649cecee38_24 = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields23zgensym_58b5c2649cecee38_24 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields23zgensym_58b5c2649cecee38_24, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft23zgensym_58b5c2649cecee38_24 := totalEncodedFields23zgensym_58b5c2649cecee38_24
	missingFieldsLeft23zgensym_58b5c2649cecee38_24 := maxFields23zgensym_58b5c2649cecee38_24 - totalEncodedFields23zgensym_58b5c2649cecee38_24

	var nextMiss23zgensym_58b5c2649cecee38_24 int32 = -1
	var found23zgensym_58b5c2649cecee38_24 [maxFields23zgensym_58b5c2649cecee38_24]bool
	var curField23zgensym_58b5c2649cecee38_24 string

doneWithStruct23zgensym_58b5c2649cecee38_24:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft23zgensym_58b5c2649cecee38_24 > 0 || missingFieldsLeft23zgensym_58b5c2649cecee38_24 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft23zgensym_58b5c2649cecee38_24, missingFieldsLeft23zgensym_58b5c2649cecee38_24, msgp.ShowFound(found23zgensym_58b5c2649cecee38_24[:]), unmarshalMsgFieldOrder23zgensym_58b5c2649cecee38_24)
		if encodedFieldsLeft23zgensym_58b5c2649cecee38_24 > 0 {
			encodedFieldsLeft23zgensym_58b5c2649cecee38_24--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField23zgensym_58b5c2649cecee38_24 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss23zgensym_58b5c2649cecee38_24 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss23zgensym_58b5c2649cecee38_24 = 0
			}
			for nextMiss23zgensym_58b5c2649cecee38_24 < maxFields23zgensym_58b5c2649cecee38_24 && (found23zgensym_58b5c2649cecee38_24[nextMiss23zgensym_58b5c2649cecee38_24] || unmarshalMsgFieldSkip23zgensym_58b5c2649cecee38_24[nextMiss23zgensym_58b5c2649cecee38_24]) {
				nextMiss23zgensym_58b5c2649cecee38_24++
			}
			if nextMiss23zgensym_58b5c2649cecee38_24 == maxFields23zgensym_58b5c2649cecee38_24 {
				// filled all the empty fields!
				break doneWithStruct23zgensym_58b5c2649cecee38_24
			}
			missingFieldsLeft23zgensym_58b5c2649cecee38_24--
			curField23zgensym_58b5c2649cecee38_24 = unmarshalMsgFieldOrder23zgensym_58b5c2649cecee38_24[nextMiss23zgensym_58b5c2649cecee38_24]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField23zgensym_58b5c2649cecee38_24)
		switch curField23zgensym_58b5c2649cecee38_24 {
		// -- templateUnmarshalMsg ends here --

		case "Chld_zid00_slc":
			found23zgensym_58b5c2649cecee38_24[0] = true
			if nbs.AlwaysNil {
				(z.Chld) = (z.Chld)[:0]
			} else {

				var zgensym_58b5c2649cecee38_25 uint32
				zgensym_58b5c2649cecee38_25, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Chld) >= int(zgensym_58b5c2649cecee38_25) {
					z.Chld = (z.Chld)[:zgensym_58b5c2649cecee38_25]
				} else {
					z.Chld = make([]Tree, zgensym_58b5c2649cecee38_25)
				}
				for zgensym_58b5c2649cecee38_22 := range z.Chld {
					bts, err = z.Chld[zgensym_58b5c2649cecee38_22].MSGPUnmarshalMsg(bts)
					if err != nil {
						return
					}
					if err != nil {
						return
					}
				}
			}
		case "Str_zid01_str":
			found23zgensym_58b5c2649cecee38_24[1] = true
			z.Str, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Par_zid02_ptr":
			found23zgensym_58b5c2649cecee38_24[2] = true
			if nbs.AlwaysNil {
				if z.Par != nil {
					z.Par.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
				}
			} else {
				// not nbs.AlwaysNil
				if msgp.IsNil(bts) {
					bts = bts[1:]
					if nil != z.Par {
						z.Par.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
					}
				} else {
					// not nbs.AlwaysNil and not IsNil(bts): have something to read

					if z.Par == nil {
						z.Par = new(S2)
					}
					bts, err = z.Par.MSGPUnmarshalMsg(bts)
					if err != nil {
						return
					}
					if err != nil {
						return
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	if nextMiss23zgensym_58b5c2649cecee38_24 != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	if p, ok := interface{}(z).(msgp.PostLoad); ok {
		p.PostLoadHook()
	}

	return
}

// fields of Tree
var unmarshalMsgFieldOrder23zgensym_58b5c2649cecee38_24 = []string{"Chld_zid00_slc", "Str_zid01_str", "Par_zid02_ptr"}

var unmarshalMsgFieldSkip23zgensym_58b5c2649cecee38_24 = []bool{false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Tree) MSGPMsgsize() (s int) {
	s = 1 + 15 + msgp.ArrayHeaderSize
	for zgensym_58b5c2649cecee38_22 := range z.Chld {
		s += z.Chld[zgensym_58b5c2649cecee38_22].MSGPMsgsize()
	}
	s += 14 + msgp.StringPrefixSize + len(z.Str) + 14
	if z.Par == nil {
		s += msgp.NilSize
	} else {
		s += z.Par.MSGPMsgsize()
	}
	return
}
