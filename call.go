package calls

//True
func True(flag bool, call func()) {
	if flag {
		call()
	}
}

//False
func False(flag bool, call func()) {
	if !flag {
		call()
	}
}

//Nil
func Nil(i interface{}, call func()) {
	if i == nil {
		call()
	}
}

//NNil
func NNil(i interface{}, call func()) {
	if i != nil {
		call()
	}
}

//Empty
func Empty(str string, call func()) {
	if str == "" {
		call()
	}
}

//NEmpty
func NEmpty(str string, call func()) {
	if str != "" {
		call()
	}
}

//Zero
func Zero(i int, call func()) {
	if i == 0 {
		call()
	}
}

//NZero
func NZero(i int, call func()) {
	if i != 0 {
		call()
	}
}

//GtZero
func GtZero(i int, call func()) {
	if i > 0 {
		call()
	}
}

//LtZero
func LtZero(i int, call func()) {
	if i < 0 {
		call()
	}
}
