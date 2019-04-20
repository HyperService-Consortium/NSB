package rlp
import(
	"fmt"
	"errors"
	"encoding/hex"
)
const(
	NIL = iota
	BYTES
	GLIST
)
// var hexmaps [256]uint64;
// //input a string and return a byte slice
// func stringtobytes(bytes string) []byte {
// 	glen := len(bytes)
// 	if  glen <= 1 || ((glen & 1) == 1) {
// 		return nil
// 	}

// 	glen >>= 1
// 	ofs := 0

// 	if bytes[1] == 'x' {
// 		ofs = 1
// 	}

// 	glen -= ofs

// 	bres := make([]byte, glen, glen)
// 	for idx := 0; idx < glen; idx++ {
// 		bres[idx] |= byte(hexmaps[bytes[(idx + ofs) << 1 ]] << 4)
// 		bres[idx] |= byte(hexmaps[bytes[(idx + ofs) << 1 | 1]])
// 	}
// 	return bres
// }
// func init() {
// 	for idx := '0'; idx <= '9'; idx++ {
// 		hexmaps[idx] = uint64(idx - '0')
// 	}
// 	for idx := 'a'; idx <= 'f'; idx++ {
// 		hexmaps[idx] = uint64(idx - 'a' + 10)
// 	}
// 	for idx := 'A'; idx <= 'F'; idx++ {
// 		hexmaps[idx] = uint64(idx - 'A' + 10)
// 	}
// }


