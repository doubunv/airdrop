package systemType

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
)

const AmountUnits int64 = 10000000000

type Amount struct {
	UnitsAmount int64
	ShowAmount  float64 // 保留10位小数
}

// 2. 为 Amount 重写 MarshaJSON 方法，在此方法中实现自定义格式的转换；
func (t Amount) MarshalJSON() ([]byte, error) {
	str := strconv.FormatFloat(t.ShowAmount, 'f', 10, 64)
	// 移除末尾的零
	str = strings.TrimRight(str, "0")

	// 如果字符串最后是小数点也移除
	str = strings.TrimRight(str, ".")

	output := fmt.Sprintf("%v", str)
	return []byte(output), nil
}

func (t *Amount) UnmarshalJSON(data []byte) (err error) {
	// 空值不进行解析
	if len(data) == 2 {
		return
	}
	if string(data) == "null" {
		return
	}
	deci, err := decimal.NewFromString(string(data))
	if err != nil {
		return
	}

	needDeci := deci.Mul(decimal.NewFromInt(AmountUnits)).Floor()
	show, _ := needDeci.Div(decimal.NewFromInt(AmountUnits)).Float64()
	*t = Amount{
		UnitsAmount: needDeci.IntPart(),
		ShowAmount:  show,
	}

	return
}

// 3. 为 Amount 实现 Value 方法，写入数据库时会调用该方法将自定义时间类型转换并写入数据库；
func (t Amount) Value() (driver.Value, error) {
	return t.UnitsAmount, nil
}

// 4. 为 Amount 实现 Scan 方法，读取数据库时会调用该方法将时间数据转换成自定义时间类型；
func (t *Amount) Scan(v interface{}) error {
	switch v.(type) {
	case int64:
		show, _ := decimal.NewFromInt(v.(int64)).Div(decimal.NewFromInt(AmountUnits)).Float64()
		*t = Amount{
			UnitsAmount: v.(int64),
			ShowAmount:  show,
		}
		return nil
	case uint8:
		show, _ := decimal.NewFromInt(int64(v.(uint8))).Div(decimal.NewFromInt(AmountUnits)).Float64()
		*t = Amount{
			UnitsAmount: int64(v.(uint8)),
			ShowAmount:  show,
		}
		return nil
	case []byte:
		if len(v.([]byte)) == 0 {
			*t = Amount{
				UnitsAmount: 0,
				ShowAmount:  0,
			}
			return nil
		}

		if len(v.([]byte)) == 1 && (v.([]byte))[0] == 48 {
			*t = Amount{
				UnitsAmount: 0,
				ShowAmount:  0,
			}
			return nil
		}

		int64V, err := strconv.ParseInt(string(v.([]byte)), 10, 64)
		if err == nil {
			show, _ := decimal.NewFromInt(int64V).Div(decimal.NewFromInt(AmountUnits)).Float64()
			*t = Amount{
				UnitsAmount: int64V,
				ShowAmount:  show,
			}
			return nil
		}
	}

	return fmt.Errorf("can not convert %v to amount", v)
}

func (t Amount) String() string {
	str := strconv.FormatFloat(t.ShowAmount, 'f', 10, 64)
	// 移除末尾的零
	str = strings.TrimRight(str, "0")

	// 如果字符串最后是小数点也移除
	str = strings.TrimRight(str, ".")

	return fmt.Sprintf("%s", str)
}

func (t Amount) GetInt64() int64 {
	return t.UnitsAmount
}

func (t Amount) GetFloat64() float64 {
	return t.ShowAmount
}

func NewAmountInt64(value int64) Amount {
	show, _ := decimal.NewFromInt(value).Div(decimal.NewFromInt(AmountUnits)).Float64()

	return Amount{
		UnitsAmount: value,
		ShowAmount:  show,
	}
}

func NewAmountFloat64(value float64) Amount {
	floor10 := decimal.NewFromFloat(value).Mul(decimal.NewFromInt(AmountUnits)).Floor()
	ok, _ := floor10.Div(decimal.NewFromInt(AmountUnits)).Float64()
	return Amount{
		UnitsAmount: floor10.IntPart(),
		ShowAmount:  ok,
	}
}

func (t Amount) GetShowFloat() float64 {
	return decimal.NewFromFloat(t.ShowAmount).RoundDown(5).InexactFloat64()
}

func AmountIcon2(amount Amount, iconRate Amount) Amount {
	nowAmount := decimal.NewFromInt(amount.GetInt64()).
		Mul(decimal.NewFromFloat(iconRate.GetFloat64())).
		Floor().
		IntPart()
	return NewAmountInt64(nowAmount)
}
