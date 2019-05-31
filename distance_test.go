package distance

import (
	"reflect"
	"testing"
)

func TestHaversineDistance(t *testing.T) {
	type args struct {
		lat1 float64
		lng1 float64
		lat2 float64
		lng2 float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{
			name: "Test 1",
			args: args{
				lat1: 37.976900,
				lng1: 23.761785,
				lat2: 37.976814,
				lng2: 23.760384,
			},
			want: 123.3073152856457,
		},
		{
			name: "Test 2",
			args: args{
				lat1: 3,
				lng1: 4,
				lat2: 1,
				lng2: 1,
			},
			want: 401182.3620220665,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HaversineDistance(tt.args.lat1, tt.args.lng1, tt.args.lat2, tt.args.lng2); got != tt.want {
				t.Errorf("HaversineDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_haversine(t *testing.T) {
	type args struct {
		theta float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := haversine(tt.args.theta); got != tt.want {
				t.Errorf("haversine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rad(t *testing.T) {
	type args struct {
		deg float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rad(tt.args.deg); got != tt.want {
				t.Errorf("rad() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSquare(t *testing.T) {
	type args struct {
		lat float64
		lng float64
		d   float64
	}
	tests := []struct {
		name  string
		args  args
		want  Point
		want1 Point
	}{
		// TODO: Add test cases.
		{
			name: "TestSquare 1",
			args: args{
				lat: 3,
				lng: 4,
				d:   10000,
			},
			want:  Point{3.089832049533689, 4.089955330187352},
			want1: Point{2.910167950466311, 3.910044669812648},
		},
		{
			name: "TestSquare 2",
			args: args{
				lat: 3,
				lng: 4,
				d:   401182.3620220665,
			},

			want:  Point{6.6039033817208725, 7.60884918410367},
			want1: Point{-0.6039033817208721, 0.3911508158963293}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Square(tt.args.lat, tt.args.lng, tt.args.d)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Square() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Square() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		lat float64
		lng float64
	}
	tests := []struct {
		name string
		args args
		want Point
	}{
		// TODO: Add test cases.
		{
			name: "Test New 1",
			args: args{
				lat: 4,
				lng: 3,
			},
			want: Point{4, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.lat, tt.args.lng); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint_DistanceFrom(t *testing.T) {
	type fields struct {
		lat float64
		lng float64
	}
	type args struct {
		pp Point
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		// TODO: Add test cases.
		{
			name: "TestPoint_DistanceFrom 1",
			fields: fields{
				lat: 37.976900,
				lng: 23.761785,
			},
			args: args{
				Point{
					lat: 37.976814,
					lng: 23.760384,
				}},
			want: 123.3073152856457,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Point{
				lat: tt.fields.lat,
				lng: tt.fields.lng,
			}
			if got := p.DistanceFrom(tt.args.pp); got != tt.want {
				t.Errorf("Point.DistanceFrom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint_Square(t *testing.T) {
	type fields struct {
		lat float64
		lng float64
	}
	type args struct {
		d float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Point
		want1  Point
	}{
		// TODO: Add test cases.
		{
			name: "TestPoint_Square 1",
			fields: fields{
				lat: 3,
				lng: 4,
			},
			args:  args{10000},
			want:  Point{3.089832049533689, 4.089955330187352},
			want1: Point{2.910167950466311, 3.910044669812648},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Point{
				lat: tt.fields.lat,
				lng: tt.fields.lng,
			}
			got, got1 := p.Square(tt.args.d)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Point.Square() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Point.Square() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPoint_Lat(t *testing.T) {
	type fields struct {
		lat float64
		lng float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		// TODO: Add test cases.
		{
			name: "TestPoint_Lat 1",
			fields: fields{
				lat: 3,
				lng: 4},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Point{
				lat: tt.fields.lat,
				lng: tt.fields.lng,
			}
			if got := p.Lat(); got != tt.want {
				t.Errorf("Point.Lat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint_lng(t *testing.T) {
	type fields struct {
		lat float64
		lng float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		// TODO: Add test cases.
		{
			name: "TestPoint_lng 1",
			fields: fields{
				lat: 3,
				lng: 4},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Point{
				lat: tt.fields.lat,
				lng: tt.fields.lng,
			}
			if got := p.Lng(); got != tt.want {
				t.Errorf("Point.Lng() = %v, want %v", got, tt.want)
			}
		})
	}
}
