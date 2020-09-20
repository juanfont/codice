# codice
A CODICE¹ parser for files downloaded from https://contrataciondelestado.es, written in Go.

codice fetches zip/7z files found in the Spanish govt procurement website², and outputs a series of CSV files out of the XMLs inside those archives. 

## Usage 

Get your binary from https://github.com/juanfont/codice/releases/

And then just run:


```
./codice zip https://contrataciondelestado.es/sindicacion/sindicacion_643/licitacionesPerfilesContratanteCompleto3_2016.zip output_prefix
```


And you'll get:

```
Entries rows written to output_prefix_entries.csv
Modifications rows written to output_prefix_modifications.csv
Financial Guarantee rows written to output_prefix_financial_guarantees.csv
```


## Disclaimer

CSV is a rather limited format. codice has an opinionated way to flatten the CODICE XMLs. 

Just ping juanfontalonso@gmail.com if you find any issue.


____________________
¹: https://contrataciondelestado.es/wps/portal/CODICEInfo

²: https://www.hacienda.gob.es/es-ES/GobiernoAbierto/Datos%20Abiertos/Paginas/licitaciones_plataforma_contratacion.aspx
