# ASCII Art Web Stylize

**ASCII Art Web Stylize** is a Go-based web application that transforms text into stylized ASCII art, building on the original **ASCII Art** project. In this version, we focus on enhancing the user interface to be more visually appealing, interactive, and user-friendly by adopting a pixelated, blocky design inspired by the popular video game Minecraft. Additionally, the web app is fully responsive, ensuring smooth functionality across desktops, tablets, and small mobile devices.

## Screenshot

![Screenshot of ascii-art-web-stylize](https://github.com/alimjeeed/ascii-art-web-stylize/blob/main/screenshot/ascii-art-web-stylize-screenshot.jpg)

## Key Enhancements

1. **Appealing and Intuitive Design**: 
   - The web interface is revamped with a Minecraft-inspired theme to create a fun and engaging experience for users.
   
2. **Improved User Interaction**: 
   - Interactive buttons allow users to easily select between ASCII banner styles like *Standard*, *Shadow*, and *Thinkertoy*.

3. **Responsive Design**:
   - The web interface is fully optimized for different screen sizes.

## Example Output

```plaintext
 _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
                                
                                
```

## Usage

1. **Clone the Repository**:
    ```bash
    git clone https://github.com/alimjeeed/ascii-art-web-stylize.git
    cd ascii-art-web-stylize
    ```

2. **Install Dependencies**:
    Ensure you have Go installed on your system. You can download it from the official [Go website](https://golang.org/).

3. **Run the Server**:
    ```bash
    go run main.go
    ```

4. **Access the Web Interface**:
    Open your web browser and go to `http://localhost:8000`.

5. **Generate ASCII Art**:
    Enter your text, choose a banner style (Standard, Shadow, Thinkertoy), and click "Generate" to see the result.

## Implementation Details

### Algorithm

The ASCII-Art-Web-Stylize application uses an algorithm to convert text into ASCII art. Here’s a brief overview:

1. **Input Processing**: Validates the input to ensure it contains only printable ASCII characters.
2. **Template Selection**: Allows users to choose from various banner styles (e.g., standard, shadow, thinkertoy). The selected template is loaded.
3. **Character Mapping**: Converts each character to its ASCII value and maps it to the corresponding art lines in the selected template.
4. **Output Generation**: Combines the ASCII art lines to produce the final output, which is displayed in the web interface.

## Authors

- **Ali** - [GitHub](https://github.com/alimjeeed)
- **Basem** - [GitHub](https://github.com/basem9999)
- **Osama** - [GitHub](https://github.com/oalmaska)

## Font License

The font used in this project is:

- **FontStruction**: Monaco
- **Designer**: Jamie Place
- **License**: Creative Commons Attribution Share Alike ([CC BY-SA 3.0](http://creativecommons.org/licenses/by-sa/3.0/))

You can view and download the font here: [Monaco FontStruction](http://fontstruct.com/fontstructions/show/753435)
