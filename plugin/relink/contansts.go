package relink

import "errors"

const (
	PcMoeApi = "http://hi.pcmoe.net/bear.php" // 萌研社API
)

type PcMoeMode string

const (
	PcMoeModeBear      PcMoeMode = "熊曰"
	PcMoeModeBuddha    PcMoeMode = "佛曰"
	PcMoeModeNewBuddha PcMoeMode = "新佛曰"
)

func (p PcMoeMode) value() (string, error) {
	switch p {
	case PcMoeModeBear:
		return "Bear", nil
	case PcMoeModeBuddha, PcMoeModeNewBuddha:
		return "Buddha", nil
	default:
		return "", errors.New("unknown mode")
	}
}

type PcMoeCode string

const (
	PcMoeCodeDecode PcMoeCode = "Decode"
	PcMoeCodeEncode PcMoeCode = "Encode"
)
