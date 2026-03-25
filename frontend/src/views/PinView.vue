<template>
  <div class="pin-page">
    <!-- ── Main card ── -->
    <div class="pin-wrap" @click.self="goBack">
      <div class="pin-card">

        <!-- Image side -->
        <div class="pin-image-side">
          <button class="btn-back" @click="goBack">
            <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M19 12H5"/><path d="M12 19l-7-7 7-7"/>
            </svg>
          </button>
          <div :class="['pin-image', pin.palette, pin.height]">
            <!-- Placeholder; replace with <img :src="pin.imageUrl" :alt="pin.title"> -->
          </div>
        </div>

        <!-- Info side -->
        <div class="pin-info-side">

          <!-- Actions -->
          <div class="pin-actions-top">
            <div class="action-left">
              <button class="btn-icon-sm" title="Поделиться" @click="share">
                <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <circle cx="18" cy="5" r="3"/><circle cx="6" cy="12" r="3"/><circle cx="18" cy="19" r="3"/>
                  <line x1="8.59" y1="13.51" x2="15.42" y2="17.49"/>
                  <line x1="15.41" y1="6.51" x2="8.59" y2="10.49"/>
                </svg>
              </button>
              <button class="btn-icon-sm" title="Ещё" @click="toggleMenu">
                <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <circle cx="5" cy="12" r="1"/><circle cx="12" cy="12" r="1"/><circle cx="19" cy="12" r="1"/>
                </svg>
              </button>
            </div>
            <div class="action-right">
              <div class="board-select">
                <span>{{ selectedBoard }}</span>
                <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <polyline points="6 9 12 15 18 9"/>
                </svg>
              </div>
              <button class="btn-save" :class="{ saved: isSaved }" @click="toggleSave">
                {{ isSaved ? 'Сохранено' : 'Сохранить' }}
              </button>
            </div>
          </div>

          <!-- Source link -->
          <a v-if="pin.sourceUrl" :href="pin.sourceUrl" target="_blank" class="source-link">
            <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M18 13v6a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h6"/>
              <polyline points="15 3 21 3 21 9"/><line x1="10" y1="14" x2="21" y2="3"/>
            </svg>
            {{ pin.sourceDomain }}
          </a>

          <!-- Title & description -->
          <h1 class="pin-title">{{ pin.title }}</h1>
          <p class="pin-desc">{{ pin.description }}</p>

          <!-- Tags -->
          <div v-if="pin.tags?.length" class="tags">
            <span v-for="tag in pin.tags" :key="tag" class="tag">#{{ tag }}</span>
          </div>

          <!-- Author -->
          <div class="author-row">
            <div class="author-avatar">{{ pin.authorInitials }}</div>
            <div class="author-info">
              <span class="author-name">{{ pin.authorName }}</span>
              <span class="author-handle">{{ pin.author }}</span>
            </div>
            <button
              class="btn-follow"
              :class="{ following: isFollowing }"
              @click="toggleFollow"
            >
              {{ isFollowing ? 'Вы подписаны' : 'Подписаться' }}
            </button>
          </div>

          <!-- Likes & stats -->
          <div class="pin-stats">
            <button class="stat-btn" :class="{ liked: isLiked }" @click="toggleLike">
              <svg width="16" height="16" viewBox="0 0 24 24" :fill="isLiked ? 'currentColor' : 'none'" stroke="currentColor" stroke-width="2">
                <path d="M20.84 4.61a5.5 5.5 0 0 0-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 0 0-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78 1.06-1.06a5.5 5.5 0 0 0 0-7.78z"/>
              </svg>
              {{ likeCount }}
            </button>
            <span class="stat-sep">·</span>
            <span class="stat-text">{{ pin.views }} просмотров</span>
          </div>

          <!-- Comments -->
          <div class="comments-section">
            <h2 class="comments-title">Комментарии <span>{{ comments.length }}</span></h2>

            <div class="comment-input-row">
              <div class="author-avatar sm">АИ</div>
              <div class="comment-input-wrap">
                <input
                  v-model="commentText"
                  class="comment-input"
                  type="text"
                  placeholder="Напишите комментарий..."
                  @keydown.enter="addComment"
                />
                <button v-if="commentText" class="btn-comment-send" @click="addComment">
                  <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <line x1="22" y1="2" x2="11" y2="13"/>
                    <polygon points="22 2 15 22 11 13 2 9 22 2"/>
                  </svg>
                </button>
              </div>
            </div>

            <div class="comments-list">
              <div v-for="comment in comments" :key="comment.id" class="comment">
                <div class="author-avatar sm">{{ comment.initials }}</div>
                <div class="comment-body">
                  <div class="comment-header">
                    <span class="comment-author">{{ comment.author }}</span>
                    <span class="comment-time">{{ comment.time }}</span>
                  </div>
                  <p class="comment-text">{{ comment.text }}</p>
                  <button class="comment-like" @click="comment.liked = !comment.liked">
                    <svg width="12" height="12" viewBox="0 0 24 24" :fill="comment.liked ? 'currentColor' : 'none'" stroke="currentColor" stroke-width="2">
                      <path d="M20.84 4.61a5.5 5.5 0 0 0-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 0 0-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78 1.06-1.06a5.5 5.5 0 0 0 0-7.78z"/>
                    </svg>
                    {{ comment.likes }}
                  </button>
                </div>
              </div>
            </div>
          </div>

        </div>
      </div>
    </div>

    <!-- ── Related pins ── -->
    <div class="related">
      <h2 class="related-title">Похожие пины</h2>
      <div class="masonry">
        <article
          v-for="(rpin, i) in relatedPins"
          :key="rpin.id"
          class="pin"
          :style="{ animationDelay: `${i * 0.05}s` }"
        >
          <div class="pin-img-wrap">
            <div :class="['pin-img', rpin.palette, rpin.height]" />
            <div class="pin-overlay">
              <button class="pin-save-btn" @click.stop>
                <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <path d="M19 21l-7-5-7 5V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2z"/>
                </svg>
                Сохранить
              </button>
            </div>
          </div>
          <div class="pin-body">
            <p class="pin-label">{{ rpin.title }}</p>
            <div class="pin-meta">
              <div class="pin-author">
                <div class="author-dot">{{ rpin.authorInitials }}</div>
                {{ rpin.author }}
              </div>
              <div class="pin-likes">
                <svg width="11" height="11" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <path d="M20.84 4.61a5.5 5.5 0 0 0-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 0 0-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78 1.06-1.06a5.5 5.5 0 0 0 0-7.78z"/>
                </svg>
                {{ rpin.likes }}
              </div>
            </div>
          </div>
        </article>
      </div>
    </div>

  </div>
