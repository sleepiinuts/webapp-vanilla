package models

// func TestCalendar(t *testing.T) {
// 	tt, _ := time.Parse(time.DateOnly, "2024-06-27")
// 	c := NewCalendar(time.Time{}, tt)

// 	ans := [][]string{
// 		{"26", "27", "28", "29", "30", "31", "01"},
// 		{"02", "03", "04", "05", "06", "07", "08"},
// 		{"09", "10", "11", "12", "13", "14", "15"},
// 		{"16", "17", "18", "19", "20", "21", "22"},
// 		{"23", "24", "25", "26", "27", "28", "29"},
// 		{"30", "01", "02", "03", "04", "05", "06"},
// 	}

// 	got := c.Generate()
// 	if !cmp.Equal(ans, got) {
// 		t.Fail()
// 		t.Logf("expected: %v,\nbut got %v\n", ans, got)
// 	}
// }
