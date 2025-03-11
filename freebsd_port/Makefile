PORTNAME=	diceware
DISTVERSION=	latest
CATEGORIES=	sysutils

MAINTAINER=	chief@8ff.org
COMMENT=	Command-line passphrase generator using Diceware
WWW=		https://github.com/8ff/diceware

LICENSE=	MIT

USE_GITHUB=	yes
GH_ACCOUNT=	8ff
GH_TAGNAME=	latest

# The port only installs a single binary:
PLIST_FILES=	bin/${PORTNAME}

# Download and set up the latest Go, then build diceware.
do-build:
	@${MKDIR} ${WRKDIR}/gobin
	@cd ${WRKDIR} && \
		${FETCH_CMD} -o go.json https://go.dev/dl/?mode=json && \
		LATEST_GO=$$(grep -m1 '"version"' go.json | ${CUT} -d'"' -f4 | ${TR} -d 'go') && \
		${FETCH_CMD} https://go.dev/dl/go$${LATEST_GO}.freebsd-amd64.tar.gz && \
		${TAR} xf go$${LATEST_GO}.freebsd-amd64.tar.gz
	@cd ${WRKSRC}/cmd/pwgen && \
		${SETENV} \
		PATH=${WRKDIR}/go/bin:${PATH} \
		GOROOT=${WRKDIR}/go \
		go build -o ${WRKDIR}/${PORTNAME} .

do-install:
	${INSTALL_PROGRAM} ${WRKDIR}/${PORTNAME} ${STAGEDIR}${PREFIX}/bin

.include <bsd.port.mk>
