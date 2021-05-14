package types

import (
	"fmt"

	"github.com/awatercolorpen/olap-sql/api/proto"
)

type MetricType string

func (m MetricType) ToEnum() proto.METRIC_TYPE {
	if v, ok := proto.METRIC_TYPE_value[string(m)]; ok {
		return proto.METRIC_TYPE(v)
	}
	return proto.METRIC_TYPE_METRIC_TYPE_UNKNOWN
}

func EnumToMetricType(m proto.METRIC_TYPE) MetricType {
	return MetricType(m.String())
}

const (
	MetricTypeUnknown       MetricType = "METRIC_TYPE_UNKNOWN"        // invalid type.
	MetricTypeValue         MetricType = "METRIC_TYPE_VALUE"          // single type. eg: 原始值指标
	MetricTypeCount         MetricType = "METRIC_TYPE_COUNT"          // single type. eg: 计数指标
	MetricTypeDistinctCount MetricType = "METRIC_TYPE_DISTINCT_COUNT" // single type. eg: 去重计数指标
	MetricTypeSum           MetricType = "METRIC_TYPE_SUM"            // single type. eg: 求和指标
	MetricTypeAdd           MetricType = "METRIC_TYPE_ADD"            // composition type eg: 相加指标
	MetricTypeSubtract      MetricType = "METRIC_TYPE_SUBTRACT"       // composition type eg: 相乘指标
	MetricTypeMultiply      MetricType = "METRIC_TYPE_MULTIPLY"       // composition type eg: 相减指标
	MetricTypeDivide        MetricType = "METRIC_TYPE_DIVIDE"         // composition type.eg: 相除指标
	MetricTypeExpression    MetricType = "METRIC_TYPE_EXPRESSION"     // single type. eg: 表达式
	MetricTypePost          MetricType = "METRIC_TYPE_POST"           // composition type, unsupported now
)

type Metric struct {
	Type           MetricType  `json:"type"`
	Table          string      `json:"table"`
	Name           string      `json:"name"`
	FieldName      string      `json:"field_name"`
	ExtensionValue interface{} `json:"extension_value"`
	Children       []*Metric   `json:"children"`
}

func (m *Metric) Statement() (string, error) {
	col, err := m.column()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v AS %v", col.Sql(), col.GetAlias()), nil
}

func (m *Metric) columns() ([]Column, error) {
	var column []Column
	for _, v := range m.Children {
		col, err := v.column()
		if err != nil {
			return nil, err
		}
		column = append(column, col)
	}
	return column, nil
}

func (m *Metric) column() (Column, error) {
	column, err := m.columns()
	if err != nil {
		return nil, err
	}
	switch m.Type {
	case MetricTypeValue:
		return &SingleCol{Table: m.Table, Name: m.FieldName, Alias: m.Name, Type: ColumnTypeValue}, nil
	case MetricTypeCount:
		return &SingleCol{Table: m.Table, Name: m.FieldName, Alias: m.Name, Type: ColumnTypeCount}, nil
	case MetricTypeDistinctCount:
		return &SingleCol{Table: m.Table, Name: m.FieldName, Alias: m.Name, Type: ColumnTypeDistinctCount}, nil
	case MetricTypeSum:
		return &SingleCol{Table: m.Table, Name: m.FieldName, Alias: m.Name, Type: ColumnTypeSum}, nil
	case MetricTypeAdd:
		return &ArithmeticCol{Column: column, Alias: m.Name, Type: ColumnTypeAdd}, nil
	case MetricTypeSubtract:
		return &ArithmeticCol{Column: column, Alias: m.Name, Type: ColumnTypeSubtract}, nil
	case MetricTypeMultiply:
		return &ArithmeticCol{Column: column, Alias: m.Name, Type: ColumnTypeMultiply}, nil
	case MetricTypeDivide:
		return &ArithmeticCol{Column: column, Alias: m.Name, Type: ColumnTypeDivide}, nil
	case MetricTypeExpression:
		return &ExpressionCol{Expression: m.FieldName, Alias: m.Name}, nil
	default:
		return nil, fmt.Errorf("not supported metric type %v", m.Type)
	}
}

func (m *Metric) ToProto() *proto.Metric {
	var children []*proto.Metric
	for _, v := range m.Children {
		children = append(children, v.ToProto())
	}
	return &proto.Metric{
		Type:           m.Type.ToEnum(),
		Table:          m.Table,
		Name:           m.Name,
		FieldName:      m.FieldName,
		Children:       children,
	}
}

func ProtoToMetric(m *proto.Metric) *Metric {
	return &Metric{}
}
