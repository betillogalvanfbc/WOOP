# WOOP
Recon Plugins For Wordpress

## How to install.
go install -v github.com/betillogalvanfbc/WOOP@latest

## How to use.
WOOP https://wordpress.org/plugins/advanced-responsive-video-embedder |  xargs -I {} echo https://example.com/{} | httpx -s -sc
