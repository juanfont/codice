# codice
A CODICE¹ parser for files downloaded from https://contrataciondelestado.es, written in Go.

codice fetches zip/7z files found in the Spanish govt procurement website², and outputs a series of CSV files out of the XMLs inside those archives. 

## Usage 

Get your binary from https://github.com/juanfont/codice/releases/

And then just run:


```
./codice zip https://contrataciondelestado.es/sindicacion/sindicacion_643/licitacionesPerfilesContratanteCompleto3_2016.zip output_prefix
```

Alternatively, you can read an already-download zip from your disk:

```
./codice zip licitacionesPerfilesContratanteCompleto3_2016.zip output_prefix
```


In both cases you'll get:

```
Entries rows written to output_prefix_entries.csv
Modifications rows written to output_prefix_modifications.csv
Financial Guarantee rows written to output_prefix_financial_guarantees.csv
```

**Note:** Sometimes an entry might appear multiple times within the same zip. This happens when the date on a procurement process is updated within the same year. You can pass the flag `--aggregate` to keep only the most updated reference to these entries.

```
./codice zip https://contrataciondelestado.es/sindicacion/sindicacion_643/licitacionesPerfilesContratanteCompleto3_2016.zip output_prefix --aggregate
```


## Disclaimer

CSV is a rather limited format. codice has an opinionated way to flatten the CODICE XMLs. 

Just ping juanfontalonso@gmail.com if you find any issue.


____________________
¹: https://contrataciondelestado.es/wps/portal/CODICEInfo

²: https://www.hacienda.gob.es/es-ES/GobiernoAbierto/Datos%20Abiertos/Paginas/licitaciones_plataforma_contratacion.aspx
