# WOOP
Recon Plugins For Wordpress

## How to install.
go install -v github.com/betillogalvanfbc/WOOP@latest

## Remember First.
curl -s -X GET https://example.com/ | grep -E 'wp-content/plugins/' | sed -E 's,href=|src=,THIIIIS,g' | awk -F "THIIIIS" '{print $2}' | cut -d "'" -f2 |

## How to use.
WOOP https://wordpress.org/plugins/advanced-responsive-video-embedder |  xargs -I {} echo https://example.com/{} | httpx -s -sc

## OR

WOOP https://wordpress.org/plugins/advanced-responsive-video-embedder |  xargs -I {} echo https://example.com/wp-content/plugins/{} | httpx -s -sc
