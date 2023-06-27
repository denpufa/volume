package volume
import (
	"math"
	"testing"
)


// CAMINHOS FELIZES
func TestCalcularVolumePiramide(t *testing.T) {
	ladoQuadradoBase := float32(4)
	alturaPiramide := float32(5)
	volumeEsperado := float32(26.666668)
	volumeCalculado, err := CalcularVolumePiramide(ladoQuadradoBase, alturaPiramide)
	if err != nil {
		t.Errorf("Erro inesperado: %s", err.Error())
	}
	if math.Abs(float64(volumeCalculado-volumeEsperado)) > 0.000001 {
		t.Errorf("Volume calculado incorreto. Esperado: %f, Obtido: %f", volumeEsperado, volumeCalculado)
	}
}

func TestCalcularVolumeCubo(t *testing.T) {
	ladoQuadradoBase := float32(3)
	volumeEsperado := float32(27)
	volumeCalculado, err := CalcularVolumeCubo(ladoQuadradoBase)
	if err != nil {
		t.Errorf("Erro inesperado: %s", err.Error())
	}
	if volumeCalculado != volumeEsperado {
		t.Errorf("Volume calculado incorreto. Esperado: %f, Obtido: %f", volumeEsperado, volumeCalculado)
	}
}

func TestCalcularVolumeParalelepipedo(t *testing.T) {
	comprimentoBase := float32(4)
	larguraBase := float32(5)
	alturaBase := float32(6)
	volumeEsperado := float32(120)
	volumeCalculado, err := CalcularVolumeParalelepipedo(comprimentoBase, larguraBase, alturaBase)
	if err != nil {
		t.Errorf("Erro inesperado: %s", err.Error())
	}
	if volumeCalculado != volumeEsperado {
		t.Errorf("Volume calculado incorreto. Esperado: %f, Obtido: %f", volumeEsperado, volumeCalculado)
	}
}

func TestCalcularVolumeEsfera(t *testing.T) {
	raio := float32(2)
	volumeEsperado := float32(33.510321)
	volumeCalculado, err := CalcularVolumeEsfera(raio)
	if err != nil {
		t.Errorf("Erro inesperado: %s", err.Error())
	}
	if math.Abs(float64(volumeCalculado-volumeEsperado)) > 0.000001 {
		t.Errorf("Volume calculado incorreto. Esperado: %f, Obtido: %f", volumeEsperado, volumeCalculado)
	}
}
// CAMINHOS INFELIZES

func TestErrorCalcularVolumePiramide(t *testing.T) {
	ladoQuadradoBase := float32(-4)
	alturaPiramide := float32(5)
	_, err := CalcularVolumePiramide(ladoQuadradoBase, alturaPiramide)
	if err == nil {
		t.Error("Esperava-se um erro, mas nenhum erro foi retornado")
	}
}

func TestErrorCalcularVolumeCubo(t *testing.T) {
	ladoQuadradoBase := float32(0)
	_, err := CalcularVolumeCubo(ladoQuadradoBase)
	if err == nil {
		t.Error("Esperava-se um erro, mas nenhum erro foi retornado")
	}
}

func TestErrorCalcularVolumeParalelepipedo(t *testing.T) {
	comprimentoBase := float32(-4)
	larguraBase := float32(5)
	alturaBase := float32(6)
	_, err := CalcularVolumeParalelepipedo(comprimentoBase, larguraBase, alturaBase)
	if err == nil {
		t.Error("Esperava-se um erro, mas nenhum erro foi retornado")
	}
}

func TestErrorCalcularVolumeEsfera(t *testing.T) {
	raio := float32(0)
	_, err := CalcularVolumeEsfera(raio)
	if err == nil {
		t.Error("Esperava-se um erro, mas nenhum erro foi retornado")
	}
}

