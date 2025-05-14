# Stage 1: Builder (with full dev toolchain)
FROM node:20-alpine AS builder
WORKDIR /app

# 1. Install dependencies (with cache optimization)
COPY package.json package-lock.json ./
RUN npm ci --frozen-lockfile --prefer-offline

# 2. Copy only files needed for build
COPY tailwind.config.js postcss.config.mjs ./
COPY src ./src
COPY public ./public
COPY next.config.ts ./

# 3. Build with caching for Tailwind
RUN --mount=type=cache,target=/app/.next/cache \
    npm run build

# Stage 2: Production Runner
FROM node:20-alpine
WORKDIR /app

# 4. Production dependencies only
COPY --from=builder /app/package.json /app/package-lock.json ./
RUN npm ci --frozen-lockfile --omit=dev --prefer-offline

# 5. Copy built assets
COPY --from=builder /app/.next ./.next
COPY --from=builder /app/public ./public
COPY --from=builder /app/next.config.js ./

# 6. Environment setup
ENV NODE_ENV=production \
    NEXT_TELEMETRY_DISABLED=1 \
    PORT=3000 \
    NODE_OPTIONS="--max-old-space-size=1024"

# 7. Non-root user
RUN addgroup -g 1001 -S nodejs && \
    adduser -S nextjs -u 1001 -G nodejs && \
    chown -R nextjs:nodejs /app

# 8. Health check (Next.js + custom endpoint)
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s \
    CMD wget -qO- http://localhost:3000/health || exit 1

# 9. Structured logging
RUN npm install --save-dev pino-pretty
ENV NODE_ENV=production \
    LOG_FORMAT=json \
    LOG_LEVEL=info

USER nextjs

EXPOSE 3000

# 10. Start with log formatting
CMD ["sh", "-c", "npm start | ./node_modules/.bin/pino-pretty"]