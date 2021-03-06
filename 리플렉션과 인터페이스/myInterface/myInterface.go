/**************************************************************************************************************************************
 * interface 타입은 구체적인 동작을 구현할 메서드의 집합을 나열하는 방식으로 정의한다.
 * 어떤 타입이 특정 인터페이스를 따르려면 그 인터페이스에서 정의한 모든 메서드를 구현해야 한다.
 * 즉, 인터페이스는 추상 타입(abstract type)으로서 그 인터페이스의 인스턴스가 되기위해 반드시 구현해야 할 함수들을 정의한 것
 * "그 인터페이스의 인스턴스가 되기위해 반드시 구현"
 *
 * 이렇게 모든 메서드를 구현했다면 그 타입은 해당 인터페이스를 따른다고 표현
 * 따라서 인터페이스는 메서드 집합과 타입으로 구성되며, 다른 타입의 동작을 정의
 * 특정한 조건을 만족하고 특정한 동작을 제공하도록 확실히 보장하고 싶다면 인터페이스를 활용해야 한다.
 *
 * Go에서 가장 많이 사용하는 인터페이스는 io.Reader와 io.Writer이다.
 * type Reader interface {
 * 		Read(p []byte) (n int, err error)
 * }
 *
 *
**************************************************************************************************************************************/

package myInterface

type Shape interface {
	Area() float64
	Perimeter() float64
}
