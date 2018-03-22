// Code generated by protoc-gen-go. DO NOT EDIT.
// source: atari.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	atari.proto

It has these top-level messages:
	Environment
	Observation
	Action
	State
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

// GameType denotes whether the environment should be from pixels or RAM.
type GameType int32

const (
	GameType_IMAGE GameType = 0
	GameType_RAM   GameType = 1
)

var GameType_name = map[int32]string{
	0: "IMAGE",
	1: "RAM",
}
var GameType_value = map[string]int32{
	"IMAGE": 0,
	"RAM":   1,
}

func (x GameType) String() string {
	return proto1.EnumName(GameType_name, int32(x))
}
func (GameType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Game is one of the available game environments. This enum is exhaustive but
// it serves the purpose of documenting the available environments as simple as
// possible.
type Game int32

const (
	Game_AIR_RAID          Game = 0
	Game_ALIEN             Game = 1
	Game_AMIDAR            Game = 3
	Game_ASSAULT           Game = 4
	Game_ASTERIX           Game = 5
	Game_ASTEROIDS         Game = 6
	Game_ATLANTIS          Game = 7
	Game_BANK_HEIST        Game = 8
	Game_BATTLE_ZONE       Game = 9
	Game_BEAM_RIDER        Game = 10
	Game_BERZERK           Game = 11
	Game_BOWLING           Game = 12
	Game_BOXING            Game = 13
	Game_BREAKOUT          Game = 14
	Game_CARNIVAL          Game = 15
	Game_CENTIPEDE         Game = 16
	Game_CHOPPER_COMMAND   Game = 17
	Game_CRAZY_CLIMBER     Game = 18
	Game_DEMON_ATTACK      Game = 19
	Game_DOUBLE_DUNK       Game = 20
	Game_ELEVATOR_ACTION   Game = 21
	Game_ENDURO            Game = 22
	Game_FISHING_DERBY     Game = 23
	Game_FREEWAY           Game = 24
	Game_FROSTBITE         Game = 25
	Game_GOPHER            Game = 26
	Game_GRAVITAR          Game = 27
	Game_ICE_HOCKEY        Game = 28
	Game_JAMES_BOND        Game = 29
	Game_JOURNEY_ESCAPE    Game = 30
	Game_KANGAROO          Game = 31
	Game_KRULL             Game = 32
	Game_KUNG_FU_MASTER    Game = 33
	Game_MONTEZUMA_REVENGE Game = 34
	Game_MS_PACMAN         Game = 35
	Game_NAME_THIS_GAME    Game = 36
	Game_PHOENIX           Game = 37
	Game_PITFALL           Game = 38
	Game_PONG              Game = 39
	Game_POOYAN            Game = 40
	Game_PRIVATE_EYE       Game = 41
	Game_QBERT             Game = 42
	Game_RIVER_RAID        Game = 43
	Game_ROAD_RUNNER       Game = 44
	Game_ROBOTANK          Game = 45
	Game_SEAQUEST          Game = 46
	Game_SKIING            Game = 47
	Game_SPACE_INVADERS    Game = 48
	Game_STAR_GUNNER       Game = 49
	Game_TENNIS            Game = 50
	Game_TIME_PILOT        Game = 51
	Game_TUTANKHAM         Game = 52
	Game_UP_N_DOWN         Game = 53
	Game_VENTURE           Game = 54
	Game_VIDEO_PINBALL     Game = 55
	Game_WIZARD_OF_WOR     Game = 56
	Game_YARS_REVENGE      Game = 57
	Game_ZAXXON            Game = 58
)

var Game_name = map[int32]string{
	0:  "AIR_RAID",
	1:  "ALIEN",
	3:  "AMIDAR",
	4:  "ASSAULT",
	5:  "ASTERIX",
	6:  "ASTEROIDS",
	7:  "ATLANTIS",
	8:  "BANK_HEIST",
	9:  "BATTLE_ZONE",
	10: "BEAM_RIDER",
	11: "BERZERK",
	12: "BOWLING",
	13: "BOXING",
	14: "BREAKOUT",
	15: "CARNIVAL",
	16: "CENTIPEDE",
	17: "CHOPPER_COMMAND",
	18: "CRAZY_CLIMBER",
	19: "DEMON_ATTACK",
	20: "DOUBLE_DUNK",
	21: "ELEVATOR_ACTION",
	22: "ENDURO",
	23: "FISHING_DERBY",
	24: "FREEWAY",
	25: "FROSTBITE",
	26: "GOPHER",
	27: "GRAVITAR",
	28: "ICE_HOCKEY",
	29: "JAMES_BOND",
	30: "JOURNEY_ESCAPE",
	31: "KANGAROO",
	32: "KRULL",
	33: "KUNG_FU_MASTER",
	34: "MONTEZUMA_REVENGE",
	35: "MS_PACMAN",
	36: "NAME_THIS_GAME",
	37: "PHOENIX",
	38: "PITFALL",
	39: "PONG",
	40: "POOYAN",
	41: "PRIVATE_EYE",
	42: "QBERT",
	43: "RIVER_RAID",
	44: "ROAD_RUNNER",
	45: "ROBOTANK",
	46: "SEAQUEST",
	47: "SKIING",
	48: "SPACE_INVADERS",
	49: "STAR_GUNNER",
	50: "TENNIS",
	51: "TIME_PILOT",
	52: "TUTANKHAM",
	53: "UP_N_DOWN",
	54: "VENTURE",
	55: "VIDEO_PINBALL",
	56: "WIZARD_OF_WOR",
	57: "YARS_REVENGE",
	58: "ZAXXON",
}
var Game_value = map[string]int32{
	"AIR_RAID":          0,
	"ALIEN":             1,
	"AMIDAR":            3,
	"ASSAULT":           4,
	"ASTERIX":           5,
	"ASTEROIDS":         6,
	"ATLANTIS":          7,
	"BANK_HEIST":        8,
	"BATTLE_ZONE":       9,
	"BEAM_RIDER":        10,
	"BERZERK":           11,
	"BOWLING":           12,
	"BOXING":            13,
	"BREAKOUT":          14,
	"CARNIVAL":          15,
	"CENTIPEDE":         16,
	"CHOPPER_COMMAND":   17,
	"CRAZY_CLIMBER":     18,
	"DEMON_ATTACK":      19,
	"DOUBLE_DUNK":       20,
	"ELEVATOR_ACTION":   21,
	"ENDURO":            22,
	"FISHING_DERBY":     23,
	"FREEWAY":           24,
	"FROSTBITE":         25,
	"GOPHER":            26,
	"GRAVITAR":          27,
	"ICE_HOCKEY":        28,
	"JAMES_BOND":        29,
	"JOURNEY_ESCAPE":    30,
	"KANGAROO":          31,
	"KRULL":             32,
	"KUNG_FU_MASTER":    33,
	"MONTEZUMA_REVENGE": 34,
	"MS_PACMAN":         35,
	"NAME_THIS_GAME":    36,
	"PHOENIX":           37,
	"PITFALL":           38,
	"PONG":              39,
	"POOYAN":            40,
	"PRIVATE_EYE":       41,
	"QBERT":             42,
	"RIVER_RAID":        43,
	"ROAD_RUNNER":       44,
	"ROBOTANK":          45,
	"SEAQUEST":          46,
	"SKIING":            47,
	"SPACE_INVADERS":    48,
	"STAR_GUNNER":       49,
	"TENNIS":            50,
	"TIME_PILOT":        51,
	"TUTANKHAM":         52,
	"UP_N_DOWN":         53,
	"VENTURE":           54,
	"VIDEO_PINBALL":     55,
	"WIZARD_OF_WOR":     56,
	"YARS_REVENGE":      57,
	"ZAXXON":            58,
}

func (x Game) String() string {
	return proto1.EnumName(Game_name, int32(x))
}
func (Game) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// Environment describes an active / to be activated environment.
type Environment struct {
	Game     Game     `protobuf:"varint,1,opt,name=game,enum=atari.Game" json:"game,omitempty"`
	GameType GameType `protobuf:"varint,2,opt,name=game_type,json=gameType,enum=atari.GameType" json:"game_type,omitempty"`
	Version  uint32   `protobuf:"varint,3,opt,name=version" json:"version,omitempty"`
}

func (m *Environment) Reset()                    { *m = Environment{} }
func (m *Environment) String() string            { return proto1.CompactTextString(m) }
func (*Environment) ProtoMessage()               {}
func (*Environment) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Environment) GetGame() Game {
	if m != nil {
		return m.Game
	}
	return Game_AIR_RAID
}

