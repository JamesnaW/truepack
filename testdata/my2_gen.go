package testdata

// NOTE: THIS FILE WAS PRODUCED BY THE
// TRUEPACK CODE GENERATION TOOL (github.com/glycerine/truepack)
// DO NOT EDIT

import (
	"github.com/glycerine/truepack/msgp"
)

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *Tr) DecodeMsg(dc *msgp.Reader) (err error) {
	var sawTopNil bool
	if dc.IsNil() {
		sawTopNil = true
		err = dc.ReadNil()
		if err != nil {
			return
		}
		dc.PushAlwaysNil()
	}

	var field []byte
	_ = field
	const maxFields5zgensym_72ba5d454ae3d9dd_6 = 6

	// -- templateDecodeMsg starts here--
	var totalEncodedFields5zgensym_72ba5d454ae3d9dd_6 uint32
	totalEncodedFields5zgensym_72ba5d454ae3d9dd_6, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft5zgensym_72ba5d454ae3d9dd_6 := totalEncodedFields5zgensym_72ba5d454ae3d9dd_6
	missingFieldsLeft5zgensym_72ba5d454ae3d9dd_6 := maxFields5zgensym_72ba5d454ae3d9dd_6 - totalEncodedFields5zgensym_72ba5d454ae3d9dd_6

	var nextMiss5zgensym_72ba5d454ae3d9dd_6 int32 = -1
	var found5zgensym_72ba5d454ae3d9dd_6 [maxFields5zgensym_72ba5d454ae3d9dd_6]bool
	var curField5zgensym_72ba5d454ae3d9dd_6 string

doneWithStruct5zgensym_72ba5d454ae3d9dd_6:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft5zgensym_72ba5d454ae3d9dd_6 > 0 || missingFieldsLeft5zgensym_72ba5d454ae3d9dd_6 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zgensym_72ba5d454ae3d9dd_6, missingFieldsLeft5zgensym_72ba5d454ae3d9dd_6, msgp.ShowFound(found5zgensym_72ba5d454ae3d9dd_6[:]), decodeMsgFieldOrder5zgensym_72ba5d454ae3d9dd_6)
		if encodedFieldsLeft5zgensym_72ba5d454ae3d9dd_6 > 0 {
			encodedFieldsLeft5zgensym_72ba5d454ae3d9dd_6--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField5zgensym_72ba5d454ae3d9dd_6 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss5zgensym_72ba5d454ae3d9dd_6 < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss5zgensym_72ba5d454ae3d9dd_6 = 0
			}
			for nextMiss5zgensym_72ba5d454ae3d9dd_6 < maxFields5zgensym_72ba5d454ae3d9dd_6 && (found5zgensym_72ba5d454ae3d9dd_6[nextMiss5zgensym_72ba5d454ae3d9dd_6] || decodeMsgFieldSkip5zgensym_72ba5d454ae3d9dd_6[nextMiss5zgensym_72ba5d454ae3d9dd_6]) {
				nextMiss5zgensym_72ba5d454ae3d9dd_6++
			}
			if nextMiss5zgensym_72ba5d454ae3d9dd_6 == maxFields5zgensym_72ba5d454ae3d9dd_6 {
				// filled all the empty fields!
				break doneWithStruct5zgensym_72ba5d454ae3d9dd_6
			}
			missingFieldsLeft5zgensym_72ba5d454ae3d9dd_6--
			curField5zgensym_72ba5d454ae3d9dd_6 = decodeMsgFieldOrder5zgensym_72ba5d454ae3d9dd_6[nextMiss5zgensym_72ba5d454ae3d9dd_6]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField5zgensym_72ba5d454ae3d9dd_6)
		switch curField5zgensym_72ba5d454ae3d9dd_6 {
		// -- templateDecodeMsg ends here --

		case "U_zid00_str":
			found5zgensym_72ba5d454ae3d9dd_6[0] = true
			z.U, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Nt_zid01_int":
			found5zgensym_72ba5d454ae3d9dd_6[1] = true
			z.Nt, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "Snot_zid02_map":
			found5zgensym_72ba5d454ae3d9dd_6[2] = true
			var zgensym_72ba5d454ae3d9dd_7 uint32
			zgensym_72ba5d454ae3d9dd_7, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Snot == nil && zgensym_72ba5d454ae3d9dd_7 > 0 {
				z.Snot = make(map[string]bool, zgensym_72ba5d454ae3d9dd_7)
			} else if len(z.Snot) > 0 {
				for key, _ := range z.Snot {
					delete(z.Snot, key)
				}
			}
			for zgensym_72ba5d454ae3d9dd_7 > 0 {
				zgensym_72ba5d454ae3d9dd_7--
				var zgensym_72ba5d454ae3d9dd_0 string
				var zgensym_72ba5d454ae3d9dd_1 bool
				zgensym_72ba5d454ae3d9dd_0, err = dc.ReadString()
				if err != nil {
					return
				}
				zgensym_72ba5d454ae3d9dd_1, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.Snot[zgensym_72ba5d454ae3d9dd_0] = zgensym_72ba5d454ae3d9dd_1
			}
		case "Notyet_zid03_map":
			found5zgensym_72ba5d454ae3d9dd_6[3] = true
			var zgensym_72ba5d454ae3d9dd_8 uint32
			zgensym_72ba5d454ae3d9dd_8, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Notyet == nil && zgensym_72ba5d454ae3d9dd_8 > 0 {
				z.Notyet = make(map[string]bool, zgensym_72ba5d454ae3d9dd_8)
			} else if len(z.Notyet) > 0 {
				for key, _ := range z.Notyet {
					delete(z.Notyet, key)
				}
			}
			for zgensym_72ba5d454ae3d9dd_8 > 0 {
				zgensym_72ba5d454ae3d9dd_8--
				var zgensym_72ba5d454ae3d9dd_2 string
				var zgensym_72ba5d454ae3d9dd_3 bool
				zgensym_72ba5d454ae3d9dd_2, err = dc.ReadString()
				if err != nil {
					return
				}
				zgensym_72ba5d454ae3d9dd_3, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.Notyet[zgensym_72ba5d454ae3d9dd_2] = zgensym_72ba5d454ae3d9dd_3
			}
		case "Setm_zid04_slc":
			found5zgensym_72ba5d454ae3d9dd_6[4] = true
			var zgensym_72ba5d454ae3d9dd_9 uint32
			zgensym_72ba5d454ae3d9dd_9, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Setm) >= int(zgensym_72ba5d454ae3d9dd_9) {
				z.Setm = (z.Setm)[:zgensym_72ba5d454ae3d9dd_9]
			} else {
				z.Setm = make([]*inn, zgensym_72ba5d454ae3d9dd_9)
			}
			for zgensym_72ba5d454ae3d9dd_4 := range z.Setm {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					if z.Setm[zgensym_72ba5d454ae3d9dd_4] != nil {
						dc.PushAlwaysNil()
						err = z.Setm[zgensym_72ba5d454ae3d9dd_4].DecodeMsg(dc)
						if err != nil {
							return
						}
						dc.PopAlwaysNil()
					}
				} else {
					// not Nil, we have something to read

					if z.Setm[zgensym_72ba5d454ae3d9dd_4] == nil {
						z.Setm[zgensym_72ba5d454ae3d9dd_4] = new(inn)
					}
					err = z.Setm[zgensym_72ba5d454ae3d9dd_4].DecodeMsg(dc)
					if err != nil {
						return
					}
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	if nextMiss5zgensym_72ba5d454ae3d9dd_6 != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	if p, ok := interface{}(z).(msgp.PostLoad); ok {
		p.PostLoadHook()
	}

	return
}

