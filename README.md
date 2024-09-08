


# Fletcher - Waifu Image Fetcher

Fletcher is a command-line tool that fetches and downloads waifu images from the [waifu.pics](https://waifu.pics) API. You can specify the number of images to download and choose between SFW or NSFW content.

## Features

- Fetch and download waifu images.
- Supports SFW and NSFW images.
- Specify the number of images to download.
- Save images to a custom directory.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/fletcher.git
   cd fletcher
   ```

2. Ensure you have `requests` installed. If not, install it via pip:

   ```bash
   pip install requests
   ```

## Usage

### Fetch Images

To fetch waifu images, use the following command:

```bash
python3 fletcher.py -c <number> [-n] [-s <save-path>]
```

### Command-Line Arguments

- `-n` or `--nsfw`: Fetch NSFW images. Omit this flag to fetch SFW images.
- `-c` or `--count`: Number of waifu images to fetch (default is 1).
- `-s` or `--save-path`: Directory to save images (default is `./waifus`).

### Examples

1. **Fetch 10 SFW Waifu Images:**

   ```bash
   python3 fletcher.py -c 10
   ```

2. **Fetch 100 NSFW Waifu Images and Save to Custom Directory:**

   ```bash
   python3 fletcher.py -n -c 100 -s ./my_waifus
   ```

3. **Fetch 3000 SFW Waifu Images and Save to Default Directory:**

   ```bash
   python3 fletcher.py -c 3000
   ```

## License

We dont have one

## Contributing

Feel free to submit pull requests or open issues if you find bugs or have suggestions for improvements!

## Contact

For any questions, please reach out to [Emaill](mailto: gzyrrr123@gmail.com).