func (m *Environment) GetGameType() GameType {
	if m != nil {
		return m.GameType
	}
	return GameType_IMAGE
}

func (m *Environment) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

// Observation is an 8bit RGB image or an array of 8 bit values in the atari RAM.
// Images have shape (250, 160, 3), and RAM values shape (128,).
type Observation struct {
	// The dimensions of the observation array.
	Shape []int32 `protobuf:"varint,1,rep,packed,name=shape" json:"shape,omitempty"`
	// The values of the observation array. These are actually 8bit, in both the
	// memory and image cases. Protobuf doesn't have a uint8 type, but does use
	// variable length encoding.
	Values []uint32 `protobuf:"varint,2,rep,packed,name=values" json:"values,omitempty"`
}

func (m *Observation) Reset()                    { *m = Observation{} }
func (m *Observation) String() string            { return proto1.CompactTextString(m) }
func (*Observation) ProtoMessage()               {}
func (*Observation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Observation) GetShape() []int32 {
	if m != nil {
		return m.Shape
	}
	return nil
}

func (m *Observation) GetValues() []uint32 {
	if m != nil {
		return m.Values
	}
	return nil
}

// Action is a discrete action.
type Action struct {
	// The index of the action. Atari environments have discrete action spaces
	// with from 3 to 16 possible actions.
	Value uint32 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *Action) Reset()                    { *m = Action{} }