</template>

<script setup>
import { ref } from 'vue'
import { RouterLink, useRouter } from 'vue-router'

const router = useRouter()

// ── State ──
const searchQuery  = ref('')
const isSaved      = ref(false)
const isFollowing  = ref(false)
const isLiked      = ref(false)
const likeCount    = ref(234)
const commentText  = ref('')
const selectedBoard = ref('Интерьеры')

// ── Pin data (in real app — fetched by route param) ──
const pin = ref({
  id: 1,
  title: 'Японский минимализм в жилом интерьере',
  description: 'Пространство, вдохновлённое принципами ваби-саби: несовершенство как красота. Нейтральные текстуры, натуральные материалы и тишина — вот что делает этот интерьер особенным.',
  tags: ['минимализм', 'интерьер', 'япония', 'вабисаби'],
  palette: 'pal-1',
  height: 'img-tall',
  sourceUrl: 'https://example.com',
  sourceDomain: 'dezeen.com',
  authorName: 'Марина Дизайнер',
  author: '@marinadiz',
  authorInitials: 'МД',
  views: '1 482',
  likes: 234,
})

const comments = ref([
  { id: 1, author: 'Алексей И.',  initials: 'АИ', text: 'Невероятное пространство, особенно нравится, как работает свет.',   time: '2 ч назад',  likes: 12, liked: false },
  { id: 2, author: 'Ольга С.',    initials: 'ОС', text: 'Сохраняю в доску «Вдохновение». Хочу нечто похожее дома.',          time: '5 ч назад',  likes: 7,  liked: false },
  { id: 3, author: 'Иван А.',     initials: 'ИА', text: 'Какой материал на стенах? Штукатурка или натуральный камень?',       time: '1 д назад',  likes: 3,  liked: false },
])

