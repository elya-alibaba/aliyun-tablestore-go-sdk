package search

import (
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore/otsprotocol"
	"github.com/golang/protobuf/proto"
)

type SearchQuery interface {
	Serialize() ([]byte, error)
}

type SearchQueryBase struct {
	Offset        int32
	Limit         int32
	Query         Query
	Collapse      *Collapse
	Sort          *Sort
	GetTotalCount bool
	Token         []byte
	Aggregations  []Aggregation
	GroupBys      []GroupBy
}

func NewSearchQuery() *SearchQueryBase {
	return &SearchQueryBase{
		Offset:        -1,
		Limit:         -1,
		GetTotalCount: false,
	}
}

func (s *SearchQueryBase) SetOffset(offset int32) *SearchQueryBase {
	s.Offset = offset
	return s
}

func (s *SearchQueryBase) SetLimit(limit int32) *SearchQueryBase {
	s.Limit = limit
	return s
}

func (s *SearchQueryBase) SetQuery(query Query) *SearchQueryBase {
	s.Query = query
	return s
}

func NewAvgAggregation(name string, fieldName string) *AvgAggregation {
	return &AvgAggregation {
		AggName: name,
		Field: fieldName,
	}
}

func NewDistinctCountAggregation(name string, fieldName string) *DistinctCountAggregation {
	return &DistinctCountAggregation {
		AggName: name,
		Field: fieldName,
	}
}

func NewMaxAggregation(name string, fieldName string) *MaxAggregation {
	return &MaxAggregation {
		AggName: name,
		Field: fieldName,
	}
}

func NewMinAggregation(name string, fieldName string) *MinAggregation {
	return &MinAggregation {
		AggName: name,
		Field: fieldName,
	}
}

func NewSumAggregation(name string, fieldName string) *SumAggregation {
	return &SumAggregation {
		AggName: name,
		Field: fieldName,
	}
}

func NewCountAggregation(name string, fieldName string) *CountAggregation {
	return &CountAggregation {
		AggName: name,
		Field: fieldName,
	}
}

//
func NewGroupByField(name string, fieldName string) *GroupByField {
	return &GroupByField {
		AggName: name,
		Field:   fieldName,
	}
}

func NewGroupByRange(name string, fieldName string) *GroupByRange {
	return &GroupByRange {
		AggName: name,
		Field:   fieldName,
	}
}

func NewGroupByFilter(name string) *GroupByFilter {
	return &GroupByFilter {
		AggName: name,
	}
}

func NewGroupByGeoDistance(name string, fieldName string, origin GeoPoint) *GroupByGeoDistance {
	return &GroupByGeoDistance {
		AggName: name,
		Field:   fieldName,
		Origin: origin,
	}
}

func (s *SearchQueryBase) Aggregation(agg ...Aggregation) *SearchQueryBase {
	for i := 0; i < len(agg); i++ {
		s.Aggregations = append(s.Aggregations, agg[i])
	}
	return s
}

func (s *SearchQueryBase) GroupBy(groupBy ...GroupBy) *SearchQueryBase {
	for i := 0; i < len(groupBy); i++ {
		s.GroupBys = append(s.GroupBys, groupBy[i])
	}
	return s
}

func (s *SearchQueryBase) SetCollapse(collapse *Collapse) *SearchQueryBase {
	s.Collapse = collapse
	return s
}

func (s *SearchQueryBase) SetSort(sort *Sort) *SearchQueryBase {
	s.Sort = sort
	return s
}

func (s *SearchQueryBase) SetGetTotalCount(getTotalCount bool) *SearchQueryBase {
	s.GetTotalCount = getTotalCount
	return s
}

func (s *SearchQueryBase) SetToken(token []byte) *SearchQueryBase {
	s.Token = token
	s.Sort = nil
	return s
}

func (s *SearchQueryBase) Serialize() ([]byte, error) {
	searchQuery := &otsprotocol.SearchQuery{}
	if s.Offset >= 0 {
		searchQuery.Offset = &s.Offset
	}
	if s.Limit >= 0 {
		searchQuery.Limit = &s.Limit
	}
	if s.Query != nil {
		pbQuery, err := s.Query.ProtoBuffer()
		if err != nil {
			return nil, err
		}
		searchQuery.Query = pbQuery
	}
	if s.Collapse != nil {
		pbCollapse, err := s.Collapse.ProtoBuffer()
		if err != nil {
			return nil, err
		}
		searchQuery.Collapse = pbCollapse
	}
	if s.Sort != nil {
		pbSort, err := s.Sort.ProtoBuffer()
		if err != nil {
			return nil, err
		}
		searchQuery.Sort = pbSort
	}
	searchQuery.GetTotalCount = &s.GetTotalCount
	if s.Token != nil && len(s.Token) > 0 {
		searchQuery.Token = s.Token
	}

	if len(s.Aggregations) > 0 {
		pbAggregations := new(otsprotocol.Aggregations)
		for _, aggregation := range s.Aggregations {
			pbAggregation, err := aggregation.ProtoBuffer()
			if err != nil {
				return nil, err
			}
			pbAggregations.Aggs = append(pbAggregations.Aggs, pbAggregation)
		}
		searchQuery.Aggs = pbAggregations
	}

	if len(s.GroupBys) > 0 {
		pbGroupBys := new(otsprotocol.GroupBys)
		for _, groupBy := range s.GroupBys {
			pbGroupBy, err := groupBy.ProtoBuffer()
			if err != nil {
				return nil, err
			}
			pbGroupBys.GroupBys = append(pbGroupBys.GroupBys, pbGroupBy)
		}
		searchQuery.GroupBys = pbGroupBys
	}

	data, err := proto.Marshal(searchQuery)
	return data, err
}
