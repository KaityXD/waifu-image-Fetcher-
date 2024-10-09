import requests
import os
import argparse
import random
import string


def fetch_waifu(nsfw: bool, save_path: str, count: int):
    category = "nsfw" if nsfw else "sfw"
    url = f"https://api.waifu.pics/{category}/waifu"
    random_letters = ''.join(random.choices(string.ascii_lowercase, k=6))

    if not os.path.exists(save_path):
        os.makedirs(save_path)

    for i in range(count):
        try:
            response = requests.get(url)
            response.raise_for_status()

            data = response.json()
            image_url = data.get('url')

            if image_url:
                image_data = requests.get(image_url).content
                file_name = os.path.join(save_path, f"waifu_{category}_{i+1}_{random_letters}.jpg")

                with open(file_name, 'wb') as file:
                    file.write(image_data)
                print(f"Waifu image {i+1}/{count} downloaded successfully: {file_name}")
            else:
                print(f"Oops! Couldn't find a waifu for you on attempt {i+1}. Please try again later.")

        except requests.exceptions.RequestException as e:
            print(f"Error fetching waifu image {i+1}/{count}: {e}")
        except Exception as e:
            print(f"An unexpected error occurred on attempt {i+1}: {e}")

def main():
    parser = argparse.ArgumentParser(description="Fetch and download waifu images.")
    parser.add_argument("-n", "--nsfw", action="store_true", help="Fetch NSFW images.")
    parser.add_argument("-c", "--count", type=int, default=1, help="Number of waifu images to fetch (default is 1).")
    parser.add_argument("-s", "--save-path", type=str, default='./waifus', help="Directory to save images (default is './waifus').")

    args = parser.parse_args()

    fetch_waifu(args.nsfw, args.save_path, args.count)

if __name__ == "__main__":
    main()
