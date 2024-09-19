# ASCII Art Web Stylize

**ASCII Art Web Stylize** is a Go-powered web app that turns your text into stunning ASCII art with a fun, Minecraft-inspired design. Its responsive interface ensures a smooth, engaging experience on any device, from desktop to mobile, making your words come alive in creative, pixelated style.

## Screenshot

![Screenshot of ascii-art-web-stylize](https://github.com/alimjeeed/ascii-art-web-stylize/blob/main/screenshot/ascii-art-web-stylize-screenshot.jpg)

## Example Output

```plaintext
 _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
                                
                                
```

## How to Use

1. **Clone the Repository**:
    ```bash
    git clone https://github.com/alimjeeed/ascii-art-web-stylize.git
    cd ascii-art-web-stylize
    ```

2. **Install Go**:  
   Make sure Go is installed. You can download it from the official [Go website](https://golang.org/).

3. **Run the Application**:
    ```bash
    go run main.go
    ```

4. **Access the Application**:
   Open your web browser and navigate to `http://localhost:8000` to access the web interface.

5. **Generate ASCII Art**:
   Input your text, select a style (Standard, Shadow, Thinkertoy), and click "Generate" to display the ASCII art.

## Project Structure

```plaintext
.
├── README.md
├── banners
│   ├── shadow.txt
│   ├── standard.txt
│   └── thinkertoy.txt
├── go.mod
├── internal
│   ├── generateasciiart.go
│   ├── getasciiart.go
│   ├── isstringsliceempty.go
│   ├── readbannerfile.go
│   └── stringtoascii.go
├── main.go
├── screenshot
│   └── ascii-art-web-stylize-screenshot.jpg
├── static
│   ├── minecraft-day.webp
│   ├── minecraft-night.webp
│   ├── monaco.ttf
│   └── style.css
├── templates
│   └── index.html
└── web
    ├── asciiart.go
    └── home.go
```
   
## How It Works (Algorithm)

The application converts text into ASCII art using the following steps:

1. **Input Validation**: The input is checked to ensure it contains only valid ASCII characters.
2. **Banner Style Selection**: Users can choose between different ASCII art styles (Standard, Shadow, Thinkertoy).
3. **ASCII Mapping**: Each character in the input text is mapped to the corresponding ASCII representation from the selected template.
4. **Output Generation**: The ASCII art is constructed and displayed on the web page.

## Dockerized Version

For a Dockerized version of this project, check out the repository here: [ASCII Art Web Dockerize](https://github.com/alimjeeed/ascii-art-web-dockerize.git).

## Authors

- **Ali** - [GitHub](https://github.com/alimjeeed)
- **Basem** - [GitHub](https://github.com/basem9999)
- **Osama** - [GitHub](https://github.com/oalmaska)

## License Information

The project uses the **Monaco** font, designed by Jamie Place, licensed under the Creative Commons Attribution Share Alike ([CC BY-SA 3.0](http://creativecommons.org/licenses/by-sa/3.0/)) license.

**Font Source**: [Monaco FontStruction](http://fontstruct.com/fontstructions/show/753435)
