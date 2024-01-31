package cycleDecimales

import "testing"

func Test1(t *testing.T) {
	res, _ := trouveCycle(1, 3)
	if res != 3 {
		t.Error("trouveCycle(1,3) devrait retourner 3 mais retourne ", res)
	}
}

func Test2(t *testing.T) {
	res, _ := trouveCycle(1, 2)
	if res != 0 {
		t.Error("trouveCycle(1,2) devrait retourner 0 mais retourne ", res)
	}
}

func Test3(t *testing.T) {
	res, _ := trouveCycle(7, 11)
	if res != 63 {
		t.Error("trouveCycle(7,11) devrait retourner 63 mais retourne ", res)
	}
}