// fields of Tr
var decodeMsgFieldOrder5zgensym_72ba5d454ae3d9dd_6 = []string{"U_zid00_str", "Nt_zid01_int", "Snot_zid02_map", "Notyet_zid03_map", "Setm_zid04_slc", ""}

var decodeMsgFieldSkip5zgensym_72ba5d454ae3d9dd_6 = []bool{false, false, false, false, false, true}

// fieldsNotEmpty supports omitempty tags
func (z *Tr) fieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 5
	}
	var fieldsInUse uint32 = 5
	isempty[0] = (len(z.U) == 0) // string, omitempty
	if isempty[0] {
		fieldsInUse--
	}
	isempty[1] = (z.Nt == 0) // number, omitempty
	if isempty[1] {
		fieldsInUse--
	}
	isempty[2] = (len(z.Snot) == 0) // string, omitempty
	if isempty[2] {
		fieldsInUse--
	}
	isempty[3] = (len(z.Notyet) == 0) // string, omitempty
	if isempty[3] {
		fieldsInUse--
	}
	isempty[4] = (len(z.Setm) == 0) // string, omitempty
	if isempty[4] {
		fieldsInUse--
	}

	return fieldsInUse
}

// EncodeMsg implements msgp.Encodable
func (z *Tr) EncodeMsg(en *msgp.Writer) (err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	// honor the omitempty tags
	var empty_zgensym_72ba5d454ae3d9dd_10 [6]bool
	fieldsInUse_zgensym_72ba5d454ae3d9dd_11 := z.fieldsNotEmpty(empty_zgensym_72ba5d454ae3d9dd_10[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zgensym_72ba5d454ae3d9dd_11)
	if err != nil {
		return err
	}

	if !empty_zgensym_72ba5d454ae3d9dd_10[0] {
		// write "U_zid00_str"
		err = en.Append(0xab, 0x55, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x73, 0x74, 0x72)
		if err != nil {
			return err
		}
		err = en.WriteString(z.U)
		if err != nil {
			return
		}
	}

	if !empty_zgensym_72ba5d454ae3d9dd_10[1] {
		// write "Nt_zid01_int"
		err = en.Append(0xac, 0x4e, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x31, 0x5f, 0x69, 0x6e, 0x74)
		if err != nil {
			return err
		}
		err = en.WriteInt(z.Nt)
		if err != nil {
			return
		}
	}

	if !empty_zgensym_72ba5d454ae3d9dd_10[2] {
		// write "Snot_zid02_map"
		err = en.Append(0xae, 0x53, 0x6e, 0x6f, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x32, 0x5f, 0x6d, 0x61, 0x70)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Snot)))
		if err != nil {
			return
		}
		for zgensym_72ba5d454ae3d9dd_0, zgensym_72ba5d454ae3d9dd_1 := range z.Snot {
			err = en.WriteString(zgensym_72ba5d454ae3d9dd_0)
			if err != nil {
				return
			}
			err = en.WriteBool(zgensym_72ba5d454ae3d9dd_1)
			if err != nil {
				return
			}
		}
	}

	if !empty_zgensym_72ba5d454ae3d9dd_10[3] {
		// write "Notyet_zid03_map"
		err = en.Append(0xb0, 0x4e, 0x6f, 0x74, 0x79, 0x65, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x33, 0x5f, 0x6d, 0x61, 0x70)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Notyet)))
		if err != nil {
			return
		}
		for zgensym_72ba5d454ae3d9dd_2, zgensym_72ba5d454ae3d9dd_3 := range z.Notyet {
			err = en.WriteString(zgensym_72ba5d454ae3d9dd_2)
			if err != nil {
				return
			}
			err = en.WriteBool(zgensym_72ba5d454ae3d9dd_3)
			if err != nil {
				return
			}
		}
	}

	if !empty_zgensym_72ba5d454ae3d9dd_10[4] {
		// write "Setm_zid04_slc"
		err = en.Append(0xae, 0x53, 0x65, 0x74, 0x6d, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x34, 0x5f, 0x73, 0x6c, 0x63)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Setm)))
		if err != nil {
			return
		}
		for zgensym_72ba5d454ae3d9dd_4 := range z.Setm {
			if z.Setm[zgensym_72ba5d454ae3d9dd_4] == nil {
				err = en.WriteNil()
				if err != nil {
					return
				}
			} else {
				err = z.Setm[zgensym_72ba5d454ae3d9dd_4].EncodeMsg(en)
				if err != nil {
					return
				}
			}
		}
	}

	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Tr) MarshalMsg(b []byte) (o []byte, err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.Msgsize())

	// honor the omitempty tags
	var empty [6]bool
	fieldsInUse := z.fieldsNotEmpty(empty[:])
	o = msgp.AppendMapHeader(o, fieldsInUse)

	if !empty[0] {
		// string "U_zid00_str"
		o = append(o, 0xab, 0x55, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x73, 0x74, 0x72)
		o = msgp.AppendString(o, z.U)
	}

	if !empty[1] {
		// string "Nt_zid01_int"
		o = append(o, 0xac, 0x4e, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x31, 0x5f, 0x69, 0x6e, 0x74)
		o = msgp.AppendInt(o, z.Nt)
	}

	if !empty[2] {
		// string "Snot_zid02_map"
		o = append(o, 0xae, 0x53, 0x6e, 0x6f, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x32, 0x5f, 0x6d, 0x61, 0x70)
		o = msgp.AppendMapHeader(o, uint32(len(z.Snot)))
		for zgensym_72ba5d454ae3d9dd_0, zgensym_72ba5d454ae3d9dd_1 := range z.Snot {
			o = msgp.AppendString(o, zgensym_72ba5d454ae3d9dd_0)
			o = msgp.AppendBool(o, zgensym_72ba5d454ae3d9dd_1)
		}
	}

	if !empty[3] {
		// string "Notyet_zid03_map"
		o = append(o, 0xb0, 0x4e, 0x6f, 0x74, 0x79, 0x65, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x33, 0x5f, 0x6d, 0x61, 0x70)
		o = msgp.AppendMapHeader(o, uint32(len(z.Notyet)))
		for zgensym_72ba5d454ae3d9dd_2, zgensym_72ba5d454ae3d9dd_3 := range z.Notyet {
			o = msgp.AppendString(o, zgensym_72ba5d454ae3d9dd_2)
			o = msgp.AppendBool(o, zgensym_72ba5d454ae3d9dd_3)
		}
	}

	if !empty[4] {
		// string "Setm_zid04_slc"
		o = append(o, 0xae, 0x53, 0x65, 0x74, 0x6d, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x34, 0x5f, 0x73, 0x6c, 0x63)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Setm)))
		for zgensym_72ba5d454ae3d9dd_4 := range z.Setm {
			if z.Setm[zgensym_72ba5d454ae3d9dd_4] == nil {
				o = msgp.AppendNil(o)
			} else {
				o, err = z.Setm[zgensym_72ba5d454ae3d9dd_4].MarshalMsg(o)
				if err != nil {
					return
				}
			}
		}
	}

	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Tr) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *Tr) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields12zgensym_72ba5d454ae3d9dd_13 = 6

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields12zgensym_72ba5d454ae3d9dd_13 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields12zgensym_72ba5d454ae3d9dd_13, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft12zgensym_72ba5d454ae3d9dd_13 := totalEncodedFields12zgensym_72ba5d454ae3d9dd_13
	missingFieldsLeft12zgensym_72ba5d454ae3d9dd_13 := maxFields12zgensym_72ba5d454ae3d9dd_13 - totalEncodedFields12zgensym_72ba5d454ae3d9dd_13

	var nextMiss12zgensym_72ba5d454ae3d9dd_13 int32 = -1
	var found12zgensym_72ba5d454ae3d9dd_13 [maxFields12zgensym_72ba5d454ae3d9dd_13]bool
	var curField12zgensym_72ba5d454ae3d9dd_13 string

