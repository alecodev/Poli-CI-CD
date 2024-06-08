FROM node:20-alpine AS base
LABEL authors="docker@alecodev.com"

RUN wget -qO- https://get.pnpm.io/install.sh | env PNPM_VERSION=9.0.6 ENV="$HOME/.shrc" SHELL="$(which sh)" sh -
ENV PNPM_HOME="/root/.local/share/pnpm"
ENV PATH="$PNPM_HOME:$PATH"

# Set base directory
WORKDIR /app



# Install dependencies only when needed
FROM base AS deps-base
# Check https://github.com/nodejs/docker-node/tree/b4117f9333da4138b03a546ec926ef50a31506c3#nodealpine to understand why libc6-compat might be needed.
RUN apk add --no-cache libc6-compat
COPY package.json pnpm-lock.yaml ./



# Install dependencies
FROM deps-base AS deps
RUN sed -i 's/"prepare": "cd .. && husky frontend\/.husky"/"prepare": ""/g' package.json
RUN CI=true pnpm install --frozen-lockfile



# Install dev dependencies
FROM deps-base AS deps-dev
RUN pnpm install



FROM base AS dev
COPY --from=deps-dev /app/node_modules ./node_modules
COPY . .



# Rebuild the source code only when needed
FROM base AS builder
COPY --from=deps /app/node_modules ./node_modules
COPY . .
ARG VITE_API_URL=""
ENV VITE_API_URL=$VITE_API_URL
ARG BASE_URL="/"
RUN pnpm build --base=$BASE_URL



# Production environment
FROM nginx:stable-alpine AS runner
LABEL authors="docker@alecodev.com"

COPY --from=builder /app/dist /usr/share/nginx/html
COPY --from=builder /app/nginx/nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]