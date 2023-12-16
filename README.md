# youtube.gaato.net

## Overview

Welcome to the YouTube Gaato Redirect Service, a user-friendly solution designed to enhance the experience of X (formerly Twitter) users. This service allows you to share YouTube videos on X more effectively by generating dynamic Twitter Cards for each shared link.

## For X Users: Enhancing Your Sharing Experience

- **Simple YouTube Video Sharing**: Share any YouTube video on X with enhanced visibility and engagement. Use `https://youtube.gaato.net/{videoID}` to share your favorite YouTube content.
- **Dynamic Twitter Cards**: Each shared link automatically generates a Twitter Card, displaying a video thumbnail and title, providing a richer and more engaging preview on X feeds.

## How to Use

1. **Find a YouTube Video**: Choose any video from YouTube that you want to share on X.
2. **Copy the Video ID**: The Video ID is the string after "watch?v=" in the YouTube URL.
3. **Create Your Link**: Append the Video ID to `https://youtube.gaato.net/`. For example, for `https://www.youtube.com/watch?v=dQw4w9WgXcQ`, your link will be `https://youtube.gaato.net/dQw4w9WgXcQ`.
4. **Share on X**: Post the `youtube.gaato.net` link on X. The Twitter Card with the video thumbnail and title will automatically appear in your post.

## For Developers

To set up the project locally for development, use Docker Compose:

```bash
docker-compose up --build
```

This command builds and runs the service in a Docker container. The application is pre-configured to listen on port 1216.

## Contributing

Feel free to contribute to the improvement of this service. Your suggestions and pull requests are welcome.

## License

This project is licensed under the [Apache License 2.0](LICENSE).

Copyright 2023 Gakuto Furuya