doneWithStruct12zgensym_72ba5d454ae3d9dd_13:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft12zgensym_72ba5d454ae3d9dd_13 > 0 || missingFieldsLeft12zgensym_72ba5d454ae3d9dd_13 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft12zgensym_72ba5d454ae3d9dd_13, missingFieldsLeft12zgensym_72ba5d454ae3d9dd_13, msgp.ShowFound(found12zgensym_72ba5d454ae3d9dd_13[:]), unmarshalMsgFieldOrder12zgensym_72ba5d454ae3d9dd_13)
		if encodedFieldsLeft12zgensym_72ba5d454ae3d9dd_13 > 0 {
			encodedFieldsLeft12zgensym_72ba5d454ae3d9dd_13--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField12zgensym_72ba5d454ae3d9dd_13 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss12zgensym_72ba5d454ae3d9dd_13 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss12zgensym_72ba5d454ae3d9dd_13 = 0
			}
			for nextMiss12zgensym_72ba5d454ae3d9dd_13 < maxFields12zgensym_72ba5d454ae3d9dd_13 && (found12zgensym_72ba5d454ae3d9dd_13[nextMiss12zgensym_72ba5d454ae3d9dd_13] || unmarshalMsgFieldSkip12zgensym_72ba5d454ae3d9dd_13[nextMiss12zgensym_72ba5d454ae3d9dd_13]) {
				nextMiss12zgensym_72ba5d454ae3d9dd_13++
			}
			if nextMiss12zgensym_72ba5d454ae3d9dd_13 == maxFields12zgensym_72ba5d454ae3d9dd_13 {
				// filled all the empty fields!
				break doneWithStruct12zgensym_72ba5d454ae3d9dd_13
			}
			missingFieldsLeft12zgensym_72ba5d454ae3d9dd_13--
			curField12zgensym_72ba5d454ae3d9dd_13 = unmarshalMsgFieldOrder12zgensym_72ba5d454ae3d9dd_13[nextMiss12zgensym_72ba5d454ae3d9dd_13]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField12zgensym_72ba5d454ae3d9dd_13)
		switch curField12zgensym_72ba5d454ae3d9dd_13 {
		// -- templateUnmarshalMsg ends here --

		case "U_zid00_str":
			found12zgensym_72ba5d454ae3d9dd_13[0] = true
			z.U, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Nt_zid01_int":
			found12zgensym_72ba5d454ae3d9dd_13[1] = true
			z.Nt, bts, err = nbs.ReadIntBytes(bts)

			if err != nil {
				return
			}
		case "Snot_zid02_map":
			found12zgensym_72ba5d454ae3d9dd_13[2] = true
			if nbs.AlwaysNil {
				if len(z.Snot) > 0 {
					for key, _ := range z.Snot {
						delete(z.Snot, key)
					}
				}

			} else {

				var zgensym_72ba5d454ae3d9dd_14 uint32
				zgensym_72ba5d454ae3d9dd_14, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Snot == nil && zgensym_72ba5d454ae3d9dd_14 > 0 {
					z.Snot = make(map[string]bool, zgensym_72ba5d454ae3d9dd_14)
				} else if len(z.Snot) > 0 {
					for key, _ := range z.Snot {
						delete(z.Snot, key)
					}
				}
				for zgensym_72ba5d454ae3d9dd_14 > 0 {
					var zgensym_72ba5d454ae3d9dd_0 string
					var zgensym_72ba5d454ae3d9dd_1 bool
					zgensym_72ba5d454ae3d9dd_14--
					zgensym_72ba5d454ae3d9dd_0, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zgensym_72ba5d454ae3d9dd_1, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.Snot[zgensym_72ba5d454ae3d9dd_0] = zgensym_72ba5d454ae3d9dd_1
				}
			}
		case "Notyet_zid03_map":
			found12zgensym_72ba5d454ae3d9dd_13[3] = true
			if nbs.AlwaysNil {
				if len(z.Notyet) > 0 {
					for key, _ := range z.Notyet {
						delete(z.Notyet, key)
					}
				}

			} else {

				var zgensym_72ba5d454ae3d9dd_15 uint32
				zgensym_72ba5d454ae3d9dd_15, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Notyet == nil && zgensym_72ba5d454ae3d9dd_15 > 0 {
					z.Notyet = make(map[string]bool, zgensym_72ba5d454ae3d9dd_15)
				} else if len(z.Notyet) > 0 {
					for key, _ := range z.Notyet {
						delete(z.Notyet, key)
					}
				}
				for zgensym_72ba5d454ae3d9dd_15 > 0 {
					var zgensym_72ba5d454ae3d9dd_2 string
					var zgensym_72ba5d454ae3d9dd_3 bool
					zgensym_72ba5d454ae3d9dd_15--
					zgensym_72ba5d454ae3d9dd_2, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zgensym_72ba5d454ae3d9dd_3, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.Notyet[zgensym_72ba5d454ae3d9dd_2] = zgensym_72ba5d454ae3d9dd_3
				}
			}
		case "Setm_zid04_slc":
			found12zgensym_72ba5d454ae3d9dd_13[4] = true
			if nbs.AlwaysNil {
				(z.Setm) = (z.Setm)[:0]
			} else {

				var zgensym_72ba5d454ae3d9dd_16 uint32
				zgensym_72ba5d454ae3d9dd_16, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Setm) >= int(zgensym_72ba5d454ae3d9dd_16) {
					z.Setm = (z.Setm)[:zgensym_72ba5d454ae3d9dd_16]
				} else {
					z.Setm = make([]*inn, zgensym_72ba5d454ae3d9dd_16)
				}
				for zgensym_72ba5d454ae3d9dd_4 := range z.Setm {
					if nbs.AlwaysNil {
						if z.Setm[zgensym_72ba5d454ae3d9dd_4] != nil {
							z.Setm[zgensym_72ba5d454ae3d9dd_4].UnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != z.Setm[zgensym_72ba5d454ae3d9dd_4] {
								z.Setm[zgensym_72ba5d454ae3d9dd_4].UnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if z.Setm[zgensym_72ba5d454ae3d9dd_4] == nil {
								z.Setm[zgensym_72ba5d454ae3d9dd_4] = new(inn)
							}
							bts, err = z.Setm[zgensym_72ba5d454ae3d9dd_4].UnmarshalMsg(bts)
							if err != nil {
								return
							}
							if err != nil {
								return
							}
						}
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
	if nextMiss12zgensym_72ba5d454ae3d9dd_13 != -1 {
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

// fields of Tr
var unmarshalMsgFieldOrder12zgensym_72ba5d454ae3d9dd_13 = []string{"U_zid00_str", "Nt_zid01_int", "Snot_zid02_map", "Notyet_zid03_map", "Setm_zid04_slc", ""}

var unmarshalMsgFieldSkip12zgensym_72ba5d454ae3d9dd_13 = []bool{false, false, false, false, false, true}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Tr) Msgsize() (s int) {
	s = 1 + 12 + msgp.StringPrefixSize + len(z.U) + 13 + msgp.IntSize + 15 + msgp.MapHeaderSize
	if z.Snot != nil {
		for zgensym_72ba5d454ae3d9dd_0, zgensym_72ba5d454ae3d9dd_1 := range z.Snot {
			_ = zgensym_72ba5d454ae3d9dd_1
			_ = zgensym_72ba5d454ae3d9dd_0
			s += msgp.StringPrefixSize + len(zgensym_72ba5d454ae3d9dd_0) + msgp.BoolSize
		}
	}
	s += 17 + msgp.MapHeaderSize
	if z.Notyet != nil {
		for zgensym_72ba5d454ae3d9dd_2, zgensym_72ba5d454ae3d9dd_3 := range z.Notyet {
			_ = zgensym_72ba5d454ae3d9dd_3
			_ = zgensym_72ba5d454ae3d9dd_2
			s += msgp.StringPrefixSize + len(zgensym_72ba5d454ae3d9dd_2) + msgp.BoolSize
		}
	}
	s += 15 + msgp.ArrayHeaderSize
	for zgensym_72ba5d454ae3d9dd_4 := range z.Setm {
		if z.Setm[zgensym_72ba5d454ae3d9dd_4] == nil {
			s += msgp.NilSize
		} else {
			s += z.Setm[zgensym_72ba5d454ae3d9dd_4].Msgsize()
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *inn) DecodeMsg(dc *msgp.Reader) (err error) {
	var sawTopNil bool
	if dc.IsNil() {
		sawTopNil = true
		err = dc.ReadNil()
		if err != nil {
			return
		}
		dc.PushAlwaysNil()
	}

	var field []byte
	_ = field
	const maxFields19zgensym_72ba5d454ae3d9dd_20 = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields19zgensym_72ba5d454ae3d9dd_20 uint32
	totalEncodedFields19zgensym_72ba5d454ae3d9dd_20, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft19zgensym_72ba5d454ae3d9dd_20 := totalEncodedFields19zgensym_72ba5d454ae3d9dd_20
	missingFieldsLeft19zgensym_72ba5d454ae3d9dd_20 := maxFields19zgensym_72ba5d454ae3d9dd_20 - totalEncodedFields19zgensym_72ba5d454ae3d9dd_20

	var nextMiss19zgensym_72ba5d454ae3d9dd_20 int32 = -1
	var found19zgensym_72ba5d454ae3d9dd_20 [maxFields19zgensym_72ba5d454ae3d9dd_20]bool
	var curField19zgensym_72ba5d454ae3d9dd_20 string

doneWithStruct19zgensym_72ba5d454ae3d9dd_20:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft19zgensym_72ba5d454ae3d9dd_20 > 0 || missingFieldsLeft19zgensym_72ba5d454ae3d9dd_20 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft19zgensym_72ba5d454ae3d9dd_20, missingFieldsLeft19zgensym_72ba5d454ae3d9dd_20, msgp.ShowFound(found19zgensym_72ba5d454ae3d9dd_20[:]), decodeMsgFieldOrder19zgensym_72ba5d454ae3d9dd_20)
		if encodedFieldsLeft19zgensym_72ba5d454ae3d9dd_20 > 0 {
			encodedFieldsLeft19zgensym_72ba5d454ae3d9dd_20--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField19zgensym_72ba5d454ae3d9dd_20 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss19zgensym_72ba5d454ae3d9dd_20 < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss19zgensym_72ba5d454ae3d9dd_20 = 0
			}
			for nextMiss19zgensym_72ba5d454ae3d9dd_20 < maxFields19zgensym_72ba5d454ae3d9dd_20 && (found19zgensym_72ba5d454ae3d9dd_20[nextMiss19zgensym_72ba5d454ae3d9dd_20] || decodeMsgFieldSkip19zgensym_72ba5d454ae3d9dd_20[nextMiss19zgensym_72ba5d454ae3d9dd_20]) {
				nextMiss19zgensym_72ba5d454ae3d9dd_20++
			}
			if nextMiss19zgensym_72ba5d454ae3d9dd_20 == maxFields19zgensym_72ba5d454ae3d9dd_20 {
				// filled all the empty fields!
				break doneWithStruct19zgensym_72ba5d454ae3d9dd_20
			}
			missingFieldsLeft19zgensym_72ba5d454ae3d9dd_20--
			curField19zgensym_72ba5d454ae3d9dd_20 = decodeMsgFieldOrder19zgensym_72ba5d454ae3d9dd_20[nextMiss19zgensym_72ba5d454ae3d9dd_20]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField19zgensym_72ba5d454ae3d9dd_20)
		switch curField19zgensym_72ba5d454ae3d9dd_20 {
		// -- templateDecodeMsg ends here --

		case "j_zid00_map":
			found19zgensym_72ba5d454ae3d9dd_20[0] = true
			var zgensym_72ba5d454ae3d9dd_21 uint32
			zgensym_72ba5d454ae3d9dd_21, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.j == nil && zgensym_72ba5d454ae3d9dd_21 > 0 {
				z.j = make(map[string]bool, zgensym_72ba5d454ae3d9dd_21)
			} else if len(z.j) > 0 {
				for key, _ := range z.j {
					delete(z.j, key)
				}
			}
			for zgensym_72ba5d454ae3d9dd_21 > 0 {
				zgensym_72ba5d454ae3d9dd_21--
				var zgensym_72ba5d454ae3d9dd_17 string
				var zgensym_72ba5d454ae3d9dd_18 bool
				zgensym_72ba5d454ae3d9dd_17, err = dc.ReadString()
				if err != nil {
					return
				}
				zgensym_72ba5d454ae3d9dd_18, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.j[zgensym_72ba5d454ae3d9dd_17] = zgensym_72ba5d454ae3d9dd_18
			}
		case "e_zid01_i64":
			found19zgensym_72ba5d454ae3d9dd_20[1] = true
			z.e, err = dc.ReadInt64()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	if nextMiss19zgensym_72ba5d454ae3d9dd_20 != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	if p, ok := interface{}(z).(msgp.PostLoad); ok {
		p.PostLoadHook()
	}

	return
}

// fields of inn
var decodeMsgFieldOrder19zgensym_72ba5d454ae3d9dd_20 = []string{"j_zid00_map", "e_zid01_i64"}

var decodeMsgFieldSkip19zgensym_72ba5d454ae3d9dd_20 = []bool{false, false}

// fieldsNotEmpty supports omitempty tags
func (z *inn) fieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 2
	}
	var fieldsInUse uint32 = 2
	isempty[0] = (len(z.j) == 0) // string, omitempty
	if isempty[0] {
		fieldsInUse--
	}
	isempty[1] = (z.e == 0) // number, omitempty
	if isempty[1] {
		fieldsInUse--
	}

	return fieldsInUse
}

// EncodeMsg implements msgp.Encodable
func (z *inn) EncodeMsg(en *msgp.Writer) (err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	// honor the omitempty tags
	var empty_zgensym_72ba5d454ae3d9dd_22 [2]bool
	fieldsInUse_zgensym_72ba5d454ae3d9dd_23 := z.fieldsNotEmpty(empty_zgensym_72ba5d454ae3d9dd_22[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zgensym_72ba5d454ae3d9dd_23)
	if err != nil {
		return err
	}

	if !empty_zgensym_72ba5d454ae3d9dd_22[0] {
		// write "j_zid00_map"
		err = en.Append(0xab, 0x6a, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x6d, 0x61, 0x70)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.j)))
		if err != nil {
			return
		}
		for zgensym_72ba5d454ae3d9dd_17, zgensym_72ba5d454ae3d9dd_18 := range z.j {
			err = en.WriteString(zgensym_72ba5d454ae3d9dd_17)
			if err != nil {
				return
			}
			err = en.WriteBool(zgensym_72ba5d454ae3d9dd_18)
			if err != nil {
				return
			}
		}
	}

	if !empty_zgensym_72ba5d454ae3d9dd_22[1] {
		// write "e_zid01_i64"
		err = en.Append(0xab, 0x65, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x31, 0x5f, 0x69, 0x36, 0x34)
		if err != nil {
			return err
		}
		err = en.WriteInt64(z.e)
		if err != nil {
			return
		}
	}

	return
}

// MarshalMsg implements msgp.Marshaler
func (z *inn) MarshalMsg(b []byte) (o []byte, err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.Msgsize())

	// honor the omitempty tags
	var empty [2]bool
	fieldsInUse := z.fieldsNotEmpty(empty[:])
	o = msgp.AppendMapHeader(o, fieldsInUse)

	if !empty[0] {
		// string "j_zid00_map"
		o = append(o, 0xab, 0x6a, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x6d, 0x61, 0x70)
		o = msgp.AppendMapHeader(o, uint32(len(z.j)))
		for zgensym_72ba5d454ae3d9dd_17, zgensym_72ba5d454ae3d9dd_18 := range z.j {
			o = msgp.AppendString(o, zgensym_72ba5d454ae3d9dd_17)
			o = msgp.AppendBool(o, zgensym_72ba5d454ae3d9dd_18)
		}
	}

	if !empty[1] {
		// string "e_zid01_i64"
		o = append(o, 0xab, 0x65, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x31, 0x5f, 0x69, 0x36, 0x34)
		o = msgp.AppendInt64(o, z.e)
	}

	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *inn) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *inn) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields24zgensym_72ba5d454ae3d9dd_25 = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields24zgensym_72ba5d454ae3d9dd_25 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields24zgensym_72ba5d454ae3d9dd_25, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft24zgensym_72ba5d454ae3d9dd_25 := totalEncodedFields24zgensym_72ba5d454ae3d9dd_25
	missingFieldsLeft24zgensym_72ba5d454ae3d9dd_25 := maxFields24zgensym_72ba5d454ae3d9dd_25 - totalEncodedFields24zgensym_72ba5d454ae3d9dd_25

	var nextMiss24zgensym_72ba5d454ae3d9dd_25 int32 = -1
	var found24zgensym_72ba5d454ae3d9dd_25 [maxFields24zgensym_72ba5d454ae3d9dd_25]bool
	var curField24zgensym_72ba5d454ae3d9dd_25 string

