/*프로그래밍 회사 A,B,C가 있습니다. A회사는 B 회사가 만든 프로그램을 사용합니다. B회사는 데이ㅓ베이스 관련 프로그램을 제작합니다.
프로그램은 GET(), SET() 두 메소드를 제공합니다. A 회사에서는 B회사 제품의 성능을 C회사의 제품과 비교해보고 싶습니다.
그래서 성능을 비교하는 프로그램을 만들었습니다.*/

//덕타이핑

type Database interface{
	get()
	set()
}



func TotalTime(db Database) int {
	db.get()
	db.set()
}

func Compare() {
	BDB := &BDatabase{}
	CDC := &CDatabase{}

	if TotalTime(BDB) > TotalTime(CDB) { //두 회사의 구조체가 달라 한 함수의 인수로 사용할 수 없음
		fmt.Println("B회사의 제품이 더 빠릅니다.")
	} else {
		fmt.Println("C회사의 제품이 더 빠릅니다.")
	}
}

