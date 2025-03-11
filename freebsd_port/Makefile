PORTNAME=	diceware
DISTVERSION=	v1.0.0
CATEGORIES=	sysutils

MAINTAINER=	chief@8ff.org
COMMENT=	Command-line passphrase generator using Diceware
WWW=		https://github.com/8ff/diceware

LICENSE=	MIT

USES=		go:modules
USE_GITHUB=	yes
GH_ACCOUNT=	8ff

GO_MODULE=	github.com/8ff/diceware
GO_TARGET=	./cmd/pwgen

PLIST_FILES=	bin/${PORTNAME}

# Rename binary from pwgen to diceware when installing
do-install:
	${INSTALL_PROGRAM} ${WRKDIR}/bin/pwgen ${STAGEDIR}${PREFIX}/bin/${PORTNAME}

.include <bsd.port.mk>