doneWithStruct24zgensym_72ba5d454ae3d9dd_25:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft24zgensym_72ba5d454ae3d9dd_25 > 0 || missingFieldsLeft24zgensym_72ba5d454ae3d9dd_25 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft24zgensym_72ba5d454ae3d9dd_25, missingFieldsLeft24zgensym_72ba5d454ae3d9dd_25, msgp.ShowFound(found24zgensym_72ba5d454ae3d9dd_25[:]), unmarshalMsgFieldOrder24zgensym_72ba5d454ae3d9dd_25)
		if encodedFieldsLeft24zgensym_72ba5d454ae3d9dd_25 > 0 {
			encodedFieldsLeft24zgensym_72ba5d454ae3d9dd_25--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField24zgensym_72ba5d454ae3d9dd_25 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss24zgensym_72ba5d454ae3d9dd_25 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss24zgensym_72ba5d454ae3d9dd_25 = 0
			}
			for nextMiss24zgensym_72ba5d454ae3d9dd_25 < maxFields24zgensym_72ba5d454ae3d9dd_25 && (found24zgensym_72ba5d454ae3d9dd_25[nextMiss24zgensym_72ba5d454ae3d9dd_25] || unmarshalMsgFieldSkip24zgensym_72ba5d454ae3d9dd_25[nextMiss24zgensym_72ba5d454ae3d9dd_25]) {
				nextMiss24zgensym_72ba5d454ae3d9dd_25++
			}
			if nextMiss24zgensym_72ba5d454ae3d9dd_25 == maxFields24zgensym_72ba5d454ae3d9dd_25 {
				// filled all the empty fields!
				break doneWithStruct24zgensym_72ba5d454ae3d9dd_25
			}
			missingFieldsLeft24zgensym_72ba5d454ae3d9dd_25--
			curField24zgensym_72ba5d454ae3d9dd_25 = unmarshalMsgFieldOrder24zgensym_72ba5d454ae3d9dd_25[nextMiss24zgensym_72ba5d454ae3d9dd_25]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField24zgensym_72ba5d454ae3d9dd_25)
		switch curField24zgensym_72ba5d454ae3d9dd_25 {
		// -- templateUnmarshalMsg ends here --

		case "j_zid00_map":
			found24zgensym_72ba5d454ae3d9dd_25[0] = true
			if nbs.AlwaysNil {
				if len(z.j) > 0 {
					for key, _ := range z.j {
						delete(z.j, key)
					}
				}

			} else {

				var zgensym_72ba5d454ae3d9dd_26 uint32
				zgensym_72ba5d454ae3d9dd_26, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.j == nil && zgensym_72ba5d454ae3d9dd_26 > 0 {
					z.j = make(map[string]bool, zgensym_72ba5d454ae3d9dd_26)
				} else if len(z.j) > 0 {
					for key, _ := range z.j {
						delete(z.j, key)
					}
				}
				for zgensym_72ba5d454ae3d9dd_26 > 0 {
					var zgensym_72ba5d454ae3d9dd_17 string
					var zgensym_72ba5d454ae3d9dd_18 bool
					zgensym_72ba5d454ae3d9dd_26--
					zgensym_72ba5d454ae3d9dd_17, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zgensym_72ba5d454ae3d9dd_18, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.j[zgensym_72ba5d454ae3d9dd_17] = zgensym_72ba5d454ae3d9dd_18
				}
			}
		case "e_zid01_i64":
			found24zgensym_72ba5d454ae3d9dd_25[1] = true
			z.e, bts, err = nbs.ReadInt64Bytes(bts)

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
	if nextMiss24zgensym_72ba5d454ae3d9dd_25 != -1 {
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

// fields of inn
var unmarshalMsgFieldOrder24zgensym_72ba5d454ae3d9dd_25 = []string{"j_zid00_map", "e_zid01_i64"}

var unmarshalMsgFieldSkip24zgensym_72ba5d454ae3d9dd_25 = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *inn) Msgsize() (s int) {
	s = 1 + 12 + msgp.MapHeaderSize
	if z.j != nil {
		for zgensym_72ba5d454ae3d9dd_17, zgensym_72ba5d454ae3d9dd_18 := range z.j {
			_ = zgensym_72ba5d454ae3d9dd_18
			_ = zgensym_72ba5d454ae3d9dd_17
			s += msgp.StringPrefixSize + len(zgensym_72ba5d454ae3d9dd_17) + msgp.BoolSize
		}
	}
	s += 12 + msgp.Int64Size
	return
}

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *u) DecodeMsg(dc *msgp.Reader) (err error) {
	var sawTopNil bool
	if dc.IsNil() {
		sawTopNil = true
		err = dc.ReadNil()
		if err != nil {
			return
		}
		dc.PushAlwaysNil()
	}

	var field []byte
	_ = field
	const maxFields29zgensym_72ba5d454ae3d9dd_30 = 3

	// -- templateDecodeMsg starts here--
	var totalEncodedFields29zgensym_72ba5d454ae3d9dd_30 uint32
	totalEncodedFields29zgensym_72ba5d454ae3d9dd_30, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft29zgensym_72ba5d454ae3d9dd_30 := totalEncodedFields29zgensym_72ba5d454ae3d9dd_30
	missingFieldsLeft29zgensym_72ba5d454ae3d9dd_30 := maxFields29zgensym_72ba5d454ae3d9dd_30 - totalEncodedFields29zgensym_72ba5d454ae3d9dd_30

	var nextMiss29zgensym_72ba5d454ae3d9dd_30 int32 = -1
	var found29zgensym_72ba5d454ae3d9dd_30 [maxFields29zgensym_72ba5d454ae3d9dd_30]bool
	var curField29zgensym_72ba5d454ae3d9dd_30 string

doneWithStruct29zgensym_72ba5d454ae3d9dd_30:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft29zgensym_72ba5d454ae3d9dd_30 > 0 || missingFieldsLeft29zgensym_72ba5d454ae3d9dd_30 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft29zgensym_72ba5d454ae3d9dd_30, missingFieldsLeft29zgensym_72ba5d454ae3d9dd_30, msgp.ShowFound(found29zgensym_72ba5d454ae3d9dd_30[:]), decodeMsgFieldOrder29zgensym_72ba5d454ae3d9dd_30)
		if encodedFieldsLeft29zgensym_72ba5d454ae3d9dd_30 > 0 {
			encodedFieldsLeft29zgensym_72ba5d454ae3d9dd_30--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField29zgensym_72ba5d454ae3d9dd_30 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss29zgensym_72ba5d454ae3d9dd_30 < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss29zgensym_72ba5d454ae3d9dd_30 = 0
			}
			for nextMiss29zgensym_72ba5d454ae3d9dd_30 < maxFields29zgensym_72ba5d454ae3d9dd_30 && (found29zgensym_72ba5d454ae3d9dd_30[nextMiss29zgensym_72ba5d454ae3d9dd_30] || decodeMsgFieldSkip29zgensym_72ba5d454ae3d9dd_30[nextMiss29zgensym_72ba5d454ae3d9dd_30]) {
				nextMiss29zgensym_72ba5d454ae3d9dd_30++
			}
			if nextMiss29zgensym_72ba5d454ae3d9dd_30 == maxFields29zgensym_72ba5d454ae3d9dd_30 {
				// filled all the empty fields!
				break doneWithStruct29zgensym_72ba5d454ae3d9dd_30
			}
			missingFieldsLeft29zgensym_72ba5d454ae3d9dd_30--
			curField29zgensym_72ba5d454ae3d9dd_30 = decodeMsgFieldOrder29zgensym_72ba5d454ae3d9dd_30[nextMiss29zgensym_72ba5d454ae3d9dd_30]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField29zgensym_72ba5d454ae3d9dd_30)
		switch curField29zgensym_72ba5d454ae3d9dd_30 {
		// -- templateDecodeMsg ends here --

		case "m_zid00_map":
			found29zgensym_72ba5d454ae3d9dd_30[0] = true
			var zgensym_72ba5d454ae3d9dd_31 uint32
			zgensym_72ba5d454ae3d9dd_31, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.m == nil && zgensym_72ba5d454ae3d9dd_31 > 0 {
				z.m = make(map[string]*Tr, zgensym_72ba5d454ae3d9dd_31)
			} else if len(z.m) > 0 {
				for key, _ := range z.m {
					delete(z.m, key)
				}
			}
			for zgensym_72ba5d454ae3d9dd_31 > 0 {
				zgensym_72ba5d454ae3d9dd_31--
				var zgensym_72ba5d454ae3d9dd_27 string
				var zgensym_72ba5d454ae3d9dd_28 *Tr
				zgensym_72ba5d454ae3d9dd_27, err = dc.ReadString()
				if err != nil {
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					if zgensym_72ba5d454ae3d9dd_28 != nil {
						dc.PushAlwaysNil()
						err = zgensym_72ba5d454ae3d9dd_28.DecodeMsg(dc)
						if err != nil {
							return
						}
						dc.PopAlwaysNil()
					}
				} else {
					// not Nil, we have something to read

					if zgensym_72ba5d454ae3d9dd_28 == nil {
						zgensym_72ba5d454ae3d9dd_28 = new(Tr)
					}
					err = zgensym_72ba5d454ae3d9dd_28.DecodeMsg(dc)
					if err != nil {
						return
					}
				}
				z.m[zgensym_72ba5d454ae3d9dd_27] = zgensym_72ba5d454ae3d9dd_28
			}
		case "s_zid01_str":
			found29zgensym_72ba5d454ae3d9dd_30[1] = true
			z.s, err = dc.ReadString()
			if err != nil {
				return
			}
		case "n_zid02_i64":
			found29zgensym_72ba5d454ae3d9dd_30[2] = true
			z.n, err = dc.ReadInt64()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	if nextMiss29zgensym_72ba5d454ae3d9dd_30 != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	if p, ok := interface{}(z).(msgp.PostLoad); ok {
		p.PostLoadHook()
	}

	return
}