const relatedPins = ref([
  { id: 10, title: 'Бетон и дерево',                   author: '@klara',        authorInitials: 'КЛ', likes: 178, height: 'h-md', palette: 'pal-7' },
  { id: 11, title: 'Скандинавская палитра',             author: '@annaN',        authorInitials: 'АН', likes: 89,  height: 'h-lg', palette: 'pal-5' },
  { id: 12, title: 'Монохромная кухня',                 author: '@tkitchen',     authorInitials: 'ТК', likes: 145, height: 'h-sm', palette: 'pal-4' },
  { id: 13, title: 'Студия с открытой планировкой',     author: '@ivanarch',     authorInitials: 'ИВ', likes: 412, height: 'h-xl', palette: 'pal-3' },
  { id: 14, title: 'Каменная терраса с видом на горы',  author: '@evgarch',      authorInitials: 'ЕА', likes: 388, height: 'h-md', palette: 'pal-10'},
  { id: 15, title: 'Тихий угол',                       author: '@olgas',        authorInitials: 'ОС', likes: 56,  height: 'h-xs', palette: 'pal-2' },
])

// ── Methods ──
function goBack() { router.back() }

function toggleSave() { isSaved.value = !isSaved.value }

function toggleFollow() { isFollowing.value = !isFollowing.value }

function toggleLike() {
  isLiked.value = !isLiked.value
  likeCount.value += isLiked.value ? 1 : -1
}

function share() { navigator.share?.({ title: pin.value.title, url: window.location.href }) }

function toggleMenu() { /* open dropdown */ }

function addComment() {
  if (!commentText.value.trim()) return
  comments.value.unshift({
    id: Date.now(),
    author: 'Алексей И.',
    initials: 'АИ',
    text: commentText.value.trim(),
    time: 'только что',
    likes: 0,
    liked: false,
  })
  commentText.value = ''
}
</script>

<style scoped>
.pin-page {
  min-height: 100vh;
  background: var(--bg);
  font-family: 'DM Sans', sans-serif;
  color: var(--ink);
}

/* ── Nav ── */
nav {
  position: sticky;
  top: 0;
  background: var(--white);
  border-bottom: 1px solid var(--border);
  z-index: 100;
  padding: 0 24px;
  display: flex;
  align-items: center;
  gap: 16px;
  height: 60px;
}

.logo {
  font-family: 'DM Serif Display', serif;
  font-size: 22px;
  color: var(--ink);
  letter-spacing: -0.5px;
  text-decoration: none;
  flex-shrink: 0;
}
.logo span { color: var(--accent); }

.nav-search {
  flex: 1;
  position: relative;
  max-width: 480px;
}

.nav-search input {
  width: 100%;
  padding: 9px 16px 9px 38px;
  border: 1.5px solid var(--border);
  border-radius: 100px;
  font-family: 'DM Sans', sans-serif;
  font-size: 14px;
  background: var(--bg);
  color: var(--ink);
  outline: none;
  transition: border-color 0.2s;
  box-sizing: border-box;
}

.nav-search input:focus { border-color: var(--ink); }
.nav-search input::placeholder { color: #C4C0BA; }

.search-icon {
  position: absolute;
  left: 12px;
  top: 50%;
  transform: translateY(-50%);
  color: var(--muted);
  pointer-events: none;
}

.nav-right {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-left: auto;
  flex-shrink: 0;
}

.btn-upload {
  padding: 8px 18px;
  background: var(--accent);
  color: var(--white);
  border: none;
  border-radius: 100px;
  font-family: 'DM Sans', sans-serif;
  font-size: 13px;
  font-weight: 500;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 6px;
  transition: opacity 0.2s;
  white-space: nowrap;
  text-decoration: none;
}
.btn-upload:hover { opacity: 0.85; }

.btn-icon {
  width: 38px;
  height: 38px;
  border: 1.5px solid var(--border);
  border-radius: 50%;
  background: transparent;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--ink);
  transition: border-color 0.2s;
  flex-shrink: 0;
}
.btn-icon:hover { border-color: var(--ink); }

