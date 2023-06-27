package volume

import (
	"math"
)

// VolumeError representa um erro relacionado ao cálculo do volume
type VolumeError struct {
	Mensagem string
}

func (e *VolumeError) Error() string {
	return e.Mensagem
}

// CalcularVolumePiramide calcula o volume de uma pirâmide
func CalcularVolumePiramide(ladoQuadradoBase, alturaPiramide float32) (float32, error) {
	if ladoQuadradoBase <= 0 || alturaPiramide <= 0 {
		return 0, &VolumeError{"Valores inválidos. O lado da base e a altura da pirâmide devem ser maiores que zero."}
	}

	volume := ladoQuadradoBase * ladoQuadradoBase * (1.0 / 3.0) * alturaPiramide

	return volume, nil
}

// CalcularVolumeCubo calcula o volume de um cubo
func CalcularVolumeCubo(ladoQuadradoBase float32) (float32, error) {
	if ladoQuadradoBase <= 0 {
		return 0, &VolumeError{"Valor inválido. O lado do cubo deve ser maior que zero."}
	}

	volume := ladoQuadradoBase * ladoQuadradoBase * ladoQuadradoBase

	return volume, nil
}

// CalcularVolumeParalelepipedo calcula o volume de um paralelepípedo
func CalcularVolumeParalelepipedo(comprimentoBase, larguraBase, alturaBase float32) (float32, error) {
	if comprimentoBase <= 0 || larguraBase <= 0 || alturaBase <= 0 {
		return 0, &VolumeError{"Valores inválidos. O comprimento, largura e altura do paralelepípedo devem ser maiores que zero."}
	}

	volume := comprimentoBase * larguraBase * alturaBase

	return volume, nil
}

// CalcularVolumeEsfera calcula o volume de uma esfera
func CalcularVolumeEsfera(raio float32) (float32, error) {
	if raio <= 0 {
		return 0, &VolumeError{"Valor inválido. O raio da esfera deve ser maior que zero."}
	}

	volume := (4.0 / 3.0) * math.Pi * raio * raio * raio

	return volume, nil
}
