package main

import (
	"io"
)

func Ok1() (n int) {
	return
}

func Ok2() (n, _ int) {
	return
}

func Ok3() (a, b int) {
	return
}

func Ok4() (int, int) {
	go func() int {
		for i := 0; i < 10; i++ {
			return 5
		}
		return 3
	}()
	return Ok3()
}

func Ok5(w io.Writer) (n int, err error) {
	return w.Write(nil)
}

func NotEnough1() int {
	/*begin*/return/*end.Not enough arguments to return|ChangeReturnsParametersFix*/
}

func NotEnough2(a int) (b, c int) {
	if a == 1 {
		return
	} else if a == 2 {
		/*begin*/return 1/*end.Not enough arguments to return|ChangeReturnsParametersFix*/
	} else if a == 3 {
		return 1, 2
	}
	/*begin*/return 3, 4, 5/*end.Too many arguments to return|ChangeReturnsParametersFix*/
}

func NotEnough3(a int) (int, int) {
	if a == 1 {
		/*begin*/return/*end.Not enough arguments to return|ChangeReturnsParametersFix*/
	} else if a == 2 {
		/*begin*/return 1/*end.Not enough arguments to return|ChangeReturnsParametersFix*/
	} else if a == 3 {
		return 1, 2
	}
	/*begin*/return 3, 4, 5/*end.Too many arguments to return|ChangeReturnsParametersFix*/
}

func TooMany1(a, b int) {
	if a > 0 {
		/*begin*/return a + b/*end.Too many arguments to return|ChangeReturnsParametersFix*/
	}
	/*begin*/return a/*end.Too many arguments to return|ChangeReturnsParametersFix*/
}

func NotMatchType() (string, int64) {
	return "ok", /*begin*/"not ok"/*end.Expression type mismatch, the expected type is int64|CastTypeFix|ChangeReturnsParametersFix*/
}

func NotMatchTypeCall() (string, int64) {
	return /*begin*/Ok2()/*end.The returned expressions don't match with the return parameters|ChangeReturnsParametersFix*/
}

func IsError() error {
	return nil
}
