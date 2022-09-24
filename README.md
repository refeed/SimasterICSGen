# SimasterICSGen -- ICS Generator from Simaster Page

[![Test and build SimasterICSGen](https://github.com/refeed/SimasterICSGen/actions/workflows/gotest.yml/badge.svg)](https://github.com/refeed/SimasterICSGen/actions/workflows/gotest.yml)
[![Build and Deploy simastericsgen-server](https://github.com/refeed/SimasterICSGen/actions/workflows/deploysmicsweb.yml/badge.svg)](https://github.com/refeed/SimasterICSGen/actions/workflows/deploysmicsweb.yml)

Easily import your exam and class schedule with SimasterICSGen! This program
will convert your Simaster schedule page from HTML to an ICS file which can
be imported by a variety of calendar apps including Google Calendar. Say good
bye to those old times where we had to input all of the schedules manually to our
favorite calendar apps.

## How to use

### Video tutorial (web based)

[![img_tag](https://img.youtube.com/vi/DvLRjajx9qs/0.jpg)](https://youtu.be/DvLRjajx9qs)

### CLI

```
Usage of SimasterICSGen:
  -input string
        (Mandatory) The HTML file of Simaster Jadwal Ujian page
  -output string
        The ICS output (default "result.ics")
```

Steps:
1. Download SimasterICSGen executable from this repo's releases page.
2. Go to [SIMASTER exam page](https://simaster.ugm.ac.id/akademik/mhs_jadwal_ujian/view)
3. Save the page (press `CTRL` + `S`)
4. Run `./SimasterICSGen -input <the html file you just saved>`, e.g
```
./SimasterICSGen -input "Simaster Jadwal Ujian.html"
```
5. The result will be saved as `result.ics`

And finally, you can import the `result.ics` file from a calendar app like Google
Calendar.

### Web based

Available at [https://refeed.github.io/SimasterICSGen/](https://refeed.github.io/SimasterICSGen/)

## TODO

- Add more docstring

## License

Available at LICENSE.txt
