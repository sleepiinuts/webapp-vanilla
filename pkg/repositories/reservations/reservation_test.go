package reservations

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

var rs *ReservationServ

type cases struct {
	name      string
	inqRooms  []int
	arrival   time.Time
	departure time.Time
	expected  map[int][]*Period
}

func TestListAvailRooms(t *testing.T) {
	cs := []cases{
		{
			name:      "basic case",
			inqRooms:  []int{1, 2, 3, 4},
			arrival:   mustParseTime("2024-06-01"),
			departure: mustParseTime("2024-06-12"),
			expected: map[int][]*Period{
				1: {
					{From: mustParseTime("2024-06-01"), To: mustParseTime("2024-06-01")},
					{From: mustParseTime("2024-06-05"), To: mustParseTime("2024-06-05")},
					{From: mustParseTime("2024-06-08"), To: mustParseTime("2024-06-12")},
				},
				2: {
					{From: mustParseTime("2024-06-01"), To: mustParseTime("2024-06-01")},
					{From: mustParseTime("2024-06-05"), To: mustParseTime("2024-06-12")},
				},
				3: {
					{From: mustParseTime("2024-06-01"), To: mustParseTime("2024-06-12")},
				},
				4: {
					{From: mustParseTime("2024-06-01"), To: mustParseTime("2024-06-12")},
				},
			},
		},
	}

	for _, c := range cs {
		t.Run(c.name, func(t *testing.T) {
			got, err := rs.ListAvailRooms(c.inqRooms, c.arrival, c.departure)
			if err != nil {
				t.Fail()
				t.Logf("expect nil error, but got %v", err)
			}

			if diff := cmp.Diff(c.expected, got); diff != "" {
				t.Fail()
				t.Logf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func init() {
	rs = New(&MockReservation{})
}
