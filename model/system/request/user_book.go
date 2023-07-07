package request

import "Library_Project/model/system"

type MyLendBook struct {
	system.Book
	system.LendList
}
