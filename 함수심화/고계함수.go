//고계함수(higher order function)
//"함수"를 넘기고 받기만 하면 고계함수 ->특별할껀 없음
//

func ReadFrom(r io.Reader, f func(line string)) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		f(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}