func (m *Action) String() string            { return proto1.CompactTextString(m) }
func (*Action) ProtoMessage()               {}
func (*Action) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Action) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

// State is the full state of the system.
type State struct {
	// The current observation.
	Observation *Observation `protobuf:"bytes,1,opt,name=observation" json:"observation,omitempty"`
	// The scalar reward received this period.
	Reward float32 `protobuf:"fixed32,2,opt,name=reward" json:"reward,omitempty"`
	// Whether the episode is done.
	Done bool `protobuf:"varint,3,opt,name=done" json:"done,omitempty"`
	// Information for debugging purposes.
	Info string `protobuf:"bytes,4,opt,name=info" json:"info,omitempty"`
}

func (m *State) Reset()                    { *m = State{} }
func (m *State) String() string            { return proto1.CompactTextString(m) }
func (*State) ProtoMessage()               {}
func (*State) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *State) GetObservation() *Observation {
	if m != nil {
		return m.Observation
	}
	return nil
}

func (m *State) GetReward() float32 {
	if m != nil {
		return m.Reward
	}
	return 0
}

func (m *State) GetDone() bool {
	if m != nil {
		return m.Done
	}
	return false
}

func (m *State) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func init() {
	proto1.RegisterType((*Environment)(nil), "atari.Environment")
	proto1.RegisterType((*Observation)(nil), "atari.Observation")
	proto1.RegisterType((*Action)(nil), "atari.Action")
	proto1.RegisterType((*State)(nil), "atari.State")
	proto1.RegisterEnum("atari.GameType", GameType_name, GameType_value)
	proto1.RegisterEnum("atari.Game", Game_name, Game_value)
}

