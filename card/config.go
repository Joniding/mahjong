package card

// 麻将牌面对应的数值定义
const (
	MAHJONG_PLACEHOLDER = iota // 空，占位，无意义
	MAHJONG_CRAK1              // 万(1 ~ 9)
	MAHJONG_CRAK2
	MAHJONG_CRAK3
	MAHJONG_CRAK4
	MAHJONG_CRAK5
	MAHJONG_CRAK6
	MAHJONG_CRAK7
	MAHJONG_CRAK8
	MAHJONG_CRAK9
	MAHJONG_CRAK_PLACEHOLDER // 10, 占位，无意义
	MAHJONG_BAM1             // 条(11 ~ 19)
	MAHJONG_BAM2
	MAHJONG_BAM3
	MAHJONG_BAM4
	MAHJONG_BAM5
	MAHJONG_BAM6
	MAHJONG_BAM7
	MAHJONG_BAM8
	MAHJONG_BAM9
	MAHJONG_BAM_PLACE_HOLDER // 20, 占位，无意义
	MAHJONG_DOT1             // 筒(21 ~ 29)
	MAHJONG_DOT2
	MAHJONG_DOT3
	MAHJONG_DOT4
	MAHJONG_DOT5
	MAHJONG_DOT6
	MAHJONG_DOT7
	MAHJONG_DOT8
	MAHJONG_DOT9
	MAHJONG_DOT_PLACEHOLDER // 30, 占位，无意义

	MAHJONG_EAST  = 31 // 东南北西 (31 ~ 34)
	MAHJONG_NORTH = 32
	MAHJONG_SOUTH = 33
	MAHJONG_WEST  = 34

	MAHJONG_GREE  = 41 // 发财、红中、白板 (35 ~ 37)
	MAHJONG_RED   = 42
	MAHJONG_WHITE = 43

	MAHJONG_SEASON1 = 51 // 春夏秋冬（41 ~ 44）
	MAHJONG_SEASON2 = 52
	MAHJONG_SEASON3 = 53
	MAHJONG_SEASON4 = 54

	MAHJONG_FLOWER1 = 61 // 花色 (45 ~ 48)
	MAHJONG_FLOWER2 = 62
	MAHJONG_FLOWER3 = 63
	MAHJONG_FLOWER4 = 64
)
