FROM node:20-alpine as base

FROM base as dev

WORKDIR /app

COPY ./next-app/package.json ./next-app/yarn.lock* ./next-app/package-lock.json* ./next-app/pnpm-lock.yaml* ./

RUN \
    if [ -f yarn.lock ]; then yarn install --frozen-lockfile; \
    elif [ -f package-lock.json ]; then npm ci; \
    elif [ -f pnpm-lock.yaml ]; then yarn global add pnpm && pnpm i; \
    else echo "Lockfileがありません" && exit 1; \
    fi

COPY ./next-app/tsconfig.json .

CMD [ "yarn", "dev" ]