.avatar-nav {
  width: 38px;
  height: 38px;
  border-radius: 50%;
  background: var(--ink);
  cursor: pointer;
  display: block;
  flex-shrink: 0;
}

/* ── Pin wrap ── */
.pin-wrap {
  padding: 32px 24px;
  display: flex;
  justify-content: center;
}

.pin-card {
  display: grid;
  grid-template-columns: minmax(320px, 520px) minmax(320px, 440px);
  background: var(--white);
  border-radius: 24px;
  overflow: hidden;
  border: 1px solid var(--border);
  max-width: 960px;
  width: 100%;
  animation: cardIn 0.4s ease both;
}

/* ── Image side ── */
.pin-image-side {
  position: relative;
  background: var(--bg);
  min-height: 520px;
}

.btn-back {
  position: absolute;
  top: 16px;
  left: 16px;
  z-index: 2;
  width: 36px;
  height: 36px;
  border-radius: 50%;
  background: var(--white);
  border: 1.5px solid var(--border);
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--ink);
  transition: border-color 0.2s;
}
.btn-back:hover { border-color: var(--ink); }

.pin-image {
  width: 100%;
  height: 100%;
  min-height: 520px;
}

/* height helpers for placeholder */
.img-tall { min-height: 560px; }

/* ── Info side ── */
.pin-info-side {
  padding: 28px 32px;
  display: flex;
  flex-direction: column;
  gap: 0;
  overflow-y: auto;
  max-height: 90vh;
}

/* Actions top row */
.pin-actions-top {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 20px;
}

.action-left { display: flex; gap: 6px; }
.action-right { display: flex; align-items: center; gap: 8px; }

.btn-icon-sm {
  width: 36px;
  height: 36px;
  border-radius: 50%;
  border: 1.5px solid var(--border);
  background: transparent;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--ink);
  transition: border-color 0.2s;
}
.btn-icon-sm:hover { border-color: var(--ink); }

.board-select {
  display: flex;
  align-items: center;
  gap: 6px;
  padding: 7px 12px;
  border: 1.5px solid var(--border);
  border-radius: 8px;
  font-size: 13px;
  font-weight: 500;
  cursor: pointer;
  transition: border-color 0.2s;
  white-space: nowrap;
}
.board-select:hover { border-color: var(--ink); }

.btn-save {
  padding: 9px 20px;
  background: var(--accent);
  color: var(--white);
  border: none;
  border-radius: 8px;
  font-family: 'DM Sans', sans-serif;
  font-size: 13px;
  font-weight: 500;
  cursor: pointer;
  transition: opacity 0.2s;
  white-space: nowrap;
}
.btn-save:hover { opacity: 0.85; }
.btn-save.saved { background: var(--ink); }

/* Source */
.source-link {
  display: inline-flex;
  align-items: center;
  gap: 5px;
  font-size: 12px;
  color: var(--muted);
  text-decoration: none;
  margin-bottom: 14px;
  transition: color 0.2s;
}
.source-link:hover { color: var(--ink); }

/* Title & desc */
.pin-title {
  font-family: 'DM Serif Display', serif;
  font-size: 26px;
  letter-spacing: -0.5px;
  line-height: 1.2;
  margin-bottom: 12px;
}

.pin-desc {
  font-size: 14px;
  color: var(--muted);
  line-height: 1.7;
  font-weight: 300;
  margin-bottom: 16px;
}

/* Tags */
.tags {
  display: flex;
  flex-wrap: wrap;
  gap: 6px;
  margin-bottom: 24px;
}