func init() { proto1.RegisterFile("atari.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 928 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x54, 0x5d, 0x77, 0x1b, 0x35,
	0x10, 0x4d, 0x1a, 0x3b, 0x1f, 0x72, 0x9c, 0x4c, 0xd5, 0x0f, 0x4c, 0x81, 0x36, 0x18, 0x0a, 0xa1,
	0x94, 0x42, 0xd3, 0xf2, 0xfd, 0x34, 0xbb, 0x3b, 0xb1, 0x55, 0xef, 0x4a, 0xdb, 0x59, 0xad, 0x13,
	0xfb, 0x45, 0xc7, 0xa5, 0x4b, 0xc9, 0x39, 0xd4, 0xce, 0x71, 0x4c, 0x38, 0x7d, 0xe0, 0x87, 0xf3,
	0xc6, 0x99, 0xb5, 0x03, 0x79, 0xe0, 0xc9, 0xba, 0xda, 0x3b, 0x73, 0xef, 0x95, 0xe4, 0x51, 0xad,
	0xc9, 0x62, 0x32, 0x3f, 0x7b, 0x72, 0x3e, 0x9f, 0x2d, 0x66, 0xba, 0x59, 0x83, 0xee, 0xa5, 0x6a,
	0xd1, 0xf4, 0xf2, 0x6c, 0x3e, 0x9b, 0xbe, 0xad, 0xa6, 0x0b, 0xfd, 0x40, 0x35, 0xde, 0x4c, 0xde,
	0x56, 0x9d, 0xf5, 0x83, 0xf5, 0xc3, 0xbd, 0xa3, 0xd6, 0x93, 0x65, 0x45, 0x6f, 0xf2, 0xb6, 0xe2,
	0xfa, 0x83, 0x7e, 0xac, 0x76, 0xe4, 0x37, 0x2c, 0xde, 0x9d, 0x57, 0x9d, 0x1b, 0x35, 0x6b, 0xff,
	0x1a, 0xcb, 0xbf, 0x3b, 0xaf, 0x78, 0xfb, 0xcd, 0x6a, 0xa5, 0x3b, 0x6a, 0xeb, 0xb2, 0x9a, 0x5f,
	0x9c, 0xcd, 0xa6, 0x9d, 0x8d, 0x83, 0xf5, 0xc3, 0x36, 0x5f, 0xc1, 0xee, 0xcf, 0xaa, 0xe5, 0x5e,
	0x5d, 0x54, 0xf3, 0xcb, 0xc9, 0xe2, 0x6c, 0x36, 0xd5, 0xb7, 0x55, 0xf3, 0xe2, 0xb7, 0xc9, 0xb9,
	0x08, 0x6f, 0x1c, 0x36, 0x79, 0x09, 0xf4, 0x5d, 0xb5, 0x79, 0x39, 0xf9, 0xfd, 0x8f, 0xea, 0xa2,
	0x73, 0xe3, 0x60, 0xe3, 0xb0, 0xcd, 0x2b, 0xd4, 0xbd, 0xaf, 0x36, 0xf1, 0x97, 0xab, 0xba, 0x7a,
	0xaf, 0x36, 0xdc, 0xe6, 0x25, 0xe8, 0xfe, 0xa5, 0x9a, 0xc5, 0x62, 0xb2, 0xa8, 0xf4, 0x73, 0xd5,
	0x9a, 0xfd, 0xa7, 0x52, 0x93, 0x5a, 0x47, 0x7a, 0xe5, 0xf7, 0x9a, 0x3e, 0x5f, 0xa7, 0x89, 0xec,
	0xbc, 0xfa, 0x73, 0x32, 0x7f, 0x5d, 0x07, 0xbc, 0xc1, 0x2b, 0xa4, 0xb5, 0x6a, 0xbc, 0x9e, 0x4d,
	0xab, 0x3a, 0xca, 0x36, 0xd7, 0x6b, 0xd9, 0x3b, 0x9b, 0xfe, 0x3a, 0xeb, 0x34, 0x0e, 0xd6, 0x0f,
	0x77, 0xb8, 0x5e, 0x3f, 0xba, 0xaf, 0xb6, 0xaf, 0xce, 0x42, 0xef, 0xa8, 0xa6, 0xc9, 0xb0, 0x47,
	0xb0, 0xa6, 0xb7, 0xd4, 0x06, 0x63, 0x06, 0xeb, 0x8f, 0xfe, 0xde, 0x54, 0x0d, 0x21, 0xe8, 0x5d,
	0xb5, 0x8d, 0x86, 0x03, 0xa3, 0x49, 0x60, 0x4d, 0xa8, 0x98, 0x1a, 0xb2, 0xb0, 0xae, 0x95, 0xda,
	0xc4, 0xcc, 0x24, 0xc8, 0xb0, 0xa1, 0x5b, 0x6a, 0x0b, 0x8b, 0x02, 0xcb, 0xd4, 0x43, 0x63, 0x09,
	0x3c, 0xb1, 0x39, 0x85, 0xa6, 0x6e, 0xab, 0x9d, 0x1a, 0x38, 0x93, 0x14, 0xb0, 0x59, 0x77, 0xf3,
	0x29, 0x5a, 0x6f, 0x0a, 0xd8, 0xd2, 0x7b, 0x4a, 0x45, 0x68, 0x07, 0xa1, 0x4f, 0xa6, 0xf0, 0xb0,
	0xad, 0xf7, 0x55, 0x2b, 0x42, 0xef, 0x53, 0x0a, 0x63, 0x67, 0x09, 0x76, 0x6a, 0x02, 0x61, 0x16,
	0xd8, 0x24, 0xc4, 0xa0, 0xa4, 0x75, 0x44, 0x3c, 0x26, 0x1e, 0x40, 0xab, 0x06, 0xee, 0x24, 0x35,
	0xb6, 0x07, 0xbb, 0xe2, 0x26, 0x72, 0xa7, 0xb2, 0x6e, 0x8b, 0x48, 0xc4, 0x84, 0x03, 0x57, 0x7a,
	0xd8, 0x13, 0x14, 0x23, 0x5b, 0x33, 0xc4, 0x14, 0xf6, 0xc5, 0x4f, 0x4c, 0xd6, 0x9b, 0x9c, 0x12,
	0x02, 0xd0, 0xb7, 0xd4, 0x7e, 0xdc, 0x77, 0x79, 0x4e, 0x1c, 0x62, 0x97, 0x65, 0x68, 0x13, 0xb8,
	0xa9, 0x6f, 0xaa, 0x76, 0xcc, 0x38, 0x1e, 0x85, 0x38, 0x35, 0x59, 0x44, 0x0c, 0x5a, 0x83, 0xda,
	0x4d, 0x28, 0x73, 0x36, 0xa0, 0xf7, 0x18, 0x0f, 0xe0, 0x96, 0x78, 0x4d, 0x5c, 0x19, 0xa5, 0x14,
	0x92, 0xd2, 0x0e, 0xe0, 0xb6, 0xb4, 0xa2, 0x94, 0x86, 0xe8, 0x1d, 0x07, 0x8c, 0xbd, 0x71, 0x16,
	0xee, 0x88, 0x2d, 0xb2, 0x49, 0xc9, 0x0e, 0xee, 0x4a, 0xdb, 0x63, 0x53, 0xf4, 0x8d, 0xed, 0x85,
	0x84, 0x38, 0x1a, 0xc1, 0x7b, 0x12, 0xe1, 0x98, 0x89, 0x4e, 0x70, 0x04, 0x1d, 0xb1, 0x76, 0xcc,
	0xae, 0xf0, 0x91, 0xf1, 0x04, 0xef, 0x4b, 0x69, 0xcf, 0xe5, 0x7d, 0x62, 0xb8, 0x27, 0x19, 0x7a,
	0x8c, 0x43, 0xe3, 0x91, 0xe1, 0x03, 0x39, 0x15, 0x13, 0x53, 0xe8, 0xbb, 0x78, 0x40, 0x23, 0xf8,
	0x50, 0xf0, 0x0b, 0xcc, 0xa8, 0x08, 0x91, 0xb3, 0x09, 0x7c, 0xa4, 0xb5, 0xda, 0x7b, 0xe1, 0x4a,
	0xb6, 0x34, 0x0a, 0x54, 0xc4, 0x98, 0x13, 0xdc, 0x97, 0x0e, 0x03, 0xb4, 0x3d, 0x64, 0xe7, 0xe0,
	0x81, 0x5c, 0xe3, 0x80, 0xcb, 0x34, 0x85, 0x03, 0x21, 0x0f, 0x4a, 0xdb, 0x0b, 0xc7, 0x65, 0xc8,
	0xea, 0x9b, 0x82, 0x8f, 0xf5, 0x1d, 0x75, 0x33, 0x73, 0xd6, 0xd3, 0xb8, 0xcc, 0x30, 0x30, 0x0d,
	0xc9, 0xf6, 0x08, 0xba, 0x62, 0x30, 0x2b, 0x42, 0x8e, 0x71, 0x86, 0x16, 0x3e, 0x91, 0x4a, 0x8b,
	0x19, 0x05, 0xdf, 0x37, 0x45, 0xe8, 0x61, 0x46, 0xf0, 0xa9, 0x04, 0xca, 0xfb, 0x8e, 0xac, 0x39,
	0x85, 0x87, 0x35, 0x30, 0xfe, 0x18, 0xd3, 0x14, 0x3e, 0xd3, 0xdb, 0xaa, 0x91, 0x3b, 0xdb, 0x83,
	0xcf, 0x25, 0x58, 0xee, 0xdc, 0x08, 0x2d, 0x1c, 0xca, 0x29, 0xe6, 0x6c, 0x86, 0xe8, 0x29, 0xd0,
	0x88, 0xe0, 0x0b, 0x71, 0xf6, 0x32, 0x22, 0xf6, 0xf0, 0x48, 0x62, 0xb1, 0x19, 0xd2, 0xea, 0xed,
	0x7d, 0x29, 0x5c, 0x76, 0x98, 0x04, 0x2e, 0xad, 0x25, 0x86, 0xc7, 0x92, 0x89, 0x5d, 0xe4, 0x3c,
	0xda, 0x01, 0x7c, 0x25, 0xa8, 0x20, 0x7c, 0x59, 0x52, 0xe1, 0xe1, 0x89, 0x88, 0x14, 0x03, 0x23,
	0xef, 0xe1, 0x6b, 0x31, 0x5a, 0xe4, 0x18, 0x53, 0x30, 0x76, 0x88, 0x09, 0x71, 0x01, 0xdf, 0x48,
	0xb3, 0xc2, 0x23, 0x87, 0xde, 0xb2, 0xd9, 0x53, 0x29, 0xf0, 0x64, 0xad, 0x29, 0xe0, 0x48, 0x94,
	0xbd, 0xc9, 0x28, 0xe4, 0x26, 0x75, 0x1e, 0x9e, 0x49, 0x70, 0x5f, 0x8a, 0x4c, 0x1f, 0x33, 0x78,
	0x2e, 0xb0, 0xcc, 0x83, 0x0d, 0x89, 0x3b, 0xb1, 0xf0, 0xad, 0xc4, 0x1c, 0x92, 0xf5, 0x25, 0x13,
	0x7c, 0x27, 0x97, 0x3c, 0x34, 0x09, 0xb9, 0x90, 0x1b, 0x1b, 0x49, 0xf2, 0xef, 0x65, 0xeb, 0xc4,
	0x8c, 0x91, 0x93, 0xe0, 0x8e, 0xc3, 0x89, 0x63, 0xf8, 0x41, 0x9e, 0xd3, 0x08, 0xb9, 0xf8, 0xf7,
	0x6c, 0x7f, 0x14, 0xf9, 0x31, 0x9e, 0x9e, 0x3a, 0x0b, 0x3f, 0x1d, 0x4d, 0x54, 0x13, 0xe5, 0xdf,
	0xaf, 0x9f, 0xd6, 0x33, 0x62, 0xbe, 0xd0, 0x57, 0xe3, 0xe0, 0xda, 0x18, 0xbc, 0xf7, 0x3f, 0x23,
	0xa2, 0xbb, 0xa6, 0x1f, 0xaa, 0x46, 0xb1, 0xa8, 0xce, 0x75, 0x7b, 0xf5, 0x75, 0x39, 0x83, 0xee,
	0xed, 0xae, 0x60, 0x3d, 0x72, 0xba, 0x6b, 0xd1, 0xd6, 0xb8, 0x59, 0x8f, 0xd8, 0x57, 0x9b, 0xf5,
	0xcf, 0xb3, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x18, 0x43, 0xb5, 0x7e, 0x78, 0x05, 0x00, 0x00,
}