// fields of u
var decodeMsgFieldOrder29zgensym_72ba5d454ae3d9dd_30 = []string{"m_zid00_map", "s_zid01_str", "n_zid02_i64"}

var decodeMsgFieldSkip29zgensym_72ba5d454ae3d9dd_30 = []bool{false, false, false}

// fieldsNotEmpty supports omitempty tags
func (z *u) fieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 3
	}
	var fieldsInUse uint32 = 3
	isempty[0] = (len(z.m) == 0) // string, omitempty
	if isempty[0] {
		fieldsInUse--
	}
	isempty[1] = (len(z.s) == 0) // string, omitempty
	if isempty[1] {
		fieldsInUse--
	}
	isempty[2] = (z.n == 0) // number, omitempty
	if isempty[2] {
		fieldsInUse--
	}

	return fieldsInUse
}

// EncodeMsg implements msgp.Encodable
func (z *u) EncodeMsg(en *msgp.Writer) (err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	// honor the omitempty tags
	var empty_zgensym_72ba5d454ae3d9dd_32 [3]bool
	fieldsInUse_zgensym_72ba5d454ae3d9dd_33 := z.fieldsNotEmpty(empty_zgensym_72ba5d454ae3d9dd_32[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zgensym_72ba5d454ae3d9dd_33)
	if err != nil {
		return err
	}

	if !empty_zgensym_72ba5d454ae3d9dd_32[0] {
		// write "m_zid00_map"
		err = en.Append(0xab, 0x6d, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x6d, 0x61, 0x70)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.m)))
		if err != nil {
			return
		}
		for zgensym_72ba5d454ae3d9dd_27, zgensym_72ba5d454ae3d9dd_28 := range z.m {
			err = en.WriteString(zgensym_72ba5d454ae3d9dd_27)
			if err != nil {
				return
			}
			if zgensym_72ba5d454ae3d9dd_28 == nil {
				err = en.WriteNil()
				if err != nil {
					return
				}
			} else {
				err = zgensym_72ba5d454ae3d9dd_28.EncodeMsg(en)
				if err != nil {
					return
				}
			}
		}
	}

	if !empty_zgensym_72ba5d454ae3d9dd_32[1] {
		// write "s_zid01_str"
		err = en.Append(0xab, 0x73, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x31, 0x5f, 0x73, 0x74, 0x72)
		if err != nil {
			return err
		}
		err = en.WriteString(z.s)
		if err != nil {
			return
		}
	}

	if !empty_zgensym_72ba5d454ae3d9dd_32[2] {
		// write "n_zid02_i64"
		err = en.Append(0xab, 0x6e, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x32, 0x5f, 0x69, 0x36, 0x34)
		if err != nil {
			return err
		}
		err = en.WriteInt64(z.n)
		if err != nil {
			return
		}
	}

	return
}

