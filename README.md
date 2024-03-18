# ssltool

ssltool is a tool to check and generate certificates.

## Table of Contents
- [Getting Started](#getting-started)
- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Getting Started

You can download the binary from the releases page or you can build it yourself.

You can run ```make```

Or run ```go build``` from the root directory

## Installation

You can just copy the binary to a location in your path or run it directly from the
directory where you downloaded it.

## Usage

### Certificate Checking
Check the certificate details:

```./ssltool details --host www.example.com```

Check the certificate details and the certificate:

```./ssltool details --host www.example.com --cert```

Check service using a tls certificate with a custom port:

```./ssltool details --host ldaps.example.com --port 636```

If it is using a self signed certificate, you can use the --insecure flag to ignore the certificate errors.

```./ssltool details --host ldaps.example.com --port 636 --insecure```

### Certificate Generation
Generate a self signed certificate:

```LOCALITY="Bowling Green" PROVINCE="Kentucky" COUNTRY="US" ORG="Example ORG" OU="Example OU" ./ssltool gen -c www.example.com```

## Contributing

If you would like to contribute, please open an issue or a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

