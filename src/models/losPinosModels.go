package models

type DatosLosPinos struct {
	Nieve float32 `json:"nieve"`
    Lluvia float32 `json:"lluvia"`
    Visibilidad float32 `json:"visibilidad"`
    NubesBajas float32 `json:"nubesBajas"`
    NubesMedias float32 `json:"nubesMedias"`
    NubesAltas float32 `json:"nubesAltas"`
    NubesTotales float32 `json:"nubesTotales"`
    Precipitacion float32 `json:"precipitacion"`
    Msg string `json:"msg"`
}