package constanttest_test

const{
	Monday = iota +1
	Tuesday
	wednesday
}

func TestConstantTry(t *testing.T)  {
	t.Log(Monday, Tuesday)
}