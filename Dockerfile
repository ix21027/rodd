FROM golang:1.25.7-bookworm AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o rodd main.go

FROM archlinux:latest

RUN pacman -Sy --noconfirm archlinux-keyring && \
    pacman -Syu --noconfirm

# 2. Встановлюємо Chromium та все необхідне
RUN pacman -S --noconfirm \
    # Сам браузер (він підтягне libnss, libgtk, libalsa і т.д.)
    chromium \
    # Віртуальний дисплей та авторизація
    xorg-server-xvfb \
    xorg-xauth \
    # --- ШРИФТИ (ОБОВ'ЯЗКОВО!) ---
    # Без них скріншоти будуть з квадратиками
    ttf-liberation \
    noto-fonts \
    ttf-dejavu \
    # --- ДОДАТКОВІ БІБЛІОТЕКИ ---
    # libxss - для обробки скрінсейверів (іноді браузер падає без неї)
    libxss \
    # libxtst - для емуляції вводу
    libxtst \
    # xdg-utils - для роботи з посиланнями та середовищем
    xdg-utils \
    # Утиліти для мережі та скачування
    wget \
    ca-certificates \
    && pacman -Scc --noconfirm

# RUN apt-get update && apt-get install -y \
#     wget \
#     unzip \
#     xvfb \
#     xauth \
#     ca-certificates \
#     libnss3 \
#     libatk-bridge2.0-0 \
#     libgtk-3-0 \
#     libgbm1 \
#     libasound2 \
#     libxcomposite1 \
#     libxdamage1 \
#     libxrandr2 \
#     libxss1 \
#     libxtst6 \
#     fonts-liberation \
#     libappindicator3-1 \
#     xdg-utils \
#     --no-install-recommends \
#     && rm -rf /var/lib/apt/lists/*


# RUN wget -q -O chrome.zip "https://www.googleapis.com/download/storage/v1/b/chromium-browser-snapshots/o/Linux_x64%2F1520176%2Fchrome-linux.zip?alt=media" && \
#     unzip chrome.zip -d /opt/ && \
#     rm chrome.zip && \
#     chmod +x /opt/chrome-linux/chrome

# ENV CHROME_PATH="/opt/chrome-linux/chrome"

ENV CHROME_PATH="./download/linux-1520176/chrome-linux/chrome"

WORKDIR /app

COPY --from=builder /app/rodd .
COPY img/ ./img/
COPY entrypoint.sh .
COPY voeru ./voeru
RUN chmod +x entrypoint.sh ./rodd 
CMD ["./entrypoint.sh"]



