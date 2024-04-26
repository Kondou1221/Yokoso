FROM node:20-alpine as base

FROM base as dev

WORKDIR /app

COPY package.json yarn.lock* package-lock.json* pnpm-lock.yaml* ./

RUN \
    if [ -f yarn.lock ]; then yarn install --frozen-lockfile; \
    elif [ -f package-lock.json ]; then npm ci; \
    elif [ -f pnpm-lock.yaml ]; then yarn global add pnpm && pnpm i; \
    else echo "Lockfileがありません" && exit 1; \
    fi

COPY next.config.mjs .
COPY tsconfig.json .

CMD [ "yarn", "dev" ]