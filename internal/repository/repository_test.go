package repository

import (
	"reflect"
	"testing"
	"time"
)

func TestAddCost(t *testing.T) {
	type args struct {
		date     time.Time
		category string
		sum      int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "base",
			args: args{
				time.Now(),
				"еда",
				34567},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AddCost(tt.args.date, tt.args.category, tt.args.sum); (err != nil) != tt.wantErr {
				t.Errorf("AddCost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetCostsForPeriod(t *testing.T) {
	type args struct {
		begin time.Time
		end   time.Time
	}
	tests := []struct {
		name    string
		args    args
		costsEx map[time.Time]map[string]int
		want    map[string]int
		wantErr bool
	}{
		{
			name: "base",
			args: args{
				time.Now(),
				time.Now().AddDate(-1, 0, 0),
			},
			costsEx: map[time.Time]map[string]int{},
			want:    map[string]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCostsForPeriod(tt.args.begin, tt.args.end)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCostsForPeriod() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCostsForPeriod() got = %v, want %v", got, tt.want)
			}
		})
	}
}