type Atom interface{}
type Glist struct {
	dat Atom
	typeId int
}
func makeGlistContentGlist() *Glist {
	return &Glist{dat: make([]Atom, 0, 65535), typeId: 2}
}
func makeGlistContentBytes(data []byte) *Glist {
	return &Glist{dat: data, typeId: 1}
}
func (g *Glist) Type() int{
	return g.typeId
}
func _Unserialize(dtr []byte) []Atom {
	atomList := make([]Atom, 0, 65535)
	if len(dtr) == 0 {
		return atomList
	}
	if(dtr[0] < 128){
		return append(append(atomList, makeGlistContentBytes(dtr[0 : 1])), _Unserialize(dtr[1 : ])...)
	}else if dtr[0] < 184 {
		return append(append(atomList, makeGlistContentBytes(dtr[1 : dtr[0] - 127])), _Unserialize(dtr[dtr[0] - 127 : ])...)
	}else if dtr[0] < 192 {
		// dlen = truelen + 1
		dlen := int(dtr[0]) - 182
		tlen := decodelength(dtr[1 : dlen])
		return append(append(atomList, makeGlistContentBytes(dtr[dlen : dlen + tlen])), _Unserialize(dtr[dlen + tlen : ])...)
		//return &Glist{dat:, typeId: 1}
	}else if dtr[0] < 248 {
		// dlen = truelen + 1
		return append(append(atomList, Unserialize(dtr[0 : dtr[0] - 191])), _Unserialize(dtr[dtr[0] - 191 : ])...)
	}else {
		// dlen = truelen + 1
		dlen := int(dtr[0]) - 246
		tlen := decodelength(dtr[1 : dlen])
		// fmt.Println("here",dlen,decodelength(dtr[1 : dlen]))
		return append(append(atomList, Unserialize(dtr[0 : dlen + tlen])), _Unserialize(dtr[dlen + tlen : ])...)
	}
	return atomList
}
func Unserialize(dtr []byte) *Glist {
	if len(dtr) == 0 {
		return makeGlistContentGlist()
	}
	if(dtr[0] < 128){
		if len(dtr) > 1 {
			err := errors.New("superflours")
			fmt.Println(err)
			return nil
		}
		return makeGlistContentBytes(dtr[0 : ])
	}else if dtr[0] < 184 {
		if len(dtr) > int(dtr[0] - 127) {
			err := errors.New("superflours")
			fmt.Println(err)
			return nil
		}
		return makeGlistContentBytes(dtr[1 : ])
	}else if dtr[0] < 192 {
		// dlen = truelen + 1
		//must be uint32, but int for convenience
		dlen := int(dtr[0]) - 182
		tlen := decodelength(dtr[1 : dlen])
		if len(dtr) > tlen + dlen {
			err := errors.New("superflours")
			fmt.Println(err)
			return nil
		}
		//fmt.Println(decodelength(dtr[1 : dlen]))
		return makeGlistContentBytes(dtr[dlen : ])
		//return &Glist{dat:, typeId: 1}
	}else if dtr[0] < 248 {
		// dlen = truelen + 1
		// fmt.Println(dtr[0] - 191)
		if len(dtr) > int(dtr[0] - 191) {
			err := errors.New("superflours")
			fmt.Println(err)
			return nil
		}
		return &Glist{dat: _Unserialize(dtr[1 : ]), typeId: 2}
	}else {
		// dlen = truelen + 1
		//must be uint32, but int for convenience
		dlen := int(dtr[0]) - 246
		tlen := decodelength(dtr[1 : dlen])
		if len(dtr) > tlen + dlen {
			err := errors.New("superflours")
			fmt.Println(err)
			return nil
		}
		return &Glist{dat: _Unserialize(dtr[dlen : ]), typeId: 2}
	}
}
func decodelength(dtr []byte) int {
	var len = 0
	for _, t := range dtr {
		len = (len<<8) | int(t)
	}
	return len
}
func PrintList(g *Glist){
	if g == nil {
		fmt.Print("[]")
		return ;
	}
	switch g.typeId {
		case 1:{
			fmt.Print(g.dat.([]byte))
			break ;
		}
		case 2:{
			fmt.Print("[")
			for i, v := range(g.dat.([]Atom)) {
				if i != 0 {
					fmt.Print(",")
				}
				PrintList(v.(*Glist))
			}
			fmt.Print("]")
			break;
		}
		default :{
			break ;
		}
	}
}
func PrintListInString(g *Glist){
	if g == nil {
		fmt.Print("~")
		return ;
	}
	switch g.typeId {
		case 1:{
			fmt.Print("\"0x"+hex.EncodeToString(g.dat.([]byte))+"\"")
			break ;
		}
		case 2:{
			fmt.Print("[")
			for i, v := range(g.dat.([]Atom)) {
				if i != 0 {
					fmt.Print(",")
				}
				PrintListInString(v.(*Glist))
			}
			fmt.Print("]")
			break;
		}
		default :{
			break ;
		}
	}
}
func (g *Glist) Get(ref int) *Glist {
	return g.dat.([]Atom)[ref].(*Glist)
}
func (g *Glist) AsBytes() []byte {
	return g.dat.([]byte)
}
func (g *Glist) AsString() string {
	return hex.EncodeToString(g.dat.([]byte))
}
func (g *Glist) Length() int {
	if g.typeId == 1 {
		return len(g.dat.([]byte))
	}else if g.typeId == 2 {
		return len(g.dat.([]Atom))
	}else {
		return -1
	}
}
func main(){
	// x := Unserialize(stringtobytes("f90191a0891757929ce116380174486268097769a2f30291732dd18582327788cd95250580a07c2a362d8127b83cbadc4363ff461ee6d72cbefc9a03e6080b739ca340d2860b80a08dad12008136bbe56bc544d8c6711b5c5a594e0f7bfd9a1851355f1fc199e131a0be512df0fee45c716a3ab87e2dffb2072ebaf0dc7f69784e6e7c7b0fae92841ca046b0b0ee9c8e7f0c37a5b8763280abf0c7e47caeb2ab4ff93f0c67fe26c8ad2ba0815bd660bd5e85681a9adab43e0a651e807ba5132f94bd210e9457ba055d538280a0663e186f98fff71d728118b5f788fa72506a4bd6e85aa2915b8ae40059f4ab86a092ef150745b4207ac88519318bd7bb007fe6157e719d459edf5b59f025fe16fda0236e8f61ecde6abfebc6c529441f782f62469d8a2cc47b7aace2c136bd3b1ff0a0d1a0d264eaf005589a3f1486b1dc51909c782272c969c6fcf362b122baf6d188a0303876dcca400618ab14b9d5ff6416ebfde7cd05bd17fd78cab31d79c2759187a0b3a005a4dd06158b44eea4bbfbcdadd5fab5f15f3fe59fe87a9da80ec08552908080"))
	// PrintListInString(x)
	// fmt.Println("")
	// x = Unserialize(append(make([]byte,0), 0x0))
	// PrintListInString(x)
	// fmt.Println("")
	// x = Unserialize(append(make([]byte,0), 0x83, 'd', 'o', 'g'))
	// PrintListInString(x)
	// fmt.Println("")
	// x = Unserialize(append(make([]byte,0), 0xc0))
	// PrintListInString(x)
	// fmt.Println("")
	// x = Unserialize(append(make([]byte,0), 0xc8, 0x83, 0x55, 0x83, 0x83, 0x83, 0x83, 0x83, 0x83))
	// PrintListInString(x)
	// fmt.Println("")
	// x = Unserialize(append(make([]byte,0), 0xc7, 0xc0, 0xc1, 0xc0, 0xc3, 0xc0, 0xc1, 0xc0))
	// PrintListInString(x)
}
/*
[[],[[]],[[],[[]]]]
[[],[[]],[[],[[]]]]
 */