// MarshalMsg implements msgp.Marshaler
func (z *u) MarshalMsg(b []byte) (o []byte, err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.Msgsize())

	// honor the omitempty tags
	var empty [3]bool
	fieldsInUse := z.fieldsNotEmpty(empty[:])
	o = msgp.AppendMapHeader(o, fieldsInUse)

	if !empty[0] {
		// string "m_zid00_map"
		o = append(o, 0xab, 0x6d, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x6d, 0x61, 0x70)
		o = msgp.AppendMapHeader(o, uint32(len(z.m)))
		for zgensym_72ba5d454ae3d9dd_27, zgensym_72ba5d454ae3d9dd_28 := range z.m {
			o = msgp.AppendString(o, zgensym_72ba5d454ae3d9dd_27)
			if zgensym_72ba5d454ae3d9dd_28 == nil {
				o = msgp.AppendNil(o)
			} else {
				o, err = zgensym_72ba5d454ae3d9dd_28.MarshalMsg(o)
				if err != nil {
					return
				}
			}
		}
	}

	if !empty[1] {
		// string "s_zid01_str"
		o = append(o, 0xab, 0x73, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x31, 0x5f, 0x73, 0x74, 0x72)
		o = msgp.AppendString(o, z.s)
	}

	if !empty[2] {
		// string "n_zid02_i64"
		o = append(o, 0xab, 0x6e, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x32, 0x5f, 0x69, 0x36, 0x34)
		o = msgp.AppendInt64(o, z.n)
	}

	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *u) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *u) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields34zgensym_72ba5d454ae3d9dd_35 = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields34zgensym_72ba5d454ae3d9dd_35 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields34zgensym_72ba5d454ae3d9dd_35, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft34zgensym_72ba5d454ae3d9dd_35 := totalEncodedFields34zgensym_72ba5d454ae3d9dd_35
	missingFieldsLeft34zgensym_72ba5d454ae3d9dd_35 := maxFields34zgensym_72ba5d454ae3d9dd_35 - totalEncodedFields34zgensym_72ba5d454ae3d9dd_35

	var nextMiss34zgensym_72ba5d454ae3d9dd_35 int32 = -1
	var found34zgensym_72ba5d454ae3d9dd_35 [maxFields34zgensym_72ba5d454ae3d9dd_35]bool
	var curField34zgensym_72ba5d454ae3d9dd_35 string

