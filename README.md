


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
