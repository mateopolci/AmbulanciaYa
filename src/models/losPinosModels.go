package models

type DatosLosPinos struct {
	Nieve int16 `json:"nieve"`
    Lluvia int16 `json:"lluvia"`
    Visibilidad int16 `json:"visibilidad"`
    NubesBajas int16 `json:"nubesBajas"`
    NubesMedias int16 `json:"nubesMedias"`
    NubesAltas int16 `json:"nubesAltas"`
    NubesTotales int16 `json:"nubesTotales"`
    Precipitacion int16 `json:"precipitacion"`
    Msg string `json:"msg"`
}