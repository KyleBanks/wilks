VERSION = 1.0
RELEASE_PKG = github.com/KyleBanks/wilks/cmd/wilks
INSTALL_PKG = $(RELEASE_PKG)

include github.com/KyleBanks/make/go/install
include github.com/KyleBanks/make/go/sanity
include github.com/KyleBanks/make/go/release
include github.com/KyleBanks/make/git/precommit

# Runs an example wilks score.
example: | install
	@wilks -lbs -bw 180 -total 1200 -gender m
.PHONY: example
