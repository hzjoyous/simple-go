package CqGoClient

import (
	"encoding/json"
	"errors"
	"strconv"
)

func ToString(_var interface{}) string {
	val, _ := ValToString(_var)
	return val
}
func ToInt(_var interface{}) int {
	val, _ := ValToInt(_var)
	return val
}
func ValToString(_var interface{}) (string, error) {
	switch t := _var.(type) {
	default:
		return "", errors.New("invalid String")
	case bool:
		if t {
			return "true", nil
		} else {
			return "false", nil
		}
	case nil:
		return "", nil
	case int:
		return strconv.FormatInt(int64(t), 10), nil
	case int16:
		return strconv.FormatInt(int64(t), 10), nil
	case int32:
		return strconv.FormatInt(int64(t), 10), nil
	case int64:
		return strconv.FormatInt(int64(t), 10), nil
	case uintptr:
		return strconv.FormatInt(int64(t), 10), nil
	case byte:
		return strconv.FormatInt(int64(t), 10), nil
	case float32:
		return strconv.FormatFloat(float64(t), 'g', 1, 64), nil
	case float64:
		return strconv.FormatFloat(t, 'g', 1, 64), nil
	case uint:
		return strconv.FormatFloat(float64(t), 'g', 1, 64), nil
	case uint16:
		return strconv.FormatFloat(float64(t), 'g', 1, 64), nil
	case uint32:
		return strconv.FormatFloat(float64(t), 'g', 1, 64), nil
	case uint64:
		return strconv.FormatUint(uint64(t), 10), nil
	case json.Number:
		num := t.String()
		return num, nil
	case string:
		return t, nil
	}
}

func ValToInt64(_var interface{}) (int64, error) {
	switch t := _var.(type) {
	default:
		return 0, errors.New("invalid num")
	case bool:
		if t {
			return 1, nil
		} else {
			return 0, nil
		}
	case nil:
		return 0, errors.New("invalid nil")
	case int:
		return int64(t), nil
	case int16:
		return int64(t), nil
	case int32:
		return int64(t), nil
	case int64:
		return t, nil
	case float32:
		return int64(t), nil
	case float64:
		return int64(t), nil
	case uint:
		return int64(t), nil
	case uint16:
		return int64(t), nil
	case uint32:
		return int64(t), nil
	case uint64:
		return int64(t), nil
	case json.Number:
		num, err := t.Int64()
		return num, err
	case string:
		return strconv.ParseInt(t, 10, 64)
	case uintptr:
		return int64(t), nil
	case byte:
		return int64(t), nil
	}
}

func ValToInt(_var interface{}) (int, error) {
	switch t := _var.(type) {
	default:
		return 0, errors.New("invalid num")
	case bool:
		if t {
			return 1, nil
		} else {
			return 0, nil
		}
	case nil:
		return 0, errors.New("invalid nil")
	case int:
		return t, nil
	case int16:
		return int(t), nil
	case int32:
		return int(t), nil
	case int64:
		return int(t), nil
	case float32:
		return int(t), nil
	case float64:
		return int(t), nil
	case uint:
		return int(t), nil
	case uint16:
		return int(t), nil
	case uint32:
		return int(t), nil
	case uint64:
		return int(t), nil
	case json.Number:
		num, err := t.Int64()
		return int(num), err
	case string:
		num, err := strconv.Atoi(t)
		return int(num), err
	case uintptr:
		return int(t), nil
	case byte:
		return int(t), nil
	}
}

func ValToFloat(_var interface{}) (float64, error) {
	switch t := _var.(type) {
	default:
		return float64(0), errors.New("invalid num")
	case bool:
		if t {
			return float64(1), nil
		} else {
			return float64(0), nil
		}
	case nil:
		return float64(0), errors.New("invalid nil")
	case int:
		return float64(t), nil
	case int16:
		return float64(t), nil
	case int32:
		// int32 is rune
		return float64(t), nil
	case int64:
		return float64(t), nil
	case float32:
		return float64(t), nil
	case float64:
		return t, nil
	case uint:
		return float64(t), nil
	case uint16:
		return float64(t), nil
	case uint32:
		return float64(t), nil
	case uint64:
		return float64(t), nil
	case json.Number:
		num, err := t.Int64()
		return float64(num), err
	case string:
		return strconv.ParseFloat(t, 64)
	case uintptr:
		return float64(t), nil
	case byte:
		return float64(t), nil
	}
}

func ValToDouble(_var interface{}) (float64, error) {
	return ValToFloat(_var)
}
