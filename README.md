# ASCII-ART-WEB

## 📌 Description

**Ascii-art-web** is a web application written in **Go** that allows users to generate ASCII art using different banners through a **web graphical user interface (GUI)**.  
It is the web version of the previous **ascii-art** project, where users can enter text, choose a banner, and instantly view the ASCII result in the browser.

The application supports the following banners:

- standard
- shadow
- thinkertoy

It works using HTTP **GET** and **POST** requests and follows proper **HTTP status codes**.

---

## 👨‍💻 Authors

- Yaakoub Elhaouari
- Hichame Ait benalla

## 🚀 Usage (How to Run)

### 1. Clone the repository:

```bash
git clone https://learn.zone01oujda.ma/git/yelhaoua/ascii-art-web.git
cd ascii-art-web
go run main.go
or
go run .
http://localhost:8080
```

### 2. Run the server:

```bash
go run main.go
or
go run .
```

### 3. Open your browser and go to:

```text
http://localhost:8080
```

### 4. Use the Web Interface:

Enter your text

Select a banner

Click Submit

The ASCII result will be displayed

## Implementation Details: Algorithm

- **Overall:**: The app converts user input text into ASCII art by loading a banner/font file (one of `standard.txt`, `shadow.txt`, `thinkertoy.txt`), parsing it into glyph rows, then mapping each printable ASCII character in the input to its glyph and assembling the output row-by-row.
- **Font file format:**: Each banner file contains a header line followed by blocks of 8 lines per glyph (with a separating blank line). `thinkertoy.txt` uses CRLF (`\r\n`) line endings; the others use LF (`\n`).
- **Parsing (`func.Splite`):**: Reads the selected font file (`os.ReadFile`) and splits its contents into lines (choosing `"\r\n"` for `thinkertoy.txt` and `"\n"` otherwise). It then iterates over the lines starting after the header and takes slices of 8 lines at a time (`all[i:i+8]`) to build a `[][]string` where each inner slice represents one character's 8 rows. The character order corresponds to ASCII codes starting at 32 (space).
- **Rendering (`func.PrintSymbole`):**: Takes the parsed glyph array and the input string. It splits the input into logical lines (by `"\r\n"`) and, for each input line, builds the ASCII art by iterating 8 rows (glyph height). For each rune in the input line, if it is a printable ASCII character between `' '` and `'~'`, the function computes the glyph index as `int(rune) - 32` and appends `arr[index][row]` to the current output row. Unsupported characters are skipped. The assembled rows are joined with newlines and returned as a single string.
- **HTTP flow (`handlers.HandleAscii` and `main.go`):**: `HandleAscii` validates the request (must be POST and include `name` and `radio` form fields), ensures the chosen banner file is allowed, calls `Splite` to get glyphs, calls `PrintSymbole` to render the art, and then executes the `ascii-art.html` template with the resulting string. `main.go` wires the routes: `/` (home), `/ascii-art` (render), and `/files/` (static file guard).
- **Notes and edge cases:**: `thinkertoy.txt` needs CRLF handling, empty input returns an empty string, and the renderer ignores characters outside printable ASCII. The mapping formula is `index = int(rune) - 32`.