.tag {
  padding: 4px 10px;
  background: var(--bg);
  border: 1px solid var(--border);
  border-radius: 100px;
  font-size: 12px;
  color: var(--muted);
  cursor: pointer;
  transition: border-color 0.2s, color 0.2s;
}
.tag:hover { border-color: var(--ink); color: var(--ink); }

/* Author */
.author-row {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 16px 0;
  border-top: 1px solid var(--border);
  border-bottom: 1px solid var(--border);
  margin-bottom: 16px;
}

.author-avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  background: linear-gradient(135deg, var(--accent), var(--ink));
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 13px;
  font-weight: 500;
  color: var(--white);
  flex-shrink: 0;
}
.author-avatar.sm { width: 32px; height: 32px; font-size: 11px; }

.author-info {
  display: flex;
  flex-direction: column;
  gap: 2px;
  flex: 1;
}

.author-name {
  font-size: 14px;
  font-weight: 500;
  color: var(--ink);
}

.author-handle {
  font-size: 12px;
  color: var(--muted);
}

.btn-follow {
  padding: 8px 16px;
  background: transparent;
  border: 1.5px solid var(--ink);
  border-radius: 8px;
  font-family: 'DM Sans', sans-serif;
  font-size: 13px;
  font-weight: 500;
  cursor: pointer;
  transition: background 0.2s, color 0.2s;
  white-space: nowrap;
}
.btn-follow:hover { background: var(--ink); color: var(--white); }
.btn-follow.following { background: var(--bg); border-color: var(--border); color: var(--muted); }

/* Stats */
.pin-stats {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-bottom: 24px;
  font-size: 13px;
  color: var(--muted);
}

.stat-btn {
  display: flex;
  align-items: center;
  gap: 5px;
  background: none;
  border: none;
  cursor: pointer;
  font-family: 'DM Sans', sans-serif;
  font-size: 13px;
  color: var(--muted);
  transition: color 0.2s;
  padding: 0;
}
.stat-btn:hover { color: var(--ink); }
.stat-btn.liked { color: var(--accent); }
.stat-sep { color: var(--border); }
.stat-text { font-size: 13px; }

/* Comments */
.comments-section { flex: 1; }

.comments-title {
  font-size: 14px;
  font-weight: 500;
  margin-bottom: 16px;
  display: flex;
  align-items: center;
  gap: 6px;
}
.comments-title span {
  font-size: 12px;
  color: var(--muted);
  font-weight: 400;
}

.comment-input-row {
  display: flex;
  gap: 10px;
  align-items: center;
  margin-bottom: 20px;
}

.comment-input-wrap {
  position: relative;
  flex: 1;
}

