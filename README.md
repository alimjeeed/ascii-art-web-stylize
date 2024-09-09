# ascii-art-web

**ascii-art-web** is a Go-based web application that transforms text into stylized ASCII art using predefined character banners. It builds on the original **ascii-art** project by incorporating a web server for easier interaction through a web browser.

## Authors

- **Ali** - [@amajeed](https://learn.reboot01.com/git/amajeed)
- **Based** - [@bajaafar](https://learn.reboot01.com/git/bajaafar)
- **Osama** - [@oalmaska](https://learn.reboot01.com/git/oalmaska)

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
    git clone https://github.com/alimjeeed/ascii-art-web.git
    cd ascii-art-web
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
    Enter your text, choose a banner style, and click "Generate ASCII art" to see the result.

## Implementation Details

### Algorithm

The ASCII-Art-Web application uses an algorithm to convert text into ASCII art. Here’s a brief overview:

1. **Input Processing**: Validates the input to ensure it contains only printable ASCII characters.
2. **Template Selection**: Allows users to choose from various banner styles (e.g., standard, shadow). The selected template is loaded.
3. **Character Mapping**: Converts each character to its ASCII value and maps it to the corresponding art lines in the selected template.
4. **Output Generation**: Combines the ASCII art lines to produce the final output, which is displayed in the web interface.
