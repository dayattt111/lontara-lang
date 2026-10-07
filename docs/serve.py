#!/usr/bin/env python3
import http.server
import socketserver
import os

PORT = 3000
DOCS_DIR = os.path.dirname(os.path.abspath(__file__))

class DocsHandler(http.server.SimpleHTTPRequestHandler):
    def __init__(self, *args, **kwargs):
        super().__init__(*args, directory=DOCS_DIR, **kwargs)

if __name__ == "__main__":
    with socketserver.TCPServer(("", PORT), DocsHandler) as httpd:
        print(f"\n==========================================")
        print(f" Dokumen Lontara-Lang aktif di:")
        print(f" http://localhost:{PORT}")
        print(f" Tekan Ctrl+C untuk berhenti")
        print(f"==========================================\n")
        try:
            httpd.serve_forever()
        except KeyboardInterrupt:
            print("\nServer dimatikan.")