doneWithStruct34zgensym_72ba5d454ae3d9dd_35:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft34zgensym_72ba5d454ae3d9dd_35 > 0 || missingFieldsLeft34zgensym_72ba5d454ae3d9dd_35 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft34zgensym_72ba5d454ae3d9dd_35, missingFieldsLeft34zgensym_72ba5d454ae3d9dd_35, msgp.ShowFound(found34zgensym_72ba5d454ae3d9dd_35[:]), unmarshalMsgFieldOrder34zgensym_72ba5d454ae3d9dd_35)
		if encodedFieldsLeft34zgensym_72ba5d454ae3d9dd_35 > 0 {
			encodedFieldsLeft34zgensym_72ba5d454ae3d9dd_35--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField34zgensym_72ba5d454ae3d9dd_35 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss34zgensym_72ba5d454ae3d9dd_35 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss34zgensym_72ba5d454ae3d9dd_35 = 0
			}
			for nextMiss34zgensym_72ba5d454ae3d9dd_35 < maxFields34zgensym_72ba5d454ae3d9dd_35 && (found34zgensym_72ba5d454ae3d9dd_35[nextMiss34zgensym_72ba5d454ae3d9dd_35] || unmarshalMsgFieldSkip34zgensym_72ba5d454ae3d9dd_35[nextMiss34zgensym_72ba5d454ae3d9dd_35]) {
				nextMiss34zgensym_72ba5d454ae3d9dd_35++
			}
			if nextMiss34zgensym_72ba5d454ae3d9dd_35 == maxFields34zgensym_72ba5d454ae3d9dd_35 {
				// filled all the empty fields!
				break doneWithStruct34zgensym_72ba5d454ae3d9dd_35
			}
			missingFieldsLeft34zgensym_72ba5d454ae3d9dd_35--
			curField34zgensym_72ba5d454ae3d9dd_35 = unmarshalMsgFieldOrder34zgensym_72ba5d454ae3d9dd_35[nextMiss34zgensym_72ba5d454ae3d9dd_35]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField34zgensym_72ba5d454ae3d9dd_35)
		switch curField34zgensym_72ba5d454ae3d9dd_35 {
		// -- templateUnmarshalMsg ends here --

		case "m_zid00_map":
			found34zgensym_72ba5d454ae3d9dd_35[0] = true
			if nbs.AlwaysNil {
				if len(z.m) > 0 {
					for key, _ := range z.m {
						delete(z.m, key)
					}
				}

			} else {

				var zgensym_72ba5d454ae3d9dd_36 uint32
				zgensym_72ba5d454ae3d9dd_36, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.m == nil && zgensym_72ba5d454ae3d9dd_36 > 0 {
					z.m = make(map[string]*Tr, zgensym_72ba5d454ae3d9dd_36)
				} else if len(z.m) > 0 {
					for key, _ := range z.m {
						delete(z.m, key)
					}
				}
				for zgensym_72ba5d454ae3d9dd_36 > 0 {
					var zgensym_72ba5d454ae3d9dd_27 string
					var zgensym_72ba5d454ae3d9dd_28 *Tr
					zgensym_72ba5d454ae3d9dd_36--
					zgensym_72ba5d454ae3d9dd_27, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					if nbs.AlwaysNil {
						if zgensym_72ba5d454ae3d9dd_28 != nil {
							zgensym_72ba5d454ae3d9dd_28.UnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != zgensym_72ba5d454ae3d9dd_28 {
								zgensym_72ba5d454ae3d9dd_28.UnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if zgensym_72ba5d454ae3d9dd_28 == nil {
								zgensym_72ba5d454ae3d9dd_28 = new(Tr)
							}
							bts, err = zgensym_72ba5d454ae3d9dd_28.UnmarshalMsg(bts)
							if err != nil {
								return
							}
							if err != nil {
								return
							}
						}
					}
					z.m[zgensym_72ba5d454ae3d9dd_27] = zgensym_72ba5d454ae3d9dd_28
				}
			}
		case "s_zid01_str":
			found34zgensym_72ba5d454ae3d9dd_35[1] = true
			z.s, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "n_zid02_i64":
			found34zgensym_72ba5d454ae3d9dd_35[2] = true
			z.n, bts, err = nbs.ReadInt64Bytes(bts)

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
	if nextMiss34zgensym_72ba5d454ae3d9dd_35 != -1 {
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

// fields of u
var unmarshalMsgFieldOrder34zgensym_72ba5d454ae3d9dd_35 = []string{"m_zid00_map", "s_zid01_str", "n_zid02_i64"}

var unmarshalMsgFieldSkip34zgensym_72ba5d454ae3d9dd_35 = []bool{false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *u) Msgsize() (s int) {
	s = 1 + 12 + msgp.MapHeaderSize
	if z.m != nil {
		for zgensym_72ba5d454ae3d9dd_27, zgensym_72ba5d454ae3d9dd_28 := range z.m {
			_ = zgensym_72ba5d454ae3d9dd_28
			_ = zgensym_72ba5d454ae3d9dd_27
			s += msgp.StringPrefixSize + len(zgensym_72ba5d454ae3d9dd_27)
			if zgensym_72ba5d454ae3d9dd_28 == nil {
				s += msgp.NilSize
			} else {
				s += zgensym_72ba5d454ae3d9dd_28.Msgsize()
			}
		}
	}
	s += 12 + msgp.StringPrefixSize + len(z.s) + 12 + msgp.Int64Size
	return
}
