# YouTube CLI (ytcli)

YouTube CLI is a feature-rich command-line interface for interacting with YouTube. Manage subscriptions, discover trending videos, and more, all from your terminal!

## Features

- View latest videos from your subscriptions
- Discover trending videos
- Interactive mode for continuous operation
- Multiple output formats (pretty, simple, JSON)
- Customizable result limits

## Installation

### Prerequisites

- Go 1.16 or higher
- A Google account with YouTube API access

### Steps

1. Clone the repository:

   ```
   git clone https://github.com/yourusername/youtube-cli.git
   cd youtube-cli
   ```

2. Install dependencies:

   ```
   go mod tidy
   ```

3. Build the application:

   ```
   go build -o ytcli
   ```

4. Set up your YouTube API credentials:
   - Go to the [Google Cloud Console](https://console.cloud.google.com/)
   - Create a new project or select an existing one
   - Enable the YouTube Data API v3
   - Create OAuth 2.0 credentials (Client ID and Client Secret)
   - Create a `.env` file in the project root and add your credentials:
     ```
     YOUTUBE_CLIENT_ID=your_client_id_here
     YOUTUBE_CLIENT_SECRET=your_client_secret_here
     ```

## Usage

### Non-interactive mode

Run commands directly:

```
./ytcli subs
./ytcli trending
./ytcli subs --limit 10
./ytcli trending -o json
```

### Interactive mode

Start the interactive session:

```
./ytcli -i
```

Then enter commands at the prompt:

```
ytcli> subs
ytcli> trending --limit 15
ytcli> exit
```

### Available Commands

- `subs` (aliases: `subscriptions`): Show latest videos from your subscriptions
- `trending` (aliases: `popular`, `tr`): Show trending videos

### Global Flags

- `--output`, `-o`: Output format (pretty/simple/json) (default "pretty")
- `--limit`, `-l`: Maximum number of results to show (default 5)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Thanks to the Go community for the excellent libraries and tools
- YouTube API for providing access to YouTube data

## Support

If you encounter any problems or have any questions, please open an issue on the GitHub repository.