.comment-input {
  width: 100%;
  padding: 10px 40px 10px 14px;
  border: 1.5px solid var(--border);
  border-radius: 100px;
  font-family: 'DM Sans', sans-serif;
  font-size: 13px;
  background: var(--bg);
  color: var(--ink);
  outline: none;
  transition: border-color 0.2s;
  box-sizing: border-box;
}
.comment-input:focus { border-color: var(--ink); }
.comment-input::placeholder { color: #C4C0BA; }

.btn-comment-send {
  position: absolute;
  right: 10px;
  top: 50%;
  transform: translateY(-50%);
  background: none;
  border: none;
  cursor: pointer;
  color: var(--accent);
  display: flex;
  align-items: center;
  padding: 0;
}

.comments-list { display: flex; flex-direction: column; gap: 16px; }

.comment {
  display: flex;
  gap: 10px;
  align-items: flex-start;
}

.comment-body { flex: 1; }

.comment-header {
  display: flex;
  align-items: baseline;
  gap: 8px;
  margin-bottom: 4px;
}

.comment-author { font-size: 13px; font-weight: 500; }

.comment-time { font-size: 11px; color: var(--muted); }

.comment-text { font-size: 13px; color: var(--ink); line-height: 1.5; margin-bottom: 6px; }

.comment-like {
  display: flex;
  align-items: center;
  gap: 4px;
  background: none;
  border: none;
  cursor: pointer;
  font-family: 'DM Sans', sans-serif;
  font-size: 11px;
  color: var(--muted);
  padding: 0;
  transition: color 0.2s;
}
.comment-like:hover { color: var(--accent); }

/* ── Related ── */
.related {
  padding: 0 24px 60px;
  max-width: 1400px;
  margin: 0 auto;
}

.related-title {
  font-family: 'DM Serif Display', serif;
  font-size: 22px;
  letter-spacing: -0.4px;
  margin-bottom: 20px;
}

.masonry { columns: 6; column-gap: 12px; }

.pin {
  break-inside: avoid;
  background: var(--white);
  border-radius: 12px;
  overflow: hidden;
  margin-bottom: 12px;
  cursor: pointer;
  border: 1px solid transparent;
  transition: transform 0.25s cubic-bezier(0.34, 1.56, 0.64, 1);
  animation: pinIn 0.4s ease both;
}
.pin:hover { transform: translateY(-4px); border-color: var(--border); }
.pin:hover .pin-overlay { opacity: 1; }
.pin:hover .pin-img { transform: scale(1.03); }

.pin-img-wrap { position: relative; overflow: hidden; }

.pin-img {
  width: 100%;
  display: block;
  transition: transform 0.3s;
}

.h-xs  { height: 120px; }
.h-sm  { height: 160px; }
.h-md  { height: 210px; }
.h-lg  { height: 260px; }
.h-xl  { height: 320px; }

.pal-1  { background: #D4C8B8; }
.pal-2  { background: #B8C8C0; }
.pal-3  { background: #C8B8A0; }
.pal-4  { background: #D8D0C8; }
.pal-5  { background: #C0C8D8; }
.pal-6  { background: #D0C0C8; }
.pal-7  { background: #C8D0C0; }
.pal-8  { background: #D0C8B0; }
.pal-9  { background: #B8C0C8; }
.pal-10 { background: #C8C0B8; }

.pin-overlay {
  position: absolute;
  inset: 0;
  background: rgba(26, 26, 24, 0.08);
  opacity: 0;
  transition: opacity 0.2s;
  display: flex;
  align-items: flex-start;
  justify-content: flex-end;
  padding: 10px;
}

.pin-save-btn {
  background: var(--white);
  border: none;
  border-radius: 100px;
  padding: 7px 14px;
  font-family: 'DM Sans', sans-serif;
  font-size: 12px;
  font-weight: 500;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 5px;
  color: var(--ink);
}
.pin-save-btn:hover { background: var(--bg); }

.pin-body { padding: 10px 12px 12px; }
.pin-label { font-size: 13px; font-weight: 500; line-height: 1.4; color: var(--ink); margin-bottom: 6px; }

.pin-meta { display: flex; align-items: center; justify-content: space-between; }

.pin-author {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 11px;
  color: var(--muted);
}

.author-dot {
  width: 18px;
  height: 18px;
  border-radius: 50%;
  background: var(--border);
  font-size: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--muted);
  font-weight: 500;
}

.pin-likes {
  display: flex;
  align-items: center;
  gap: 3px;
  font-size: 11px;
  color: var(--muted);
}

@keyframes cardIn {
  from { opacity: 0; transform: translateY(24px); }
  to   { opacity: 1; transform: translateY(0); }
}
@keyframes pinIn {
  from { opacity: 0; transform: translateY(12px); }
  to   { opacity: 1; transform: translateY(0); }
}

@media (max-width: 1200px) { .masonry { columns: 5; } }
@media (max-width: 960px)  { .masonry { columns: 4; } }
@media (max-width: 768px) {
  .pin-card { grid-template-columns: 1fr; }
  .pin-image-side { min-height: 300px; }
  .pin-image { min-height: 300px; }
  .pin-info-side { max-height: none; }
  .masonry { columns: 3; }
  .pin-wrap { padding: 16px; }
}
@media (max-width: 480px) { .masonry { columns: 2; } }
</style>
